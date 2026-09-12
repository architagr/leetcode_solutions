---
meta_title: "The line that stops the recursion also fixes the arithmetic"
meta_description: "Max Area of Island is yesterday's flood fill with a return value. Sinking the cell before recursing is what stops every cell being counted four times."
---

![Day 38](HERO.png)

## 365 Days of LeetCode Challenge — Day 38/365

**[695. Max Area of Island](https://leetcode.com/problems/max-area-of-island/)** (Medium)

Same grid as yesterday, same definition of an island. Return the size of the largest one.

## This is yesterday's solution with a return value

I mean that literally. The diff against day 37 is that the fill returns an `int` and the scan keeps a maximum instead of a count.

Yesterday the fill was called for its side effect: sink an island so the scan would not count it twice. Today it has to answer a question as well, and the question has an obvious recursive shape.

The area reachable from a cell is that cell, plus the area reachable from each of its four neighbours.

```go
func dfs(i, j int, grid *[][]int) int {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] != 1 {
		return 0
	}
	(*grid)[i][j] = 0

	return 1 + dfs(i+1, j, grid) + dfs(i-1, j, grid) + dfs(i, j+1, grid) + dfs(i, j-1, grid)
}
```

The `1` is this cell. The four calls are everything it can reach. The guard returning `0` is what makes the sum come out right with no special cases: out of bounds, water and already-sunk land all contribute nothing, and all three are the same `return 0`.

## Watching it run

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

A lone cell: all four neighbours are rejected, so the call returns `1 + 0 + 0 + 0 + 0`.

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

## Why this does not count cells four times

That sum looks wrong the first time you read it. Each of four neighbours recurses into its own four neighbours, and every one of those includes the cell you just came from.

The line above the sum is what prevents it:

```go
(*grid)[i][j] = 0
```

The cell is sunk **before** the four calls run. By the time any neighbour tries to recurse back into it, the guard sees `0` and returns `0`.

Every cell contributes its `1` exactly once, on the single call that first reached it.

Yesterday this same line was what stopped the recursion running forever. Today it is doing that *and* making the arithmetic correct, from one assignment. That is the thing I like about this problem: the fix for termination and the fix for double counting turn out to be the same fix.

## Counting without a set

An earlier version of this kept a `visited` map per island and used its length as the area. Same answers, and a map allocated for every island to count something a return value can carry.

When a traversal is already visiting exactly the cells you want to count, the count can ride back up the call stack. That is all `1 +` is doing.

## Complexity

- **Time: O(m x n)**. Every cell scanned once, entered by a fill at most once more.
- **Space: O(m x n)** worst case, the recursion stack on an all-land grid.

## Builds on

- [Day 37: Number of Islands](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/number_of_islands/) — the same scan-and-sink flood fill; the only change is that the fill now reports how big it was

Full code and the step-by-step walkthrough:
[max_area_of_island](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/601_700/max_area_of_island/SOLUTION.md)

#DSA #LeetCode #Golang #Graphs #DFS #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
