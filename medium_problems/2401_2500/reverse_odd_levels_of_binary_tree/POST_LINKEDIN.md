365 Days of LeetCode Challenge — Day 72/365

Reverse Odd Levels of Binary Tree (Medium)
https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/

BFS with a slice per level works. But reversing a level is the same thing as swapping every mirror pair on it, and mirror pairs are something recursion hands you directly. Descend two nodes at a time and swap when the depth is odd. No buffers, O(log n) stack.

The catch is that the two recursive calls have to cross: left.Left pairs with right.Right, left.Right with right.Left. Pair them the tidy-looking way and you permute the level instead of reversing it.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #BinaryTree #Recursion #DFS #Golang #Algorithms
