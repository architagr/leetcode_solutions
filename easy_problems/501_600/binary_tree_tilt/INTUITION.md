## Intuition

Computing a single node's tilt needs the sum of *all* values in its left subtree and
the sum of *all* values in its right subtree — not just its immediate children. A naive
approach would re-sum each subtree from scratch for every node, which is wasteful:
the sum of a subtree rooted lower down gets recomputed over and over as you climb back
up the tree.

The fix is to compute each subtree's sum exactly once, bottom-up, and hand it back to
the caller. That's a **postorder** traversal: visit the left subtree, visit the right
subtree, *then* do work at the current node — because the current node's tilt (and its
own contribution to its parent's subtree sum) depends on both children's totals already
being known.

So the recursive helper does two jobs at once on every call: it accumulates into a
running tilt total (a side effect, via a pointer to an int so every recursive call
shares the same accumulator), and it returns the sum of the subtree rooted at the
current node (so the parent can use it). Each node's tilt is simply the absolute
difference between what its left recursive call returned and what its right recursive
call returned.

**Complexity:**
- Time: O(n) — every node is visited exactly once, and all work at a node is O(1).
- Space: O(h) for the recursion stack, where h is the tree's height (O(log n) for a
  balanced tree, O(n) for a completely skewed one).
