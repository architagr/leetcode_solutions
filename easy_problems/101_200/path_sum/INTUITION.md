## Intuition

The key constraint is **root-to-leaf**, not "any path" — a match only counts if it
runs all the way from the root down to a node with no children. That rules out the
simplest-looking shortcut: you can't just check "does any node along the way equal
`targetSum` so far?", because a match at an internal node (one with children) doesn't
count. The sum has to be evaluated exactly once, at the leaf, using everything
accumulated above it.

That naturally suggests carrying the running total *down* the tree as an extra
argument, rather than summing back *up* through return values: at each node, add its
value to the sum accumulated by its ancestors, then hand that running total to
whichever child(ren) exist. When a node has no children left to hand it to — i.e. it's
a leaf — that's the only point where the running total is compared against
`targetSum`.

The branching also has to be explicit about *why* a node is being treated as a leaf.
A node with only a left child (no right child) is not a leaf, even though
`root.Right == nil` — it still needs to recurse into its one real child rather than
being scored on the spot. So the check isn't just "is `Left` nil" or "is `Right` nil"
in isolation; it's "does this node have *any* child to descend into," and only when
the answer is no for both sides does the running total get compared to the target.

**Complexity:**
- Time: O(n) — every node is visited at most once.
- Space: O(h) for the recursion stack, where h is the tree's height (O(log n) for a
  balanced tree, O(n) for a completely skewed one).
