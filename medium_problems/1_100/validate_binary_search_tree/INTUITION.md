## Intuition

The definition in the problem is recursive and local-sounding: left subtree smaller, right
subtree larger, and both subtrees also valid. Implement it literally and it's easy to write
the classic wrong answer — check each node against its two children only. That passes
`[5,1,4,null,null,3,6]`, where `5 > 1` and `5 < 4` is false at the root, but also passes
trees where a node deep in a left subtree is larger than an ancestor several levels up. A
BST constrains a node against every ancestor, not just its parent.

The way out is a property this batch has already used twice.

## Builds on

- [Day 30: Minimum Absolute Difference in BST](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/minimum_absolute_difference_in_bst/) — in-order traversal of a BST yields values in sorted order, which that problem relied on to compare only adjacent values
- [Day 29: Find Mode in Binary Search Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/find_mode_in_binary_search_tree/) — the same property, noted there as the follow-up's constant-space route

In-order traversal of a BST yields its values in ascending order. That's not a side effect,
it's equivalent to the definition — and equivalences run both ways. If in-order comes out
sorted, the tree is a BST; if it doesn't, it isn't.

So validation becomes: traverse in-order into a slice, then check the slice is strictly
increasing. No ancestor bounds to thread down, no min/max pairs, and the "deep node
violating a distant ancestor" case is handled without ever being thought about — that node
simply lands in the wrong place in the sorted order.

Strictly increasing, not merely non-decreasing. The problem says strictly less and strictly
greater, so duplicates are invalid, which is why the check is `arr[i] >= arr[i+1]` rather
than `>`.

The file also keeps a second approach, `IsValidBstApproch1`, which is the honest but slow
version of the local definition: for each node, scan its entire left subtree for a maximum
and its entire right subtree for a minimum. That is correct — it compares against whole
subtrees rather than immediate children — but it re-scans subtrees at every node, making it
O(n^2) on a skewed tree. It's kept as the contrast rather than as the answer.

**Complexity:**
- Time: O(n) — one traversal, then one linear scan.
- Space: O(n) for the collected slice plus O(h) for the recursion stack. The textbook
  improvement is to compare each value against the previous one during the traversal
  instead of materialising the slice, which drops it to O(h). Day 30 already does exactly
  that, so the pieces for it are on the shelf.
