# Clone Graph — solution walkthrough

From `main.go`:

```go
func cloneGraph(node *Node) *Node {
	visited := make(map[int]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[int]*Node) *Node {
	if node == nil {
		return node
	}
	newNode, ok := visited[node.Val]
	if ok {
		return newNode
	}
	newNode = new(Node)
	newNode.Val = node.Val
	visited[node.Val] = newNode

	if len(node.Neighbors) > 0 {
		newNode.Neighbors = make([]*Node, len(node.Neighbors))
		for i, nNode := range node.Neighbors {
			newNode.Neighbors[i] = clone(nNode, visited)
		}
	}

	return newNode
}
```

## Why the tree version fails

Copying a binary tree is: make a node, recurse into left, recurse into right. It works because a tree cannot cycle and every node has exactly one parent.

A graph gives up both guarantees, and each one breaks the naive copy in its own way.

**Cycles.** If A and B are neighbours, A's copy needs B's copy, which needs A's copy. Recursion never bottoms out.

**Shared nodes.** Even with no cycle, a node reachable by two routes gets copied twice, producing a graph with the right values and the wrong shape: two separate copies where the original had one shared node.

## One map solves both

```go
visited := make(map[int]*Node)
```

A map from original value to the copy of that node.

![Step 1](images/walkthrough-1.png)

```go
newNode, ok := visited[node.Val]
if ok {
	return newNode
}
```

Before cloning anything, check whether this node has been cloned already. If so, hand back the existing copy.

That check does two jobs. It terminates cycles, because the second arrival at a node returns instead of descending. And it preserves sharing, because every original maps to exactly one copy, so two neighbours referring to the same node end up referring to the same copy.

A plain `visited` set would only do the first job. The map has to hold the copies, because the copies are what the reconstruction needs to wire together.

## The order is the trick

```go
newNode = new(Node)
newNode.Val = node.Val
visited[node.Val] = newNode

if len(node.Neighbors) > 0 {
	newNode.Neighbors = make([]*Node, len(node.Neighbors))
	for i, nNode := range node.Neighbors {
		newNode.Neighbors[i] = clone(nNode, visited)
	}
}
```

![Step 2](images/walkthrough-2.png)

The new node is recorded in the map **before** its neighbours are cloned. At that moment it is a node with a value and no edges: half built.

That is deliberate and it is the whole termination argument. Recurse first and record afterwards, and a cycle comes back to a node that is not in the map yet, clones it again, and recurses again.

The map has to contain the unfinished node for the cycle to close.

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

When the walk reaches node 1 for the second time, from node 4, the lookup succeeds and the existing copy comes straight back. The recursion stops there and the cycle is wired up correctly.

This is day 37's `(*grid)[i][j] = '2'` in different clothes. Both mark a node as in-progress before descending, and both work because a partially finished node is enough to stop a re-entrant call.

## Keying the map by `Val`

```go
visited[node.Val]
```

The map is keyed by the node's value, not by the node pointer.

That is safe here only because the problem guarantees `Node.val` is unique for each node. On a graph where two distinct nodes could hold the same value, this would collapse them into one and silently produce the wrong graph.

Keying by `*Node` instead works regardless, since pointers are comparable in Go and identity is exactly what is being tracked. If you are adapting this code to a problem without that uniqueness guarantee, change the key first.

## The nil check

```go
if node == nil {
	return node
}
```

Covers the empty-graph case, where the whole input is nil. Since the graph is stated to be connected, this cannot happen partway through a traversal; it is only ever the top-level call.

## The neighbour slice is sized up front

```go
newNode.Neighbors = make([]*Node, len(node.Neighbors))
for i, nNode := range node.Neighbors {
	newNode.Neighbors[i] = clone(nNode, visited)
}
```

Allocated at full length and filled by index, rather than appended to. Same result; this avoids the regrowth `append` would do, and it keeps neighbour order identical to the original, which matters because the expected output compares adjacency lists positionally.

## Complexity

- **Time: O(V + E).** Every node is cloned exactly once. Every edge is traversed once from each endpoint, and the second traversal ends immediately at the map lookup.
- **Space: O(V)** for the map, plus a recursion stack that can reach O(V) on a path-shaped graph.

## Test

`main_test.go` builds the four-node cycle from the problem statement and checks the clone has the same structure while sharing no nodes with the original. The shared-node check is the one that fails if the map is dropped.
