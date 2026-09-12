365 Days of LeetCode Challenge — Day 20/365

876. Middle of the Linked List (Easy)
https://leetcode.com/problems/middle-of-the-linked-list/

Obvious answer: count the nodes, then walk to position count/2. Two passes, and perfectly fine.

Better answer: run two pointers from the head, one moving a node at a time and one moving two. When the fast one reaches the end, the slow one is exactly halfway. The length never exists as a number anywhere in the program.

Both conditions in that loop are load-bearing, and they guard different cases. One catches even-length lists, where the fast pointer lands exactly on nil. The other catches odd-length lists, where it lands on the last node. Drop either and half of all inputs panic.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #LinkedList #TwoPointers #CodingInterview #Algorithms
