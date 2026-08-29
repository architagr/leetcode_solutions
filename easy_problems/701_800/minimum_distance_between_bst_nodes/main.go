package minimumdistancebetweenbstnodes

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	arr := inOrder(root)
	minVal := math.MaxInt
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

func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(inOrder(root.Left), append([]int{root.Val}, inOrder(root.Right)...)...)
}
