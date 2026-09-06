# 365 Days of LeetCode Challenge — Day 32/365

## Reverse Odd Levels of Binary Tree

**LeetCode:** [#2415 - Reverse Odd Levels of Binary Tree](https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/)  
**Solution:** [github.com/architagr/leetcode_solutions/blob/main/medium_problems/2401_2500/reverse_odd_levels_of_binary_tree/](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/2401_2500/reverse_odd_levels_of_binary_tree/)

---

## Related Easy Problems

- [Day 8: Binary Tree Preorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/) — tree traversal patterns
- [Day 17: Two Sum IV - Input is a BST](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/two_sum_iv_input_is_a_bst/) — comparing nodes in a tree

---

## Intuition

In a perfect binary tree, nodes at mirrored positions (left of left subtree vs right of right subtree) are at the same level. Compare nodes at mirrored positions recursively. When the level number is odd, swap their values. Then recurse into the children, but swap the child pointers to maintain the mirror relationship.

Key: don't rearrange the tree structure, just swap node values.

---

## Solution Walkthrough

### Approach: Level-Wise Recursion with Mirroring

Traverse the tree level-by-level, comparing nodes at mirrored positions. For odd levels, swap their values.

### Code

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
    // Recurse into children, swapping the pointers to maintain mirror
    rev(l.Left, r.Right, d+1)
    rev(l.Right, r.Left, d+1)
}
```

**Algorithm:**
1. Start with the left and right children of root (level 1)
2. If the level is odd, swap their values
3. Recurse: pass left.Left and right.Right (continuing the mirror)
4. Recurse: pass left.Right and right.Left (swapped to maintain mirror relationship)

### Complexity

- **Time:** O(n) — visit each node once
- **Space:** O(h) — recursion depth

---

#BinaryTree #PerfectBinaryTree #Recursion #LeetCode #DSA #Golang

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
