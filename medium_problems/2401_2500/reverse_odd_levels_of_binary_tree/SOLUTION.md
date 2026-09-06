# Solution Walkthrough

Traverse level-by-level comparing mirrored nodes. Swap values at odd levels only.

## Approach: Level-Wise Recursion with Mirroring

Perfect binary tree nodes at mirrored positions are at same level. Process recursively: for odd levels, swap mirrored node values.

## Code

```go
func reverseOddLevels(root *TreeNode) *TreeNode {
    rev(root.Left, root.Right, 1)
    return root
}

func rev(l, r *TreeNode, depth int) {
    if l == nil {
        return
    }
    // At odd depths, swap the mirrored nodes' values
    if depth%2 == 1 {
        l.Val, r.Val = r.Val, l.Val
    }
    // Recurse into children, swapping pointers to maintain mirror
    rev(l.Left, r.Right, depth+1)
    rev(l.Right, r.Left, depth+1)
}
```

## Execution Trace

**Key insight:** Perfect binary tree nodes at mirrored positions maintain their mirror relationship through recursion. At odd depths, swap values of mirrored pairs.

### Tree Structure (Perfect Binary Tree)

```
         1          (depth 0 - even, root)
       /   \
      2     3      (depth 1 - odd, swap: 2 <-> 3)
     / \   / \
    4   5 6   7    (depth 2 - even, no swap)
```

### Processing Steps

### Step 1: Start recursion with left=2, right=3 (depth 1 - odd)

- Current depth = 1 (odd) → **Swap values: 2 ↔ 3**
- After swap: tree has 3 on left, 2 on right
- Recurse into next level

### Step 2: Process depth 2 (even - no swap)

- Left pair: recurse with 4 (from left.Left of 3) and 7 (from right.Right of 2)
  - Depth 2 is even → no swap
  - Recurse deeper
- Right pair: recurse with 5 (from left.Right of 3) and 6 (from right.Left of 2)
  - Depth 2 is even → no swap
  - Recurse deeper

### Step 3: Process depth 3 (odd - swap)

Nodes at depth 3 form leaf pairs:

- Nodes from left subtree at even positions pair with nodes from right subtree at even positions
- All depth 3 nodes are leaves, so swapping leaf values completes the level reversal

### Step 4: Maintain Mirror Through Recursion

Key mechanism:
- `rev(l.Left, r.Right, d+1)`: compares left's left child with right's right child (continues mirror)
- `rev(l.Right, r.Left, d+1)`: compares left's right child with right's left child (crosses pointers to maintain mirror)

This ensures every node at each level gets paired with its mirror.

### Result

- Odd-level nodes have reversed values (2↔3 at level 1, leaf nodes at level 3 etc.)
- Even-level nodes keep original values (4,5,6,7 at level 2 unchanged)

## Algorithm Steps

1. Start with left and right children of root (depth 1)
2. If depth is odd, swap their values
3. Recurse on (left.Left, right.Right) to continue mirror
4. Recurse on (left.Right, right.Left) with swapped pointers to maintain mirror pairs
5. Base case: stop when l == nil (tree ends)

## Complexity

- **Time:** O(n) — visit each node once (all nodes in perfect binary tree)
- **Space:** O(h) — recursion stack depth (where h is height)
  - Perfect binary tree height = log(n), so O(log n)
