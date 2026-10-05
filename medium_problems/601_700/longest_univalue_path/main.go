package longestunivaluepath

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func longestUnivaluePath(root *TreeNode) int {
	res = 0
	process(root)
	return maxVal(res-1, 0)
}

func process(node *TreeNode) int {
	if node == nil {
		return 0
	}

	ans, result := 1, 1

	arr := make([]*TreeNode, 0, 2)
	if node.Left != nil {
		arr = append(arr, node.Left)
	}
	if node.Right != nil {
		arr = append(arr, node.Right)
	}

	for _, n := range arr {
		x := process(n)
		if node.Val == n.Val {
			result = maxVal(result, x+1)
			ans += x
		}
	}
	res = maxVal(res, ans)
	return result
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
