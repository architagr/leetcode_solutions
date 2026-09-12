---
meta_title: "A grid is a graph you never have to build"
meta_description: "Number of Islands is component counting with the adjacency list deleted - a cell's neighbours are arithmetic, not a lookup."
---

![Day 37](HERO.png)

## 365 Days of LeetCode Challenge — Day 37/365

**[200. Number of Islands](https://leetcode.com/problems/number-of-islands/)** (Medium)

Count the connected groups of `1`s in a grid, where connected means touching horizontally or vertically.

## Nothing here says graph

That is the point of the problem.

A cell is a node. Its neighbours are the cells above, below, left and right. There is no edge list and no adjacency list, because in a grid the edges are implied by position: a neighbour is arithmetic, not a lookup.

Yesterday half the work was converting an edge list into something usable. Here that half is free. Everything after it is the same traversal.

Recognising a grid as a graph is worth more than this one problem. Mazes, flood fill, shortest path on a map, image segmentation: same shape.

## Counting components

Yesterday's BFS answered "what can I reach from here?" Run that from every unvisited node and count how many times you had to start, and you have counted connected components.

```go
if grid[i][j] != '1' {
	continue
}
count++
dfs(&grid, i, j)
```

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

Finding land increments the count **once**, then sinks the entire island.

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

The count is not counting islands directly. It counts how many times the scan found land that no earlier traversal had already consumed, which is the same number.

## One guard, four unguarded calls

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

Out of bounds, water, and already-sunk are one condition at the top of the function rather than four checks at each call site. Throw any coordinate at `dfs` and it decides for itself whether there is work to do.

The order inside that condition matters: bounds first, then the cell value. Go short-circuits `||`, so the index is only read once it is known to be in range. Swap them and out-of-range coordinates panic.

## Sinking before recursing is not an optimisation

```go
(*grid)[i][j] = '2'
dfs(grid, i-1, j)
```

Day 36 marked nodes on the way in to avoid duplicate work. Here marking first is a *termination* requirement.

Two adjacent land cells each list the other as a neighbour. Without sinking first, `dfs(A)` calls `dfs(B)` calls `dfs(A)`, forever. Sinking first means the re-entrant call sees `'2'` and returns.

## The side effect worth naming

This destroys the caller's grid. It returns a count and leaves the input full of `'2'`.

For a submission, fine. In code with a second caller, it is exactly what day 23 went out of its way to avoid. The fixes are the same: copy first, or keep a separate visited set and pay the allocation.

## Complexity

- **Time: O(m x n)**. Every cell scanned once, and entered by a fill at most once more.
- **Space: O(m x n)** worst case, which is the recursion stack on a grid that is entirely land.

That stack depth is the argument for an iterative BFS on a very large grid: same time, bounded memory.

## Builds on

- [Day 1: Maximum Depth of Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/maximum_depth_of_binary_tree/) — the same recursion shape: visit, recurse into every child, let the base case stop you

Full code and the step-by-step walkthrough:
[number_of_islands](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/number_of_islands/SOLUTION.md)

#DSA #LeetCode #Golang #Graphs #DFS #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
