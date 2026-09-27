## Intuition

This starts the third binary-tree arc with an easy win. Bottom-up level order is ordinary
level order, reversed. The levels themselves are the same, and so is the left-to-right
order inside each level; only the order of the levels flips.

So the solution does exactly that: a breadth-first traversal that groups nodes by level
(top to bottom), then one pass to reverse the list of levels.

## Builds on

- [Day 29: Binary Tree Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_level_order_traversal/) — grouping values by level, top to bottom
- [Day 30: Binary Tree Zigzag Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_zigzag_level_order_traversal/) — the BFS with a nil marker at the end of each level, which this reuses

The tempting alternative is to build the answer bottom-up directly by inserting each new
level at the *front* of the result. In Go that means copying the whole slice every time,
which adds up to O(levels²) copies. Appending and reversing once at the end is O(levels)
for the reversal, and it keeps the traversal code identical to the top-down version.

The traversal uses a `nil` marker in the queue to know where one level ends. When a
`nil` is popped and the queue still has nodes, those nodes are the next level: start a new
inner slice and push another marker. When it's popped with the queue empty, the traversal
is done.

**Complexity:**
- Time: O(n) for the traversal, plus O(levels) to reverse.
- Space: O(n) for the queue and the result.
