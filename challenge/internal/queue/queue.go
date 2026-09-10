// Package queue manages challenge/queue.yaml, the ordered, resumable
// record of which already-solved LeetCode questions have had content
// generated and/or posted, in the fixed "365 Days" order.
package queue

import (
	"os"
	"sort"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	StatusContentReady = "content_ready"
	StatusPosted       = "posted"
)

// The destinations an entry can be posted to. Posting state is tracked
// per destination in Entry.PostedAt, so adding a new place to post means
// adding a key here — no schema migration, no change to Status.
const (
	DestinationDiscord         = "discord"
	DestinationLinkedInMain    = "linkedin_main_account"
	DestinationLinkedInCompany = "linkedin_company_page"
	DestinationLinkedInGroup   = "linkedin_group"
	DestinationX               = "x"
	DestinationSubstack        = "substack"
)

// Entry is one question's record in the challenge queue.
//
// Status describes content generation only ("content_ready"); it says
// nothing about posting. Whether this entry has gone out to a given
// destination lives in PostedAt, keyed by destination, because one
// status field can't describe N independent destinations.
type Entry struct {
	Day        int    `yaml:"day"`
	Number     int    `yaml:"number"`
	Title      string `yaml:"title"`
	Difficulty string `yaml:"difficulty"`
	Folder     string `yaml:"folder"`
	Batch      string `yaml:"batch"`
	Status     string `yaml:"status"`
	// BuildsOn lists earlier questions whose technique this one reuses,
	// by LeetCode number. It drives the "Builds on" section in the
	// write-ups, so a reader arriving at a medium is pointed back at the
	// easies that taught the pieces.
	//
	// This field has to exist here for the values to survive at all.
	// Without it yaml.Unmarshal drops the key on load and Marshal writes
	// the file back without it, so every queue write silently deletes
	// the whole chain - which is exactly what happened to the 51 entries
	// added in 9292bce.
	BuildsOn []int              `yaml:"builds_on,omitempty"`
	PostedAt map[string]*string `yaml:"posted_at"`
}

// HasContent reports whether this entry's content has been generated.
//
// An entry is queued the moment its day number is assigned, which can be
// well before anything is written for it, so the queue routinely holds
// days with no POST_*.md on disk at all. Selecting one of those is not a
// missing-file edge case, it is the normal state of a day that has not
// been written yet — and it used to take the whole run down with it,
// because a day queued but unwritten sits in front of every written day
// behind it and every selection found it first.
func (e Entry) HasContent() bool {
	return e.Status == StatusContentReady || e.Status == StatusPosted
}

// IsPosted reports whether this entry has already gone out to
// destination. An absent key and an explicit null both mean "not yet".
func (e Entry) IsPosted(destination string) bool {
	if e.PostedAt == nil {
		return false
	}
	at, ok := e.PostedAt[destination]
	return ok && at != nil && *at != ""
}

// Queue is the full contents of challenge/queue.yaml.
type Queue struct {
	NextDay int     `yaml:"next_day"`
	Entries []Entry `yaml:"entries"`
}

// Load reads a Queue from path. A missing file is not an error — it
// returns an empty Queue starting at day 1, so a fresh repo works with
// no setup.
func Load(path string) (*Queue, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Queue{NextDay: 1}, nil
		}
		return nil, err
	}
	var q Queue
	if err := yaml.Unmarshal(data, &q); err != nil {
		var legacy legacyQueue
		if legacyErr := yaml.Unmarshal(data, &legacy); legacyErr != nil {
			return nil, err
		}
		q = legacy.normalize()
	}
	if q.NextDay == 0 {
		q.NextDay = 1
	}
	for i := range q.Entries {
		if q.Entries[i].PostedAt == nil {
			q.Entries[i].PostedAt = newPostedAt()
			continue
		}
		// Fill in destinations added since this entry was written, so an
		// older entry doesn't hide the newer destinations from view.
		for _, d := range KnownDestinations {
			if _, ok := q.Entries[i].PostedAt[d]; !ok {
				q.Entries[i].PostedAt[d] = nil
			}
		}
	}
	return &q, nil
}

// legacyQueue reads the pre-destination schema, where posted_at was a
// single nullable timestamp rather than a per-destination map. Entries
// written before this repo posted anywhere but Discord can only have
// meant Discord, so that's where a bare timestamp lands.
type legacyQueue struct {
	NextDay int `yaml:"next_day"`
	Entries []struct {
		Day        int     `yaml:"day"`
		Number     int     `yaml:"number"`
		Title      string  `yaml:"title"`
		Difficulty string  `yaml:"difficulty"`
		Folder     string  `yaml:"folder"`
		Batch      string  `yaml:"batch"`
		Status     string  `yaml:"status"`
		PostedAt   *string `yaml:"posted_at"`
	} `yaml:"entries"`
}

