Day 81/365 · Ugly Number II (Medium)

The heap solution needs a set, because 6 arrives as 2x3 and again as 3x2.

Three pointers dedupe for free: when both produce 6, both advance. Three ifs, not an else-if. O(n), no heap.

https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/ugly_number_ii/SOLUTION.md

#golang
