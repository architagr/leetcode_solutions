package isthishierarchyvalid

import "math"

// The price index, as loaded from its snapshot at startup. Every lookup in
// episodes 23 and 25 assumes the whole left subtree of a node is cheaper
// than it and the whole right subtree dearer. A snapshot written by an older
// build, or patched by hand, may not be.
type Node struct {
	Price       int
	Left, Right *Node
}

// ValidByChildren is the check that looks right: every node is dearer than
// its left child and cheaper than its right one. It is fast and it is wrong,
// because the rule is about whole subtrees, not children.
func ValidByChildren(n *Node) bool {
	if n == nil {
		return true
	}
	if n.Left != nil && n.Left.Price >= n.Price {
		return false
	}
	if n.Right != nil && n.Right.Price <= n.Price {
		return false
	}
	return ValidByChildren(n.Left) && ValidByChildren(n.Right)
}

// ValidBySubtrees is what I would write once the children check is known to
// be wrong: state the rule exactly. Every node must be dearer than everything
// on its left and cheaper than everything on its right.
func ValidBySubtrees(n *Node) bool {
	if n == nil {
		return true
	}
	if n.Left != nil && maxOf(n.Left) >= n.Price {
		return false
	}
	if n.Right != nil && minOf(n.Right) <= n.Price {
		return false
	}
	return ValidBySubtrees(n.Left) && ValidBySubtrees(n.Right)
}

func maxOf(n *Node) int {
	m := n.Price
	if n.Left != nil {
		m = max(m, maxOf(n.Left))
	}
	if n.Right != nil {
		m = max(m, maxOf(n.Right))
	}
	return m
}

func minOf(n *Node) int {
	m := n.Price
	if n.Left != nil {
		m = min(m, minOf(n.Left))
	}
	if n.Right != nil {
		m = min(m, minOf(n.Right))
	}
	return m
}

// ValidBySorting is day 53's approach: a tree is a BST exactly when its
// in-order sequence is strictly increasing, so collect it and scan it once.
func ValidBySorting(n *Node) bool {
	var seq []int
	var walk func(n *Node)
	walk = func(n *Node) {
		if n == nil {
			return
		}
		walk(n.Left)
		seq = append(seq, n.Price)
		walk(n.Right)
	}
	walk(n)
	for i := 1; i < len(seq); i++ {
		if seq[i-1] >= seq[i] {
			return false
		}
	}
	return true
}

// ValidByBounds carries down the range every node must fall in. Going left,
// the node's price becomes the new upper bound; going right, the new lower
// bound. Each node is checked once, against every ancestor at the same time.
func ValidByBounds(root *Node) bool {
	var check func(n *Node, lo, hi int) bool
	check = func(n *Node, lo, hi int) bool {
		if n == nil {
			return true
		}
		// Strictly inside: lo and hi are the tightest ancestors on each side.
		if n.Price <= lo || n.Price >= hi {
			return false
		}
		return check(n.Left, lo, n.Price) && check(n.Right, n.Price, hi)
	}
	return check(root, math.MinInt, math.MaxInt)
}
