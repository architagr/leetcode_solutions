package prunethedeadconfig

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func key(k string, used bool, kids ...*Node) *Node { return &Node{k, used, kids} }

func clone(x *Node) *Node {
	if x == nil {
		return nil
	}
	c := &Node{Key: x.Key, Used: x.Used}
	for _, k := range x.Kids {
		c.Kids = append(c.Kids, clone(k))
	}
	return c
}

func size(x *Node) int {
	if x == nil {
		return 0
	}
	s := 1
	for _, k := range x.Kids {
		s += size(k)
	}
	return s
}

// config is a random tree of n entries, at most fan children each and at
// most depth deep; fan to the power depth has to comfortably exceed n. A leaf is used with probability used; a section's own
// name never is.
func config(r *rand.Rand, n, fan, depth int, used float64) *Node {
	type at struct {
		n *Node
		d int
	}
	root := key("root", false)
	nodes := []at{{root, 0}}
	for i := 1; i < n; i++ {
		var p at
		for {
			p = nodes[r.Intn(len(nodes))]
			if len(p.n.Kids) < fan && p.d < depth {
				break
			}
		}
		c := key(fmt.Sprintf("k%d", i), false)
		p.n.Kids = append(p.n.Kids, c)
		nodes = append(nodes, at{c, p.d + 1})
	}
	var mark func(*Node)
	mark = func(x *Node) {
		if len(x.Kids) == 0 {
			x.Used = r.Float64() < used
		}
		for _, k := range x.Kids {
			mark(k)
		}
	}
	mark(root)
	return root
}

// rules is a pricing rule chain: each rule has a condition and a price, and
// an else holding the next rule. branches rules deep. Only the fallback at the
// bottom is still read by the code.
func rules(branches int) *Node {
	root := key("pricing", false)
	cur := root
	for i := 0; i < branches; i++ {
		next := key("else", false)
		cur.Kids = append(cur.Kids,
			key(fmt.Sprintf("when%d", i), false), key(fmt.Sprintf("price%d", i), false), next)
		cur = next
	}
	cur.Kids = append(cur.Kids, key("fallback", true))
	return key("root", false, key("service", true), root)
}

// PruneOnTheWayDown is the first version I wrote: drop any key nobody reads,
// deciding before looking underneath it.
func PruneOnTheWayDown(n *Node) *Node {
	kept := n.Kids[:0]
	for _, k := range n.Kids {
		if k.Used {
			kept = append(kept, PruneOnTheWayDown(k))
		}
	}
	n.Kids = kept
	return n
}

func TestOnTheWayDownDropsWhatIsRead(t *testing.T) {
	cfg := key("root", false,
		key("payments", false, key("timeout_ms", true), key("legacy_gateway", false)),
		key("old_banner", false))
	want := key("root", false, key("payments", false, key("timeout_ms", true)))
	if got := PruneOnTheWayDown(clone(cfg)); reflect.DeepEqual(got, want) {
		t.Fatal("expected deciding on the way down to lose payments.timeout_ms")
	}
	for name, fn := range map[string]func(*Node) *Node{"asking": PruneByAsking, "bottom-up": PruneBottomUp} {
		if got := fn(clone(cfg)); !reflect.DeepEqual(got, want) {
			t.Fatalf("%s kept the wrong config", name)
		}
	}
}

func TestBothAgree(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 2000; i++ {
		c := config(r, 1+r.Intn(300), 2+r.Intn(5), 9+r.Intn(4), r.Float64())
		a, b := PruneByAsking(clone(c)), PruneBottomUp(clone(c))
		if !reflect.DeepEqual(a, b) {
			t.Fatalf("case %d: asking and bottom-up disagree", i)
		}
	}
	for _, s := range shapes {
		if !reflect.DeepEqual(PruneByAsking(clone(s.cfg)), PruneBottomUp(clone(s.cfg))) {
			t.Fatalf("%s: disagree", s.name)
		}
	}
}

func TestSteps(t *testing.T) {
	for _, s := range shapes {
		visits := 0
		var used func(*Node) bool
		used = func(n *Node) bool {
			visits++
			if n.Used {
				return true
			}
			for _, k := range n.Kids {
				if used(k) {
					return true
				}
			}
			return false
		}
		var walk func(*Node)
		walk = func(n *Node) {
			if !used(n) {
				return
			}
			for _, k := range n.Kids {
				walk(k)
			}
		}
		walk(clone(s.cfg))
		t.Logf("%-9s %7d entries, %7d kept  |  entries the questions read: %10d",
			s.name, size(s.cfg), size(PruneBottomUp(clone(s.cfg))), visits)
	}
}
