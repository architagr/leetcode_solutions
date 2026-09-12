---
meta_title: "The solution that never looks for a cycle"
meta_description: "A tree is connected and acyclic. Union-find answers the first, and a tree on n nodes having n-1 edges turns the second into arithmetic."
tags: [golang, union-find, graphs, dsa]
---

![Day 43](HERO.png)

*365 Days of LeetCode Challenge — Day 43/365*

**[261. Graph Valid Tree](https://leetcode.com/problems/graph-valid-tree/)** (Medium)

Given `n` nodes labelled `0` to `n-1` and a list of undirected edges, determine whether the graph is a valid tree.

## What a tree actually requires

Two conditions, and it is easy to remember only one of them.

**Connected.** Every node reachable from every other. One component, not several.

**Acyclic.** No cycle anywhere.

Both are load-bearing, and each one fails independently. A forest of two separate trees is perfectly acyclic and is not a tree. A triangle is perfectly connected and is not a tree.

So a correct solution has to establish both.

## The fact that removes half the work

Here is the piece of knowledge this problem is really testing:

> A tree on `n` nodes has exactly `n - 1` edges.

That is not a coincidence or a heuristic, it follows from the definition. A tree is minimally connected: every edge is essential, remove any one and it falls into two pieces. Adding any edge to a tree creates a cycle, because the two endpoints were already connected by exactly one path and now there are two.

Which means, once you know the graph is connected, the edge count answers the cycle question by itself:

```go
if set.count > 1 {
	return false
}
return len(edges) == n-1
```

Walk through the cases:

- Connected with exactly `n - 1` edges. A tree.
- Connected with **more** than `n - 1` edges. There must be a cycle. It is not possible to join `n` nodes with more than `n - 1` edges without creating one, so no search is required to be certain.
- **Fewer** than `n - 1` edges. Cannot be connected at all, so the first check has already returned false.

Read the solution again and notice what is absent. There is no cycle detection anywhere in it. It counts components, then counts edges. That is the whole thing.

I like this problem for exactly that reason. The obvious approach — traverse, track visited, watch for an edge back to somewhere you have been — is correct and is more code than necessary, and the shorter version comes from knowing a property rather than writing a better search.

## A structure that merges instead of walking

Every problem in this arc so far has done the same thing: build an adjacency list, then traverse it. Day 36 built one from an edge list, day 39 from a matrix, day 37 skipped building one only because a grid implies its own.

This solution does not build an adjacency list at all.

Union-find consumes the edge list directly, in the order it arrives, and maintains one number — how many separate components exist right now. That number *is* the connectivity answer, available at any moment without recomputation.

It is worth pausing on how different that is in shape. A traversal explores: it starts somewhere, fans out, and discovers what is reachable. Union-find never explores anything. It is handed relationships one at a time and merges groups accordingly, and at the end it can tell you how many groups survived.

## Initialisation

```go
func InitDisjoinSet(n int) *DisJointSet {
	root := make([]int, n)
	rank := make([]int, n)

	for i := range n {
		root[i] = i
		rank[i] = 1
	}
	return &DisJointSet{count: n, root: root, rank: rank}
}
```

![Step 1](images/walkthrough-1.png)

Every node is its own root, which means `n` components, one per node. Nothing has been joined yet.

## `Find`, and the line that keeps it fast

```go
func (set *DisJointSet) Find(node int) int {
	if node == set.root[node] {
		return node
	}
	set.root[node] = set.Find(set.root[node])
	return set.root[node]
}
```

To find which component a node belongs to, follow `root[node]` upward until you reach a node that is its own root. That root is the component's identity.

The middle line is doing more than it looks:

```go
set.root[node] = set.Find(set.root[node])
```

It does not merely *return* the root. It **writes it back**. As the recursion unwinds, every node on the path from `node` to the root is repointed directly at that root.

So the walk that found the root also flattens the path it walked. A second lookup on any of those nodes is a single step.

This is path compression, and without it the structure degenerates. A sequence of unions can build a long chain, and every `Find` then walks the whole chain, and the whole thing slides toward linear per operation.

## `Union`

```go
func (set *DisJointSet) Union(node1, node2 int) {
	root1 := set.Find(node1)
	root2 := set.Find(node2)

	if root1 == root2 {
		return
	}
	if set.rank[root1] > set.rank[root2] {
		set.root[root2] = root1
	} else if set.rank[root1] < set.rank[root2] {
		set.root[root1] = root2
	} else {
		...
	}
	set.count--
}
```

![Step 2](images/walkthrough-2.png)

Find both roots. If they are the same, the two nodes were already in one component and there is nothing to merge, so the function returns without touching `count`.

Otherwise the edge genuinely joins two separate groups, one root is attached under the other, and `count` drops by one.

That decrement is what makes `count` meaningful at all times. It is never recomputed and never scanned for; it is maintained incrementally, one edge at a time.

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

## The cycle detector that is already written

Look again at this:

```go
if root1 == root2 {
	return
}
```

When an edge's two endpoints already share a root, they were already connected by some earlier path. Adding this edge creates a second path between them, which is a cycle.

This solution does not use that. The edge count catches cycles more cheaply, so the early return exists only to avoid a pointless merge.

But it is worth seeing, because a closely related problem — *which* edge creates the cycle, as in Redundant Connection — is answered by returning the edge at that exact line. The information is already in the structure; this problem just does not need it.

## Union by rank

When two roots merge, which one should adopt the other?

Rank tracks approximate depth. The shallower tree is hung under the deeper one, because attaching the deeper under the shallower produces a result one level taller, and every future `Find` in that component pays for the extra level.

## What the optimisations do not change

Both path compression and union by rank are performance work and nothing else.

Delete both and this solution still returns the same answers on every input. `Find` would walk longer paths and the trees would be taller, and the results would be identical.

What they buy together is an amortised cost of O(α(n)) per operation, where α is the inverse Ackermann function. That function grows so slowly that it is less than 5 for any `n` that will ever be stored in a computer, which is why union-find is usually just called effectively constant time.

## One constraint this leans on

The solution relies on the problem's guarantee that there are no self-loops and no duplicate edges.

A self-loop would have both endpoints sharing a root, so `Union` would ignore it — but `len(edges)` would still count it, and the arithmetic at the end would be wrong. Worth knowing before lifting this code into a problem with looser constraints.

## Complexity

- **Time: O(E x α(n)).** One pass over the edges, each doing two near-constant `Find` calls. No adjacency list is ever built.
- **Space: O(n)** for `root` and `rank`.

## Builds on

- [Day 36: Find if Path Exists in Graph](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1901_2000/find_if_path_exists_in_graph/) — the same edge-list input, handled here without ever building an adjacency list

Full code and the step-by-step walkthrough:
[graph_valid_tree](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/graph_valid_tree/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
