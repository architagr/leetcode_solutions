## Intuition

A node survives if its own subtree contains a `1` anywhere. Written as a recursion, that's
almost the definition read aloud: this node keeps its place if it is a `1`, or if either
of its subtrees keeps anything.

The direction matters. A node cannot answer that question on the way down — it doesn't yet
know what's beneath it. So this is post-order: both children report first, and the node
combines their answers with its own value.

## Builds on

- [Day 9: Binary Tree Postorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_postorder_traversal/) — the children-before-parent order that makes a bottom-up answer possible
- [Day 69: Evaluate Boolean Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/2301_2400/evaluate_boolean_binary_tree/) — the same shape, where each node returned a boolean built from its children's booleans

The part that's easy to get wrong is *who does the deleting*. A node can't remove itself —
it has no reference to its own parent, and nulling a local variable changes nothing the
caller can see. So the pruning is done by the parent, immediately after each recursive
call returns:

```go
right := parse(node.Right)
if !right {
    node.Right = nil
}
```

The child reports "nothing worth keeping down here", and the parent is the one holding the
pointer that has to be cleared.

Which leaves the root, because the root has no parent. That's why `PruneTree` exists as a
wrapper at all: it calls `parse`, and if the whole tree reports back false, it sets `root`
to nil itself. Without that, a tree of all zeroes would come back intact instead of empty.

The nil base case returning `false` is the other quiet piece. An absent subtree contains no
`1`, so it contributes nothing to the `||` — and because it's already nil, the parent's
`node.Right = nil` is a harmless no-op.

One stylistic note: this returns a boolean and mutates the tree as a side effect. The more
common formulation returns `*TreeNode` and has the caller reassign — `root.Left =
prune(root.Left)` — which is the shape Day 56 used for deletion. Both work. The reassigning
version needs no wrapper, since returning nil for the root handles the whole-tree case
naturally.

**Complexity:**
- Time: O(n) — every node is visited once.
- Space: O(h) for the recursion stack, where h is the tree's height.
