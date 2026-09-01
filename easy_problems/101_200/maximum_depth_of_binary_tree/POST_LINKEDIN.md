**365 Days of LeetCode Challenge — Day 1/365**

**Maximum Depth of Binary Tree** (Easy)
🔗 https://leetcode.com/problems/maximum-depth-of-binary-tree/

Finding a tree's max depth with a single BFS queue — no per-level size tracking needed.
The trick: drop a `nil` sentinel into the queue right after the root to mark "end of
this level." Pop the sentinel → level's done, bump the counter, and if there's more
tree left, push the next sentinel.

Full breakdown in today's newsletter article ⬇️
