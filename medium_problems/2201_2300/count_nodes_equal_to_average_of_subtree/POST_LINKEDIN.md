365 Days of LeetCode Challenge — Day 31/365

Count Nodes Equal to Average of Subtree (Medium)
https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/

An average needs a sum and a count, so have the recursion return both. Post-order means every node already has its children's numbers when it runs, and the whole thing collapses to one pass instead of re-walking a subtree per node.

Watch the integer division. One node in LeetCode's own example has average 11/2, which counts as 5 and matches, but only because it floors. Use floats and you quietly lose it.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #BinaryTree #Recursion #DFS #Golang #CodingInterview
