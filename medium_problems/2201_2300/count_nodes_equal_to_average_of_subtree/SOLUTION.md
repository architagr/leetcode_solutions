# Solution Walkthrough

## Approach: Post-Order Traversal with Sum and Count

To count nodes equal to their subtree's average, compute each subtree's sum and count. A recursive function returns both values, allowing us to check each node.

## Diagram

![Count Nodes Equal to Average](images/walkthrough-2265.png)

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

**Post-Order Traversal:** Process left and right subtrees before the current node.

1. Recursively get sum and count from left subtree
2. Recursively get sum and count from right subtree  
3. Compute current subtree's sum: left_sum + node.val + right_sum
4. Compute current subtree's count: left_count + right_count + 1
5. Check if node.val equals sum/count (integer division)
6. Return sum and count for parent to use

## Complexity

- **Time:** O(n) — visit each node once
- **Space:** O(h) — recursion depth
