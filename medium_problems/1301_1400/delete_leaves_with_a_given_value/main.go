package deleteleaveswithagivenvalue

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// removeLeafNodes returns the subtree with target leaves removed, or nil
// if the node itself goes. The parent stores the result, so returning nil
// is how a node deletes itself, the root included.
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	// Children first (post-order), so every deletion below has happened by
	// the time this node looks at itself.
	if !isLeaf(root) {
		if root.Left != nil {
			root.Left = removeLeafNodes(root.Left, target)
		}
		if root.Right != nil {
			root.Right = removeLeafNodes(root.Right, target)
		}
	}
	// isLeaf is asked again on purpose: it now sees the children after
	// deletion. A parent whose children were all removed becomes a leaf
	// here, which is the whole "repeat until you cannot" cascade in one pass.
	if isLeaf(root) && root.Val == target {
		return nil
	}
	return root
}

// isLeaf is nil-safe: a nil node is not a leaf.
func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}
