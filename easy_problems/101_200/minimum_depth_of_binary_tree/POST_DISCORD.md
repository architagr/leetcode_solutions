**365 Days of LeetCode Challenge — Day 7/365**
**Minimum Depth of Binary Tree** (Easy)
🔗 https://leetcode.com/problems/minimum-depth-of-binary-tree/

This looks like Maximum Depth with `min` swapped in, but a leaf needs no children, not just a missing one. A node with only one child still isn't a leaf, so a nil child's depth (`0`) must never win a naive `min()` comparison, or the code reports the tree bottoming out one step too early. Easy trap if you're pattern-matching off memory.

```go
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left != 0 && right != 0 {
		return minVal(left, right) + 1
	} else if left != 0 {
		return left + 1
	}
	return right + 1
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/minimum_depth_of_binary_tree/SOLUTION.md
