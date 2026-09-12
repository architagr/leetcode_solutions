# Graph Valid Tree — intuition

## Builds on

- [Day 36: Find if Path Exists in Graph](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1901_2000/find_if_path_exists_in_graph/) — the same edge-list input, handled here without ever building an adjacency list

## The problem in one line

Given `n` nodes and a list of undirected edges, is the graph a tree?

## What makes a graph a tree

Two conditions, and both are needed:

1. **Connected.** One component, not several.
2. **Acyclic.** No cycle anywhere.

Miss either and you do not have a tree. A forest of two trees is acyclic but not connected. A triangle is connected but not acyclic.

## The counting shortcut

There is a fact about trees worth knowing: a tree on `n` nodes has exactly `n - 1` edges.

That gives a much shorter test. Check connectivity, then check the edge count:

- Connected and `n - 1` edges: a tree.
- Connected with more than `n - 1` edges: there must be a cycle, because you cannot connect `n` nodes with more than `n - 1` edges without creating one.
- Connected with fewer: impossible, since `n - 1` is the minimum to connect `n` nodes at all.

So once connectivity is established, `len(edges) == n-1` settles the cycle question without ever looking for a cycle.

## Why union-find rather than a traversal

You could BFS from node 0, check every node was reached, and count edges. That works, and it means building an adjacency list first.

Union-find skips the adjacency list entirely. It consumes the edge list directly, and it maintains one number, the component count, that answers the connectivity question for free.

This is the first genuinely new data structure in the arc, and its shape is different from anything so far: it does not traverse. It merges.

## How it works

Every node starts as its own root, so there are `n` components.

For each edge, find both endpoints' roots. If they differ, the edge joins two separate components, so point one root at the other and decrement the count. If they are the same, both endpoints were already connected, and this edge is a cycle — though this solution does not need to notice that, because the edge count catches it.

When all the edges are processed, the count is the number of components. One means connected.

## The two optimisations that make it fast

**Path compression.** When `Find` walks up to a root, it repoints every node it passed directly at that root, so the next lookup is immediate.

**Union by rank.** When merging, the shallower tree is hung under the deeper one, which stops the structure degenerating into a long chain.

Neither changes what the structure computes. Both are about keeping `Find` near constant time.

## Complexity

- **Time: O(E x α(n))**, where α is the inverse Ackermann function. Effectively linear in the number of edges.
- **Space: O(n)** for the two arrays.
