package symmetrictree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric checks whether the right half, mirrored, equals the left
// half. Note that invertTree works in place: the caller's right subtree is
// left inverted after this returns. A direct mirror comparison (a.Left
// against b.Right, a.Right against b.Left) would avoid the mutation.
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(root.Left, invertTree(root.Right))
}

// isSameTree is the lockstep comparison from Same Tree.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Both-nil first, so the || below can only mean "exactly one".
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	return left && right

}

// invertTree is the in-place mirror from Invert Binary Tree.
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	root.Right, root.Left = root.Left, root.Right
	return root
}
