**365 Days of LeetCode Challenge — Day 17/365**
**Univalued Binary Tree** (Easy)
🔗 https://leetcode.com/problems/univalued-binary-tree/

A tree is uni-valued exactly when every parent-child edge connects two equal values. No node ever needs the root's value. Each one just compares its own children to itself as the recursion walks down, and an empty subtree is trivially uni-valued so `nil` returns `true`. A mismatch found deep down rides upward through each ancestor's check with no re-work.

```go
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left, right := true, true
	if root.Left != nil {
		left = isUnivalTree(root.Left) && root.Left.Val == root.Val
	}
	if root.Right != nil {
		right = isUnivalTree(root.Right) && root.Right.Val == root.Val
	}
	return left && right
}
```

O(n) time, O(h) space (tree height).

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/901_1000/univalued_binary_tree/SOLUTION.md
