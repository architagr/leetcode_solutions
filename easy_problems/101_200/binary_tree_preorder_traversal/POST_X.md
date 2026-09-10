Day 8/365 · Binary Tree Preorder Traversal (Easy)

Root, left, right — the recursion writes itself.

The Go part doesn't. append can reallocate, so the slice threads through as argument and return. Mutate in place and the caller keeps a stale header.

https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/SOLUTION.md
