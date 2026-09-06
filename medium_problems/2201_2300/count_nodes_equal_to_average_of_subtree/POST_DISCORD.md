**365 Days of LeetCode Challenge — Day 31/365**
**Count Nodes Equal to Average of Subtree** (Medium)
🔗 https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/

An average needs a sum and a count, so return both from the recursion. Post-order means every node already has its children's numbers when it runs, so one pass does it instead of re-walking a subtree per node.

Watch the integer division. A node in LeetCode's own example has average 11/2, which counts as 5 and matches, but only because it floors. Use floats and you quietly lose it.

```go
sumAndCountOfNodes = func(node *TreeNode) (sum, cnt int) {
	sum, cnt = 0, 0
	if node == nil {
		return
	}
	ls, lc := sumAndCountOfNodes(node.Left)
	rs, rc := sumAndCountOfNodes(node.Right)
	sum = ls + node.Val + rs
	cnt = lc + rc + 1
	if sum/cnt == node.Val {
		count++
	}
	return
}
```

O(n) time, O(h) stack.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/2201_2300/count_nodes_equal_to_average_of_subtree/SOLUTION.md
