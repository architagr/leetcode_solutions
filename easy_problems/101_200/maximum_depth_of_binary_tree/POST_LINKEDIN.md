**365 Days of LeetCode Challenge — Day 1/365**

**Maximum Depth of Binary Tree** (Easy)
🔗 https://leetcode.com/problems/maximum-depth-of-binary-tree/

Find the max depth of a binary tree — but do it with a single queue and no per-level
size tracking. The trick: push a `nil` sentinel right after the root to mark "end of
this level." Every time you pop that sentinel, you've finished a level — bump the depth
counter, and if there's still work left, drop in a fresh sentinel for the next level.

```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	queue := []*TreeNode{root, nil}
	// ... pop/push with the nil-sentinel trick, see main.go for the full walkthrough
}
```

Full solution + walkthrough: `easy_problems/101_200/maximum_depth_of_binary_tree/`
