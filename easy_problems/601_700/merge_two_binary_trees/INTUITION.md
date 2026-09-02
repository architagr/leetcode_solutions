## Intuition

This is a tree recursion where the "combine" step is trivial (just add two numbers) —
the real work is deciding what to do when one side of the merge is missing.

At every pair of positions `(root1, root2)` in the two trees, there are exactly three
cases:
- Both nodes exist → the merged node's value is their sum, and both children still
  need to be merged the same way, one level down.
- Only `root1` exists (`root2` is `nil` at this position) → there's nothing to add,
  so the merged subtree from here on is just whatever `root1` already is.
- Only `root2` exists (symmetric case) → the merged subtree is just `root2`.

Because a `nil` subtree simply "loses" to whatever the other tree has at that
position, you never need to build brand-new nodes for the parts where only one tree
has content — you can hand back the existing subtree from whichever tree is non-nil
and let it be reused as-is inside the merged result. New nodes only get allocated at
positions where *both* trees are actually present and their values need summing.

The recursion naturally proceeds in **preorder**: settle the current node's value
first (so the merged parent exists before its children are attached), then recurse
into the left pair and the right pair to build the merged subtrees, and finally wire
those results in as `Left`/`Right`.

**Complexity:**
- Time: O(min(m, n)) — the recursion only descends as far as both trees still have
  nodes; once one side runs out (`nil`), that branch returns immediately without
  visiting further, whatever the other tree's remaining size below there.
- Space: O(min(m, n)) for the recursion stack in the worst case (skewed trees), same
  bound as the time complexity for the same reason.
