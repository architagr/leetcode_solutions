package deleteleaveswithagivenvalue

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if !isLeaf(root) {
		if root.Left != nil {
			root.Left = removeLeafNodes(root.Left, target)
		}
		if root.Right != nil {
			root.Right = removeLeafNodes(root.Right, target)
		}
	}
	if isLeaf(root) && root.Val == target {
		return nil
	}
	return root
}
func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}
