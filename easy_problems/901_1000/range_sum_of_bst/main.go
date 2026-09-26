package rangesumofbst

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rangeSumBST sums every value in [low, high]. It visits every node and
// doesn't use the BST ordering, so it would work on any binary tree.
// Guarding the left call with root.Val > low and the right call with
// root.Val < high would skip subtrees that cannot reach the range.
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	sum := 0
	sum += rangeSumBST(root.Left, low, high)
	sum += rangeSumBST(root.Right, low, high)
	// Both bounds are inclusive.
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}
	return sum
}
