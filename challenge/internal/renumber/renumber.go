package renumber

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"leetcode_solutions/challenge/internal/hero"
	"leetcode_solutions/challenge/internal/queue"
)

// Entry is one question's intended position. Entries already in the
// queue are matched by Number; entries absent from it are inserted, so
// one plan both reorders and inserts and the queue is never left half
// reordered.
type Entry struct {
	Day        int    `json:"day"`
	Number     int    `json:"number"`
	Title      string `json:"title"`
	Difficulty string `json:"difficulty"`
	Folder     string `json:"folder"`
	Batch      string `json:"batch"`
	BuildsOn   []int  `json:"buildsOn,omitempty"`
}

// Plan is the full intended ordering. It is complete rather than a
// diff: a partial plan cannot be checked for duplicate days, and the
// whole point of validating up front is that nothing is written until
// the entire target state is known to be legal.
type Plan struct {
	Entries []Entry `json:"entries"`
}

// DayMap returns old day → new day for every entry that actually moves,
// keyed through the LeetCode number, which is the only identity that
// survives a renumber. Days that do not move are absent, so callers can
// treat "in the map" as "needs rewriting".
func DayMap(q *queue.Queue, p Plan) map[int]int {
	oldDayByNumber := make(map[int]int, len(q.Entries))
	for _, e := range q.Entries {
		oldDayByNumber[e.Number] = e.Day
	}
	m := make(map[int]int)
	for _, e := range p.Entries {
		old, ok := oldDayByNumber[e.Number]
		if !ok || old == e.Day {
			continue
		}
		m[old] = e.Day
	}
	return m
}

// Validate reports the first reason the plan must not be applied.
//
// Every check runs before anything is written. A renumber that fails
// halfway leaves post files, hero images and queue.yaml disagreeing
// about what day a question is on, and there is no operation that puts
// that back.
func Validate(q *queue.Queue, p Plan, repoRoot string) error {
	seenDay := map[int]int{}
	seenNumber := map[int]bool{}
	for _, e := range p.Entries {
		if prev, ok := seenDay[e.Day]; ok {
			return fmt.Errorf("day %d is assigned twice, to questions %d and %d", e.Day, prev, e.Number)
		}
		seenDay[e.Day] = e.Number
		if seenNumber[e.Number] {
			return fmt.Errorf("question %d appears twice in the plan", e.Number)
		}
		seenNumber[e.Number] = true
	}

	// A posted day number is public on LinkedIn and Discord. Renumber
	// around it, never through it.
	for _, e := range q.Entries {
		posted := false
		for _, d := range queue.KnownDestinations {
			if e.IsPosted(d) {
				posted = true
				break
			}
		}
		if !posted {
			continue
		}
		for _, pe := range p.Entries {
			if pe.Number == e.Number && pe.Day != e.Day {
				return fmt.Errorf("question %d is posted on day %d and cannot move to day %d", e.Number, e.Day, pe.Day)
			}
		}
	}

	dayByNumber := make(map[int]int, len(p.Entries))
	for _, e := range p.Entries {
		dayByNumber[e.Number] = e.Day
	}
	// Entries outside the plan keep the day they already have, and a
	// plan entry may legitimately build on one of them.
	for _, e := range q.Entries {
		if _, ok := dayByNumber[e.Number]; !ok {
			dayByNumber[e.Number] = e.Day
		}
	}
	for _, e := range p.Entries {
		for _, dep := range e.BuildsOn {
			depDay, ok := dayByNumber[dep]
			if !ok {
				return fmt.Errorf("question %d on day %d has builds_on %d, which is not in the queue", e.Number, e.Day, dep)
			}
			if depDay >= e.Day {
				return fmt.Errorf("question %d on day %d has builds_on %d, which lands on day %d — a forward reference", e.Number, e.Day, dep, depDay)
			}
		}
	}

	for _, e := range p.Entries {
		if e.Folder == "" {
			return fmt.Errorf("question %d on day %d has no folder", e.Number, e.Day)
		}
		info, err := os.Stat(filepath.Join(repoRoot, e.Folder))
		if err != nil || !info.IsDir() {
			return fmt.Errorf("question %d on day %d: folder %s does not exist", e.Number, e.Day, e.Folder)
		}
	}
	return nil
}

