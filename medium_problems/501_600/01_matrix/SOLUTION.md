# 01 Matrix — solution walkthrough

From `main.go`:

```go
func updateMatrix(mat [][]int) [][]int {
	result := make([][]int, len(mat))
	n := len(mat)
	m := len(mat[0])
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = mat[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if result[i][j] == 0 {
				continue
			}
			minNeighbor := m * n
			if i > 0 {
				minNeighbor = minValue(minNeighbor, result[i-1][j])
			}
			if j > 0 {
				minNeighbor = minValue(minNeighbor, result[i][j-1])
			}
			result[i][j] = minNeighbor + 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if result[i][j] == 0 {
				continue
			}
			minNeighbor := m * n
			if i < n-1 {
				minNeighbor = minValue(minNeighbor, result[i+1][j])
			}
			if j < m-1 {
				minNeighbor = minValue(minNeighbor, result[i][j+1])
			}
			result[i][j] = minValue(result[i][j], minNeighbor+1)
		}
	}
	return result
}
```

## This is yesterday's problem, and not yesterday's solution

Rotting Oranges asked: how many minutes until this orange rots, with every rotten orange spreading at once. This asks: how far is this cell from the nearest zero, with every zero a source.

Same question. Yesterday's multi-source BFS would solve it directly, and reaching for it here is a correct instinct.

This solution does something different, and that is the reason to read it.

## Breaking a circular definition

The distance from a cell to its nearest zero is one more than the smallest distance among its four neighbours.

True, and useless as written, because every cell's answer depends on its neighbours' answers and theirs depend on it. There is no order in which to evaluate that definition.

The circle breaks if you split it by direction. Any shortest path leaves the cell in one of four directions. Group them: up-or-left, and down-or-right.

Sweep the grid top-left to bottom-right and, at the moment you reach a cell, every cell above it and to its left is already finished. So for that cell you can correctly compute the best distance *among paths whose first step goes up or left*. Sweep the other way and you get the other two directions.

![Step 1](images/walkthrough-1.png)

## Pass one: up and left

```go
for i := 0; i < n; i++ {
	for j := 0; j < m; j++ {
		if result[i][j] == 0 {
			continue
		}
		minNeighbor := m * n
		if i > 0 {
			minNeighbor = minValue(minNeighbor, result[i-1][j])
		}
		if j > 0 {
			minNeighbor = minValue(minNeighbor, result[i][j-1])
		}
		result[i][j] = minNeighbor + 1
	}
}
```

![Step 2](images/walkthrough-2.png)

Only `result[i-1][j]` and `result[i][j-1]` are read: the cell above and the cell to the left. Both were finalised earlier in this same sweep, which is exactly why the sweep goes in this direction.

`minNeighbor` starts at `m * n`, standing in for infinity. It is safe because no real distance in an `m x n` grid can reach `m * n`, so a cell with no usable neighbour ends up with a value large enough to be beaten by anything real in the second pass.

Zeros are skipped by `continue`, which leaves them at 0 permanently. They are the sources and their distance is already correct.

![Step 3](images/walkthrough-3.png)

After this pass, a cell whose nearest zero lies up or left is correct. Everything else is too large.

## Pass two: down and right

```go
for i := n - 1; i >= 0; i-- {
	for j := m - 1; j >= 0; j-- {
		...
		result[i][j] = minValue(result[i][j], minNeighbor+1)
	}
}
```

![Step 4](images/walkthrough-4.png)

The loops run backwards, so `result[i+1][j]` and `result[i][j+1]` have already been finalised by this sweep.

The last line is the one that matters:

```go
result[i][j] = minValue(result[i][j], minNeighbor+1)
```

`minValue(result[i][j], ...)` rather than a plain assignment. Pass one's answer is kept when it was already better. Overwrite instead and you throw away every correct distance the first pass found.

That single `minValue` is what makes two passes sufficient rather than merely sequential.

## Why exactly two

Every shortest path in a four-directional grid is a sequence of steps that never needs to backtrack. Pass one covers all shortest paths whose steps go up and left; pass two covers the rest, and because it takes a minimum, it can also improve on a value pass one computed.

There is no third case, so there is no third pass.

## The copy at the top

```go
result[i][j] = mat[i][j]
```

The input is copied rather than modified, so `updateMatrix` leaves `mat` untouched.

Worth calling out because days 37, 38 and 40 all mutated the grid they were given. Those problems returned a number and quietly ate their input. This one returns a grid, so allocating is unavoidable, and the function ends up with no side effects almost by accident.

## BFS or this?

In an interview, multi-source BFS. It is the natural framing of a distance problem, it generalises to weighted edges and higher dimensions, and it is easier to explain out loud.

This version is worth knowing because it allocates no queue and touches each cell a fixed number of times, and because the underlying pattern generalises: when a definition is circular, look for an evaluation order that makes part of it acyclic, then cover the rest with a second pass in the opposite order.

## Complexity

- **Time: O(m x n).** Two sweeps, constant work per cell.
- **Space: O(m x n)** for the output, which the problem requires. No queue, no visited set.

## Test

`main_test.go` covers both worked examples, including the grid whose bottom row is all ones, which is the case that fails if pass two overwrites instead of taking a minimum.
