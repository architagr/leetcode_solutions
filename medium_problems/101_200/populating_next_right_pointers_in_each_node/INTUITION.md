## Intuition

Every node needs a pointer to the node on its right at the same level, and the rightmost
node of each level needs nil. "Same level" makes it a BFS problem, and knowing where a
level ends makes it the nil-sentinel BFS from Day 16.

## Builds on

- [Day 16: Binary Tree Zigzag Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_zigzag_level_order_traversal/) — the nil sentinel in the queue, marking where one level stops and the next begins
- [Day 15: Binary Tree Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_level_order_traversal/) — level order as the underlying shape

The obvious way to write it is to walk each level left to right, remembering the previous
node, and set `prev.Next = current`. That works.

This implementation does something neater. It pushes children **right before left**, so the
BFS visits every level from right to left. And once the walk runs backwards, the node to
your right is simply the node visited just before you:

```go
current.Next = prev
prev = current
```

No lookahead, no writing into a node you've already passed. The assignment reads exactly
like the requirement — "my next is the one before me" — because the traversal was reversed
to make that true.

The level boundary falls out too. When the sentinel is popped, `prev` is set to the
sentinel itself, which is nil. So the first node of the next level — the *rightmost* one,
since the walk is backwards — gets `Next = nil`, which is precisely what the rightmost node
of a level needs. The rule that would otherwise be a special case is just the general rule
applied at a boundary.

Worth noting what this solution doesn't use: the tree is perfect, and it never relies on
that. A perfect tree allows the well-known O(1)-space answer, where you walk each level
using the `Next` pointers you established on the level above, threading the level below as
you go, and never allocate a queue at all. This one is the general BFS, which would work
unchanged on the follow-up problem (117) where the tree isn't perfect.

**Complexity:**
- Time: O(n) — every node is pushed and popped once.
- Space: O(w) for the queue, where w is the width of the widest level. In a perfect tree
  that's about n/2, which is what the O(1) approach exists to avoid.
