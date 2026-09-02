package evaluatebooleanbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Mirrors LeetCode's encoding: leaves carry FALSE/TRUE, internal nodes carry
// the OR/AND operator to apply to their two children's evaluations.
const (
	FALSE = 0
	TRUE  = 1
	OR    = 2
	AND   = 3
)

func evaluateTree(root *TreeNode) bool {
	// The tree is guaranteed "full" (every node has 0 or 2 children), so
	// checking that both children are nil is enough to identify a leaf -
	// there's no partial-children case to handle separately.
	if root.Left == nil && root.Right == nil {
		return root.Val == TRUE
	}

	// Post-order: evaluate both subtrees before this node can combine
	// anything, since the operator needs both children's results.
	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)
	if root.Val == OR {
		return left || right
	}
	// Only AND is left once OR and the leaf case are ruled out, since the
	// problem guarantees non-leaf nodes are only ever OR or AND.
	return left && right
}
