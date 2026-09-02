package minimumdepthofbinarytree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	// An empty subtree has no depth. This 0 sentinel also lets a parent
	// tell "no subtree here" apart from "a real subtree of depth 1".
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	// Both children exist (neither depth came back as the nil-sentinel 0):
	// this is a real branching node, so the shortest path takes the
	// shallower side.
	if left != 0 && right != 0 {
		return minVal(left, right) + 1
	} else if left != 0 {
		// Only the left child exists; a 0 on the right means "no subtree",
		// not "a zero-length path", so it must never win the comparison.
		return left + 1
	}
	// Falls through when only the right child exists, or when both are
	// nil (a true leaf), in which case right is 0 and this yields 1.
	return right + 1
}

// minVal returns the smaller of a and b.
func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
