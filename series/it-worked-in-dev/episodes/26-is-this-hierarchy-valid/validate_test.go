package isthishierarchyvalid

import "testing"

// balanced builds a perfectly balanced BST over prices lo..hi step 10.
func balanced(lo, hi int) *Node {
	if lo > hi {
		return nil
	}
	mid := lo + ((hi-lo)/20)*10
	return &Node{Price: mid, Left: balanced(lo, mid-10), Right: balanced(mid+10, hi)}
}

func count(n *Node) int {
	if n == nil {
		return 0
	}
	return 1 + count(n.Left) + count(n.Right)
}

// chain is a valid BST that is one long right spine: a snapshot of an index
// built from a sorted import.
func chain(n int) *Node {
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

// corruptDeep takes a valid tree and changes the dearest price in the root's
// left subtree to something dearer than the root. Against its own parent and
// children it still looks right - it is a right child and bigger than its
// parent, and still bigger than its own left child - but the root needs
// everything on its left to be cheaper. This is what a hand-patched snapshot
// looks like.
func corruptDeep(root *Node) *Node {
	n := root.Left
	for n.Right != nil {
		n = n.Right
	}
	n.Price = root.Price + 5
	return root
}

// corruptEarly breaks the rule at the root's left child, the first thing any
// top-down check looks at.
func corruptEarly(root *Node) *Node {
	root.Left.Price = root.Price + 1
	return root
}

func TestTheChildrenCheckAcceptsACorruptTree(t *testing.T) {
	tr := corruptDeep(balanced(0, 10230))
	if !ValidByChildren(tr) {
		t.Fatal("the children check caught it; the episode says it does not")
	}
	for name, ok := range map[string]bool{
		"subtrees": ValidBySubtrees(tr), "sorting": ValidBySorting(tr), "bounds": ValidByBounds(tr),
	} {
		if ok {
			t.Fatalf("%s accepted a corrupt tree", name)
		}
	}
	t.Logf("%d products, one price on the wrong side of the root: children check says valid", count(tr))
}

func TestCorrectVersionsAgree(t *testing.T) {
	trees := []*Node{nil, {Price: 5}, balanced(0, 990), chain(300),
		corruptDeep(balanced(0, 9990)), corruptEarly(balanced(0, 990)),
		{Price: 5, Left: &Node{Price: 5}}, // a duplicate is not strictly smaller
	}
	for _, s := range shapes {
		if count(s.root) <= 5000 {
			trees = append(trees, s.root)
		}
	}
	for i, tr := range trees {
		a, b, c := ValidBySubtrees(tr), ValidBySorting(tr), ValidByBounds(tr)
		if a != b || b != c {
			t.Fatalf("tree %d: subtrees %v, sorting %v, bounds %v", i, a, b, c)
		}
	}
}
