# Intuition

## Related Easy Problems

- [Day 1: Maximum Depth of Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/maximum_depth_of_binary_tree/) — subtree traversal
- [Day 6: Minimum Depth of Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/minimum_depth_of_binary_tree/) — understanding subtree properties

---

To check if a node equals its subtree's average, we need to know the subtree's sum and node count. Use post-order traversal (visit children before parent) to compute these values bottom-up. As we return from recursion, we have all the information needed to check each node and count matches.

The key: process children first, so by the time we're at a node, we already know its subtree's sum and count.
