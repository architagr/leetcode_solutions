## Intuition

This problem is really two smaller problems stacked on top of each other:

1. **"Are two trees identical?"** — same shape, same values at every position. That's
   the classic Same Tree check.
2. **"Does `root` contain a node where, if you rooted a tree right there, it would be
   identical to `subRoot`?"** — that's just problem 1, tried at every possible anchor
   point in `root`.

So the approach is: walk every node of `root`, and at each one ask "if I treat this
node as the root of its own little tree, is that tree identical to `subRoot`?" As soon
as one node answers yes, the whole thing is a match. If no node ever answers yes, it
isn't.

A cheap pruning trick makes this practical: don't even bother running the full
structural comparison unless the current node's value already matches `subRoot`'s
value — two trees can't be identical if their roots don't match, so checking
`root.Val == subRoot.Val` first before doing the expensive full comparison avoids a lot
of wasted work.

**Complexity:**
- Time: O(m·n) in the worst case, where `m` is the number of nodes in `root` and `n` is
  the number of nodes in `subRoot` — for each of the `m` candidate anchor nodes, the
  full-equality check can touch up to `n` nodes of `subRoot`.
- Space: O(h1 + h2) for the recursion stack, where `h1` and `h2` are the heights of
  `root` and `subRoot` respectively.
