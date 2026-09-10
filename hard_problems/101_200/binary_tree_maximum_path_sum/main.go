package binarytreemaximumpathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	// Below anything reachable: values are >= -1000 and a path is
	// non-empty, so the first real candidate always replaces this. A 0
	// seed would be wrong on an all-negative tree, where the answer is
	// the least negative single node.
	max := -2000
	recurssivePathSum(root, &max)
	return max
}

// recurssivePathSum computes two different things per node, and keeping
// them apart is the whole problem.
//
// As the TOP of a path, a node may use both children: left+right+val is a
// candidate answer, recorded into max.
//
// As a LINK in some ancestor's path, it may use only one child, because
// the ancestor attaches above and a path cannot contain a node with three
// neighbours. That one-sided value is what gets returned.
//
// Returning the two-sided value instead is the classic bug here: the
// parent then builds a "path" that forks.
func recurssivePathSum(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := recurssivePathSum(root.Left, max)
	right := recurssivePathSum(root.Right, max)

	// Candidate for the answer. The nesting is equivalent to the shorter
	// root.Val + max(0, left) + max(0, right) - each comparison against
	// root.Val alone is what lets a negative branch be dropped.
	(*max) = maxValue(*max,
		maxValue(root.Val,
			maxValue(
				maxValue(
					maxValue(right+root.Val, root.Val),
					maxValue(left+root.Val, root.Val),
				),
				left+right+root.Val,
			),
		),
	)

	// Equivalent to root.Val + max(0, left, right): take the better single
	// branch, or neither if both hurt.
	return maxValue(maxValue(left, right)+root.Val, root.Val)

}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
