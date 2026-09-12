365 Days of LeetCode Challenge — Day 36/365

1971. Find if Path Exists in Graph (Easy)
https://leetcode.com/problems/find-if-path-exists-in-graph/

New topic. Graphs start today.

The part people skip when they say a graph problem "is just BFS": the input is not a graph. You get a flat list of edge pairs. Asking that list for a node's neighbours means scanning all of it, and doing that once per node turns an O(V + E) algorithm into O(V x E).

So the first half of the problem is building an adjacency list. The traversal is the easy half. Choosing the representation is the half that is actually a decision.

One bug worth naming, because it passes a lot of tests before it fails one: an undirected edge has to be recorded under both endpoints. Record it once and you have built a directed graph by accident, and the traversal refuses to walk an edge backwards.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Graphs #BFS #CodingInterview #Algorithms
