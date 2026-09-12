365 Days of LeetCode Challenge — Day 31/365

Binary Tree Right Side View (Medium)
🔗 https://leetcode.com/problems/binary-tree-right-side-view/

Standing on the right and seeing what isn't hidden has a plainer description: from each
level, you see exactly one node — the rightmost.

That sounds like BFS. This solution is a depth-first walk with two small changes: record
only the first node reached at each depth, and visit the right child before the left one.

The second change is what makes the first one correct. Neither does anything alone, and
swapping those two lines back gives you the left side view.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #DFS #Golang
