**365 Days of LeetCode Challenge — Day 8/365**

**Binary Tree Preorder Traversal** (Easy)
🔗 https://leetcode.com/problems/binary-tree-preorder-traversal/

"Preorder" is the whole spec: visit the node itself, then its left subtree, then
its right subtree. Recursion mirrors that definition directly — the trick in Go
is threading the output slice through as both argument and return value, since
`append` can reallocate its backing array.

Full breakdown in today's newsletter article ⬇️
