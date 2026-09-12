---
meta_title: "The graph problem is not the traversal"
meta_description: "You are handed a list of edges, not a graph. Building an adjacency list is the real decision, and recording an undirected edge once is the classic bug."
tags: [golang, graphs, bfs, dsa]
---

![Day 36](HERO.png)

*365 Days of LeetCode Challenge — Day 36/365*

**[1971. Find if Path Exists in Graph](https://leetcode.com/problems/find-if-path-exists-in-graph/)** (Easy)

Graphs start today, and this is the gentlest one in the set: given an undirected graph and two nodes, is there a path between them?

Before the algorithm, a word about what carries over. The linked list arc that just finished kept returning to one idea, that a structure is defined entirely by what points at what. A graph is that idea with the restriction removed. A list node points at one thing, a tree node at two, a graph node at as many as it likes, including back at things that already point to it. Everything that made linked lists fiddly is still here, with fewer guarantees.

## The input is not a graph

I want to start here because it is the part that gets waved away when someone says a graph problem "is just BFS".

What arrives is `edges`: a flat array of pairs, each naming two nodes that touch. That is a perfectly good way to *store* a graph and a terrible way to *use* one.

Ask that array "what are node 3's neighbours?" and there is no answer except scanning the entire thing. Inside a traversal, which asks that question once per node, an O(V + E) algorithm quietly becomes O(V x E).

So the first half of this problem is a conversion:

```go
at := make(map[int]map[int]bool)
for i := 0; i < len(edges); i++ {
	val, ok := at[edges[i][0]]
	if !ok {
		val = make(map[int]bool)
	}
	val[edges[i][1]] = true
	at[edges[i][0]] = val

	val, ok = at[edges[i][1]]
	if !ok {
		val = make(map[int]bool)
	}
	val[edges[i][0]] = true
	at[edges[i][1]] = val
}
```

One pass over the edges, and afterwards every neighbour lookup is immediate.

The traversal that follows is the easy half. Choosing this representation is the half that is actually a decision, and it is the decision that determines the complexity.

## Recording an undirected edge twice

Look at what that loop does with each pair. It writes the edge under `edges[i][0]` **and** under `edges[i][1]`.

That is not redundancy. In an undirected graph an edge is a mutual relationship, and an adjacency list stores relationships from one side at a time, so a mutual one has to be written from both sides.

Record it only once and you have built a directed graph without intending to. Your traversal will then decline to walk an edge "backwards", and it will report no path on inputs where a path exists but happens to approach from the other end.

I single this out because of how it fails. It does not crash, and it does not fail every test. It fails the ones where the path runs against the order the edges happened to be listed in, which means you can watch several examples pass and conclude the code is fine.

## The traversal

```go
queue.Enqueue(source)
visited[source] = true
for !queue.IsEmpty() {
	x := queue.Dequeue().Data
	val := at[x]
	for key := range val {
		if !visited[key] {
			queue.Enqueue(key)
			visited[key] = true
		}
	}
}
return visited[destination]
```

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

The second example is the instructive one. The graph is in two pieces, and a walk starting in one of them cannot reach the other no matter how long it runs. `visited[5]` is still false when the queue empties, so the answer is false.

That is worth holding onto, because "which nodes can this walk reach" is the same question as "which component is this node in", and day 39 is that question asked directly.

## Mark on the way in

```go
if !visited[key] {
	queue.Enqueue(key)
	visited[key] = true
}
```

The node is marked when it goes **into** the queue, not when it comes out.

It is easy to write it the other way. You dequeue a node, mark it, look at its neighbours. It still terminates and it still gives the right answer, so it is not wrong exactly. What it does is let a node be enqueued several times before it is first processed: five neighbours all see it as unvisited, all push it, and the queue now holds five copies of the same node, each of which will expand its neighbours again when its turn comes.

On a sparse graph you will not notice. On a dense one it is the difference between linear and quadratic.

## BFS or DFS, and when it starts to matter

This problem asks about reachability, not distance. Either traversal answers it. A recursive DFS here would be shorter than the BFS and exactly as correct, and you should feel free to write it that way.

The distinction becomes real when a problem asks for the *shortest* path, because BFS reaches every node by a shortest route and DFS reaches it by whatever route it wandered down first. Day 40 is where that starts to matter, and it is worth noticing now that the two traversals are interchangeable, so that it is obvious later when they stop being.

## One thing the code could do better

The loop drains the entire queue and only then checks the destination.

It could compare on each dequeue and return `true` the moment it arrives. For a false answer this changes nothing, since the whole component has to be explored to prove a negative. For a true answer on a large component it can stop far earlier.

The version here is slightly wasteful and noticeably easier to read, which is a defensible trade on an Easy and worth remembering when the graph is large.

## A note on the queue

The file defines its own `Queue` as a linked list of nodes, which looks like a lot of code for something the standard library would normally provide. Go does not provide it. The usual Go idiom is a slice with `q = q[1:]` to pop, which is what days 39 and 40 use.

Both work. The slice version is far shorter. The linked-list version does not hold on to the backing array of everything already popped, which matters only when the queue gets very long.

## Complexity

- **Time: O(V + E).** One pass over the edges to build the adjacency list, then a traversal that visits each node once and considers each edge from both of its ends.
- **Space: O(V + E).** The adjacency list dominates; the visited array and the queue are each O(V).

Full code and the step-by-step walkthrough:
[find_if_path_exists_in_graph](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1901_2000/find_if_path_exists_in_graph/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
