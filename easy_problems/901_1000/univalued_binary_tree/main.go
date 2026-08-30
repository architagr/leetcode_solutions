package univaluedbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left, right := true, true
	if root.Left != nil {
		left = isUnivalTree(root.Left) && root.Left.Val == root.Val
	}
	if root.Right != nil {
		right = isUnivalTree(root.Right) && root.Right.Val == root.Val
	}
	return left && right
}
