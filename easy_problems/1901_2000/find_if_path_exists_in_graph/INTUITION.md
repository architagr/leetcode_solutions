# Find if Path Exists in Graph — intuition

## The problem in one line

Given an undirected graph as a list of edges, is `destination` reachable from `source`?

## The first thing to notice: the input is not a graph

You get `edges`, a flat list of pairs. To ask "what are node 3's neighbours?" against that list you would scan all of it, every time, which turns a cheap question into an O(E) one.

So the first half of this problem is building an adjacency list: a map from each node to the nodes it touches. Do it once, in one pass over the edges, and every neighbour lookup afterwards is immediate.

This is the step that gets skipped when people say a graph problem "is just BFS". The traversal is the easy part. Choosing a representation that makes the traversal cheap is the part that is actually a decision.

## Undirected means both directions

```go
val[edges[i][1]] = true   // u knows about v
val[edges[i][0]] = true   // and v knows about u
```

An edge in an undirected graph has to be recorded twice, once under each endpoint. Record it once and you have quietly built a directed graph, and the traversal will then miss paths that run "backwards" along an edge.

## Then it is a traversal

Push the source, mark it visited, and repeatedly take a node off the queue and push its unvisited neighbours. When the queue empties, `visited` holds exactly the nodes reachable from the source, and the answer is whether the destination is among them.

## Marking on the way in, not on the way out

Nodes are marked visited when they are *enqueued*, not when they are dequeued.

If you mark on dequeue, a node with several neighbours can be pushed several times before it is ever processed, and in a dense graph that is a lot of duplicated work. Marking on enqueue guarantees every node enters the queue at most once.

## BFS or DFS makes no difference here

The question is reachability, not distance, so either traversal answers it. BFS is what this solution uses; a recursive DFS would be shorter to write and equally correct.

That changes the moment a problem asks for the *shortest* path, because BFS reaches every node by a shortest route and DFS does not. Day 40 is the problem where that distinction starts paying.

## Complexity

- **Time: O(V + E).** One pass to build the adjacency list, and a traversal that visits each node once and each edge twice.
- **Space: O(V + E).** The adjacency list dominates; the queue and visited array are O(V).
