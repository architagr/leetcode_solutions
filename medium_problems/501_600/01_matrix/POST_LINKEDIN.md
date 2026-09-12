365 Days of LeetCode Challenge — Day 41/365

542. 01 Matrix (Medium)
https://leetcode.com/problems/01-matrix/

This is yesterday's problem. Rotting Oranges asked how long until an orange rots with every rotten one spreading at once; this asks how far a cell is from the nearest zero with every zero a source. Same question, and multi-source BFS solves it directly.

Today's solution does it without a queue, and that is why it is worth reading.

The definition is circular: a cell's distance is one more than the smallest among its four neighbours, but every neighbour depends on the cell. There is no order to evaluate that in.

Direction breaks the circle. Sweep top-left to bottom-right and, at each cell, everything above and left is already final - so you can compute the best distance among paths that start by going up or left. Sweep the other way for the other two directions. Take the minimum.

The line that makes it work is min with the existing value rather than an assignment. Overwrite instead and you throw away everything the first pass got right.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #DynamicProgramming #Grids #CodingInterview #Algorithms
