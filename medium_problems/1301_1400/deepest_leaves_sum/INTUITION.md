## Intuition

The obvious solution is two passes: find the maximum depth, then walk again adding up
everything at that depth. Both halves are things this batch has already built — Day 1 for
the depth, Day 14 for the per-level accumulation.

## Builds on

- [Day 1: Maximum Depth of Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/maximum_depth_of_binary_tree/) — finding the deepest level, which the two-pass version would do first and this one discovers as it goes
- [Day 14: Average of Levels in Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/average_of_levels_in_binary_tree/) — accumulating a number per level during a depth-first walk

One pass is enough, and the trick is being willing to throw work away.

Carry the depth down as a parameter, and keep two values: the deepest level seen so far and
the running sum for that level. At every leaf, compare:

- deeper than anything seen before? Then everything accumulated so far belonged to a
  shallower level and is now worthless. **Replace** the sum with this leaf's value and
  record the new depth.
- exactly equal to the deepest so far? **Add** to the sum.
- shallower? Ignore it entirely.

The replacement is the whole idea. You never need to know the final depth in advance,
because any discovery of a deeper leaf invalidates the previous answer and starts over.
Whatever survives to the end was accumulated at the true maximum depth.

Two structural details are worth noticing, because they're unusual for this batch.

The base case is a **leaf**, not nil. Almost every recursion in this series bottoms out at
`if node == nil`. This one checks `node.Left == nil && node.Right == nil` instead, and never
recurses into a nil child — the calls are guarded. That's consistent: since only leaves
contribute to the answer, a leaf is the meaningful terminal case, and nil never needs to be
represented.

The consequence is that `foo` dereferences `node` immediately without checking it. That's
safe only because the constraints guarantee at least one node, so `root` is never nil, and
every recursive call is guarded. Change the constraint to allow an empty tree and this
panics on the first line.

**Complexity:**
- Time: O(n) — every node is visited once, in one pass rather than two.
- Space: O(h) for the recursion stack. Nothing per-level is stored: two ints carry the
  entire answer, regardless of how wide the tree is.
