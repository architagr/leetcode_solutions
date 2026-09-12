365 Days of LeetCode Challenge — Day 57/365

Closest Binary Search Tree Value II (Hard)
🔗 https://leetcode.com/problems/closest-binary-search-tree-value-ii/

Second hard of the challenge, and like the first, the traversal isn't the hard part. It's
one claim you have to notice and trust:

In a sorted array, the k values closest to a target are always contiguous.

That turns "pick k things" into "pick a window" — and once it's a window, the algorithm is
just: flatten the BST in-order, find the single closest value, and expand outward taking
whichever neighbour is nearer.

Greedy works there because distance from the target grows monotonically as you move away
in either direction. It's the merge step of a merge sort, run outward from a centre.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinarySearchTree #TwoPointers #Golang
