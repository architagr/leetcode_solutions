package binarytreeupsidedown

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// upsideDownBinaryTree does two unrelated jobs on one recursion.
//
// One is finding the new root - the deepest node on the left spine - and
// passing it upward. The other is local pointer surgery at each level.
// They share a traversal and otherwise have nothing to do with each other.
//
// Only the left child is recursed into. Right children are moved but never
// descended, which is safe because the problem guarantees every right child
// is a leaf, so reattaching one cannot drag a subtree with it.
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	// The only place a value is ever produced.
	if root == nil || root.Left == nil {
		return root
	}
	// x is a pass-through: discovered at the bottom, returned untouched
	// through every frame above. Nothing below modifies it.
	x := upsideDownBinaryTree(root.Left)
	// Named for what each node is about to become. The rewiring happens
	// after the recursive call because the child subtree has to be turned
	// over before its old parent can hang underneath it.
	newRight := root
	newNode := root.Left
	newLeft := root.Right
	newNode.Left = newLeft
	newNode.Right = newRight
	// Only takes effect for the ORIGINAL root. For every other node the
	// parent's frame immediately overwrites both pointers as part of its
	// own rewiring. The original root becomes a leaf in the flipped tree
	// and genuinely needs clearing, so this is written uniformly and
	// matters exactly once.
	root.Left = nil
	root.Right = nil
	return x
}
