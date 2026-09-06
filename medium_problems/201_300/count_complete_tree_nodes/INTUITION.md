# Intuition

## Related Easy Problems

Before diving in, check out these foundational tree problems:

- [Day 1: Maximum Depth of Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/maximum_depth_of_binary_tree/) — understanding height/depth
- [Day 5: Balanced Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/balanced_binary_tree/) — recognizing tree structure

---

The straightforward approach is to simply traverse every node and count them. While this runs in O(n) time, a complete binary tree has special structure we can exploit to do better.

The key insight: if we traverse down the left spine (always going left), and separately traverse down the right spine (always going right), we can tell whether the left subtree is a "perfect" binary tree (all levels completely filled).

If both spines have the same height, the left subtree is a perfect binary tree of that height, so it has exactly 2^h - 1 nodes. We can skip traversing it and recursively count only the right subtree.

If the left spine is taller, then the right subtree is a perfect binary tree, and we recursively count only the left subtree.

This binary search on height lets us do O(log^2 n) in the best case, since we skip entire subtrees.

The current implementation uses the simple O(n) approach, which is correct and sometimes more readable, especially for an interview when the constraint wasn't immediately clear.
