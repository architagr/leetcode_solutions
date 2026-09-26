package maximumdifferencebetweennodesandancestor

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// res is package-level and reset on every call, which is fine sequentially
// but would be shared by concurrent calls.
var res int

// maxAncestorDiff: any ancestor/descendant pair lies on one root-to-leaf
// path, and the widest gap on a path is its max minus its min.
func maxAncestorDiff(root *TreeNode) int {
	res = 0
	// Extreme seeds, replaced by the root's own value on the first fold.
	diff(root, math.MinInt, math.MaxInt)
	return res
}

// diff carries the path's max and min down as parameters, so each branch
// keeps the range of its own path only.
func diff(node *TreeNode, ancestorMax, ancestorMin int) {
	ancestorMax, ancestorMin = maxVal(node.Val, ancestorMax), minVal(node.Val, ancestorMin)
	// Scoring only at leaves is enough: going down, the max never falls and
	// the min never rises, so a leaf's gap is the widest on its path.
	if node.Left == nil && node.Right == nil {
		res = maxVal(res, ancestorMax-ancestorMin)
		return
	}

	// diff reads node.Val first, so nil children must be filtered here.
	if node.Left != nil {
		diff(node.Left, ancestorMax, ancestorMin)
	}
	if node.Right != nil {
		diff(node.Right, ancestorMax, ancestorMin)
	}
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
