# Solution Walkthrough

The implementation uses reverse in-order traversal with a running sum to convert each node.

## Approach: Reverse In-Order Traversal with Running Sum

To convert each node, add the sum of all greater nodes to it. We achieve this using reverse in-order traversal (right-to-left), maintaining a running cumulative sum.

## Diagram

![BST to Greater Sum Tree](images/walkthrough-1038.png)

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

**Key insight:** Visit nodes right-to-left (largest to smallest). When we visit a node, add the accumulated sum from all greater nodes.

**Algorithm:**
1. Recursively traverse to rightmost node
2. On return, accumulate node's value into running sum
3. Add that sum to current node's value
4. Continue to left subtree with updated sum

## Complexity

- **Time:** O(n) — visit each node once
- **Space:** O(h) — recursion stack depth (where h is height)

Note: This is O(log n) space for balanced BSTs, O(n) for skewed trees.
