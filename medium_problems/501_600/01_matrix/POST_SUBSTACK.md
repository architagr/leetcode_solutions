---
meta_title: "Breaking a circular definition with two sweeps"
meta_description: "01 Matrix is yesterday's BFS problem solved without a queue. Every cell depends on its neighbours, and direction is what breaks the circle."
tags: [golang, dynamic-programming, grids, dsa]
---

![Day 41](HERO.png)

*365 Days of LeetCode Challenge — Day 41/365*

**[542. 01 Matrix](https://leetcode.com/problems/01-matrix/)** (Medium)

Given a grid of zeros and ones, return a grid of the same shape where every cell holds its distance to the nearest zero.

## You have already solved this

Yesterday: how many minutes until this orange rots, given that every rotten orange spreads at once.

Today: how far is this cell from the nearest zero, given that every zero is a source.

Those are the same question with different nouns. Yesterday's multi-source BFS, seeded with every zero instead of every rotten orange, solves this directly and correctly. If that is where your mind went reading the problem, your instinct is good and you should trust it.

So why write about a different solution?

Because this one contains an idea that BFS does not, and the idea transfers to problems where no queue is available.

## The definition is circular

Start from what is obviously true:

> The distance from a cell to its nearest zero is one more than the smallest distance among its four neighbours.

That statement is correct. It is also unusable, and it is worth being precise about why.

Every cell's value depends on its four neighbours' values. Those neighbours' values depend on *their* neighbours, one of which is the original cell. There is no cell you can evaluate first, because every cell is waiting on something that is waiting on it.

A recursive definition needs a base case and an order of evaluation. This has a base case, the zeros. What it lacks is an order.

## Direction supplies the order

Here is the move.

Any shortest path from a cell to a zero has to leave the cell in one of four directions. Split those four into two groups: **up or left**, and **down or right**.

Now sweep the grid from the top-left corner to the bottom-right. When the sweep arrives at a cell, it has already finished every cell above it and every cell to its left. Those values are final.

So for that cell, you can correctly compute the best distance *among the paths whose first step goes up or left*. Not the true answer yet, but a correct answer to a restricted question, computed with no circularity at all, because the cells you depend on are behind you.

Then sweep the other way, bottom-right to top-left, and the cells below and to the right are the finished ones. That covers the other two directions.

Take the minimum of the two results and every direction has been accounted for.

![Step 1](images/walkthrough-1.png)

## Pass one

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

Note which cells are read: `result[i-1][j]` and `result[i][j-1]`. Up and left. Nothing else. Both were finalised earlier in this same sweep, which is precisely why the sweep runs in this direction and not some other.

`minNeighbor` starts at `m * n`, standing in for infinity. That works because no genuine distance inside an `m x n` grid can be as large as `m * n`, so a cell with no usable neighbour ends up holding a value that anything real will beat in the second pass.

The `continue` on zeros leaves them at zero forever. They are the sources; their distance is already right.

![Step 3](images/walkthrough-3.png)

After this pass, any cell whose nearest zero happens to lie up or left of it is correct. Every other cell holds a number that is too large. Not garbage, just an overestimate, which is exactly the right kind of wrong to be.

## Pass two, and the line that matters

```go
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
```

![Step 4](images/walkthrough-4.png)

The loops run backwards, so now `result[i+1][j]` and `result[i][j+1]` are the finished ones.

And then:

```go
result[i][j] = minValue(result[i][j], minNeighbor+1)
```

This is the line the whole solution turns on, and it is the one line that differs structurally from pass one.

Pass one assigned. Pass two takes a **minimum with what is already there**.

If it assigned instead, pass two would simply overwrite pass one's work, and the final answer would only ever reflect paths going down and right. Every correct distance the first sweep found would be destroyed on the way back.

Taking the minimum is what makes the two passes *combine* rather than one of them winning. It is a small syntactic difference and the entire correctness argument lives inside it.

## Why exactly two and not more

It is fair to wonder whether two sweeps really cover everything, or whether some awkward path needs a third.

Two is enough because of a property of the grid: a shortest path between two cells in a four-directional grid never needs to backtrack. It is a monotone staircase. So every shortest path is either composed of up-and-left steps, or down-and-right steps, or a mix — and a mix is handled because pass two takes a minimum against pass one's partial answer rather than starting fresh.

There is no third category of direction, so there is no third pass.

## A difference worth noticing

```go
result[i][j] = mat[i][j]
```

This copies the input before touching anything, so `mat` comes back untouched.

Days 37, 38 and 40 all consumed the grid they were given. Each of them returned a number and quietly rearranged its caller's data on the way.

This one does not, and I want to be honest that it is not because it is more principled. It has to return a grid, so it had to allocate one regardless. Having no side effects fell out of the signature rather than out of a decision.

Still worth noticing, because it is the same distinction day 23 turned on, and here the shape of the problem made the right thing free.

## Which one should you actually write?

In an interview: the BFS. It is the natural framing of a distance problem, it extends without modification to weighted moves, to eight-directional grids, to three dimensions, and it is much easier to explain out loud while typing.

This version earns its place for two reasons. It allocates nothing beyond the output and touches each cell a fixed number of times, so it is faster in practice. And the underlying pattern is general: *when a definition is circular, look for an evaluation order that makes part of it acyclic, then cover the remainder with a second pass in the opposite order*. That shows up well beyond grids.

## Complexity

- **Time: O(m x n).** Two sweeps, constant work per cell.
- **Space: O(m x n)** for the output, which the problem requires anyway. No queue, no visited set, nothing else.

## Builds on

- [Day 40: Rotting Oranges](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/901_1000/rotting_oranges/) — the same question, distance from every cell to the nearest source, answered there with a queue and here without one

Full code and the step-by-step walkthrough:
[01_matrix](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/501_600/01_matrix/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