// sortedEntries returns the plan's entries in day order.
func sortedEntries(p Plan) []Entry {
	out := append([]Entry(nil), p.Entries...)
	sort.Slice(out, func(i, j int) bool { return out[i].Day < out[j].Day })
	return out
}

// Change is one question's move, for the report.
type Change struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	Folder string `json:"folder"`
	OldDay int    `json:"oldDay,omitempty"`
	NewDay int    `json:"newDay"`
}

// Result reports what a run did, or what a dry run would have done.
type Result struct {
	DryRun         bool     `json:"dryRun"`
	Moved          []Change `json:"moved"`
	Inserted       []Change `json:"inserted"`
	FilesRewritten []string `json:"filesRewritten"`
	HeroesRendered []string `json:"heroesRendered"`
	Warnings       []string `json:"warnings,omitempty"`
}

// HeroRenderer renders a hero image for one day. Apply takes it as a
// function so the tests can run without Chromium — the screenshot path
// depends on an external browser binary and is verified by hand.
type HeroRenderer func(folder string, day int, e Entry) error

// Apply validates the plan, rewrites every day reference in every
// markdown file of every queued folder, re-renders the heroes of moved
// days, and writes queue.yaml last.
//
// heroTemplate empty means skip hero rendering — the tests use that, and
// so does a run that only needs to see the file diff.
func Apply(repoRoot, queuePath string, p Plan, heroTemplate string, dryRun bool) (Result, error) {
	return ApplyWith(repoRoot, queuePath, p, dryRun, heroRendererFor(repoRoot, heroTemplate))
}

