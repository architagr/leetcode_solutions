package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	// current starts at 0: no value has been accumulated before we've even
	// looked at the root.
	return sum(root, targetSum, 0)
}

// sum carries the running total accumulated from the root down to root's
// parent in current, and only compares it against target once it reaches a
// genuine leaf (a node with no children) -- that's the only point along a
// root-to-leaf path where a sum "counts".
func sum(root *TreeNode, target, current int) bool {
	if root == nil {
		// Empty subtree: also covers the fully empty tree, and makes sure a
		// nonexistent child is never treated as a valid path.
		return false
	} else if root.Left != nil && root.Right != nil {
		// Two children: not a leaf, so don't compare yet. Add this node's
		// value to the running total and try both children, short-circuiting
		// on || the moment either side finds a matching path.
		return sum(root.Left, target, current+root.Val) || sum(root.Right, target, current+root.Val)
	} else if root.Left != nil {
		// Only a left child: still not a leaf (root.Right == nil doesn't
		// make this node a leaf), so descend into the one real child instead
		// of falling through to the leaf check below.
		return sum(root.Left, target, current+root.Val)
	} else if root.Right != nil {
		// Mirror of the left-only case above.
		return sum(root.Right, target, current+root.Val)
	}
	// Both children are nil: root is a genuine leaf, so this is the only
	// place the accumulated sum is actually checked against target.
	return target == root.Val+current
}
