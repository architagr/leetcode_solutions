365 Days of LeetCode Challenge — Day 17/365

Construct String from Binary Tree (Medium)
🔗 https://leetcode.com/problems/construct-string-from-binary-tree/

The traversal is plain pre-order. What makes it a medium is one clause in the formatting
rules: empty parentheses are omitted, except when a node has a right child and no left
child, where you have to emit `()` anyway.

That exception is what keeps the string reversible. Without it, `1(2)` describes both a
left child and a right child.

The fix is to make the code asymmetric the way the rules are — and then no branch anywhere
has to test for the special case at all.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #Recursion #Golang
