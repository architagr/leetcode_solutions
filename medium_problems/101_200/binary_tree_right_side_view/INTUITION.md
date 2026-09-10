## Intuition

The picture the problem paints — standing to the right, seeing what isn't hidden — has a
plainer description: from each level, you see exactly one node, the rightmost one.

So the answer has one entry per level, top to bottom. That framing makes it sound like a
BFS problem, and the queue-based version is perfectly good: walk each level, keep the last
node you popped.

## Builds on

- [Day 15: Binary Tree Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_level_order_traversal/) — the same depth-as-an-index idea, one slot per level, reached by a depth-first walk rather than a queue

This solution does it with a depth-first walk instead, and gets there with two small
changes to the Day 15 shape.

The first is what gets stored. Day 15 accumulated every node at a depth; here only one
node per depth is wanted, so the append is guarded: `if len(arr) == level`. Since `arr`
holds one entry per level filled so far, `len(arr)` is the next level that hasn't been
recorded yet. The condition therefore means "this is the first node I've reached at this
depth" — and only that node gets appended. Every later arrival at the same depth finds
`len(arr)` already past it and is skipped.

The second change is the one that makes it correct, and it's a single line:

```go
arr = findRight(head.Right, arr, level+1)
arr = findRight(head.Left, arr, level+1)
```

Right before left. Reverse the usual order, and the first node reached at any depth is the
rightmost one at that depth. The guard stores first arrivals; the traversal order decides
which arrival is first. Neither piece does anything on its own.

Swap those two lines back and the same function computes the left side view, which is a
nice way to check you've understood why it works.

There's one subtlety worth stating because it's easy to assume it's wrong: the rightmost
node at a level is not always in the root's right subtree. In example 2 the deepest
visible node hangs off the left branch, because the right branch simply ran out first. The
traversal handles that without a special case — the right subtree gets first refusal at
every depth, and where it has nothing, the left subtree's node is the first arrival and
gets recorded.

**Complexity:**
- Time: O(n) — every node is visited once and does constant work.
- Space: O(h) for the recursion stack, where h is the tree's height, plus O(h) for the
  output, which holds one value per level. A BFS solution would instead peak at the width
  of the widest level.
