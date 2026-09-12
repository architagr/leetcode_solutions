---
meta_title: "The day BFS and DFS stop being interchangeable"
meta_description: "Rotting Oranges asks when each orange rots, which is a distance. Multi-source BFS answers it in one pass, and DFS answers it wrongly."
---

![Day 40](HERO.png)

## 365 Days of LeetCode Challenge — Day 40/365

**[994. Rotting Oranges](https://leetcode.com/problems/rotting-oranges/)** (Medium)

Rot spreads from every rotten orange to its four neighbours once per minute. How many minutes until nothing fresh remains, or `-1` if something never rots?

## The promise from day 36 comes due

Four days ago I said reachability can be answered with BFS or DFS equally well, and that the distinction starts mattering the moment a problem asks about **distance**. This is that problem.

BFS reaches every cell by a shortest path, because it expands outward one ring at a time. DFS reaches a cell by whatever route it wandered down first, which can be far longer.

"When does this orange rot?" is a shortest-path question: an orange rots as soon as the **nearest** rotten one gets to it. Solve this with DFS and you get wrong answers, not slow ones.

## Multi-source BFS

The usual BFS has one start. Here every already-rotten orange starts spreading at the same instant, so they all go into the queue before the walk begins.

```go
for i := 0; i < len(grid); i++ {
	for j := 0; j < len(grid[i]); j++ {
		if grid[i][j] == 2 {
			q = append(q, node{row: i, col: j, t: 0})
		}
	}
}
```

![Step 1](images/walkthrough-1.png)

That is the whole trick, and it is less clever than it sounds. A BFS seeded with many sources expands as one combined wavefront, and every cell ends up with its distance to the **nearest** source. Nothing inside the loop needs to know there was more than one start. The queue does not care where its contents came from.

The alternative, one BFS per rotten orange taking the minimum at each cell, is also correct and costs O(sources x cells) instead of O(cells).

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

An empty cell is rejected by the guard, so the rot goes around it rather than through.

![Step 5](images/walkthrough-5.png)

## One line that looks sloppy and is not

```go
ans = x.t
```

Not `if x.t > ans`. It overwrites every single pop.

That is correct because BFS dequeues in non-decreasing order of `t`: everything at minute 1 comes out before anything at minute 2. The last value assigned is therefore the largest, and a comparison would be work done to reach the same answer.

It depends entirely on the queue being FIFO. Swap it for a stack and this line becomes silently wrong.

## The final scan cannot be skipped

```go
if grid[i][j] == 1 {
	return -1
}
```

BFS only visits what it can reach. An orange sealed behind empty cells is never enqueued, never rots, and the queue drains with no indication that it exists.

So `-1` cannot be detected during the traversal. It has to be a scan afterwards.

This is the same fact as day 36's two-component example, where a walk from node 0 simply never learned that nodes 3, 4 and 5 were there. Unreachable things do not announce themselves; you have to go and look.

## Complexity

- **Time: O(m x n)**. Each cell is pushed at most once, plus two scans.
- **Space: O(m x n)** for the queue, worst case a fully rotten grid.

## Builds on

- [Day 29: Binary Tree Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_level_order_traversal/) — BFS visiting a whole level before the next; here a level is a minute
- [Day 37: Number of Islands](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/number_of_islands/) — a grid as a graph, with neighbours computed rather than looked up

Full code and the step-by-step walkthrough:
[rotting_oranges](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/901_1000/rotting_oranges/SOLUTION.md)

#DSA #LeetCode #Golang #Graphs #BFS #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
