Day 24/365 · Find Leaves of Binary Tree (Medium)

Don't simulate the stripping — that's O(n·h).

A node's round is its HEIGHT, not its depth: max(left, right) + 1, bottom-up. nil returns -1, so leaves land at 0.

https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/301_400/find_leaves_of_binary_tree/SOLUTION.md

#golang #leetcode
