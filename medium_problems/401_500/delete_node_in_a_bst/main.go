package deletenodeinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// Search half: one comparison picks the only direction the key could
	// be in. The result is reassigned rather than discarded because a
	// deletion can change which node roots a subtree - that reassignment
	// is also what makes the leaf case below actually take effect.
	// delete from the right subtree
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val { // delete from the left subtree
		root.Left = deleteNode(root.Left, key)
	} else { // delete the current node
		// the node is a leaf
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil { // the node is not a leaf and has a right child
			// The node is never removed. Its value is overwritten with the
			// in-order successor - the only value besides the predecessor
			// that can fill the hole and keep the in-order sequence sorted
			// - and the duplicate below is deleted instead.
			//
			// This terminates because the successor is the leftmost node of
			// the right subtree and therefore has no left child, so each
			// step recurses into a strictly easier case.
			root.Val = successor(root)
			root.Right = deleteNode(root.Right, root.Val)
		} else { // the node is not a leaf, has no right child, and has a left child
			// Needed as its own branch: successor() does root = root.Right
			// unconditionally and would panic here. The predecessor is the
			// mirror answer and equally valid.
			root.Val = predecessor(root)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

func successor(root *TreeNode) int {
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func predecessor(root *TreeNode) int {
	root = root.Left
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}
