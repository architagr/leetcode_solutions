package sumofleftleaves

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := sumOfLeftLeaves(root.Left)
	r := sumOfLeftLeaves(root.Right)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		l += root.Left.Val
	}
	return l + r
}
