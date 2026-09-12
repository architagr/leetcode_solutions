Day 55/365 · Binary Search Tree Iterator (Medium)

An iterator hides WHEN the work happens. That's the only real decision here.

This flattens in-order up front, so Next() is a slice read: O(1) calls, O(n) held. The follow-up wants a stack, O(h).

https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_search_tree_iterator/SOLUTION.md

#golang
