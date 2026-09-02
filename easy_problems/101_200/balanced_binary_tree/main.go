package balancedbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	// Only the balance flag matters here; the overall tree height is discarded.
	_, ok := validateTree(root)
	return ok
}

// validateTree does a single post-order pass that returns, for the subtree rooted at
// node, both its height and whether it (and everything beneath it) is height-balanced.
// Computing both in one pass avoids recomputing subtree heights repeatedly, which is
// what makes a naive "check height() at every node" approach O(n^2) in the worst case.
func validateTree(node *TreeNode) (height int, ok bool) {
	// Base case: an empty subtree has height 0 and is trivially balanced. This is what
	// lets leaf nodes resolve cleanly with no special-casing above.
	height = 0
	ok = true
	if node == nil {
		return
	}
	rightHeight := 0
	rightHeight, ok = validateTree(node.Right)
	if !ok {
		// Short-circuit: the right subtree is already unbalanced, so this subtree
		// (and everything above it) can never be balanced either. Skip checking the
		// left subtree entirely.
		return
	}
	leftHeight := 0
	leftHeight, ok = validateTree(node.Left)
	if !ok {
		// Same short-circuit, triggered by the left subtree instead.
		return
	}
	// Both children are individually balanced (ok == true so far); now check that
	// their heights don't differ by more than 1 at this node.
	diff := leftHeight - rightHeight
	if diff > 1 || diff < -1 {
		ok = false
		return
	}
	// This node passes its own check: its height is one more than its taller child.
	height = maxValue(leftHeight, rightHeight) + 1
	return
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
