365 Days of LeetCode Challenge — Day 42/365

133. Clone Graph (Medium)
https://leetcode.com/problems/clone-graph/

Copying a binary tree is four lines: make a node, recurse left, recurse right. It works because a tree cannot cycle and every node has one parent.

A graph gives up both guarantees, and each breaks the naive copy differently. Cycles mean A's copy needs B's copy needs A's copy, forever. Shared nodes mean a node reachable two ways gets copied twice, giving you the right values in the wrong shape.

One map from original to copy fixes both. Look it up before cloning; if the copy exists, return it. Cycles terminate because the second arrival returns instead of descending, and sharing survives because every original maps to exactly one copy.

Worth dwelling on: a plain visited SET would only fix the cycle. This has to map to the copies, because the copies are what the reconstruction wires together.

And the order matters. Record the new node in the map BEFORE cloning its neighbours, while it is still half built. Recurse first and a cycle returns to a node that is not in the map yet.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Graphs #DFS #CodingInterview #Algorithms
