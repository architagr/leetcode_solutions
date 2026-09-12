# Rotting Oranges — intuition

## Builds on

- [Day 29: Binary Tree Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_level_order_traversal/) — BFS visiting a whole level before the next; here a level is a minute
- [Day 37: Number of Islands](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/number_of_islands/) — a grid as a graph, with neighbours computed rather than looked up

## The problem in one line

Rot spreads from every rotten orange to its four neighbours once per minute. How many minutes until nothing fresh is left, or `-1` if something never rots?

## This is the day BFS stops being interchangeable

Day 36 said reachability can be answered by BFS or DFS, and promised that would change when a problem asked about distance. Here it is.

BFS reaches every cell by a shortest path. DFS reaches it by whatever route it wandered down first. This problem asks *when* each orange rots, which is a distance, so DFS gives wrong answers rather than slower ones.

## Multi-source: every rotten orange starts at minute 0

The usual BFS has one start node. Here every already-rotten orange starts rotting its neighbours simultaneously, so they all go into the queue before the walk begins.

That is the entire trick, and it is smaller than it sounds. A BFS seeded with many sources expands as one combined wavefront, and each cell ends up with its distance to the *nearest* source rather than to any particular one. Nothing in the loop has to know there was more than one start.

Seeding one source at a time and taking the minimum would also work, and would be O(sources x cells) instead of O(cells).

## Carrying the minute on the node

Each queued cell carries the minute it rotted. A neighbour rots at `t + 1`. The answer is the largest `t` ever dequeued.

Because BFS dequeues in non-decreasing order of `t`, the last value seen is the maximum, so `ans = x.t` on every pop lands on the right number without a comparison.

## Marking on push, not on pop

A cell is turned rotten at the moment it is pushed:

```go
grid[row][col] = 2
q = append(q, node{...})
```

Day 36's rule again, and here it also removes the need for a separate visited structure: the grid itself records what has been reached.

## The final scan is not optional

BFS only visits what it can reach. An orange walled off by empty cells is never enqueued and never rots, and the queue empties without ever noticing it.

So after the walk, the grid has to be checked for any remaining `1`. That is what produces `-1`, and it cannot be detected during the traversal.

## Complexity

- **Time: O(m x n).** Every cell is enqueued at most once, plus two full scans.
- **Space: O(m x n)** for the queue in the worst case.
