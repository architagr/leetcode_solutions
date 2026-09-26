package invertbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree mirrors the tree in place and returns the same root.
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// Post-order: both subtrees come back already mirrored, so the swap
	// below only has to put the two halves on opposite sides.
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	// Swapping pointers moves whole subtrees, children and all; no node is
	// copied. The right-hand side is evaluated first, so no temp is needed.
	root.Right, root.Left = root.Left, root.Right
	return root
}
