**365 Days of LeetCode Challenge — Day 37/365**
**Number of Islands** (Medium)
🔗 https://leetcode.com/problems/number-of-islands/

Nothing in this problem says "graph", and that is the point.

A cell is a node; its neighbours are up, down, left, right. There is no adjacency list to build because in a grid the edges are implied by position - a neighbour is arithmetic, not a lookup. Yesterday half the work was converting an edge list. Here that half is free.

```go
func dfs(grid *[][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] != '1' {
		return
	}
	(*grid)[i][j] = '2'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
```

Two details worth stealing.

One guard at the top covers bounds, water and already-sunk, so the four recursive calls need no checks at all. Bounds come first in that condition because Go short-circuits `||` - swap them and out-of-range coordinates panic.

Sink the cell BEFORE recursing. Two adjacent land cells each list the other as a neighbour, so without it dfs(A) calls dfs(B) calls dfs(A) forever. Not an optimisation - termination.

O(m*n) time. Space is the recursion stack, O(m*n) when the grid is all land.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/number_of_islands/SOLUTION.md
