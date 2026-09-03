package binarysearchtreetogreatersumtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	foo(root, 0)
	return root
}

func foo(node *TreeNode, parentSum int) int {
	if node == nil {
		return parentSum + 0
	}
	right := foo(node.Right, parentSum)
	node.Val += right
	return foo(node.Left, node.Val)
}
