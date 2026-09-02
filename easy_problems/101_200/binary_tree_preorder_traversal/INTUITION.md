## Intuition

"Preorder" just names the order in which the three things at each node get visited:
the node **itself** first, then its **left** subtree, then its **right** subtree
(root → left → right). That ordering is the entire problem — there's no searching,
comparing, or bookkeeping beyond "visit this node, then recurse left, then recurse
right."

The natural way to express that is recursion that mirrors the definition directly: a
helper function visits a node by appending its value, then calls itself on `Left`,
then calls itself on `Right`. The recursion's call stack does the traversal's
bookkeeping for free — there's no need for an explicit stack, since the position in
the recursive calls already encodes "where am I in the tree, and what's left to
visit."

The one wrinkle in Go specifically is that `append` can reallocate the underlying
array, so the slice being built has to be threaded through as both an argument and a
return value — each recursive call returns the (possibly grown) slice so the caller
picks up wherever the callee left off, rather than each call silently mutating a
slice the caller no longer has a valid reference to.

The base case is the empty subtree (`nil`): visiting nothing contributes nothing, so
the accumulated slice is simply handed back unchanged.

**Complexity:**
- Time: O(n) — every node is visited exactly once, and each visit does O(1) work
  (one append).
- Space: O(h) for the recursion stack, where h is the tree's height (O(log n) for a
  balanced tree, O(n) for a completely skewed one), plus O(n) for the output slice
  itself.
