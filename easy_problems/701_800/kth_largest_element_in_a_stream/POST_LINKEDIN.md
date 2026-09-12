# 365 Days of LeetCode Challenge — Day 75/365

## Kth Largest Element in a Stream

[LeetCode 703 — Kth Largest Element in a Stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/) · Difficulty: Easy

You don't need to sort the whole stream to know the kth largest value. You only need to
track the top k scores, and the kth largest is always the weakest one in that group. A
min-heap capped at size k gets you that answer in O(log k) per insert instead of
re-sorting on every call. The naive fix here is obvious, the efficient one just needs
you to notice you're tracking a much smaller problem than the one you were handed.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #Heap #DataStructures #Golang #CodingInterview
