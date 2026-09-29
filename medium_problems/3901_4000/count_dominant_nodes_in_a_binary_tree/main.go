package countdominantnodesinabinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countDominantNodes(root *TreeNode) int {
	result := 0

	var parse func(*TreeNode) int
	var maxVal func(a, b int) int

	parse = func(n *TreeNode) (max int) {
		if n == nil {
			return 0
		}
		left := parse(n.Left)
		right := parse(n.Right)
		max = maxVal(maxVal(left, right), n.Val)
		if n.Val == max {
			result++
		}
		return max
	}

	maxVal = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	parse(root)
	return result
}
