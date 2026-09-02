package secondminimumnodeinabinarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	ans int // best candidate found so far for the second-minimum value; math.MaxInt64 means "none yet"
	min int // global minimum value in the tree; by the tree's invariant this always equals root.Val
)

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	if min < root.Val && root.Val < ans {
		// root.Val is strictly between the known min and our best candidate so far,
		// so it's a new best candidate. Because root.val = min(left.val, right.val)
		// holds throughout the tree, this node's value is also the minimum of its
		// entire subtree, so nothing further down can beat this candidate — prune
		// by not recursing into root.Left/root.Right.
		ans = root.Val
	} else if min == root.Val {
		// Haven't branched away from the global minimum yet: the second-minimum
		// value, if any, must be further down, so keep searching both children.
		dfs(root.Left)
		dfs(root.Right)
	}
	// else: root.Val >= ans already, so this subtree can't improve ans either — dead end.
}

func findSecondMinimumValue(root *TreeNode) int {
	min = root.Val   // the tree invariant guarantees the root always holds the global minimum
	ans = math.MaxInt64 // sentinel: no second-minimum candidate found yet
	dfs(root)
	if ans < math.MaxInt64 {
		return ans
	}
	return -1 // every value in the tree equals min, so there is no second minimum
}
