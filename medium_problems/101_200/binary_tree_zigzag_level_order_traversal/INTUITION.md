## Intuition

Yesterday's level order came out of a depth-first walk, with no queue and no idea when a
level ended. That worked because grouping by depth doesn't need level boundaries — a node
appends into its own slot whenever it happens to arrive.

Zigzag needs the boundary. To decide whether a level gets reversed you have to know the
level is finished, and "finished" is exactly the thing a DFS never learns. So this one goes
back to BFS.

## Builds on

- [Day 15: Binary Tree Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_level_order_traversal/) — the same output grouped the same way, solved without level boundaries, which is what makes the contrast here worth drawing
- [Day 1: Maximum Depth of Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/maximum_depth_of_binary_tree/) — the nil sentinel pushed into the queue to mark where a level ends

The level boundary comes from the same trick as Day 1: push a `nil` into the queue right
behind the root. Everything ahead of that `nil` is the current level. When the `nil`
surfaces, the level just drained, and a fresh `nil` goes on the back to close the next one
— but only if the queue still holds nodes, otherwise you'd loop forever on a sentinel with
nothing after it.

With the boundary in hand, the zigzag is almost an afterthought. Collect each level left
to right as usual, and reverse every second one before it goes into the result. A
`leftToRight` flag flips at each boundary and decides.

Reversing after collection is one of three ways to do this, and worth knowing the others
exist: you could push children in alternating order so the queue itself delivers each
level pre-zigzagged, or prepend rather than append when building the level. Reversing is
the one that keeps the traversal and the zigzag as separate ideas, which is why it reads
easily — the BFS half is identical to a non-zigzag BFS, and the alternation is four lines
bolted on top.

One detail that isn't decoration: the level is copied into a fresh slice before being
appended to the result. `arr` is reused across levels, and `reverseArr` mutates in place,
so without the copy every level in the result would alias the same backing array and the
answer would come out as several copies of the last level.

**Complexity:**
- Time: O(n) — every node enters and leaves the queue once, and the reversals total O(n)
  across all levels since each element is swapped at most once.
- Space: O(w) for the queue, where w is the width of the widest level, plus O(n) for the
  output. That width, rather than the tree's height, is the real memory cost of BFS.
