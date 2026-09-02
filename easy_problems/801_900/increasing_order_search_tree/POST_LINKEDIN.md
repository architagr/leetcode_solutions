**365 Days of LeetCode Challenge — Day 23/365**

**Increasing Order Search Tree** (Easy)
🔗 https://leetcode.com/problems/increasing-order-search-tree/

The shape the problem wants — no left child, one right child everywhere — is just a
sorted linked list wearing `TreeNode`s. And "sorted node order" is exactly what an
in-order traversal of a BST already gives you for free, so the whole problem collapses
into: collect nodes in-order, then chain them with `.Right`.

Full breakdown in today's newsletter article ⬇️
