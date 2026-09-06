package countcompletetreenodes

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	rightCount := countNodes(root.Right)
	leftCount := countNodes(root.Left)
	return rightCount + leftCount + 1
}
