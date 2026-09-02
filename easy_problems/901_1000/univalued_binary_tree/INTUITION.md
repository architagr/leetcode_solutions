## Intuition

A tree is uni-valued exactly when *every* parent-child edge connects two nodes with
the same value — if that holds everywhere, all values in the tree must be equal to the
root's value by transitivity, and if it fails anywhere, the tree isn't uni-valued.

So instead of collecting every value into a set and checking they're all equal, it's
enough to walk the tree and compare each node to its own parent as we go. A node
doesn't need to know the root's value directly — it just needs to know its parent's
value, which is available to the parent at the moment it looks at its child.

That gives a simple recursion: a subtree rooted at `node` is uni-valued if both its
children (when present) share `node`'s value, *and* the subtrees hanging off those
children are themselves uni-valued. An empty subtree is trivially uni-valued — there's
nothing to disagree with — so `nil` is the base case that returns `true`.

**Complexity:**
- Time: O(n) — every node is visited exactly once.
- Space: O(h) for the recursion stack, where h is the tree's height (O(log n) for a
  balanced tree, O(n) for a completely skewed one).
