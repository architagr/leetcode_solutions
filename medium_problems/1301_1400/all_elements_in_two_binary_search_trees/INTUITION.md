## Intuition

This is two problems from earlier in the series stuck together, and I think it's worth
seeing it that way rather than as a new one.

Read a BST in order and the values come out sorted. So each tree hands you a sorted list
for free. Two sorted lists into one sorted list is the merge step from merge sort, the same
two-cursor walk as merging two sorted linked lists.

## Builds on

- [Day 47: Minimum Absolute Difference in BST](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/minimum_absolute_difference_in_bst/) — an in-order walk of a BST visits the values in sorted order
- [Day 22: Merge Two Sorted Lists](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/merge_two_sorted_lists/) — the two-cursor merge: take the smaller head, advance that side, repeat

The obvious alternative is to dump both trees into one slice in any order and sort it. That
works and is O((n+m) log(n+m)). The BST property is the whole reason this problem exists,
and using it drops the sort: two linear walks plus one linear merge.

The merge loop runs while both lists have something left. The moment one runs out, the
other's remainder is already sorted and already bigger than everything taken so far, so it
gets copied across as is. Only one of the two tail loops ever does any work, and writing
both is simpler than working out which.

On ties the code takes from the second tree (`root1Inorder[i] < root2Inorder[j]` is false
for equal values). For plain integers that choice is invisible. It would matter if the
values carried other data and you needed a stable merge.

One trade-off worth knowing about. This builds both in-order slices in full before merging,
so it holds n + m values on top of the output. [Day 56's BST iterator](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_search_tree_iterator/)
yields one value at a time with an O(h) stack, and two of those could feed the merge
directly with only O(h1 + h2) extra memory. For this problem the output is n + m anyway,
so the simpler version costs nothing that matters.

**Complexity:**
- Time: O(n + m). Each in-order walk is linear, and the merge touches every value once.
- Space: O(n + m) for the two in-order slices, plus O(h) recursion stack per walk. The
  result slice is preallocated with capacity n + m, so the appends never reallocate.
