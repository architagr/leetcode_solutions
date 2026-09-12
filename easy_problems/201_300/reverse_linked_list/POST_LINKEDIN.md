365 Days of LeetCode Challenge — Day 19/365

206. Reverse Linked List (Easy)
https://leetcode.com/problems/reverse-linked-list/

New topic today. Eighteen days of binary trees, and the last one, flattening a tree into a linked list, was already doing what today's problem is entirely about: rewriting Next pointers to change a shape without moving anything.

Reversing a list moves no data. Every node stays exactly where it is in memory. You walk it once and turn each arrow around.

The catch is that the obvious two-line loop body cannot work. Pointing a node backwards destroys the only reference to the node you were going to visit next, so you have to save it first. Four lines, and there is no other order they can go in.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #LinkedList #CodingInterview #SoftwareEngineering #Algorithms
