package binarytreemaximumpathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	max := -2000
	recurssivePathSum(root, &max)
	return max
}

func recurssivePathSum(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := recurssivePathSum(root.Left, max)
	right := recurssivePathSum(root.Right, max)

	(*max) = maxValue(*max,
		maxValue(root.Val,
			maxValue(
				maxValue(
					maxValue(right+root.Val, root.Val),
					maxValue(left+root.Val, root.Val),
				),
				left+right+root.Val,
			),
		),
	)

	return maxValue(maxValue(left, right)+root.Val, root.Val)

}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
