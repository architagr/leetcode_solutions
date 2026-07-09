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
	var prev *TreeNode
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		helper(root.Left)

		if prev != nil {
			res = min(res, root.Val-prev.Val)
		}
		prev = root

		helper(root.Right)
	}

	helper(root)

	return res
}
