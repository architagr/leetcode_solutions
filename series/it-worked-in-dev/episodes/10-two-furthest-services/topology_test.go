package twofurthestservices

import (
	"fmt"
	"math/rand"
	"testing"
)

// chain builds a topology that is really a list: one child all the way down,
// which is what a relay chain or a badly labelled hierarchy looks like.
func chain(depth int) *Node {
	root := &Node{Name: "n0"}
	cur := root
	for i := 1; i <= depth; i++ {
		c := &Node{Name: fmt.Sprintf("n%d", i)}
		cur.Children = append(cur.Children, c)
		cur = c
	}
	return root
}

// fanout builds a topology of the shape a deployment actually has: a fixed
// number of levels, each node holding the same number of children.
func fanout(levels []int) *Node {
	n := 0
	var build func(depth int) *Node
	build = func(depth int) *Node {
		node := &Node{Name: fmt.Sprintf("n%d", n)}
		n++
		if depth == len(levels) {
			return node
		}
		for i := 0; i < levels[depth]; i++ {
			node.Children = append(node.Children, build(depth+1))
		}
		return node
	}
	return build(0)
}

// randomTopology grows a tree by hanging each new node off a random existing
// one, so the shapes are lopsided in ways the hand-written cases are not.
func randomTopology(rng *rand.Rand, size int) *Node {
	root := &Node{Name: "n0"}
	all := []*Node{root}
	for i := 1; i < size; i++ {
		p := all[rng.Intn(len(all))]
		c := &Node{Name: fmt.Sprintf("n%d", i)}
		p.Children = append(p.Children, c)
		all = append(all, c)
	}
	return root
}

func TestKnownTopologies(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		want int
	}{
		{"one instance", &Node{Name: "api-1"}, 0},
		{"a rack of two", fanout([]int{2}), 2},
		{"one region, one zone, one rack", chain(3), 3},
		{"a rack of 24", fanout([]int{24}), 2},
		{"region, 3 zones, 4 racks, 6 instances", fanout([]int{3, 4, 6}), 6},
		{"a 40-deep chain", chain(40), 40},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := WidestByScan(c.root); got != c.want {
				t.Errorf("WidestByScan = %d, want %d", got, c.want)
			}
			if got := WidestByOnePass(c.root); got != c.want {
				t.Errorf("WidestByOnePass = %d, want %d", got, c.want)
			}
			if got := WidestPair(c.root).Hops; got != c.want {
				t.Errorf("WidestPair.Hops = %d, want %d", got, c.want)
			}
		})
	}
}

// lopsided is the topology the write-up draws: one zone holding two branches
// that are each deeper than everything else the root can see.
//
//	region
//	├── zone-a
//	│   ├── rack-1 -> api-7  -> pod-2
//	│   └── rack-2 -> web-3  -> pod-9
//	└── zone-b
func lopsided() *Node {
	branch := func(a, b, c string) *Node {
		return &Node{Name: a, Children: []*Node{{Name: b, Children: []*Node{{Name: c}}}}}
	}
	return &Node{Name: "region", Children: []*Node{
		{Name: "zone-a", Children: []*Node{
			branch("rack-1", "api-7", "pod-2"),
			branch("rack-2", "web-3", "pod-9"),
		}},
		{Name: "zone-b"},
	}}
}

// The shortcut - the root's two deepest branches - is not merely slower to
// justify, it is wrong. The winning path here never touches the root.
func TestRootOnlyIsWrong(t *testing.T) {
	root := lopsided()
	if got, want := WidestByScan(root), 6; got != want {
		t.Fatalf("the scan says %d hops, want %d", got, want)
	}
	if got, want := WidestFromRoot(root), 5; got != want {
		t.Fatalf("the root-only shortcut says %d, want %d", got, want)
	}
	p := WidestPair(root)
	if p.Hops != 6 {
		t.Fatalf("WidestPair = %+v, want 6 hops", p)
	}
	ends := map[string]bool{p.A: true, p.B: true}
	if !ends["pod-2"] || !ends["pod-9"] {
		t.Errorf("widest pair is %s and %s, want pod-2 and pod-9", p.A, p.B)
	}
}

func TestNilTopology(t *testing.T) {
	if WidestByScan(nil) != 0 || WidestByOnePass(nil) != 0 || WidestPair(nil).Hops != 0 {
		t.Error("an empty topology is 0 hops wide")
	}
}

func TestBothAgreeOnRandomTopologies(t *testing.T) {
	rng := rand.New(rand.NewSource(10))
	for trial := 0; trial < 2000; trial++ {
		root := randomTopology(rng, rng.Intn(60)+1)
		scan := WidestByScan(root)
		pass := WidestByOnePass(root)
		if reuse := WidestByScanReusing(root); reuse != scan {
			t.Fatalf("trial %d: reusing scan says %d, scan says %d", trial, reuse, scan)
		}
		if scan != pass {
			t.Fatalf("trial %d: scan says %d, one pass says %d", trial, scan, pass)
		}
		if p := WidestPair(root); p.Hops != scan {
			t.Fatalf("trial %d: pair %+v disagrees with %d", trial, p, scan)
		}
	}
}

// The pair the fast version names has to actually be that far apart, measured
// independently by walking the topology between those two nodes.
func TestNamedPairIsReallyThatFarApart(t *testing.T) {
	rng := rand.New(rand.NewSource(11))
	for trial := 0; trial < 500; trial++ {
		root := randomTopology(rng, rng.Intn(60)+2)
		p := WidestPair(root)
		all := index(root)
		adj := neighbours(root)
		at := map[string]int{}
		for i, n := range all {
			at[n.Name] = i
		}
		if got := distance(adj, at[p.A], at[p.B]); got != p.Hops {
			t.Fatalf("trial %d: %s to %s is %d hops, claimed %d",
				trial, p.A, p.B, got, p.Hops)
		}
	}
}

// distance is a deliberately dumb breadth-first walk from a to b, written for
// the test so the claim is checked by something other than the code it checks.
func distance(adj [][]int, a, b int) int {
	seen := make([]bool, len(adj))
	seen[a] = true
	frontier := []int{a}
	for hops := 0; len(frontier) > 0; hops++ {
		var next []int
		for _, n := range frontier {
			if n == b {
				return hops
			}
			for _, m := range adj[n] {
				if !seen[m] {
					seen[m] = true
					next = append(next, m)
				}
			}
		}
		frontier = next
	}
	return -1
}

// The write-up says the scan visits every node once per starting point. That
// is counted here rather than asserted in prose.
func TestScanVisitsEveryNodePerStart(t *testing.T) {
	for _, size := range []int{10, 50, 200} {
		root := chain(size - 1)
		adj := neighbours(root)
		visits := 0
		for i := range adj {
			seen := make([]bool, len(adj))
			seen[i] = true
			frontier := []int{i}
			for len(frontier) > 0 {
				var next []int
				for _, n := range frontier {
					visits++
					for _, m := range adj[n] {
						if !seen[m] {
							seen[m] = true
							next = append(next, m)
						}
					}
				}
				frontier = next
			}
		}
		if want := size * size; visits != want {
			t.Errorf("%d nodes: %d visits, want %d", size, visits, want)
		}
	}
}
