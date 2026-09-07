**365 Days of LeetCode Challenge — Day 11/365**
**Path Sum** (Easy)
🔗 https://leetcode.com/problems/path-sum/

"Root-to-leaf" means the check only happens at a genuine leaf. A node with children is never compared against the target, even when the running sum matches it exactly along the way. So the running total rides down the recursion as an extra argument, and only gets checked once you hit a node with nothing left underneath it.

```go
func hasPathSum(root *TreeNode, targetSum int) bool {
	return sum(root, targetSum, 0)
}

func sum(root *TreeNode, target, current int) bool {
	if root == nil {
		return false
	} else if root.Left != nil && root.Right != nil {
		return sum(root.Left, target, current+root.Val) || sum(root.Right, target, current+root.Val)
	} else if root.Left != nil {
		return sum(root.Left, target, current+root.Val)
	} else if root.Right != nil {
		return sum(root.Right, target, current+root.Val)
	}
	return target == root.Val+current
}
```

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/path_sum/SOLUTION.md
