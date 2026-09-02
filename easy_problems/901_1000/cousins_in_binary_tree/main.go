package cousionsinbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	// Look up depth + parent independently for each target. found is
	// discarded here since the problem guarantees both x and y exist.
	depthX, parentX, _ := foo(root, 0, x)
	depthY, parentY, _ := foo(root, 0, y)
	// Cousins: same depth, but different parents (siblings share a parent
	// and are excluded by this check).
	return depthX == depthY && parentX != parentY
}

// foo walks the subtree rooted at node looking for searchVal, and reports
// the depth it was found at plus its parent's value. parent is reported as
// 0 ("no parent") both when node itself is the root-level match and when
// nothing is found down this path — safe because node values are
// constrained to be >= 1, so 0 never collides with a real value.
func foo(node *TreeNode, currentDepth, searchVal int) (depth int, parent int, found bool) {
	// Defaults for "not found down this path".
	depth = currentDepth
	parent = 0
	found = false
	if node == nil {
		return currentDepth, 0, false
	}
	if node.Val == searchVal {
		return currentDepth, 0, true
	}
	if node.Left != nil {
		// Check the child directly first: this is the only place parent
		// can become non-zero for a non-root match, since it's the one
		// vantage point where the target is visible from its own parent.
		if node.Left.Val == searchVal {
			return currentDepth + 1, node.Val, true
		}
		depth, parent, found = foo(node.Left, currentDepth+1, searchVal)
		if found {
			// Prune: no need to search the right subtree once the left
			// one already produced a match.
			return
		}
	}
	if node.Right != nil {
		if node.Right.Val == searchVal {
			return currentDepth + 1, node.Val, true
		}
		// Last statement in the function, so no "if found" guard needed
		// here — it falls straight through to the return below either way.
		depth, parent, found = foo(node.Right, currentDepth+1, searchVal)
	}
	return
}
