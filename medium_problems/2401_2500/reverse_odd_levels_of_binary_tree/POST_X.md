Day 72/365 · Reverse Odd Levels of Binary Tree (Medium)

BFS works, but allocates a slice per level and the widest holds half the tree.

Restate it: reversing a level is swapping mirror pairs. Descend two mirrored nodes, swap in place.

https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/2401_2500/reverse_odd_levels_of_binary_tree/SOLUTION.md

#golang #leetcode
