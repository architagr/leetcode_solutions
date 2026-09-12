**365 Days of LeetCode Challenge — Day 69/365**
**Evaluate Boolean Binary Tree** (Easy)
🔗 https://leetcode.com/problems/evaluate-boolean-binary-tree/

The tree itself is the boolean expression: leaves hold `True`/`False` literals, internal nodes hold `OR`/`AND`. Post-order recursion evaluates both children before combining them with the parent's operator, the same way you'd evaluate a nested expression from the inside out.

```go
const (
	FALSE = 0
	TRUE  = 1
	OR    = 2
	AND   = 3
)

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == TRUE
	}

	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)
	if root.Val == OR {
		return left || right
	}
	return left && right
}
```

O(n) time, O(h) space (tree height).

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/2301_2400/evaluate_boolean_binary_tree/SOLUTION.md
