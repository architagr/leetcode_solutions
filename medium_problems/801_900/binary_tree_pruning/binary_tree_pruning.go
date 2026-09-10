package binary_tree_pruning

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// PruneTree exists as a wrapper because the root has no parent, and the
// pruning is done by parents. Without this, a tree of all zeroes would
// come back intact instead of empty.
func PruneTree(root *TreeNode) *TreeNode {
	temp := root
	curr := parse(temp)
	if !curr {
		root = nil
	}
	return root
}

// parse reports whether node's subtree contains a 1, and prunes as a side
// effect on the way back up.
//
// Post-order: a node cannot answer this on the way down, since it does not
// yet know what is beneath it.
//
// An alternative shape returns *TreeNode and has the caller reassign, the
// way Day 39's deletion does. That version needs no wrapper, because
// returning nil for the root handles the whole-tree case naturally.
func parse(node *TreeNode) bool {
	// An absent subtree contains no 1, so it contributes nothing to the
	// || below, and the parent's node.X = nil is then a no-op.
	if node == nil {
		return false
	}
	curr := node.Val == 1
	// The PARENT prunes. A node cannot remove itself: it has no reference
	// to its parent, and nulling a local would change nothing the caller
	// can see. The child reports, the parent clears the pointer it holds.
	right := parse(node.Right)
	if !right {
		node.Right = nil
	}
	left := parse(node.Left)
	if !left {
		node.Left = nil
	}
	// Survives if it is itself a 1, or if either subtree kept anything -
	// which is why "prune every 0" would be the wrong rule.
	return curr || right || left
}
