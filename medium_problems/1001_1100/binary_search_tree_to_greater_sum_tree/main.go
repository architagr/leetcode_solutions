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
	// Traverse right subtree first to process larger values
	right := parse(node.Right, parentSum)
	// Add the accumulated sum of all greater nodes
	node.Val += right
	// Traverse left subtree with updated sum
	return parse(node.Left, node.Val)
}
