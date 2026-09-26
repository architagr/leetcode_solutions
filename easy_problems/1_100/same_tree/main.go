package sametree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Both-nil has to be checked first: the || below is also true when
	// both are nil, so this order is what lets it mean "exactly one".
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		// One side has a node where the other has nothing. This is how a
		// shape difference shows up, even when the values all match.
		return false
	}
	if p.Val != q.Val {
		return false
	}
	// Both calls run before the &&, so the right side is compared even
	// when the left already failed. Inlining them into the return would
	// let && short-circuit.
	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	return left && right
}
