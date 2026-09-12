**365 Days of LeetCode Challenge — Day 42/365**
**Clone Graph** (Medium)
🔗 https://leetcode.com/problems/clone-graph/

Copying a binary tree is four lines. It works because a tree cannot cycle and every node has one parent. A graph gives up both guarantees.

Cycles: A's copy needs B's copy needs A's copy, forever. Shared nodes: a node reachable two ways gets copied twice, giving the right values in the wrong shape.

One map from original to copy fixes both.

```go
newNode, ok := visited[node.Val]
if ok {
	return newNode
}
newNode = new(Node)
newNode.Val = node.Val
visited[node.Val] = newNode

for i, nNode := range node.Neighbors {
	newNode.Neighbors[i] = clone(nNode, visited)
}
```

A plain visited SET would only fix the cycle. This maps to the copies, because the copies are what the reconstruction wires together.

The order is the trick: the node goes into the map BEFORE its neighbours are cloned, while it is still half built. Recurse first and a cycle returns to a node that is not in the map yet and clones it again.

One caveat before reusing this: the map is keyed by `Val`, which is safe only because this problem guarantees unique values. Key by `*Node` otherwise.

O(V + E) time, O(V) space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/clone_graph/SOLUTION.md
