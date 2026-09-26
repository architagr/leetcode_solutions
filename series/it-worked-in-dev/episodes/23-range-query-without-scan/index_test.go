package rangequerywithoutscan

import (
	"math/rand"
	"reflect"
	"testing"
)

// insert is an ordinary unbalanced BST insert. Duplicate prices go right.
func insert(root *Node, p int) *Node {
	if root == nil {
		return &Node{Price: p}
	}
	n := root
	for {
		if p < n.Price {
			if n.Left == nil {
				n.Left = &Node{Price: p}
				return root
			}
			n = n.Left
		} else {
			if n.Right == nil {
				n.Right = &Node{Price: p}
				return root
			}
			n = n.Right
		}
	}
}

// catalog is n distinct prices inserted in random order, which gives a tree
// about 2 log n deep - what an index of products added over time looks like.
func catalog(n int, seed int64) *Node {
	r := rand.New(rand.NewSource(seed))
	var root *Node
	for _, i := range r.Perm(n) {
		root = insert(root, i*10) // prices 0, 10, 20, ... paise
	}
	return root
}

// sortedImport is the same prices inserted in increasing order, the way a
// bulk import from a sorted CSV builds it: a chain, every node a right child.
func sortedImport(n int) *Node {
	var root, last *Node
	for i := 0; i < n; i++ {
		nd := &Node{Price: i * 10}
		if root == nil {
			root = nd
		} else {
			last.Right = nd
		}
		last = nd
	}
	return root
}

func depth(n *Node) int {
	if n == nil {
		return 0
	}
	return 1 + max(depth(n.Left), depth(n.Right))
}

// visits counts nodes touched, so the claim is a count rather than a guess.
func visits(root *Node, lo, hi int, prune bool) int {
	c := 0
	var walk func(n *Node)
	walk = func(n *Node) {
		if n == nil {
			return
		}
		c++
		if !prune || n.Price > lo {
			walk(n.Left)
		}
		if !prune || n.Price < hi {
			walk(n.Right)
		}
	}
	walk(root)
	return c
}

func TestBothAgree(t *testing.T) {
	small := catalog(500, 1)
	chain := sortedImport(300)
	cases := []struct {
		root   *Node
		lo, hi int
	}{
		{nil, 0, 10}, {small, 0, 0}, {small, 4990, 4990}, {small, -5, -1}, {small, 5000, 9000},
		{small, 1000, 1500}, {small, 1005, 1015}, {small, 0, 4990}, {chain, 100, 200}, {chain, 2900, 2990},
	}
	for _, s := range shapes {
		cases = append(cases, struct {
			root   *Node
			lo, hi int
		}{s.root, s.lo, s.hi})
	}
	for _, c := range cases {
		a, b := RangeByScan(c.root, c.lo, c.hi), RangeByPruning(c.root, c.lo, c.hi)
		if !reflect.DeepEqual(a, b) {
			t.Fatalf("[%d, %d]: scan found %d prices, pruning found %d", c.lo, c.hi, len(a), len(b))
		}
	}
}

func TestVisits(t *testing.T) {
	for _, s := range shapes {
		t.Logf("%-16s depth %6d, %6d in range  |  nodes visited: scan %7d, pruning %7d",
			s.name, depth(s.root), len(RangeByPruning(s.root, s.lo, s.hi)),
			visits(s.root, s.lo, s.hi, false), visits(s.root, s.lo, s.hi, true))
	}
}
