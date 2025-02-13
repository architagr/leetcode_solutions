package binarytreeupsidedown

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	x := upsideDownBinaryTree(root.Left)
	newRight := root
	newNode := root.Left
	newLeft := root.Right
	newNode.Left = newLeft
	newNode.Right = newRight
	root.Left = nil
	root.Right = nil
	return x
}
