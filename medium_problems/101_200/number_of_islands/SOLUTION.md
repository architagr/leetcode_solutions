# Number of Islands — solution walkthrough

From `main.go`:

```go
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != '1' {
				continue
			}
			count++
			dfs(&grid, i, j)
		}
	}
	return count
}

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

## The grid is the graph

There is no adjacency list here and no edge list, because in a grid the edges are implied by position. Node `(i, j)`'s neighbours are `(i±1, j)` and `(i, j±1)`, computed rather than looked up.

That is the only structural difference from day 36. The traversal is the same traversal.

## The scan

![Step 1](images/walkthrough-1.png)

```go
if grid[i][j] != '1' {
	continue
}
count++
dfs(&grid, i, j)
```

![Step 2](images/walkthrough-2.png)

Finding land increments the count **once** and then removes the entire island from the grid.

The count is not counting islands directly. It counts how many times the scan encountered land that no earlier traversal had already consumed, and that is the same number.

![Step 3](images/walkthrough-3.png)

Every cell of the first island is now `'2'`, so when the scan reaches those positions later it skips them.

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

Three starts, three islands.

## The guard does two jobs

```go
if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] != '1' {
	return
}
```

Out of bounds, water, and already-sunk land are all handled by one condition at the top of the function.

That is what lets the four recursive calls be unguarded:

```go
dfs(grid, i-1, j)
dfs(grid, i+1, j)
dfs(grid, i, j-1)
dfs(grid, i, j+1)
```

Any coordinate can be thrown at `dfs`, valid or not, and it decides for itself whether there is anything to do. Check before recursing instead and the bounds test appears four times at every call site.

Note the order of the condition: the bounds tests come before `(*grid)[i][j]`, and Go short-circuits `||`, so the index is only evaluated once it is known to be in range. Reverse them and out-of-range coordinates panic.

## Marking before recursing

```go
(*grid)[i][j] = '2'
dfs(grid, i-1, j)
```

The cell is sunk **before** the recursive calls, not after.

This is the same rule as day 36's "mark on the way in", and here it is not an efficiency question but a termination one. Two adjacent land cells each list the other as a neighbour, so without marking first, `dfs(A)` calls `dfs(B)` which calls `dfs(A)` forever.

Sinking first means the guard at the top of the re-entrant call sees `'2'`, not `'1'`, and returns.

## Why `'2'` and not `'0'`

Either works. Writing `'2'` leaves the finished grid distinguishable: `'0'` is original water, `'2'` is land that was consumed. Nothing in this solution reads that distinction, but it makes the state visible while debugging.

## The side effect

This destroys the caller's grid. `numIslands` returns a count and silently leaves its input full of `'2'`.

For a LeetCode submission, fine. In code with more than one caller it is the same problem day 23 took care to avoid, and the fixes are the same: copy the grid first, or keep a separate visited structure and pay the allocation.

## Complexity

- **Time: O(m x n).** The scan touches every cell once. Across all fills, each cell is entered at most once more, because a sunk cell is rejected by the guard.
- **Space: O(m x n)** worst case. That is the recursion stack, and the worst case is a grid that is entirely land, where the recursion is as deep as there are cells.

That stack depth is the argument for the iterative BFS version on a very large grid: same time, bounded memory.

## Test

`main_test.go` runs the two worked examples: the 4x5 grid that is one large island, and the grid of three separate ones.
