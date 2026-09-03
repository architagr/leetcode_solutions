package binarysearchtreetogreatersumtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	parse(root, 0)
	return root
}

func parse(node *TreeNode, parentSum int) int {
	if node == nil {
		return parentSum + 0
	}
	right := parse(node.Right, parentSum)
	node.Val += right
	return parse(node.Left, node.Val)
}
