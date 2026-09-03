**365 Days of LeetCode Challenge — Day 1/365**

**Maximum Depth of Binary Tree** (Easy)
🔗 https://leetcode.com/problems/maximum-depth-of-binary-tree/

Counting a tree's max depth with one BFS queue, no per-level size tracking. The trick:
drop a `nil` sentinel into the queue right after the root to mark "end of this level."
Pop the sentinel, the level's done, bump the counter, and push the next sentinel if
there's more tree left to walk.

Full breakdown in today's newsletter article.

#DSA #LeetCode #100DaysOfCode #BinaryTree #BFS #Golang #SoftwareEngineering
