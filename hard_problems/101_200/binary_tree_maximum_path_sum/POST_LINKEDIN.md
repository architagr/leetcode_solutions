365 Days of LeetCode Challenge — Day 18/365

Binary Tree Maximum Path Sum (Hard)
🔗 https://leetcode.com/problems/binary-tree-maximum-path-sum/

First hard of the challenge, and the traversal isn't the hard part.

Each node has to compute two different things. As the top of a path it can use both
children. As a link inside some ancestor's path it can use only one, because a path that
forked would contain a node with three neighbours.

Only one of those can be the return value. The other gets recorded as a side effect — and
returning the wrong one is the bug this problem is really testing for.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #Recursion #Golang
