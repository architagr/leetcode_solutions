# 365 Days of LeetCode Challenge — Day 20/365

## Kth Largest Element in a Stream

[LeetCode 703 — Kth Largest Element in a Stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/) · Difficulty: Easy

You don't need to sort the whole stream to know the kth largest value — you only need
to track the top k scores, and the kth largest is always the weakest one in that group.
A min-heap capped at size k gives you that answer in O(log k) per insert instead of
re-sorting on every call.

Full breakdown in today's newsletter article ⬇
