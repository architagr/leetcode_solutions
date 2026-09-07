**365 Days of LeetCode Challenge — Day 36/365**
**Binary Search Tree to Greater Sum Tree** (Medium)
🔗 https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/

"Every key becomes itself plus all greater keys" reads like a search, but a BST already knows what greater means. Walk it right-to-left and you hit keys in descending order, so every greater key has already gone past. Carry a running total and add it in.

Neat bit: after `node.Val += right`, the node's value *is* the running sum, so there's no accumulator variable at all.

```go
func parse(node *TreeNode, parentSum int) int {
	if node == nil {
		return parentSum + 0
	}
	right := parse(node.Right, parentSum)
	node.Val += right
	return parse(node.Left, node.Val)
}
```

O(n) time, O(h) stack, converted in place.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1001_1100/binary_search_tree_to_greater_sum_tree/SOLUTION.md
