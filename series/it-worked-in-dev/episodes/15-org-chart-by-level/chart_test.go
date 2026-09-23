package orgchartbylevel

import (
	"fmt"
	"reflect"
	"testing"
)

// org builds a chart `depth` levels deep where everybody who is not a leaf has
// `branch` direct reports. Names carry their level and their position so a
// wrong level shows up as a wrong name rather than as a silent reordering.
func org(depth, branch int) *Person {
	n := 0
	var build func(level int) *Person
	build = func(level int) *Person {
		n++
		p := &Person{Name: fmt.Sprintf("L%d-%d", level, n)}
		if level == depth {
			return p
		}
		p.Reports = make([]*Person, 0, branch)
		for i := 0; i < branch; i++ {
			p.Reports = append(p.Reports, build(level+1))
		}
		return p
	}
	return build(0)
}

// chain is the reorg nobody meant to ship: everyone has exactly one report.
func chain(depth int) *Person { return org(depth, 1) }

// flat is the other end: one person at the top and everybody else under them.
func flat(width int) *Person {
	root := &Person{Name: "L0-1"}
	for i := 0; i < width; i++ {
		root.Reports = append(root.Reports, &Person{Name: fmt.Sprintf("L1-%d", i+1)})
	}
	return root
}

func count(p *Person) int {
	if p == nil {
		return 0
	}
	n := 1
	for _, r := range p.Reports {
		n += count(r)
	}
	return n
}

func TestBothWalksAgree(t *testing.T) {
	for _, tree := range []*Person{
		nil, org(0, 3), org(1, 4), org(3, 3), org(5, 2), chain(40), flat(500),
	} {
		a := LevelsByDepthKeyedWalk(tree)
		b := LevelsByQueue(tree)
		if len(a) == 0 && len(b) == 0 {
			continue
		}
		if !reflect.DeepEqual(a, b) {
			t.Fatalf("the two walks disagree:\n queue %v\n walk  %v", b, a)
		}
	}
}

// Every person appears exactly once, at exactly one level. A level-order that
// drops or duplicates somebody is the failure that renders as a missing row.
func TestEveryPersonAppearsOnce(t *testing.T) {
	root := org(4, 3)
	levels := LevelsByQueue(root)
	seen := map[string]int{}
	for _, level := range levels {
		for _, name := range level {
			seen[name]++
		}
	}
	if len(seen) != count(root) {
		t.Fatalf("%d distinct names for %d people", len(seen), count(root))
	}
	for name, n := range seen {
		if n != 1 {
			t.Fatalf("%s appears %d times", name, n)
		}
	}
}

// The streaming version stops when the renderer stops asking, and the answer
// it has produced so far is identical to the first k levels of the whole.
func TestStreamingStopsAndMatches(t *testing.T) {
	root := org(6, 3)
	all := LevelsByQueue(root)
	for k := 1; k <= len(all); k++ {
		got := FirstLevelsByQueue(root, k)
		if !reflect.DeepEqual(got, all[:k]) {
			t.Fatalf("k=%d: streaming gave %d levels, want the first %d",
				k, len(got), k)
		}
	}
}

// The number the episode is about. Not a duration - a count of people each
// walk has to look at before level k is finished.
func TestVisitsBeforeLevelComplete(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *Person
		k    int
	}{
		{"org 5k, level 1", org(6, 4), 1},
		{"org 5k, level 2", org(6, 4), 2},
		{"org 56k, level 1", org(6, 6), 1},
		{"org 56k, level 2", org(6, 6), 2},
		{"flat 50k, level 1", flat(50000), 1},
		{"chain 2000, level 1", chain(2000), 1},
	} {
		lastSeen, declared, queued := VisitsBeforeLevelComplete(tc.root, tc.k)
		t.Logf("%-20s  walk has it at %6d, can say so at %6d  |  queue %6d",
			tc.name, lastSeen, declared, queued)
	}
}
