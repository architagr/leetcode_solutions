## 365 Days of LeetCode Challenge — Day 13/365

### Diameter of Binary Tree

[LeetCode 543 — Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/) · Difficulty: **Easy**

The longest path in a binary tree doesn't have to pass through the root, it
can be tucked away inside any subtree, which is the part that makes this one
fun. The trick: at every node, the longest path through that node is just
height(left) + height(right). One post-order DFS computes heights and tracks
that running max at the same time, no second pass needed.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #TechCareer #BinaryTree #Recursion #Golang
