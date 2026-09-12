# Number of Islands — intuition

## Builds on

- [Day 1: Maximum Depth of Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/maximum_depth_of_binary_tree/) — the same recursion shape: visit, recurse into every child, let the base case stop you

## The problem in one line

Count the connected groups of `1`s in a grid, where connected means touching horizontally or vertically.

## The grid is a graph

Nothing here says "graph", and that is the whole point of the problem.

A cell is a node. Its neighbours are the cells above, below, left and right. There is no edge list because the edges are implied by position: you compute a neighbour with arithmetic instead of looking it up. Everything else about the traversal is identical to yesterday's.

Recognising a grid as a graph is worth more than this problem is. Once you see it, maze problems, flood fills, shortest-path-on-a-map problems and image segmentation are all the same shape.

## Counting components

Yesterday's BFS answered "what can I reach from here?" If you run that from every unvisited node and count how many times you had to start, you get the number of connected components. That is exactly what this is.

```go
for every cell:
    if it is land:
        count++
        sink everything reachable from it
```

The count is not counting islands directly. It is counting how many times the scan found land that no previous traversal had already consumed.

## Sinking instead of tracking visited

The traversal marks cells by writing `'2'` into the grid itself rather than keeping a separate visited structure.

That costs nothing and removes an allocation, and it is safe because every cell that gets sunk has already been counted as part of an island. The scan can never start a second island on a cell it has already swallowed.

The trade is that it destroys the caller's grid. For a LeetCode submission that is fine. In code someone else calls, it is the kind of side effect day 23 was careful about, and if it matters you either copy the grid first or keep a separate visited set.

## The guard that makes the recursion short

```go
if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] != '1' {
	return
}
```

Bounds and "not land" are checked in the same place, at the top of the call rather than before it. That means the four recursive calls need no guarding at all: throw any coordinate at `dfs` and it decides for itself whether there is work to do.

Checking before recursing instead gives you four copies of the bounds check at every call site. This version has one.

## Complexity

- **Time: O(m x n).** The scan visits every cell once. The fills also visit each cell at most once in total, because a sunk cell is never re-entered.
- **Space: O(m x n)** in the worst case, which is the recursion stack when the whole grid is one island.
