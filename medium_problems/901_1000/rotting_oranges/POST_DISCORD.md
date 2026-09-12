**365 Days of LeetCode Challenge — Day 40/365**
**Rotting Oranges** (Medium)
🔗 https://leetcode.com/problems/rotting-oranges/

Day 36 said BFS and DFS answer reachability equally well, and that it changes when a problem asks about distance. This is that problem.

BFS reaches every cell by a shortest path; DFS reaches it by whatever route it wandered down first. "When does this orange rot?" is shortest-path - it rots when the NEAREST rotten one arrives. DFS here is wrong, not slow.

The trick is multi-source BFS. Every already-rotten orange goes in the queue before the walk starts:

```go
for i := range grid {
	for j := range grid[i] {
		if grid[i][j] == 2 {
			q = append(q, node{row: i, col: j, t: 0})
		}
	}
}
```

Seeded with many sources, BFS expands as one combined wavefront and every cell ends up with its distance to the nearest source. Nothing in the loop knows there was more than one start.

Two details worth stealing. `ans = x.t` with no comparison is correct because BFS dequeues in non-decreasing t, so the last value is the largest - but it breaks silently if you swap the queue for a stack.

And the final scan is not optional: BFS only visits what it can reach, so an orange sealed behind empty cells never rots and the queue drains without noticing. -1 has to be looked for afterwards.

O(m*n) time and space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/901_1000/rotting_oranges/SOLUTION.md
