# 365 Days of LeetCode Challenge — Day 18/365

## Second Minimum Node In a Binary Tree (Easy)

🔗 https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/

This tree has a weird rule: every node's value is the smaller of its two children's
values. That single invariant means the root is always the global minimum — no search
needed — and it also means a whole subtree can be pruned the instant you find a value
above the minimum, since nothing under it can ever be smaller. One DFS pass, no
collect-and-sort required.

Full breakdown in today's newsletter article ⬇
