## Intuition

A mirrored tree is one where every node's left and right children have traded places. Not
just the root's: every node's. So the job is to visit every node once and swap its two
child pointers.

The recursive framing makes that feel smaller than it is. To invert a tree, invert the left
subtree, invert the right subtree, then swap them. Each subtree comes back already
mirrored, and one pointer swap at the top puts the two mirrored halves on the correct
sides.

## Builds on

- [Day 68: Merge Two Binary Trees](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/merge_two_binary_trees/) — the "recurse, then store the returned subtree back into the child pointer" shape this uses

The swap moves whole subtrees, not values. When the root swaps its children in Example 1,
the entire `7` subtree moves to the left with its own (already swapped) children still
attached. No node is created or copied; the tree is rearranged in place, and the function
returns the same root it was given.

Does the order matter? Here the swap happens after both recursive calls, which is
post-order. Swapping first and then recursing into the two (now swapped) children also
works, because each child gets inverted exactly once either way. What would break is
recursing into one child, swapping, and then recursing into "the other" child, which is now
the one you already did.

This is also the second of the two small functions that Symmetric Tree is built from in
this repo. That one mirrors half a tree and compares it with the other half.

**Complexity:**
- Time: O(n), each node visited once and one swap per node.
- Space: O(h) for the recursion stack.
