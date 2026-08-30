// Package queue manages challenge/queue.yaml, the ordered, resumable
// record of which already-solved LeetCode questions have had content
// generated and/or posted, in the fixed "365 Days" order.
package queue

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	StatusContentReady = "content_ready"
	StatusPosted       = "posted"
)

// Entry is one question's record in the challenge queue.
type Entry struct {
	Day        int     `yaml:"day"`
	Number     int     `yaml:"number"`
	Title      string  `yaml:"title"`
	Difficulty string  `yaml:"difficulty"`
	Folder     string  `yaml:"folder"`
	Batch      string  `yaml:"batch"`
	Status     string  `yaml:"status"`
	PostedAt   *string `yaml:"posted_at"`
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
		return nil, err
	}
	if q.NextDay == 0 {
		q.NextDay = 1
	}
	return &q, nil
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

// Append assigns the next Day number to e, sets its status to
// content_ready, appends it, and returns the assigned day.
func (q *Queue) Append(e Entry) int {
	if q.NextDay == 0 {
		q.NextDay = 1
	}
	e.Day = q.NextDay
	e.Status = StatusContentReady
	q.Entries = append(q.Entries, e)
	q.NextDay++
	return e.Day
}
