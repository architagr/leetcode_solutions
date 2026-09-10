Day 8/365 · Binary Tree Preorder Traversal (Easy)

Root, left, right — recursion writes itself.

The Go part doesn't. append can reallocate, so the slice threads through as arg and return. Mutate in place and the caller goes stale.

https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/SOLUTION.md

#golang #leetcode
