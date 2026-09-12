# Clone Graph — intuition

## Builds on

- [Day 37: Number of Islands](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/number_of_islands/) — recursive DFS where marking a node before recursing is what stops it looping back

## The problem in one line

Deep copy a connected undirected graph: new nodes, same structure, no shared references with the original.

## Why a tree copy does not work here

Copying a binary tree is four lines. Make a node, recurse left, recurse right, done. It works because a tree has no cycles: recursion strictly descends and every node is reached exactly once.

A graph has neither guarantee. Two nodes can point at each other, so a naive recursive copy walks A to B to A to B forever. And a node can be reached by several routes, so even without a cycle you would copy it once per route and end up with a structure that has the right values and the wrong shape.

## One map fixes both problems

Keep a map from original node to its copy.

Before cloning a node, look it up. If it is there, return the copy that already exists instead of making another.

That single lookup solves both failures at once:

- **The cycle** stops, because the second time you reach a node it is already in the map and the recursion returns instead of descending.
- **The shape** is preserved, because every original node maps to exactly one copy, so two neighbours pointing at the same node produce two references to the same copy.

The second point is the one worth dwelling on. If the map only prevented infinite recursion you could use a plain visited set. It has to be a map to the *copies*, because the copies are what the reconstruction needs.

## Recording before recursing

The order inside the function is the whole trick:

```go
newNode = new(Node)
newNode.Val = node.Val
visited[node.Val] = newNode      // record it...
for _, n := range node.Neighbors {
	... = clone(n, visited)       // ...then recurse
}
```

The copy is put in the map **while it is still empty**, before its neighbours are filled in.

If you recursed first and recorded afterwards, a cycle would return to a node that is not in the map yet and recurse again. The map has to contain the half-built node for the cycle to terminate.

This is day 37's "sink the cell before recursing" in a different costume. Both work because a partially finished node in the structure is enough to stop a re-entrant call.

## Complexity

- **Time: O(V + E).** Each node is cloned once, and each edge is followed once from each end.
- **Space: O(V)** for the map, plus recursion depth up to O(V).
