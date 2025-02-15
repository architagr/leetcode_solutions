package secondminimumnodeinabinarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	ans int
	min int
)

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	if min < root.Val && root.Val < ans {
		ans = root.Val
	} else if min == root.Val {
		dfs(root.Left)
		dfs(root.Right)
	}
}

func findSecondMinimumValue(root *TreeNode) int {
	min = root.Val
	ans = math.MaxInt64
	dfs(root)
	if ans < math.MaxInt64 {
		return ans
	}
	return -1
}
