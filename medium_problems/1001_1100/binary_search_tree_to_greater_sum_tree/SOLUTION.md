# Solution Walkthrough

The implementation uses reverse in-order traversal with a running sum to convert each node.

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

## Execution Trace

**Key insight:** Visit nodes right-to-left (largest to smallest). When we visit a node, add the accumulated sum from all greater nodes.

Walking through a BST with values [4, 1, 6, 0, 2, 5, 7]:

**Step 1:** Start at root (4) - traverse right subtree first to process larger values (6, 7)

- Visit node 6
- Visit node 7 (rightmost)
- Node 7: no greater nodes, stays 7. Accumulated sum = 7.
- Back to node 6: add sum of greater nodes (7) → 6 + 7 = 13. Accumulated sum = 7 + 13 = 20.

**Step 2:** Back at node 4 (root)

- Node 4: add sum of all greater nodes (20 from right subtree) → 4 + 20 = 24. Accumulated sum = 20 + 24 = 44.

**Step 3:** Visit left subtree of 4 (nodes 1, 0, 2)

- Node 2: add accumulated sum (44) → 2 + 44 = 46. Accumulated sum = 44 + 46 = 90.
- Node 1: add accumulated sum (90) → 1 + 90 = 91. Accumulated sum = 90 + 91 = 181.
- Node 0: add accumulated sum (181) → 0 + 181 = 181. Accumulated sum = 181 + 181 = 362.

**Result:** Each node now equals the sum of original values greater than or equal to itself.

## Algorithm Steps

1. Recursively traverse to rightmost node (largest value)
2. On return, accumulate current node's value
3. Add that accumulated sum to current node's value
4. Continue to left subtree with updated sum

## Complexity

- **Time:** O(n) — visit each node exactly once
- **Space:** O(h) — recursion stack depth (where h is tree height)
  - O(log n) for balanced BSTs
  - O(n) for skewed trees
