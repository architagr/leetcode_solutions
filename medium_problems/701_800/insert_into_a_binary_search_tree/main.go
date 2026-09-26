package insertintoabinarysearchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// insertIntoBST returns the root of the subtree after inserting val, so the
// caller can store it straight back into the child it recursed into.
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// Falling off the tree is the insertion point: an empty slot on the search
	// path is the only place a new leaf keeps every ancestor's ordering. This
	// also covers the empty tree, where the new node is the whole answer.
	if root == nil {
		return &TreeNode{Val: val}
	}
	// Only one side is ever visited. val is guaranteed absent, so the >= never
	// actually sees an equal value.
	if root.Val >= val {
		// Reassigning the child is what attaches the new node. At every level
		// but the last it writes back the pointer that was already there.
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
