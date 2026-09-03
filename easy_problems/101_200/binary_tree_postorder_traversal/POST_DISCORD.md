**365 Days of LeetCode Challenge — Day 9/365**
**Binary Tree Postorder Traversal** (Easy)
🔗 https://leetcode.com/problems/binary-tree-postorder-traversal/

Postorder means children before parent: recurse left, recurse right, then append the current node. The recursion guarantees every descendant is already recorded before a node appends its own value, so there's no extra state to track. Nice part is you don't manage any of that yourself, it just comes out right by construction.

```go
func PostorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	return traversal(root, arr)
}

func traversal(A *TreeNode, arr []int) []int {
	if A == nil {
		return arr
	}

	arr = traversal(A.Left, arr)
	arr = traversal(A.Right, arr)
	arr = append(arr, A.Val)
	return arr
}
```

O(n) time, O(h) space for the recursion stack plus O(n) for the output slice.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_postorder_traversal/SOLUTION.md
