## Intuition

The diameter of a binary tree is the length (in edges) of the longest path between
any two nodes. The tricky part is that this path does **not** have to pass through
the root — it can be entirely inside a left or right subtree.

The key insight: the longest path *through* any given node `root` is simply

```
height(root.Left) + height(root.Right)
```

— walk down to the deepest leaf on the left, up through `root`, then down to the
deepest leaf on the right. So if we compute the height of every node's left and
right subtree and take `left + right` at every node, the maximum of all those
values across the whole tree is the diameter. It doesn't matter which node ends up
being the "center" of the winning path — we just need to try every node as a
candidate center and keep the best one.

This turns into a single post-order DFS:

- At each node, first recurse into the left and right children to get their
  heights.
- Combine those heights to get a diameter candidate (`left + right`) and update a
  running maximum if it's bigger.
- Return `max(left, right) + 1` up to the caller — the height of the current
  subtree — so the parent can use it the same way.

Because the height computation and the diameter computation are the same
traversal, we get the answer in one pass instead of computing heights separately
for every node (which would be `O(n^2)`).

### Complexity

- **Time:** `O(n)` — every node is visited exactly once.
- **Space:** `O(h)` — the recursion stack depth is the height of the tree, `h`,
  which is `O(n)` in the worst case (a skewed tree) and `O(log n)` for a balanced
  one.
