365 Days of LeetCode Challenge — Day 38/365

695. Max Area of Island (Medium)
https://leetcode.com/problems/max-area-of-island/

This is yesterday's solution with a return value. The diff is that the flood fill returns an int and the scan keeps a maximum instead of a count.

The area reachable from a cell is that cell, plus the area reachable from each of its four neighbours. So the fill returns 1 plus its four recursive calls, and the guard returning 0 handles out of bounds, water and already-sunk land without a single special case.

The sum looks wrong at first. Four neighbours each recursing into their own four neighbours, every one of which includes the cell you came from - surely cells get counted repeatedly?

They do not, because the cell is sunk before the four calls run. Any neighbour that recurses back hits the guard and gets 0.

Yesterday that same line was what stopped the recursion running forever. Today it does that AND makes the arithmetic correct. One assignment, two jobs.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Graphs #DFS #CodingInterview #Algorithms
