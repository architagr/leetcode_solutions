**365 Days of LeetCode Challenge — Day 21/365**

**Minimum Distance Between BST Nodes** (Easy)
🔗 https://leetcode.com/problems/minimum-distance-between-bst-nodes/

The trick: an in-order traversal of a BST always visits values in sorted order, so the
minimum difference between *any* two nodes can only ever show up between two values that
are neighbors once sorted. Sort once (for free, via the traversal), then scan for the
smallest adjacent gap — no need to check every pair.

Full breakdown in today's newsletter article ⬇️
