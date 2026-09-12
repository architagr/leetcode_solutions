Day 77/365 · Kth Largest Element in an Array (Medium)

Push all n, pop k: that is O(n log n), same as sorting. A heap performing a sort, interrupted early.

A min-heap capped at k is O(n log k). Min, because you evict the smallest you kept.

https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/kth_largest_element_in_an_array/SOLUTION.md

#golang
