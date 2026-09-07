**365 Days of LeetCode Challenge — Day 6/365**
**Count Complete Tree Nodes** (medium)
🔗 https://leetcode.com/problems/count-complete-tree-nodes/

A complete binary tree is nearly full: all levels complete except possibly the last, which fills left-to-right. This solution recursively counts nodes in both subtrees. The O(log^2 n) optimization uses the tree's structure—if left and right heights match, the left subtree is a perfect binary tree and we can compute its count directly without traversing.

```go
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    rightCount := countNodes(root.Right)
    leftCount := countNodes(root.Left)
    return rightCount + leftCount + 1
}
```

Full walkthrough with diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/count_complete_tree_nodes/SOLUTION.md
