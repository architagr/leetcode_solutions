package isthistreeasubtree

import (
	"math/rand"
	"strings"
	"testing"
)

func n(kind string, kids ...*Node) *Node { return &Node{kind, kids} }

var kinds = []string{"div", "div", "div", "span", "img", "a", "button", "ul", "li", "p", "h2", "icon"}

// random is a component tree of about size nodes, at most 6 children each.
func random(r *rand.Rand, size int) *Node {
	root := n(kinds[r.Intn(len(kinds))])
	nodes := []*Node{root}
	for len(nodes) < size {
		p := nodes[r.Intn(len(nodes))]
		if len(p.Kids) >= 6 {
			continue
		}
		c := n(kinds[r.Intn(len(kinds))])
		p.Kids = append(p.Kids, c)
		nodes = append(nodes, c)
	}
	return root
}

func clone(x *Node) *Node {
	c := &Node{Kind: x.Kind}
	for _, k := range x.Kids {
		c.Kids = append(c.Kids, clone(k))
	}
	return c
}

// lastLeaf is the last component in preorder, the one a walk reaches last.
func lastLeaf(x *Node) *Node {
	for len(x.Kids) > 0 {
		x = x.Kids[len(x.Kids)-1]
	}
	return x
}

func size(x *Node) int {
	s := 1
	for _, k := range x.Kids {
		s += size(k)
	}
	return s
}

// grid is a page of count product cards built from one template, so every
// card has the same components and only the snippet differs, in its last one.
func grid(r *rand.Rand, count, cardSize int) (*Node, *Node) {
	card := random(r, cardSize)
	card.Kind = "card"
	page := n("main")
	for i := 0; i < count; i++ {
		page.Kids = append(page.Kids, clone(card))
	}
	snip := clone(card)
	lastLeaf(snip).Kind = "badge" // the new variant: one component different
	return page, snip
}

// thread is a page holding one reply chain depth comments long, and the new
// render of its last snipDepth comments, in which the bottom one was edited.
func thread(depth, snipDepth int) (*Node, *Node) {
	chain := func(d int) (*Node, *Node) {
		root := n("comment", n("avatar"), n("body"))
		cur := root
		for i := 1; i < d; i++ {
			c := n("comment", n("avatar"), n("body"))
			cur.Kids = append(cur.Kids, c)
			cur = c
		}
		return root, cur
	}
	page, _ := chain(depth)
	snip, last := chain(snipDepth)
	last.Kids = append(last.Kids, n("edited"))
	return n("main", page), snip
}

func TestFlatTextFalsePositive(t *testing.T) {
	// card(img, p) is not in the page, which has card(img(p)).
	page := n("main", n("card", n("img", n("p"))))
	snip := n("card", n("img"), n("p"))
	flat := func(x *Node) string {
		var b strings.Builder
		var w func(*Node)
		w = func(x *Node) {
			b.WriteString(x.Kind + ",")
			for _, k := range x.Kids {
				w(k)
			}
		}
		w(x)
		return b.String()
	}
	if !strings.Contains(flat(page), flat(snip)) {
		t.Fatal("expected the kinds-only text to match")
	}
	for name, fn := range funcs {
		if fn(page, snip) {
			t.Fatalf("%s: claims card(img, p) is in a page with only card(img(p))", name)
		}
	}
}

var funcs = map[string]func(*Node, *Node) bool{
	"walk": ContainsByWalk, "text": ContainsByText, "hash": ContainsByHash,
}

func TestAllAgree(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 2000; i++ {
		page := random(r, 1+r.Intn(200))
		var snip *Node
		if r.Intn(2) == 0 {
			// a real subtree of the page, sometimes altered
			var all []*Node
			var w func(*Node)
			w = func(x *Node) {
				all = append(all, x)
				for _, k := range x.Kids {
					w(k)
				}
			}
			w(page)
			snip = clone(all[r.Intn(len(all))])
			if r.Intn(2) == 0 {
				lastLeaf(snip).Kind = "badge"
			}
		} else {
			snip = random(r, 1+r.Intn(4))
		}
		want := ContainsByWalk(page, snip)
		for name, fn := range funcs {
			if fn(page, snip) != want {
				t.Fatalf("case %d: %s says %v, walk says %v", i, name, !want, want)
			}
		}
	}
	for _, s := range shapes {
		for name, fn := range funcs {
			if fn(s.page, s.snip) != s.want {
				t.Fatalf("%s: %s got %v", s.name, name, !s.want)
			}
		}
	}
}

func TestSteps(t *testing.T) {
	for _, s := range shapes {
		cmp := 0
		var sameCount func(a, b *Node) bool
		sameCount = func(a, b *Node) bool {
			cmp++
			if a.Kind != b.Kind || len(a.Kids) != len(b.Kids) {
				return false
			}
			for i := range a.Kids {
				if !sameCount(a.Kids[i], b.Kids[i]) {
					return false
				}
			}
			return true
		}
		anchors := 0
		var w func(*Node)
		w = func(x *Node) {
			if x.Kind == s.snip.Kind {
				anchors++
				sameCount(x, s.snip)
			}
			for _, k := range x.Kids {
				w(k)
			}
		}
		w(s.page)
		t.Logf("%-8s page %7d, snippet %5d, anchors %6d  |  nodes the walk compares: %10d",
			s.name, size(s.page), size(s.snip), anchors, cmp)
	}
}
