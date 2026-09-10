365 Days of LeetCode Challenge — Day 10/365

Flatten Binary Tree to Linked List (Medium)
🔗 https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

The shape being asked for is a linked list wearing TreeNode. The clause that makes it a
real problem is the ordering: not sorted, not level by level, specifically pre-order.

Which means the traversal that produces the answer is the one from Day 8. Collect the
values in that order, then rebuild the tree as a right-leaning chain — two halves that
don't interact.

The follow-up wants it in-place with O(1) space, and that version is a genuinely
different problem.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #Recursion #Golang
