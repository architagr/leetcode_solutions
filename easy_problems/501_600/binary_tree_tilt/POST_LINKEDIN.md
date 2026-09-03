**365 Days of LeetCode Challenge — Day 13/365**

**Binary Tree Tilt** (Easy)
🔗 https://leetcode.com/problems/binary-tree-tilt/

Each node's tilt depends on the sum of everything under it, not just its two kids.
Summing from scratch at every node redoes the same work over and over, so the trick
is to compute each subtree's sum once, bottom-up, and feed every node's tilt into a
shared accumulator as you go. The neat part is one recursive call quietly does both
jobs at once.

Full breakdown in today's newsletter article, linked below.

#DSA #LeetCode #100DaysOfCode #BinaryTree #Recursion #Golang #CodingInterview
