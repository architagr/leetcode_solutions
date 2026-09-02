package searchinabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	// Ran off the tree without finding val.
	if root == nil {
		return nil
	}
	// BST property: everything in root.Right is >= root.Val, so if root.Val
	// is already too big, val (if present) can only live in the left subtree.
	// The right subtree is ruled out without ever visiting it.
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	// Mirror image: root.Val is too small, so only the right subtree is
	// still worth checking.
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	// Neither comparison fired, so root.Val == val. Return root itself (not
	// just a bool) so its Left/Right pointers carry the whole subtree along,
	// which is what the problem asks for.
	return root
}
