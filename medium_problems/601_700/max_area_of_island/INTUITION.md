# Max Area of Island — intuition

## Builds on

- [Day 37: Number of Islands](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/number_of_islands/) — the same scan-and-sink flood fill; the only change is that the fill now reports how big it was

## The problem in one line

Same grid, same definition of an island. Instead of counting islands, return the size of the largest one.

## Yesterday's solution with a return value

That is genuinely all this is.

Yesterday the fill was called for its side effect: it sank an island so the scan would not count it twice, and it returned nothing. Today the fill has to answer a question as well, and the question has an obvious recursive shape.

The area of the island reachable from a cell is: this cell, plus the area reachable from each of its four neighbours.

```go
return 1 + dfs(i+1, j, grid) + dfs(i-1, j, grid) + dfs(i, j+1, grid) + dfs(i, j-1, grid)
```

The `1` is the current cell. The four calls are everything it can reach. The guard returning `0` for water and out-of-bounds is what makes the sum come out right without any special cases.

## Why nothing gets counted twice

The sum looks suspicious the first time you see it. Four neighbours each recursing into their own four neighbours sounds like it should count cells repeatedly.

It does not, because the cell is sunk *before* the recursion:

```go
(*grid)[i][j] = 0
return 1 + dfs(...) + ...
```

By the time any of those four calls runs, the current cell reads as water, so whichever of them tries to come back gets `0` from the guard. Every cell contributes its `1` exactly once, on the single call that first reached it.

This is the same line that made yesterday's version terminate. Here it does double duty: it stops the infinite recursion *and* it makes the arithmetic correct.

## Counting with a return value rather than a set

An earlier version of this solution kept a `visited` map per island and used its length as the area. That works, and it allocates a map for every island to count something a return value can carry.

The recursive sum needs no storage at all. When the recursion is already visiting exactly the cells you want to count, the count can ride back up the call stack.

## Complexity

- **Time: O(m x n).** The scan visits each cell once, and the fills enter each cell at most once more.
- **Space: O(m x n)** worst case, which is the recursion stack on a grid that is entirely land.
