**365 Days of LeetCode Challenge — Day 29/365**
**Binary Tree Level Order Traversal** (Medium)
🔗 https://leetcode.com/problems/binary-tree-level-order-traversal/

"Level order" is the textbook use for BFS. This solution doesn't use one — it's a plain DFS that produces level-ordered output anyway, for the same reason Day 27 worked.

A DFS arrives at nodes branch by branch, in an order with nothing to do with levels. That doesn't matter, because the grouping isn't decided by arrival. Each call carries its own depth, and every node appends into `arr[level]` — the slot for its own depth. Descending into Left before Right is what makes each level read left to right.

```go
func traversal(head *TreeNode, arr [][]int, level int) [][]int {
	if head == nil {
		return arr
	}
	if (len(arr) - 1) < level {
		a := make([]int, 0)
		arr = append(arr, a)
	}
	arr[level] = append(arr[level], head.Val)
	arr = traversal(head.Left, arr, level+1)
	arr = traversal(head.Right, arr, level+1)
	return arr
}
```

O(n) time. Peak memory is the tree's height rather than the widest level, which is the real trade against BFS.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_level_order_traversal/SOLUTION.md
