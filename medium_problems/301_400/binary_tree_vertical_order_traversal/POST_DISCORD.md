**365 Days of LeetCode Challenge — Day 47/365**
**Binary Tree Vertical Order Traversal** (Medium)
🔗 https://leetcode.com/problems/binary-tree-vertical-order-traversal/

Give every node a column number: root is 0, left child is parent-1, right child is parent+1. Group by it. That's a coordinate carried down the traversal, same as Day 15's depth — it just goes negative too.

The interesting part is that here the traversal **has** to be BFS, and Day 15's did not.

Day 15 grouped by depth, and left-to-right within a level came from recursing Left before Right. Arrival order across branches never mattered.

This wants each column top to bottom, and left to right among ties. Those are exactly what BFS produces. Append as you go and the lists are already correct, no sorting.

A DFS breaks it: it drives one branch to the bottom first, so a deep node from the left subtree gets appended to a column ahead of a shallower node from the right that belongs above it. Same grouping, wrong order inside it.

```go
if n.Left != nil {
	push(n.Left, n.order-1)
}
if n.Right != nil {
	push(n.Right, n.order+1)
}
```

One shortcut worth naming: the assembly loop runs `i` from -101 to 101, leaning on the at-most-100-nodes constraint. Sorting the map's keys wouldn't depend on that.

O(n) time and space.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/301_400/binary_tree_vertical_order_traversal/SOLUTION.md
