package univaluedbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	// An empty subtree has nothing that could disagree with its parent's
	// value, so it's trivially uni-valued.
	if root == nil {
		return true
	}
	// Assume each side is fine unless a present child proves otherwise.
	// A side with no child simply stays true.
	left, right := true, true
	if root.Left != nil {
		// A child can only be compared to its parent's value from the
		// parent's own call, since the child has no way to know which
		// node called it. Also recurse so the left child's own subtree
		// is checked for internal uni-valuedness.
		left = isUnivalTree(root.Left) && root.Left.Val == root.Val
	}
	if root.Right != nil {
		// Mirror of the left-child check above.
		right = isUnivalTree(root.Right) && root.Right.Val == root.Val
	}
	// The whole subtree is uni-valued only if both sides checked out.
	return left && right
}
