# Solution Walkthrough

The implementation compares nodes at mirrored positions recursively, swapping their values at odd levels only.

## Approach: Level-Wise Recursion with Mirroring

Traverse the tree level-by-level, comparing nodes at mirrored positions. For odd levels, swap their values.

## Diagram

![Reverse Odd Levels](images/walkthrough-2415.png)

## Code

```go
func reverseOddLevels(root *TreeNode) *TreeNode {
    rev(root.Left, root.Right, 1)
    return root
}

func rev(l, r *TreeNode, d int) {
    if l == nil {
        return
    }
    // At odd levels, swap the mirrored nodes' values
    if d%2 == 1 {
        l.Val, r.Val = r.Val, l.Val
    }
    // Recurse into children, swapping pointers to maintain mirror
    rev(l.Left, r.Right, d+1)
    rev(l.Right, r.Left, d+1)
}
```

**Key insight:** Perfect binary tree nodes at mirrored positions are at the same level. Swap values at odd levels only, maintain mirror relationship by swapping child pointers.

**Algorithm:**
1. Start with left and right children of root (level 1)
2. If level is odd, swap their values
3. Recurse: pass left.Left and right.Right (continuing mirror)
4. Recurse: pass left.Right and right.Left (swapped to maintain mirror)
5. Base case: stop when l == nil

## Complexity

- **Time:** O(n) — visit each node once
- **Space:** O(h) — recursion stack depth (where h is height)

Note: Perfect binary tree has height log(n), so space is O(log n).
