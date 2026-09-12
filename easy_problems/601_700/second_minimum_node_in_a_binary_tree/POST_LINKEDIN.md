# 365 Days of LeetCode Challenge — Day 70/365

## Second Minimum Node In a Binary Tree (Easy)

🔗 https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/

This tree comes with a strange rule attached: every node's value is the smaller of its
two children's values. That one invariant does a lot of work. It means the root is
always the global minimum, no searching required. And it means a whole subtree can be
pruned the moment you hit a value above the minimum, since nothing underneath it can
ever be smaller. One DFS pass, no collect-and-sort needed.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #BinaryTree #DFS #Golang
