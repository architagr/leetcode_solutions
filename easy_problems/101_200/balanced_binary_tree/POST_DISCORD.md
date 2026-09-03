**365 Days of LeetCode Challenge — Day 5/365**
**Balanced Binary Tree** (Easy)
🔗 https://leetcode.com/problems/balanced-binary-tree/

Checking height separately at every node is the trap here. It recomputes the same subtree heights over and over, which costs O(n²) on a skewed tree. The fix: compute height and balance in one bottom-up post-order pass, and short-circuit the moment any subtree turns out unbalanced.

```go
func isBalanced(root *TreeNode) bool {
	_, ok := validateTree(root)
	return ok
}

func validateTree(node *TreeNode) (height int, ok bool) {
	height = 0
	ok = true
	if node == nil {
		return
	}
	rightHeight := 0
	rightHeight, ok = validateTree(node.Right)
	if !ok {
		return
	}
	leftHeight := 0
	leftHeight, ok = validateTree(node.Left)
	if !ok {
		return
	}
	diff := leftHeight - rightHeight
	if diff > 1 || diff < -1 {
		ok = false
		return
	}
	height = maxValue(leftHeight, rightHeight) + 1
	return
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/balanced_binary_tree/SOLUTION.md
