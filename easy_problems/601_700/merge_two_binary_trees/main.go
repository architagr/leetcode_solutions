package mergetwobinarytrees

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// If one side is missing at this position, the merged subtree is just
	// whatever the other tree already has here — reused as-is, no new nodes.
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// Both nodes exist: this is the only case where a new node is allocated,
	// and its value is simply the sum of the two overlapping values.
	root := &TreeNode{Val: root1.Val + root2.Val}
	// Recurse in preorder so the merged parent exists before its children
	// are attached; each recursive call resolves to either a brand new
	// summed node or a reused subtree from whichever side is non-nil.
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}
