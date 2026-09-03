**365 Days of LeetCode Challenge — Day 8/365**

**Binary Tree Preorder Traversal** (Easy)
🔗 https://leetcode.com/problems/binary-tree-preorder-traversal/

Preorder just means root before left before right, and that's the entire
spec, no searching or comparing involved. Recursion mirrors it almost one to
one: visit, then recurse left, then recurse right. The only real gotcha is
Go-specific — `append` can reallocate its backing array, so the output slice
has to come back as a return value instead of getting mutated in place.

Full breakdown in today's newsletter article ⬇️

#DSA #LeetCode #100DaysOfCode #CodingInterview #Algorithms #BinaryTree #Recursion
