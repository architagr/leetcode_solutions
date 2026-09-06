**365 Days of LeetCode Challenge — Day 30/365**
**Binary Search Tree to Greater Sum Tree** (medium)
🔗 https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/

Transform each node by adding the sum of all greater nodes. Use reverse in-order traversal (right→node→left) to visit nodes in descending order, accumulating a running sum as you go.

```go
func parse(node *TreeNode, parentSum int) int {
    if node == nil { return parentSum }
    right := parse(node.Right, parentSum)
    node.Val += right
    return parse(node.Left, node.Val)
}
```

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1001_1100/binary_search_tree_to_greater_sum_tree/SOLUTION.md
