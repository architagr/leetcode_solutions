365 Days of LeetCode Challenge — Day 47/365

Binary Tree Vertical Order Traversal (Medium)
🔗 https://leetcode.com/problems/binary-tree-vertical-order-traversal/

Give every node a column number — root 0, left child one less, right child one more — and
group by it. That's a coordinate carried down the traversal, same as Day 15's depth, except
this one goes negative.

The interesting part is that this one has to be BFS, and Day 15's did not.

Day 15 grouped by depth, and the required left-to-right order within a level came from
recursing Left before Right. Arrival order across branches never mattered. This problem
wants each column top to bottom — so it does matter, and only a breadth-first walk gives it
for free.

Two problems, same grouping idea, opposite answers on the traversal.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #BFS #Golang