func (l legacyQueue) normalize() Queue {
	q := Queue{NextDay: l.NextDay}
	for _, e := range l.Entries {
		postedAt := map[string]*string{}
		if e.PostedAt != nil && *e.PostedAt != "" {
			at := *e.PostedAt
			postedAt[DestinationDiscord] = &at
		}
		q.Entries = append(q.Entries, Entry{
			Day:        e.Day,
			Number:     e.Number,
			Title:      e.Title,
			Difficulty: e.Difficulty,
			Folder:     e.Folder,
			Batch:      e.Batch,
			Status:     e.Status,
			PostedAt:   postedAt,
		})
	}
	return q
}

// NextUnposted returns up to limit entries that have not yet gone out to
// destination, oldest day first. Selection goes by day rather than file
// order, so a hand-edited queue can't change which days come next.
func (q *Queue) NextUnposted(destination string, limit int) []Entry {
	if limit <= 0 {
		return nil
	}
	pending := make([]Entry, 0, len(q.Entries))
	for _, e := range q.Entries {
		if !e.IsPosted(destination) && e.HasContent() {
			pending = append(pending, e)
		}
	}
	sort.Slice(pending, func(i, j int) bool { return pending[i].Day < pending[j].Day })
	if len(pending) > limit {
		pending = pending[:limit]
	}
	return pending
}

// MarkContentReady flips the entry for number from pending_content to
// content_ready, reporting whether such an entry was found.
//
// An entry can be queued before its content is written — that is how a
// day number gets reserved so the queue's easy/medium interleaving can
// be decided up front. Finishing that content later needs some way to
// say so, and queue-append is not it: appending again would assign a
// second day number to a question that already has one.
//
// Already-ready and already-posted entries are left alone rather than
// rewritten, so running this twice is harmless.
func (q *Queue) MarkContentReady(number int) bool {
	for i := range q.Entries {
		if q.Entries[i].Number != number {
			continue
		}
		if q.Entries[i].Status != StatusPosted {
			q.Entries[i].Status = StatusContentReady
		}
		return true
	}
	return false
}

// MarkPosted records that the entry for number went out to destination
// at time at, leaving every other destination untouched. It reports
// whether such an entry was found.
func (q *Queue) MarkPosted(number int, destination string, at time.Time) bool {
	for i := range q.Entries {
		if q.Entries[i].Number != number {
			continue
		}
		if q.Entries[i].PostedAt == nil {
			q.Entries[i].PostedAt = map[string]*string{}
		}
		stamp := at.UTC().Format(time.RFC3339)
		q.Entries[i].PostedAt[destination] = &stamp
		return true
	}
	return false
}

// Save writes the Queue to path as YAML.
func (q *Queue) Save(path string) error {
	data, err := yaml.Marshal(q)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// Has reports whether an entry for the given question number already
// exists in the queue.
func (q *Queue) Has(number int) bool {
	for _, e := range q.Entries {
		if e.Number == number {
			return true
		}
	}
	return false
}

// KnownDestinations are the places an entry can be posted, seeded into
// every new entry as explicit nulls so challenge/queue.yaml documents
// its own shape rather than showing an opaque empty map.
//
// A destination missing from this list still works — IsPosted and
// MarkPosted treat any key the same — so adding one here is about
// discoverability, not capability.
var KnownDestinations = []string{
	DestinationDiscord,
	DestinationLinkedInMain,
	DestinationLinkedInCompany,
	DestinationLinkedInGroup,
	DestinationX,
	DestinationSubstack,
}

func newPostedAt() map[string]*string {
	m := make(map[string]*string, len(KnownDestinations))
	for _, d := range KnownDestinations {
		m[d] = nil
	}
	return m
}

// Append assigns the next Day number to e, sets its status to
// content_ready, appends it, and returns the assigned day.
func (q *Queue) Append(e Entry) int {
	if q.NextDay == 0 {
		q.NextDay = 1
	}
	e.Day = q.NextDay
	e.Status = StatusContentReady
	if e.PostedAt == nil {
		e.PostedAt = newPostedAt()
	}
	q.Entries = append(q.Entries, e)
	q.NextDay++
	return e.Day
}
