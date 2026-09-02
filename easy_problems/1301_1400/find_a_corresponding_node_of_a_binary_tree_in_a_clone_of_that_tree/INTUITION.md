## Intuition

`cloned` is an exact structural copy of `original` — same shape, same values, just
different node objects in memory. That guarantee is what makes this problem easy: since
the two trees are structurally identical, you can walk them **in lockstep**, one step
down `original` and the matching step down `clones` at the same time, and the two
pointers will always be looking at "the same" logical position in the tree.

So instead of searching `cloned` on its own (which would need some way to recognize
`target`'s position independently), the algorithm advances both pointers together with
a single recursive call. At each pair of nodes visited, it asks: *is this the node?* —
by comparing `clones.Val == target.Val`. The moment that's true, `clones` is standing
exactly where `target` stands in `original`, so it's returned directly as the answer.

If it's not a match yet, the search recurses left in both trees together
(`original.Left`, `clones.Left`); if that returns something, it's propagated up
immediately. Otherwise it recurses right in both trees together
(`original.Right`, `clones.Right`) and returns whatever that finds. Because `clones`
mirrors `original` node-for-node, `clones == nil` alone is enough as the base case —
wherever `original` runs out, `clones` runs out too.

This is a plain DFS/pre-order traversal; the only twist is that it's a **paired**
traversal over two trees at once, which is what lets it "search by structure" instead
of needing to compare node identities or maintain any extra bookkeeping.

**Complexity:**
- Time: O(n) — in the worst case (e.g. `target` is the last node visited, or the tree
  is degenerate) every node pair is visited once.
- Space: O(h) for the recursion stack, where h is the tree's height (O(log n) for a
  balanced tree, O(n) for a completely skewed one).
