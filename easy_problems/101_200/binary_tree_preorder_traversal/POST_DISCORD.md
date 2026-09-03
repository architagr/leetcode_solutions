**365 Days of LeetCode Challenge — Day 8/365**
**Binary Tree Preorder Traversal** (Easy)
🔗 https://leetcode.com/problems/binary-tree-preorder-traversal/

Preorder is root first, then left subtree, then right subtree, full stop. Recursion mirrors that directly: visit, recurse left, recurse right. The part that's actually fiddly is Go-specific. `append` can reallocate its backing array, so the output slice gets threaded through as both an argument and a return value rather than mutated in place and trusted to update.

```go
func PreorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	return traversal(root, arr)
}

func traversal(A *TreeNode, arr []int) []int {
	if A == nil {
		return arr
	}
	arr = append(arr, A.Val)
	arr = traversal(A.Left, arr)

	arr = traversal(A.Right, arr)
	return arr
}
```

O(n) time, O(h) space for the recursion stack plus O(n) for the output slice.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/SOLUTION.md