// ApplyWith is Apply with an injectable hero renderer.
func ApplyWith(repoRoot, queuePath string, p Plan, dryRun bool, render HeroRenderer) (Result, error) {
	q, err := queue.Load(queuePath)
	if err != nil {
		return Result{}, fmt.Errorf("loading queue %s: %w", queuePath, err)
	}
	if err := Validate(q, p, repoRoot); err != nil {
		return Result{}, err
	}

	res := Result{DryRun: dryRun}
	dayMap := DayMap(q, p)

	inQueue := map[int]queue.Entry{}
	for _, e := range q.Entries {
		inQueue[e.Number] = e
	}
	for _, e := range sortedEntries(p) {
		existing, ok := inQueue[e.Number]
		switch {
		case !ok:
			res.Inserted = append(res.Inserted, Change{Number: e.Number, Title: e.Title, Folder: e.Folder, NewDay: e.Day})
		case existing.Day != e.Day:
			res.Moved = append(res.Moved, Change{Number: e.Number, Title: e.Title, Folder: e.Folder, OldDay: existing.Day, NewDay: e.Day})
		}
	}

	// Rewrite references across every queued folder, not only the moved
	// ones: the stale reference lives in the folder that cites the moved
	// day, which is usually some other question entirely.
	if len(dayMap) > 0 {
		folders := map[string]bool{}
		for _, e := range q.Entries {
			folders[e.Folder] = true
		}
		for _, e := range p.Entries {
			folders[e.Folder] = true
		}
		names := make([]string, 0, len(folders))
		for f := range folders {
			names = append(names, f)
		}
		sort.Strings(names)

		for _, folder := range names {
			matches, err := filepath.Glob(filepath.Join(repoRoot, folder, "*.md"))
			if err != nil {
				return Result{}, err
			}
			sort.Strings(matches)
			for _, path := range matches {
				raw, err := os.ReadFile(path)
				if err != nil {
					return Result{}, fmt.Errorf("reading %s: %w", path, err)
				}
				updated := RewriteDayRefs(string(raw), dayMap)
				if updated == string(raw) {
					continue
				}
				rel, _ := filepath.Rel(repoRoot, path)
				res.FilesRewritten = append(res.FilesRewritten, rel)
				if dryRun {
					continue
				}
				if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
					return Result{}, fmt.Errorf("writing %s: %w", path, err)
				}
			}
		}
	}

	// Heroes carry the day number on the card, and the palette is
	// derived from the day, so a moved day's image is genuinely wrong
	// until it is re-rendered.
	if render != nil {
		for _, mv := range res.Moved {
			var entry Entry
			for _, e := range p.Entries {
				if e.Number == mv.Number {
					entry = e
					break
				}
			}
			rel := filepath.Join(mv.Folder, "HERO.png")
			if _, err := os.Stat(filepath.Join(repoRoot, rel)); err != nil {
				res.Warnings = append(res.Warnings, fmt.Sprintf("day %d (%d): no HERO.png to re-render", mv.NewDay, mv.Number))
				continue
			}
			res.HeroesRendered = append(res.HeroesRendered, rel)
			if dryRun {
				continue
			}
			if err := render(mv.Folder, mv.NewDay, entry); err != nil {
				return Result{}, fmt.Errorf("re-rendering hero for day %d (%d): %w", mv.NewDay, mv.Number, err)
			}
		}
	}

	if dryRun {
		return res, nil
	}

	dayByNumber := make(map[int]int, len(p.Entries))
	for _, e := range p.Entries {
		dayByNumber[e.Number] = e.Day
	}
	q.Reorder(dayByNumber)

	// The plan is the full intended state, not a diff, so a chain it
	// describes wins over whatever the queue currently holds. Without
	// this, correcting a builds_on means hand-editing the yaml the
	// renumber was supposed to own.
	planBuildsOn := make(map[int][]int, len(p.Entries))
	for _, e := range p.Entries {
		planBuildsOn[e.Number] = e.BuildsOn
	}
	for i := range q.Entries {
		if deps, ok := planBuildsOn[q.Entries[i].Number]; ok {
			q.Entries[i].BuildsOn = deps
		}
	}

	for _, ins := range res.Inserted {
		var e Entry
		for _, pe := range p.Entries {
			if pe.Number == ins.Number {
				e = pe
				break
			}
		}
		q.Entries = append(q.Entries, queue.Entry{
			Day:        e.Day,
			Number:     e.Number,
			Title:      e.Title,
			Difficulty: e.Difficulty,
			Folder:     e.Folder,
			Batch:      e.Batch,
			Status:     queue.StatusPendingContent,
			BuildsOn:   e.BuildsOn,
			PostedAt:   nil,
		})
	}
	q.Reorder(dayByNumber)
	if err := q.Save(queuePath); err != nil {
		return Result{}, fmt.Errorf("writing queue %s: %w", queuePath, err)
	}
	return res, nil
}

// heroRendererFor returns a renderer backed by the hero package, or nil
// when no template was given.
//
// This calls hero.RenderHTML and hero.Screenshot directly rather than
// cli.HeroGenerate, which wraps the same two steps: cli imports this
// package for the queue-renumber subcommand, so importing cli back here
// would be a cycle.
func heroRendererFor(repoRoot, heroTemplate string) HeroRenderer {
	if heroTemplate == "" {
		return nil
	}
	return func(folder string, day int, e Entry) error {
		html, err := hero.RenderHTML(heroTemplate, hero.Data{
			Day:        day,
			Total:      365,
			Topic:      e.Batch,
			Difficulty: e.Difficulty,
			Title:      e.Title,
		})
		if err != nil {
			return fmt.Errorf("rendering hero HTML: %w", err)
		}
		tmp, err := os.CreateTemp("", fmt.Sprintf("hero-day-%d-*.html", day))
		if err != nil {
			return err
		}
		tmpPath := tmp.Name()
		if _, err := tmp.WriteString(html); err != nil {
			tmp.Close()
			os.Remove(tmpPath)
			return err
		}
		if err := tmp.Close(); err != nil {
			os.Remove(tmpPath)
			return err
		}
		if err := hero.Screenshot(tmpPath, filepath.Join(repoRoot, folder, "HERO.png")); err != nil {
			return fmt.Errorf("screenshotting hero (HTML kept at %s): %w", tmpPath, err)
		}
		os.Remove(tmpPath)
		return nil
	}
}
