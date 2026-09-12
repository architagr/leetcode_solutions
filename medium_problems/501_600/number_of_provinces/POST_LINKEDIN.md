365 Days of LeetCode Challenge — Day 39/365

547. Number of Provinces (Medium)
https://leetcode.com/problems/number-of-provinces/

The counting is day 37's, unchanged: scan every node, and each time you find one no earlier traversal reached, add one and consume everything it can reach.

What changes is the input, and this is the third graph representation in four days. Day 36 gave an edge list. Day 37 gave a grid where the edges were implied by position. Today gives a matrix. All three become the same adjacency list before anything interesting happens.

The trap: this matrix is not a grid. In day 37, cell (i,j) was a place. Here cell (i,j) is a relationship between city i and city j - an n x n matrix describing n nodes, not n squared. So you traverse the cities and consult the matrix, not the other way round.

Read only above the diagonal. The matrix is symmetric, so every relationship is stored twice, and i+1 also skips the useless diagonal where a city is connected to itself.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Graphs #BFS #CodingInterview #Algorithms
