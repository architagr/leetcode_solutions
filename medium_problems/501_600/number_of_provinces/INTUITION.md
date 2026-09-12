# Number of Provinces — intuition

## Builds on

- [Day 36: Find if Path Exists in Graph](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1901_2000/find_if_path_exists_in_graph/) — building an adjacency list, then BFS from a start node
- [Day 37: Number of Islands](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/number_of_islands/) — counting components by counting how many times the scan had to start

## The problem in one line

Given an `n x n` matrix where `isConnected[i][j] == 1` means cities `i` and `j` are directly joined, count the connected groups.

## It is day 37 with a different input format

Day 37 counted connected groups in a grid. This counts connected groups in a graph given as a matrix. The counting logic is identical:

```
for every node:
    if not visited:
        count++
        traverse everything reachable from it
```

What changes is only how you find a node's neighbours, and that is worth paying attention to because the matrix is a third representation, after day 36's edge list and day 37's implicit grid.

## The matrix is not a grid

This is the trap. `isConnected` looks like day 37's grid and means something completely different.

In day 37, cell `(i, j)` was a place. Here, cell `(i, j)` is a *relationship* between city `i` and city `j`. The matrix has `n²` cells describing `n` nodes.

So you do not traverse the matrix. You traverse the cities, and consult the matrix to find out who is adjacent to whom.

## Reading only above the diagonal

```go
for j := i + 1; j < len(isConnected[i]); j++ {
```

The matrix is symmetric: if `i` is joined to `j` then `j` is joined to `i`, so the same fact is stored twice. Reading only the upper triangle sees each relationship once, and the loop adds both directions to the adjacency list itself.

It also skips the diagonal, where `isConnected[i][i]` is always 1 and means nothing useful.

## Why convert to an adjacency list at all

You could BFS straight off the matrix, asking `isConnected[x][k]` for every `k`. That is O(n) per node regardless of how many neighbours it actually has, so the traversal is O(n²) even on a sparse graph.

Converting first costs one O(n²) pass and then lets the traversal cost O(V + E). For a dense graph that is no gain at all. For a sparse one it is the difference between scanning every city and walking the few that are actually joined.

## Complexity

- **Time: O(n²).** The matrix has to be read, and reading it is n² whatever you do afterwards.
- **Space: O(n + E)** for the adjacency list, plus the visited set and queue.
