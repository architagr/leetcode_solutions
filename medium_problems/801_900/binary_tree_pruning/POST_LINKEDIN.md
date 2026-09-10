365 Days of LeetCode Challenge — Day 44/365

Binary Tree Pruning (Medium)
🔗 https://leetcode.com/problems/binary-tree-pruning/

A node survives if its subtree contains a 1 anywhere, which as a recursion is almost the
definition read aloud. Post-order, because a node can't answer that on the way down.

The part people get wrong isn't the condition — it's who does the deleting. A node can't
remove itself: it has no reference to its parent, and nulling a local variable changes
nothing the caller can see.

So the child reports "nothing worth keeping here", and the parent clears the pointer. Which
leaves the root, because the root has no parent — and that's the only reason the wrapper
function exists.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #Recursion #Golang
