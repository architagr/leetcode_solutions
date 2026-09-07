**365 Days of LeetCode Challenge — Day 35/365**

**Increasing Order Search Tree** (Easy)
🔗 https://leetcode.com/problems/increasing-order-search-tree/

The shape the problem wants (no left child, one right child everywhere) is just a
sorted linked list wearing `TreeNode`s. An in-order traversal of a BST already gives
you sorted node order for free, so the whole thing collapses into two steps: collect
nodes in-order, then chain them with `.Right`.

Full breakdown in today's newsletter article below.

#DSA #LeetCode #100DaysOfCode #BinarySearchTree #InOrderTraversal #CodingInterview
