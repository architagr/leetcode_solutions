**365 Days of LeetCode Challenge — Day 43/365**
**Graph Valid Tree** (Medium)
🔗 https://leetcode.com/problems/graph-valid-tree/

A graph is a tree when it is connected AND acyclic. A forest is acyclic but disconnected; a triangle is connected but cyclic.

The shortcut: a tree on n nodes has exactly n-1 edges. Once you know the graph is connected, the edge count settles the cycle question.

```go
func validTree(n int, edges [][]int) bool {
	set := InitDisjoinSet(n)
	for _, e := range edges {
		set.Union(e[0], e[1])
	}
	if set.count > 1 {
		return false
	}
	return len(edges) == n-1
}
```

This never searches for a cycle. It counts.

Union-find is the first structure in this arc that does not traverse. Every problem so far built an adjacency list and walked it - this consumes the edge list directly and maintains one number, the component count, which is exactly the connectivity answer.

Two things worth knowing. Path compression is one line - Find writes the root back into every node it passed, so the next lookup is one step. And when an edge's endpoints already share a root, that edge closes a cycle; this solution does not need that, but it is sitting right there if a problem asks which edge.

O(E * alpha(n)) time, effectively linear.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/graph_valid_tree/SOLUTION.md
