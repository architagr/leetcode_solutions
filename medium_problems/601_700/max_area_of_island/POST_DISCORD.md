**365 Days of LeetCode Challenge — Day 38/365**
**Max Area of Island** (Medium)
🔗 https://leetcode.com/problems/max-area-of-island/

Yesterday's solution with a return value. That is the whole diff: the flood fill returns an int, and the scan keeps a maximum instead of a count.

```go
func dfs(i, j int, grid *[][]int) int {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] != 1 {
		return 0
	}
	(*grid)[i][j] = 0

	return 1 + dfs(i+1, j, grid) + dfs(i-1, j, grid) + dfs(i, j+1, grid) + dfs(i, j-1, grid)
}
```

The area reachable from a cell is that cell plus what its four neighbours reach. The guard returning 0 covers out of bounds, water and already-sunk land with no special cases.

The sum looks wrong at first - four neighbours each recursing into their own four, every one of which includes the cell you came from. Surely cells get counted twice?

No, because the cell is sunk BEFORE the four calls run. Any neighbour recursing back hits the guard and gets 0. Every cell contributes its 1 exactly once.

Yesterday that same line stopped the recursion running forever. Today it does that and fixes the arithmetic. One assignment, two jobs.

O(m*n) time, O(m*n) stack worst case.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/601_700/max_area_of_island/SOLUTION.md
