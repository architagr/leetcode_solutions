365 Days of LeetCode Challenge — Day 54/365

Binary Search Tree to Greater Sum Tree (Medium)
https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/

"Every key becomes itself plus all greater keys" sounds like a search problem. It isn't. Walk the BST right-to-left instead of left-to-right and you visit keys in descending order, which means every greater key has already gone past you. Carry a running total and add it in.

The part I liked: after `node.Val += right`, the node's value *is* the running sum, so there's no accumulator variable to maintain at all.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #BinarySearchTree #Recursion #Golang #CodingInterview #Algorithms
