package safedeleteorder

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

// hierarchy builds a fixed-shape tree: levels[0] children under the root,
// levels[1] under each of those, and so on.
func hierarchy(levels []int) *Record {
	n := 0
	var build func(depth int) *Record
	build = func(depth int) *Record {
		r := &Record{ID: fmt.Sprintf("r%d", n)}
		n++
		if depth == len(levels) {
			return r
		}
		for i := 0; i < levels[depth]; i++ {
			r.Children = append(r.Children, build(depth+1))
		}
		return r
	}
	return build(0)
}

// thread builds a hierarchy that is a single chain - a comment on a comment on
// a comment, or a task with one subtask all the way down.
func thread(depth int) *Record {
	root := &Record{ID: "r0"}
	cur := root
	for i := 1; i <= depth; i++ {
		c := &Record{ID: fmt.Sprintf("r%d", i)}
		cur.Children = append(cur.Children, c)
		cur = c
	}
	return root
}

func randomHierarchy(rng *rand.Rand, size int) *Record {
	root := &Record{ID: "r0"}
	all := []*Record{root}
	for i := 1; i < size; i++ {
		p := all[rng.Intn(len(all))]
		c := &Record{ID: fmt.Sprintf("r%d", i)}
		p.Children = append(p.Children, c)
		all = append(all, c)
	}
	return root
}

func normalise(w Waves) []string {
	out := make([]string, 0, len(w))
	for _, wave := range w {
		ids := append([]string(nil), wave...)
		sort.Strings(ids)
		out = append(out, fmt.Sprint(ids))
	}
	return out
}

func sameWaves(a, b Waves) bool {
	x, y := normalise(a), normalise(b)
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// mixed is the hierarchy the write-up draws: a workspace with one deep project
// and one board that is already empty.
//
//	workspace
//	├── project-a
//	│   ├── board-1 -> card-9
//	│   └── board-3        (empty, deletable right now)
//	└── board-7            (empty, deletable right now)
func mixed() *Record {
	return &Record{ID: "workspace", Children: []*Record{
		{ID: "project-a", Children: []*Record{
			{ID: "board-1", Children: []*Record{{ID: "card-9"}}},
			{ID: "board-3"},
		}},
		{ID: "board-7"},
	}}
}

func TestKnownPlans(t *testing.T) {
	got := WavesByOnePass(mixed())
	want := Waves{{"card-9", "board-3", "board-7"}, {"board-1"}, {"project-a"}, {"workspace"}}
	if !sameWaves(got, want) {
		t.Errorf("one pass = %v, want %v", got, want)
	}
	if scan := WavesByRepeatedScan(mixed()); !sameWaves(scan, want) {
		t.Errorf("repeated scan = %v, want %v", scan, want)
	}
}

// Grouping by depth is safe and is not an answer to "what can I delete first".
// board-7 is empty and could go immediately; the depth plan makes it wait two
// round trips because it happens to sit near the top.
func TestDepthAnswersADifferentQuestion(t *testing.T) {
	depth := WavesByDepth(mixed())
	if len(depth[0]) != 1 || depth[0][0] != "card-9" {
		t.Fatalf("first depth wave is %v, expected just card-9", depth[0])
	}
	first := WavesByOnePass(mixed())[0]
	if len(first) != 3 {
		t.Fatalf("first real wave is %v, expected three deletable records", first)
	}
	// Both plans are safe, and both take the same number of round trips here.
	if got, want := len(depth), len(WavesByOnePass(mixed())); got != want {
		t.Errorf("depth plan has %d waves, one pass has %d", got, want)
	}
}

func TestBothPlansAgree(t *testing.T) {
	cases := []struct {
		name string
		root *Record
	}{
		{"one record", &Record{ID: "only"}},
		{"a board of cards", hierarchy([]int{20})},
		{"a workspace", hierarchy([]int{3, 5, 20})},
		{"a 40-deep thread", thread(40)},
		{"the mixed hierarchy", mixed()},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if !sameWaves(WavesByRepeatedScan(c.root), WavesByOnePass(c.root)) {
				t.Errorf("plans disagree:\n scan %v\n pass %v",
					WavesByRepeatedScan(c.root), WavesByOnePass(c.root))
			}
		})
	}
}

func TestBothPlansAgreeOnRandomHierarchies(t *testing.T) {
	rng := rand.New(rand.NewSource(16))
	for trial := 0; trial < 2000; trial++ {
		root := randomHierarchy(rng, rng.Intn(60)+1)
		if !sameWaves(WavesByRepeatedScan(root), WavesByOnePass(root)) {
			t.Fatalf("trial %d disagreed", trial)
		}
	}
}

// The plan has to be safe: when a record is deleted, every record pointing at
// it must already be gone. Checked by replaying the plan against the tree.
func TestEveryPlanIsSafeToApply(t *testing.T) {
	rng := rand.New(rand.NewSource(17))
	plans := map[string]func(*Record) Waves{
		"repeated scan": WavesByRepeatedScan,
		"one pass":      WavesByOnePass,
		"by depth":      WavesByDepth,
	}
	for trial := 0; trial < 300; trial++ {
		root := randomHierarchy(rng, rng.Intn(60)+1)
		children := map[string][]string{}
		var walk func(r *Record)
		walk = func(r *Record) {
			for _, c := range r.Children {
				children[r.ID] = append(children[r.ID], c.ID)
				walk(c)
			}
		}
		walk(root)
		for name, plan := range plans {
			gone := map[string]bool{}
			for _, wave := range plan(root) {
				for _, id := range wave {
					for _, child := range children[id] {
						if !gone[child] {
							t.Fatalf("%s, trial %d: deleted %s while %s still referenced it",
								name, trial, id, child)
						}
					}
				}
				for _, id := range wave {
					gone[id] = true
				}
			}
			if len(gone) != Count(root) {
				t.Fatalf("%s, trial %d: plan deleted %d of %d records",
					name, trial, len(gone), Count(root))
			}
		}
	}
}

// The write-up says the scan re-walks everything once per wave, so the visits
// come to waves x records. Counted rather than claimed.
func TestScanWalksEverythingPerWave(t *testing.T) {
	for _, depth := range []int{10, 50, 200} {
		root := thread(depth)
		records := Count(root)
		visits := 0
		gone := map[*Record]bool{}
		waves := 0
		for {
			var removed []*Record
			var visit func(r *Record)
			visit = func(r *Record) {
				visits++
				for _, c := range r.Children {
					visit(c)
				}
				if gone[r] {
					return
				}
				for _, c := range r.Children {
					if !gone[c] {
						return
					}
				}
				removed = append(removed, r)
			}
			visit(root)
			if len(removed) == 0 {
				break
			}
			for _, r := range removed {
				gone[r] = true
			}
			waves++
		}
		// one walk per wave, plus the final walk that finds nothing left
		if want := (waves + 1) * records; visits != want {
			t.Errorf("depth %d: %d visits over %d waves, want %d",
				depth, visits, waves, want)
		}
	}
}
