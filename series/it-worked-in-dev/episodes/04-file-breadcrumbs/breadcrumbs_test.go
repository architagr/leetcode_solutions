package filebreadcrumbs

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// tree builds a folder tree `depth` levels below the root, each folder holding
// `width` children, and returns the root plus every node in it.
func tree(depth, width int) (*Node, []*Node) {
	root := &Node{Name: "root"}
	all := []*Node{root}
	frontier := []*Node{root}
	for d := 0; d < depth; d++ {
		var next []*Node
		for _, p := range frontier {
			for w := 0; w < width; w++ {
				c := &Node{Name: fmt.Sprintf("d%dw%d", d, w), Parent: p}
				p.Children = append(p.Children, c)
				all = append(all, c)
				next = append(next, c)
			}
		}
		frontier = next
	}
	return root, all
}

func TestBothAgree(t *testing.T) {
	cases := []struct{ name string; depth, width int }{
		{"root only", 0, 0},
		{"one level", 1, 3},
		{"balanced 3x3", 3, 3},
		{"wide and shallow", 1, 50},
		{"deep chain", 40, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			root, all := tree(c.depth, c.width)
			up := PathsByWalkingUp(all)
			down := PathsByCarryingDown(root)
			if !reflect.DeepEqual(up, down) {
				t.Fatalf("disagreed on %d nodes", len(all))
			}
		})
	}
}

func TestKnownPaths(t *testing.T) {
	root := &Node{Name: "docs"}
	y := &Node{Name: "2026", Parent: root}
	q := &Node{Name: "q3", Parent: y}
	f := &Node{Name: "report.pdf", Parent: q}
	root.Children = []*Node{y}
	y.Children = []*Node{q}
	q.Children = []*Node{f}

	want := map[*Node]string{
		root: "docs",
		y:    "docs/2026",
		q:    "docs/2026/q3",
		f:    "docs/2026/q3/report.pdf",
	}
	if got := PathsByCarryingDown(root); !reflect.DeepEqual(got, want) {
		t.Errorf("carrying down produced %v", got)
	}
	if got := PathsByWalkingUp([]*Node{root, y, q, f}); !reflect.DeepEqual(got, want) {
		t.Errorf("walking up produced %v", got)
	}
}

func TestBothAgreeOnRandomTrees(t *testing.T) {
	rng := rand.New(rand.NewSource(29))
	for trial := 0; trial < 500; trial++ {
		root := &Node{Name: "r"}
		all := []*Node{root}
		for i := 0; i < rng.Intn(80)+1; i++ {
			parent := all[rng.Intn(len(all))]
			c := &Node{Name: fmt.Sprintf("n%d", i), Parent: parent}
			parent.Children = append(parent.Children, c)
			all = append(all, c)
		}
		if up, down := PathsByWalkingUp(all), PathsByCarryingDown(root); !reflect.DeepEqual(up, down) {
			t.Fatalf("trial %d disagreed", trial)
		}
	}
}

// The write-up claims the upward version copies a segment once per ancestor,
// so the total segment count is the sum of every node's depth. Counted here
// rather than claimed in prose.
func TestUpwardSegmentCount(t *testing.T) {
	for _, d := range []int{10, 50, 200} {
		_, all := tree(d, 1) // a chain: depths are 0,1,2,...,d
		segments := 0
		for _, n := range all {
			for cur := n; cur != nil; cur = cur.Parent {
				segments++
			}
		}
		// nodes are at depth 0..d, each contributing depth+1 segments
		want := (d + 1) * (d + 2) / 2
		if segments != want {
			t.Errorf("depth %d: %d segments, want %d", d, segments, want)
		}
	}
}
