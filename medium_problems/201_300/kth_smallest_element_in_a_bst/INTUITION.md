## Intuition

Read a BST in order and the values come out sorted, so the `k`th smallest value is simply
the `k`th value the in-order walk produces. This solution collects the whole walk into a
slice and returns `arr[k-1]`.

## Builds on

- [Day 47: Minimum Absolute Difference in BST](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/minimum_absolute_difference_in_bst/) — an in-order walk of a BST visits the values in sorted order
- [Day 57: All Elements in Two Binary Search Trees](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1301_1400/all_elements_in_two_binary_search_trees/) — the same inOrder helper that threads one slice through the recursion

It's correct and O(n). The obvious improvement is to stop as soon as the `k`th value has
been seen: count down `k` as nodes are visited and return when it reaches zero. For small
`k` on a big tree that skips almost everything; the work becomes O(h + k) instead of O(n),
and no slice is needed. I checked that version against this one on a few hundred random
BSTs.

The follow-up asks what to do if the tree changes often and the query runs often. Then
walking at all is too slow. The standard answer is to store, in every node, the size of its
left subtree. At each node, if `k` is at most that size the answer is on the left; if it's
one more, it's this node; otherwise go right with `k` reduced. That's O(h) per query, and
the sizes can be kept up to date during inserts and deletes.

**Complexity (as written):**
- Time: O(n).
- Space: O(n) for the slice, plus O(h) recursion.
