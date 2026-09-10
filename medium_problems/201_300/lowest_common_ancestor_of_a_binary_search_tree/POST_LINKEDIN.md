365 Days of LeetCode Challenge — Day 27/365

Lowest Common Ancestor of a Binary Search Tree (Medium)
🔗 https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

The general version of this problem needs a search. The BST version doesn't, for the same
reason Day 26 didn't: a comparison in a BST is a direction, not a verdict.

Both targets smaller than the node? The answer is further left. Both larger? Further
right. Anything else, and this node is the answer.

That last case is doing quiet double duty — it covers the targets sitting on opposite
sides, and one of them being this node. Neither needs detecting separately.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinarySearchTree #Recursion #Golang
