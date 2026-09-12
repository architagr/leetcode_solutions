365 Days of LeetCode Challenge — Day 37/365

200. Number of Islands (Medium)
https://leetcode.com/problems/number-of-islands/

Nothing in this problem says "graph", and that is the point.

A cell is a node. Its neighbours are the cells above, below, left and right. There is no adjacency list to build, because in a grid the edges are implied by position - a neighbour is arithmetic, not a lookup. Yesterday half the work was converting an edge list into something usable. Here that half is free.

Then it is component counting. Scan for land, and every time you find a cell no earlier traversal has consumed, add one and sink the whole island. The count is not counting islands directly; it counts how many times the scan had to start.

One detail worth stealing: sink the cell BEFORE recursing. Two adjacent land cells each list the other as a neighbour, so without it the recursion bounces between them forever. That is not an optimisation, it is termination.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Graphs #DFS #CodingInterview #Algorithms
