**365 Days of LeetCode Challenge — Day 6/365**

**Minimum Depth of Binary Tree** (Easy)
🔗 https://leetcode.com/problems/minimum-depth-of-binary-tree/

This one looks like a copy-paste of Maximum Depth with `min` swapped in, until you
remember a leaf needs no children, not just a missing one. A node with only one
child still counts toward depth, so a naive `min()` at every node quietly breaks on
lopsided trees. Took me a second look to catch it.

Full breakdown in today's newsletter article below.

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #Recursion #Golang
