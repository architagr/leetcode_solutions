**365 Days of LeetCode Challenge — Day 16/365**
**Binary Tree Zigzag Level Order Traversal** (Medium)
🔗 https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

Yesterday's level order came out of a DFS with no queue and no idea when a level ended. Zigzag needs that boundary — you can't decide whether to reverse a level until you know it's finished — so this one goes back to BFS.

The boundary comes from Day 1's trick: push a `nil` behind the root. Everything ahead of it is the current level. When it surfaces, the level drained; push a fresh one, but only if the queue still holds nodes, or the loop spins on a sentinel forever.

```go
if node == nil {
	if len(queue) > 0 {
		push(nil)
	}
	x := make([]int, len(arr))
	copy(x, arr)
	if !leftToRight {
		x = reverseArr(x)
	}
	result = append(result, x)
	arr = make([]int, 0)
	leftToRight = !leftToRight
	continue
}
```

That copy isn't defensive — `arr` is reused and `reverseArr` mutates in place, so without it every row of the result would alias one array.

O(n) time, O(w) space for the queue where w is the widest level.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_zigzag_level_order_traversal/SOLUTION.md
