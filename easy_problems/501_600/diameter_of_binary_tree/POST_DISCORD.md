**365 Days of LeetCode Challenge — Day 21/365**
**Diameter of Binary Tree** (Easy)
🔗 https://leetcode.com/problems/diameter-of-binary-tree/

The longest path between two nodes doesn't have to pass through the root, it can sit entirely inside a subtree. The trick: for any node, the longest path through it is `height(left) + height(right)`. So one post-order DFS computes heights normally and updates a running max diameter at every node on the way back up. One pass, no extra work.

```go
var dia = 0

func diameterOfBinaryTree(root *TreeNode) int {
	dia = 0
	calc(root)
	return dia
}

func calc(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := calc(root.Left)
	right := calc(root.Right)
	dia = maxVal(left+right, dia)
	return maxVal(left, right) + 1
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

`calc`'s top-level return value gets discarded because only the side effect matters. `dia` is package-level, so it's reset at the start of every call to keep a previous run's result from leaking in. O(n) time, O(h) space.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/diameter_of_binary_tree/SOLUTION.md
