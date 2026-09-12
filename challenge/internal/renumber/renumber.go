package renumber

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

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
