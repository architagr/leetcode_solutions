# Intuition

## Related Easy Problems

- [Day 3: Binary Tree Paths](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/binary_tree_path/) — tree traversal
- [Day 11: Minimum Absolute Difference in BST](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/minimum_absolute_difference_in_bst/) — understanding BST properties

---

The key insight: traverse the BST in reverse in-order (right → node → left). This visits nodes from largest to smallest. Keep a running sum as we go. When we visit a node, add the running sum to it, then update the sum.

This works because:
- In a BST, the right subtree has all values greater than the current node
- By going right first, we process all greater values before the current node
- We accumulate them in a sum that gets applied as we traverse back through smaller nodes
