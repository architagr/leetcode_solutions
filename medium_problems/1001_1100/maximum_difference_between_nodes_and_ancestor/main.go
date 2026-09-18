package maximumdifferencebetweennodesandancestor

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func maxAncestorDiff(root *TreeNode) int {
	res = 0
	diff(root, math.MinInt, math.MaxInt)
	return res
}

func diff(node *TreeNode, ancestorMax, ancestorMin int) {
	ancestorMax, ancestorMin = maxVal(node.Val, ancestorMax), minVal(node.Val, ancestorMin)
	if node.Left == nil && node.Right == nil {
		res = maxVal(res, ancestorMax-ancestorMin)
		return
	}

	if node.Left != nil {
		diff(node.Left, ancestorMax, ancestorMin)
	}
	if node.Right != nil {
		diff(node.Right, ancestorMax, ancestorMin)
	}
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
