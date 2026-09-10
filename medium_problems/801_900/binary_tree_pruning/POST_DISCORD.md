**365 Days of LeetCode Challenge — Day 44/365**
**Binary Tree Pruning** (Medium)
🔗 https://leetcode.com/problems/binary-tree-pruning/

A node survives if its subtree contains a `1` anywhere. As a recursion that's almost the definition read aloud: keep this node if it's a 1, or if either subtree kept anything.

Post-order, because a node can't answer that on the way down — it doesn't know what's beneath it yet.

The part people get wrong isn't the condition, it's *who does the deleting*. A node can't remove itself: no parent pointer, and nulling a local changes nothing the caller sees. So the child reports, and the parent clears the pointer it holds.

```go
right := parse(node.Right)
if !right {
	node.Right = nil
}
...
return curr || right || left
```

Which leaves the root, because the root has no parent — hence the wrapper. Without it, a tree of all zeroes comes back intact instead of empty.

Worth noting a 0 can survive: node 0 with a 1 below it stays, which is why "prune every 0" is the wrong rule.

O(n) time, O(h) space.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/801_900/binary_tree_pruning/SOLUTION.md
