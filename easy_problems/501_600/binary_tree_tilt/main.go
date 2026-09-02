package binarytreetilt

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// res is shared across every recursive call via pointer, so each node
	// can add its own tilt directly into one running total instead of
	// returning a (subtreeSum, tiltSum) pair that would need merging.
	res := 0
	sum(root, &res)
	return res
}

// sum does double duty: it accumulates every node's tilt into *res as a
// side effect, and it returns the sum of the subtree rooted at node so the
// caller (the parent) can use it when computing its own tilt.
func sum(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	// Postorder: both children must be fully resolved before this node's
	// tilt (which depends on both subtree sums) can be computed.
	l := sum(node.Left, res)
	r := sum(node.Right, res)
	// This node's tilt is the absolute difference between its left and
	// right subtree sums; add it straight into the shared accumulator.
	*res += absDiff(l, r)
	// Hand the parent this whole subtree's total as a single number.
	return l + r + node.Val
}

func absDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff
}
