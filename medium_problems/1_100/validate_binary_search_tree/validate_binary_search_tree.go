package validate_binary_search_tree

import "leetcode_solutions/common"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBst(root *TreeNode) bool {
	arr := InorderTraversal(root)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}
func InorderTraversal(A *TreeNode) []int {
	arr := make([]int, 0)
	return inorderTraversal(A, arr)
}

func inorderTraversal(A *TreeNode, arr []int) []int {
	if A == nil {
		return arr
	}

	arr = inorderTraversal(A.Left, arr)
	arr = append(arr, A.Val)
	arr = inorderTraversal(A.Right, arr)
	return arr
}

func IsValidBstApproch1(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return check(root) && IsValidBstApproch1(root.Right) && IsValidBstApproch1(root.Left)
}

func check(root *TreeNode) bool {
	if root == nil {
		return true
	}
	max := common.MinInt()
	min := common.MaxInt()

	max = getMax(root.Left, max)
	min = getMin(root.Right, min)
	if root.Val > max && root.Val < min {
		return true
	}
	return false
}
func getMax(root *TreeNode, max int) int {
	if root == nil {
		return max
	}
	max = maxValue(max, root.Val)
	max = maxValue(max, getMax(root.Left, max))
	return maxValue(max, getMax(root.Right, max))
}

func getMin(root *TreeNode, min int) int {
	if root == nil {
		return min
	}
	min = minValue(min, root.Val)
	min = minValue(min, getMin(root.Left, min))
	return minValue(min, getMin(root.Right, min))
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
