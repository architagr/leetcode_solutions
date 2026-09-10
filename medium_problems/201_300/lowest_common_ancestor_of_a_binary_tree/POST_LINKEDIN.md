365 Days of LeetCode Challenge — Day 28/365

Lowest Common Ancestor of a Binary Tree (Medium)
🔗 https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

Yesterday's problem with one guarantee removed. Worth doing back to back, because the gap
between the two solutions is exactly what a BST's ordering buys you.

Yesterday, a comparison told you which subtree a target was in, so the walk went down one
path. Here a comparison tells you nothing, so the only way to know whether a target is
below a node is to look.

The shape flips: instead of deciding on the way down, every subtree reports upward how
many of the two targets it contains. The first node to reach two is the answer.

O(n) instead of O(h). That's the price.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #Recursion #Golang
