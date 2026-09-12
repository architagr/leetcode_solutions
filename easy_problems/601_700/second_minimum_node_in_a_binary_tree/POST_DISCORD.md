**365 Days of LeetCode Challenge — Day 70/365**
**Second Minimum Node In a Binary Tree** (Easy)
🔗 https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/

Every node here has 0 or 2 children, and its value is the smaller of the two. That rule alone means the root holds the global minimum, and any node's value is the minimum of its own subtree. So the moment we hit a node above the minimum we record it and stop descending, because nothing underneath it can be smaller. The pruning is the fun bit.

```go
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
```

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/second_minimum_node_in_a_binary_tree/SOLUTION.md
