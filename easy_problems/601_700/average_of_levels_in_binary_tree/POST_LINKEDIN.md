**365 Days of LeetCode Challenge — Day 16/365**

**Average of Levels in Binary Tree** (Easy)
🔗 https://leetcode.com/problems/average-of-levels-in-binary-tree/

Most people reach for BFS here — process one level at a time with a queue. This
solution does it with plain DFS instead: carry the depth as a parameter and use it as
an index into running sum/count arrays, so nodes from different branches still land in
the same slot as long as they share a level.

Full breakdown in today's newsletter article ⬇️
