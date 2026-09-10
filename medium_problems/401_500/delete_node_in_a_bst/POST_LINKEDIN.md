365 Days of LeetCode Challenge — Day 39/365

Delete Node in a BST (Medium)
🔗 https://leetcode.com/problems/delete-node-in-a-bst/

The problem splits itself: find the node, then delete it. Finding it is the walk from Day
26. Deleting it is where it gets interesting, because a node with two children can't just
be unhooked.

Only two values in the whole tree can fill the hole — the in-order predecessor and the
successor. That falls straight out of the property that in-order traversal of a BST is
sorted: the replacement has to sit between everything on the left and everything on the
right.

So the solution never removes an internal node at all. It overwrites the value and deletes
the duplicate below, which is always a strictly easier case.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinarySearchTree #Recursion #Golang
