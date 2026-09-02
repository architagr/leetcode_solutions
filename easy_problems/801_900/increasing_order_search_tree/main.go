package increaseordersearchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	// In-order traversal of a BST visits nodes in ascending value order,
	// so this is already the exact node order the final vine needs.
	inorder := inOrder(root)
	// Rewire consecutive nodes with .Right only, turning the sorted list
	// order into the "no left child, one right child" chain the problem wants.
	for i := 1; i < len(inorder); i++ {
		inorder[i-1].Right = inorder[i]
	}
	// The smallest node (first in ascending order) becomes the new root.
	return inorder[0]
}

func inOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	res := make([]*TreeNode, 0)
	// Recurse on the original left/right children before root is touched,
	// so each call still sees the tree's original shape to descend into.
	left := inOrder(root.Left)
	right := inOrder(root.Right)

	// Detach root from its old children now that both subtrees have already
	// been walked and collected. This leaves every node fully unlinked when
	// it's handed back, so increasingBST's relinking pass starts clean with
	// no stale pointers left over from the original tree shape.
	root.Left = nil
	root.Right = nil
	// Standard in-order assembly: everything from the left subtree, then
	// root itself, then everything from the right subtree.
	res = append(left, root)
	res = append(res, right...)
	return res
}
