# Solution Walkthrough

Compute each subtree's sum and count via post-order traversal. Check if node equals average of its subtree.

## Approach: Post-Order Traversal with Sum and Count

Process children first (post-order), so by the time we reach a node, we know its subtree's total sum and node count.

![Example tree from LeetCode](images/example.jpg)

## Code

```go
func averageOfSubtree(root *TreeNode) int {
    var count = 0
    var sumAndCountOfNodes func(node *TreeNode) (sum, nodes int)
    sumAndCountOfNodes = func(node *TreeNode) (sum, nodes int) {
        if node == nil {
            return 0, 0
        }
        
        leftSum, leftCount := sumAndCountOfNodes(node.Left)
        rightSum, rightCount := sumAndCountOfNodes(node.Right)
        
        currentSum := leftSum + node.Val + rightSum
        currentCount := leftCount + rightCount + 1
        
        // Check if current node equals average of its subtree
        if currentSum/currentCount == node.Val {
            count++
        }
        
        return currentSum, currentCount
    }
    sumAndCountOfNodes(root)
    return count
}
```

## Execution Trace

**Key insight:** Process nodes bottom-up (post-order). When at a node, we already know child subtree sums and counts. Add node value, compute average, check equality.

### Tree Structure

```
       4
      / \
     8   5
    / \   \
   0   1   6
```

### Step 1: Process left leaf node (0)

- Node 0 (leaf): sum = 0, count = 1
- Average = 0 ÷ 1 = 0
- Node value = 0 → **Match! count = 1**

### Step 2: Process second left leaf node (1)

- Node 1 (leaf): sum = 1, count = 1
- Average = 1 ÷ 1 = 1
- Node value = 1 → **Match! count = 2**

### Step 3: Process left subtree root (8)

- Combine children: sum = 0 + 1 = 1 (from children)
- Add current node: sum = 1 + 8 = 9, count = 2 + 1 = 3
- Average = 9 ÷ 3 = 3
- Node value = 8 → No match

### Step 4: Process right leaf node (6)

- Node 6 (leaf): sum = 6, count = 1
- Average = 6 ÷ 1 = 6
- Node value = 6 → **Match! count = 3**

### Step 5: Process right subtree root (5)

- Combine children: sum = 6 (from child)
- Add current node: sum = 6 + 5 = 11, count = 1 + 1 = 2
- Average = 11 ÷ 2 = 5 (integer division)
- Node value = 5 → **Match! count = 4**

### Step 6: Process tree root (4)

- Combine subtrees: sum = 9 + 11 = 20 (from both subtrees), count = 3 + 2 = 5
- Add current node: sum = 20 + 4 = 24, count = 5 + 1 = 6
- Average = 24 ÷ 6 = 4
- Node value = 4 → **Match! count = 5**

### Result

Nodes matching their subtree average: 0, 1, 6, 5, 4. **Total = 5**

## Algorithm Steps

1. Recursively compute left subtree: sum and count
2. Recursively compute right subtree: sum and count
3. Current subtree sum = left_sum + node.val + right_sum
4. Current subtree count = left_count + right_count + 1
5. If node.val == (current_sum / current_count), increment count
6. Return sum and count to parent

## Complexity

- **Time:** O(n) — visit each node exactly once
- **Space:** O(h) — recursion stack depth (where h is tree height)
  - O(log n) for balanced trees
  - O(n) for skewed trees
