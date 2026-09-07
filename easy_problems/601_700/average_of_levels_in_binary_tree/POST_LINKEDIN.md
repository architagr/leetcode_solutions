**365 Days of LeetCode Challenge — Day 14/365**

**Average of Levels in Binary Tree** (Easy)
🔗 https://leetcode.com/problems/average-of-levels-in-binary-tree/

Most people reach for BFS here, processing one level at a time with a queue. I used
plain DFS instead: carry the depth as a parameter and use it as an index into running
sum/count arrays, so nodes from totally different branches still land in the same slot
as long as they're on the same level. Kind of neat that it works without ever grouping
nodes by level explicitly.

Full breakdown in today's newsletter article.

#DSA #LeetCode #100DaysOfCode #BinaryTree #DFS #SoftwareEngineering #TechCareer
