Day 25/365 · Binary Tree Maximum Path Sum (Hard)

Each node computes two things, and mixing them is the bug.

As top of a path: left+right+val — record it.
As a link in an ancestor's path: one side only, since paths can't fork — return that.

https://github.com/architagr/leetcode_solutions/blob/main/hard_problems/101_200/binary_tree_maximum_path_sum/SOLUTION.md

#golang
