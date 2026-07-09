package binarytreetilt

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	sum(root, &res)
	return res
}

func sum(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	l := sum(node.Left, res)
	r := sum(node.Right, res)
	*res += absDiff(l, r)
	return l + r + node.Val
}

func absDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff
}
