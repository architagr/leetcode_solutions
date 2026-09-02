package minimumabsolutedifferenceinbst

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt
	// prev is the previously-visited node in in-order sequence, i.e. the
	// current node's in-order predecessor. nil until the first node is visited.
	var prev *TreeNode
	// helper performs a standard in-order traversal (left, node, right), which
	// for a BST visits values in strictly increasing sorted order. That means
	// the minimum gap between ANY two nodes must occur between two in-order
	// adjacent nodes, so comparing each node only to its immediate
	// predecessor is enough — no need to compare every pair.
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		helper(root.Left)

		// Skip the comparison on the very first node visited (no predecessor
		// yet). Since traversal order is sorted, root.Val >= prev.Val always,
		// so no abs() is needed.
		if prev != nil {
			res = min(res, root.Val-prev.Val)
		}
		// This node becomes the predecessor for whichever node is visited next.
		prev = root

		helper(root.Right)
	}

	helper(root)

	return res
}
