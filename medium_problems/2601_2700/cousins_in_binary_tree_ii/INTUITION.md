## Intuition

A node's cousins are everything on its level except its own family: itself and its
sibling. So "the sum of my cousins" is "the sum of my level, minus me and my sibling".

That splits the work into two passes. First, a level-order traversal records each level's
total. Second, a walk from the top where each parent looks at its own children: their
combined value is the family's share, and each child's new value is the next level's total
minus that share. Both siblings get the same number, which makes sense: they have the same
cousins.

## Builds on

- [Day 28: Cousins in Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/901_1000/cousins_in_binary_tree/) — what a cousin is: same depth, different parent
- [Day 29: Binary Tree Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_level_order_traversal/) — visiting a tree level by level, which gives each level's sum here

The order inside the second pass matters. A parent must read its children's *original*
values to compute the family sum, and only then overwrite them. The code sums
`node.Left.Val + node.Right.Val` first and assigns afterwards. Recursing into a child
after its value has changed is safe, because that child only ever reads its own children,
which haven't been touched yet.

The root has no parent and no cousins, so it's simply set to 0.

**Complexity:**
- Time: O(n): two traversals.
- Space: O(n) for the queue and the per-level sums.
