# Solution Walkthrough

The implementation uses post-order traversal to compute each subtree's sum and count, then checks if the node equals its subtree's average.

## Approach: Post-Order Traversal with Sum and Count

To count nodes equal to their subtree's average, compute each subtree's sum and count. A recursive function returns both values, allowing us to check each node.

## Code

```go
func averageOfSubtree(root *TreeNode) int {
    var count = 0
    var sumAndCountOfNodes func(node *TreeNode) (currentNodeSum, countNodes int)
    sumAndCountOfNodes = func(node *TreeNode) (currentNodeSum, countNodes int) {
        currentNodeSum, countNodes = 0, 0
        if node == nil {
            return
        }

        leftSubtreeNodeSum, leftSubTreeNodeCount := sumAndCountOfNodes(node.Left)
        rightSybTreeNodeSum, rightSubTreeNodeCount := sumAndCountOfNodes(node.Right)

        currentNodeSum = leftSubtreeNodeSum + node.Val + rightSybTreeNodeSum
        countNodes = leftSubTreeNodeCount + rightSubTreeNodeCount + 1
        if currentNodeSum/countNodes == node.Val {
            count++
        }
        return
    }
    sumAndCountOfNodes(root)
    return count
}
```

**Key insight:** Process children first (post-order), so by the time we reach a node, we already know its subtree's sum and count.

**Algorithm:**
1. Recursively compute left subtree sum and count
2. Recursively compute right subtree sum and count
3. Compute current subtree sum: left_sum + node.val + right_sum
4. Compute current subtree count: left_count + right_count + 1
5. Check if node.val equals sum/count (integer division)
6. Return sum and count for parent to use

## Complexity

- **Time:** O(n) — visit each node once
- **Space:** O(h) — recursion stack depth (where h is height)
