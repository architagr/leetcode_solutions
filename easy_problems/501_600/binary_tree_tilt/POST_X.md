Day 14/365 · Binary Tree Tilt (Easy)

Tilt needs whole-subtree sums, not child values. Re-summing per node is O(n²).

Postorder: each call returns its sum and adds its tilt on the way past. Every sum computed once.

https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/binary_tree_tilt/SOLUTION.md

#golang #leetcode
