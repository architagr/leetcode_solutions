package subtreeofanothertree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSubtree searches root for a node that, taken as its own tree, is
// structurally identical to subRoot.
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// An empty subRoot is trivially a subtree of anything.
	if subRoot == nil {
		return true
	}
	// root is empty but subRoot isn't, so there's nothing left to match against.
	if root == nil {
		return subRoot == nil
	}

	// Try root itself as the anchor: only pay for the full structural
	// comparison (equalBinaryTree) if the root values even agree first.
	if root.Val == subRoot.Val && equalBinaryTree(root, subRoot) {
		return true
	}
	// root wasn't a match; look for an anchor deeper in either subtree.
	// || short-circuits, so the right subtree is only searched if the
	// left subtree search doesn't already find a match.
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// equalBinaryTree reports whether root and subRoot have the exact same
// shape and node values, walked in lockstep.
func equalBinaryTree(root *TreeNode, subRoot *TreeNode) bool {
	// Both sides empty at the same position -> equal so far.
	if root == nil {
		return subRoot == nil
	}
	// root is empty but subRoot isn't (root == nil is already false here).
	if subRoot == nil {
		return root == nil
	}
	// Values differ at this position: not equal, stop immediately.
	if root.Val != subRoot.Val {
		return false
	}
	// Both children must match, left-to-left and right-to-right.
	return equalBinaryTree(root.Left, subRoot.Left) && equalBinaryTree(root.Right, subRoot.Right)
}
