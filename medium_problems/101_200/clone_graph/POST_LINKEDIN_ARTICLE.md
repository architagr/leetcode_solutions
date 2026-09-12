---
meta_title: "Cloning a graph means recording the copy before you finish it"
meta_description: "A tree copy is four lines. A graph has cycles and shared nodes, and one map keyed to the copies fixes both - as long as you record before recursing."
---

![Day 42](HERO.png)

## 365 Days of LeetCode Challenge — Day 42/365

**[133. Clone Graph](https://leetcode.com/problems/clone-graph/)** (Medium)

Deep copy a connected undirected graph. New nodes, same structure, nothing shared with the original.

## Why the tree version does not port

Copying a binary tree is four lines: make a node, recurse left, recurse right. It works because a tree cannot cycle and every node has exactly one parent.

A graph gives up both guarantees, and each one breaks the naive copy differently.

**Cycles.** If A and B are neighbours, A's copy needs B's copy, which needs A's copy. The recursion never bottoms out.

**Shared nodes.** Even with no cycle, a node reachable by two routes gets copied twice. You end up with a graph that has all the right values and the wrong shape: two separate copies where the original had one shared node.

## One map fixes both

```go
visited := make(map[int]*Node)
```

A map from each original node to its copy. Before cloning anything, look it up. If the copy exists, return it.

That single check does two jobs:

- **Cycles terminate**, because the second arrival at a node returns instead of descending.
- **Sharing is preserved**, because every original maps to exactly one copy, so two neighbours pointing at the same node produce two references to the same copy.

The second is the one worth dwelling on. A plain `visited` set would only do the first job. This has to map to the **copies**, because the copies are what the reconstruction needs to wire together.

## The order is the trick

```go
newNode = new(Node)
newNode.Val = node.Val
visited[node.Val] = newNode      // record it...

for i, nNode := range node.Neighbors {
	newNode.Neighbors[i] = clone(nNode, visited)   // ...then recurse
}
```

The copy goes into the map **before** its neighbours are cloned. At that instant it is a node with a value and no edges. Half built.

That is the whole termination argument. Recurse first and record afterwards, and a cycle returns to a node not yet in the map, clones it again, and recurses again.

The map has to contain the unfinished node for the cycle to close.

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

Reaching node 1 for the second time, from node 4, the lookup succeeds and the existing copy comes straight back.

This is day 37's `(*grid)[i][j] = '2'` wearing different clothes. Both mark a node as in-progress before descending, and both work because a partially finished node is enough to stop a re-entrant call.

## One line to check before you reuse this

```go
visited[node.Val]
```

The map is keyed by the node's **value**, not by the node pointer.

That is safe here only because the problem guarantees values are unique. On a graph where two distinct nodes could hold the same value, this collapses them into one and silently produces the wrong graph.

Keying by `*Node` works regardless — pointers are comparable in Go, and identity is exactly what is being tracked. If you adapt this code to a problem without that guarantee, change the key first.

## Complexity

- **Time: O(V + E)**. Every node cloned once; every edge followed once from each end, with the second traversal stopping at the map lookup.
- **Space: O(V)** for the map, plus recursion depth up to O(V).

## Builds on

- [Day 37: Number of Islands](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/number_of_islands/) — recursive DFS where marking a node before recursing is what stops it looping back

Full code and the step-by-step walkthrough:
[clone_graph](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/clone_graph/SOLUTION.md)

#DSA #LeetCode #Golang #Graphs #DFS #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
