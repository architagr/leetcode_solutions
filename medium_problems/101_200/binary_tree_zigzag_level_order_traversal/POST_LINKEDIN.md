365 Days of LeetCode Challenge — Day 16/365

Binary Tree Zigzag Level Order Traversal (Medium)
🔗 https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

Yesterday's level order came out of a depth-first walk with no queue and no idea when a
level ended. That was fine, because grouping by depth never needed the boundary.

Zigzag does. You can't decide whether a level gets reversed until you know the level is
finished, and "finished" is exactly what a DFS never learns. So this one goes back to BFS,
with a nil sentinel in the queue marking where each level stops.

With the boundary in hand the zigzag itself is four lines and a flag.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #BFS #Golang
