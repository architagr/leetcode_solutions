package minimumdistancebetweenbstnodes

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	// No pair of nodes to compare in an empty tree or a single-node tree
	// (constraints guarantee n >= 2, so this mainly guards root == nil).
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	// In-order traversal of a BST visits every value in strictly increasing
	// order, so the minimum difference between ANY two nodes can only occur
	// between two values that end up adjacent here.
	arr := inOrder(root)
	minVal := math.MaxInt
	// Only adjacent pairs in the sorted list need checking.
	for i := 1; i < len(arr); i++ {
		minVal = min(minVal, abs(arr[i]-arr[i-1]))
	}
	return minVal
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

// inOrder returns this subtree's values in sorted order (left, node, right).
func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// Concatenate: left subtree values, then this node's value, then right
	// subtree values — the classic in-order recipe.
	return append(inOrder(root.Left), append([]int{root.Val}, inOrder(root.Right)...)...)
}
