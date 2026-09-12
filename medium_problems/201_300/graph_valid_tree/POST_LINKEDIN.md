365 Days of LeetCode Challenge — Day 43/365

261. Graph Valid Tree (Medium)
https://leetcode.com/problems/graph-valid-tree/

A graph is a tree when it is connected and acyclic. Both are required - a forest is acyclic but disconnected, a triangle is connected but cyclic.

The shortcut: a tree on n nodes has exactly n-1 edges. So once you know the graph is connected, the edge count settles the cycle question. More than n-1 edges and connected means a cycle must exist, because n nodes cannot be joined by more than n-1 edges without one.

This solution never searches for a cycle. It counts.

For connectivity it uses union-find, which is the first structure in this arc that does not traverse anything. Every problem so far built an adjacency list and walked it. Union-find consumes the edge list directly and maintains one number, the component count, which is exactly the answer.

Worth noticing what it leaves on the table: when an edge's endpoints already share a root, that edge closes a cycle. This solution does not need that, but if a problem asks WHICH edge creates the cycle, it is sitting right there.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #UnionFind #Graphs #CodingInterview #Algorithms
