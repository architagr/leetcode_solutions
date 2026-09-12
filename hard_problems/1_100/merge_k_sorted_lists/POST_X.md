Day 80/365 · Merge k Sorted Lists (Hard)

This is day 22's merge with the if replaced by a heap.

The heap holds one node per list, never every node - a list's second element cannot be next while its first is unplaced. So it is O(N log k).

https://github.com/architagr/leetcode_solutions/blob/main/hard_problems/1_100/merge_k_sorted_lists/SOLUTION.md

#golang
