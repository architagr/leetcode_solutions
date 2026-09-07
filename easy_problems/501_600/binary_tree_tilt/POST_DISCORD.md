**365 Days of LeetCode Challenge — Day 22/365**
**Binary Tree Tilt** (Easy)
🔗 https://leetcode.com/problems/binary-tree-tilt/

Each node's tilt is the absolute difference between everything under its left child and everything under its right child, not just the two direct kids. Redoing that sum from scratch at every node wastes work, so compute each subtree's sum once, bottom-up, and add every node's tilt into a shared running total as you go. The fun bit is that one recursive call does both jobs at once, accumulate and return, without them getting tangled up.

```go
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
```

O(n) time, O(h) space for the recursion stack.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/binary_tree_tilt/SOLUTION.md
