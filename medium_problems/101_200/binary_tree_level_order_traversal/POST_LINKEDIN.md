365 Days of LeetCode Challenge — Day 29/365

Binary Tree Level Order Traversal (Medium)
🔗 https://leetcode.com/problems/binary-tree-level-order-traversal/

"Level order" is the textbook use for BFS. This solution doesn't use one.

It's a depth-first recursion that produces level-ordered output anyway, because the
grouping never depended on arrival order — each call carries its own depth, and every node
appends into the slot for that depth. Whichever branch a node comes from, it lands in the
right place.

Descending into the left child first is what keeps each level reading left to right.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #DFS #Golang
