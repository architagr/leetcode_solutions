**365 Days of LeetCode Challenge — Day 36/365**
**Find if Path Exists in Graph** (Easy)
🔗 https://leetcode.com/problems/find-if-path-exists-in-graph/

New topic: graphs.

The part that gets skipped when people say a graph problem "is just BFS" - the input is not a graph. You get a flat list of edge pairs. Asking that list for a node's neighbours means scanning all of it, and doing that once per node turns O(V + E) into O(V x E).

So half the problem is building an adjacency list. The traversal is the easy half.

```go
queue.Enqueue(source)
visited[source] = true
for !queue.IsEmpty() {
	x := queue.Dequeue().Data
	for key := range at[x] {
		if !visited[key] {
			queue.Enqueue(key)
			visited[key] = true
		}
	}
}
return visited[destination]
```

Two details worth stealing:

An undirected edge goes into the list under BOTH endpoints. Record it once and you have built a directed graph by accident - the traversal then refuses to walk an edge backwards, and that passes a lot of tests before it fails one.

Mark a node visited when you ENQUEUE it, not when you dequeue it. Otherwise a node with five neighbours gets pushed five times before it is ever processed.

O(V + E) time and space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1901_2000/find_if_path_exists_in_graph/SOLUTION.md
