Day 38/365 · Max Area of Island (Medium)

The fill returns 1 + its four recursive calls. Looks like it counts cells twice.

It does not: the cell is sunk before the calls run. The line that stops infinite recursion also fixes the sum.

https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/601_700/max_area_of_island/SOLUTION.md

#golang
