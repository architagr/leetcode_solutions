365 Days of LeetCode Challenge — Day 22/365

21. Merge Two Sorted Lists (Easy)
https://leetcode.com/problems/merge-two-sorted-lists/

The idea is the easy part. The smallest node not yet placed is always the head of one list or the head of the other, because both are sorted and nothing behind a head is smaller than it. Compare heads, take the smaller, repeat.

The annoying part is the first node. Until something has been placed there is no previous node to attach to, so the straightforward version writes the pick-the-smaller logic twice: once to choose the head, once inside the loop.

A dummy node deletes that. Point tail at a throwaway node and it refers to something real from the first iteration, so the loop never asks whether this is the first element. Return dummy.Next and discard it.

That trick is the thing worth keeping here, and it is not specific to merging. Any time you build a linked list front to back, a dummy head removes the special case.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #LinkedList #MergeSort #CodingInterview #Algorithms
