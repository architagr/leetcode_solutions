365 Days of LeetCode Challenge — Day 73/365

Binary Tree Upside Down (Medium)
🔗 https://leetcode.com/problems/binary-tree-upside-down/

Thirteen lines doing two unrelated jobs on one recursion.

The first is finding the new root — the deepest node on the left spine. That happens in the
base case, and it's the only place a value is ever produced. Every frame above simply
returns it untouched, so the return value is a pass-through.

The second is local pointer surgery at each level, on the way back up.

They share a traversal and otherwise have nothing to do with each other, and reading the
function as two things instead of one is most of understanding it.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #Recursion #Golang
