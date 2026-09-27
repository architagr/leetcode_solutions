## Intuition

Two problems from this series, run back to back. Yesterday, an in-order walk flattened a BST
into a sorted slice. On Day 2, a sorted slice became a height-balanced BST by making the
middle value the root and recursing on each half.

The values are what matter; the current shape is exactly the thing we want to discard. So:
flatten in order (sorted, shape gone), then build from the middle (balanced by
construction). Each recursive call splits its slice into two halves that differ in size by
at most one, which is what keeps every node's two subtrees within one level of each other.

## Builds on

- [Day 2: Convert Sorted Array to Binary Search Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/convert_sorted_array_to_binary_search_tree/) — building a height-balanced BST by making the middle value the root, recursively
- [Day 127: Kth Smallest Element in a BST](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/kth_smallest_element_in_a_bst/) — flattening a BST into a sorted slice with an in-order walk

This builds brand-new nodes rather than rearranging the old ones, so it uses O(n) extra
memory for the slice and the new tree. Rebalancing in place is possible (the
Day-Stout-Warren algorithm first straightens the tree into a chain with rotations, then
rotates it into balance), but it's much fiddlier for the same O(n) time. The two-step version
is what I'd write unless memory were tight.

"Balanced" has more than one valid answer. With `mid = len/2`, a slice of even length puts
the upper middle at the root, so Example 1 comes out as `[3,2,4,1]` rather than the
`[2,1,3,null,null,null,4]` shown; both are accepted.

**Complexity:**
- Time: O(n).
- Space: O(n) for the slice and the new nodes, plus O(log n) recursion when building.
