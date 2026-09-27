## Intuition

With a sorted *array*, this is Day 2: the middle element is the root, and each half builds a
subtree. The list version is harder only because a linked list has no index. You can't jump
to the middle; you have to walk to it.

This solution sidesteps that completely: walk the list once, copy the values into a slice,
and hand the slice to the Day 2 function. One O(n) walk, O(n) extra space, and the rest is
code that's already known to work.

## Builds on

- [Day 2: Convert Sorted Array to Binary Search Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/convert_sorted_array_to_binary_search_tree/) — the middle-as-root build this reuses once the list is a slice
- [Day 20: Middle of the Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/middle_of_the_linked_list/) — finding a list's middle without indexing, the fast/slow pointer alternative to copying

There are two well-known ways to avoid the copy, and they're worth knowing because this
problem is often asked specifically to see them:

- **Fast/slow pointers per call.** Find the middle of the current sublist the Day 20 way,
  make it the root, and recurse on the two halves. No extra array, but every level of
  recursion walks its sublists again: O(n log n) time, O(log n) space.
- **Build in order.** Count the list's length first. Then build the tree by index range
  recursively, left subtree first; each time a node is created, take the list's current head
  and advance it. Because an in-order build visits positions in sorted order, the list is
  consumed exactly in order. O(n) time and O(log n) space, with no copy.

The copy-then-build version is the one I'd write first. It's O(n) time, easy to verify, and
the only cost is memory.

**Complexity:**
- Time: O(n).
- Space: O(n) for the slice, plus O(log n) recursion.
