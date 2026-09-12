365 Days of LeetCode Challenge — Day 52/365

Validate Binary Search Tree (Medium)
🔗 https://leetcode.com/problems/validate-binary-search-tree/

The definition sounds local: left subtree smaller, right subtree larger, both subtrees also
valid. Implement it literally and you get the classic wrong answer — checking each node
against its two immediate children, which accepts trees where a deep node violates a
distant ancestor.

A BST constrains a node against every ancestor, not just its parent.

The way out is a property from earlier in this batch. In-order traversal of a BST yields
ascending values, and that equivalence runs both ways. So there are no bounds to thread
down at all: traverse, then check the sequence is strictly increasing.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinarySearchTree #Recursion #Golang
