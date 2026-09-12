365 Days of LeetCode Challenge — Day 35/365

Populating Next Right Pointers in Each Node (Medium)
🔗 https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

The obvious version walks each level left to right, remembers the previous node, and sets
prev.Next = current. That works.

This one pushes children right before left, so the BFS walks every level backwards — and
once it does, the node to your right is simply the node you visited just before. The
assignment becomes current.Next = prev, which reads exactly like the requirement.

The boundary falls out too. Popping the level sentinel resets prev to nil, so the rightmost
node of the next level gets nil without a branch testing for it.

One inversion, and two special cases stop existing.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #BFS #Golang
