Day 38/365 · Kth Largest Element in a Stream (Easy)

You never need the full sorted order — only the kth largest, now.

So keep only the top k. Among those, the kth largest is the smallest. Min-heap capped at k: O(log k) per add.

https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/701_800/kth_largest_element_in_a_stream/SOLUTION.md

#golang #leetcode
