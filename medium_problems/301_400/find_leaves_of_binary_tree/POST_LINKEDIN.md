365 Days of LeetCode Challenge — Day 24/365

Find Leaves of Binary Tree (Medium)
🔗 https://leetcode.com/problems/find-leaves-of-binary-tree/

The problem describes a process — strip the leaves, strip the new leaves, repeat — and
simulating it literally costs a full traversal per round.

You don't have to, because every node's round is already determined before anything is
removed. It just isn't the node's depth.

Every per-level problem in this batch so far grouped by depth from the root. This one
groups by the opposite measurement: height, the distance down to the deepest leaf beneath
a node. One post-order pass computes it.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #Recursion #Golang
