## Intuition

Yesterday's labels came straight from position: a node's label was where it sits in a
heap-style numbering, and moving to the parent was just halving. Here the tree has the same
shape, but every other row is numbered backwards. So the labels no longer tell you the
position directly, but they're one flip away from it.

Split it into two small conversions:

- **Label to position.** Row `L` holds labels `2^(L-1)` to `2^L - 1`. On an odd row,
  position from the left is `label - 2^(L-1)`. On an even row it's measured from the other
  end: `(2^(L-1) - 1) - (label - 2^(L-1))`.
- **Up one row.** Positions follow the ordinary tree, so the parent's position is the
  child's position halved. Convert that position back to a label for the parent's row,
  flipping again if that row is even.

Repeat until the root, filling the answer from the back.

## Builds on

- [Day 125: Find Elements in a Contaminated Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1201_1300/find_elements_in_a_contaminated_binary_tree/) — labels that come from a node's position, and a parent's position being its child's halved
- [Day 30: Binary Tree Zigzag Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_zigzag_level_order_traversal/) — the alternating left-to-right, right-to-left order of levels

There's a compact formula hiding in this. For a node on a row spanning `lo..hi`, the parent
label is `(lo + hi - label) / 2`: reflecting the label within its row undoes the zigzag,
and halving climbs one row. For 14 (row 8..15) that's `(8 + 15 - 14) / 2 = 4`, and for 4
(row 4..7) it's `(4 + 7 - 4) / 2 = 3`. The repo solution spells the two steps out
separately, which is longer but easier to check line by line.

`pow2` goes through `math.Pow` on floats. For exponents up to 20 that's exact; a bit shift,
`1 << x`, would avoid floats entirely.

**Complexity:**
- Time: O(log label): one step per row.
- Space: O(log label) for the path.
