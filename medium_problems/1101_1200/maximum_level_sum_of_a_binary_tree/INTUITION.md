## Intuition

Two things have to be true at once here, and only one of them is about sums.

The first is the obvious part: add up each level, keep the biggest total. That's the same
per-level aggregation as Day 27, and the same nil-sentinel BFS as Day 30 — push a `nil`
behind the root, and every time it surfaces, a level has just finished and its running
total is complete.

## Builds on

- [Day 30: Binary Tree Zigzag Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_zigzag_level_order_traversal/) — the nil sentinel in the queue, used the same way to know when a level's total is final
- [Day 27: Average of Levels in Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/average_of_levels_in_binary_tree/) — aggregating one number per level rather than collecting the nodes
- [Day 29: Binary Tree Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_level_order_traversal/) — level order as the underlying shape, here without needing to keep the values

The second requirement is in the wording, and it's easy to skim: return the *smallest*
level whose sum is maximal. Ties go to the shallower level.

That's not handled with any extra code. It falls out of one character — the comparison is
`sum > maxSum`, strictly greater. A later level matching the current best doesn't beat it,
so `maxLevel` keeps the earlier value. Write `>=` instead and the function still returns a
level with the maximum sum, but the wrong one whenever there's a tie, and the tests that
catch it are the ones with repeated totals.

There's a second small decision doing real work: `maxSum` starts at `root.Val`, not at
zero. Node values can be negative, and a tree whose every level sums to something negative
would otherwise never beat a zero starting point, leaving `maxLevel` at its initial value
by luck rather than by comparison. Seeding from the root means the first real comparison is
against an actual level sum.

The level counter starts at 1 rather than 0, because the problem numbers levels from 1.
Worth noticing only because the previous few days in this series all indexed from 0, and
this is the kind of off-by-one that survives testing on symmetric examples.

**Complexity:**
- Time: O(n) — every node is pushed and popped once and contributes one addition.
- Space: O(w) for the queue, where w is the width of the widest level. Note that nothing
  stores the levels themselves: only a running sum, the best sum so far, and two counters.
