# Solution Walkthrough

## Approach: Reverse In-Order Traversal with Running Sum

To convert each node, add the sum of all greater nodes to it. We achieve this using reverse in-order traversal (right-to-left), maintaining a running cumulative sum.

## Code

```go
func bstToGst(root *TreeNode) *TreeNode {
    parse(root, 0)
    return root
}

func parse(node *TreeNode, parentSum int) int {
    if node == nil {
        return parentSum + 0
    }
    // Visit right subtree first (larger values)
    right := parse(node.Right, parentSum)
    // Add accumulated sum to current node
    node.Val += right
    // Visit left subtree (smaller values)
    return parse(node.Left, node.Val)
}
```

**Traversal order:** Right → Node → Left (reverse in-order)

1. Recursively traverse to the rightmost node
2. On the way back, accumulate the current node's value into a running sum
3. Add that sum to each node we visit
4. Continue to the left subtree with the updated sum

## Complexity

- **Time:** O(n) — visit each node once
- **Space:** O(h) — recursion depth
