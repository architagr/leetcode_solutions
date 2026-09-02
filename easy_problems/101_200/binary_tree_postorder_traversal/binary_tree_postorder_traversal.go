package binary_tree_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	return traversal(root, arr)
}

// traversal walks the tree left, then right, then appends the current node's
// value, which is exactly what "postorder" means: a node is only recorded
// once everything beneath it already has been.
func traversal(A *TreeNode, arr []int) []int {
	// Base case: an empty subtree contributes nothing. This is also why a
	// leaf resolves immediately — both of its recursive calls land here.
	if A == nil {
		return arr
	}

	// Recurse into the left subtree first so its values are appended before
	// this node's own value.
	arr = traversal(A.Left, arr)
	// Then the right subtree, threading through the slice returned by the
	// left recursion so nothing already appended is lost.
	arr = traversal(A.Right, arr)
	// Both subtrees are fully recorded now, so the current node goes last.
	arr = append(arr, A.Val)
	return arr
}
