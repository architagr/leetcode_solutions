## Intuition

Every tree problem so far has taken a tree as input. This one produces trees, all of them, and
the recursion that builds them is the same recursion that defines them.

A full binary tree is either a single node, or a root with two full binary subtrees. The root
uses one node, so the other `n - 1` are shared out between the two subtrees. A full tree always
has an odd number of nodes (a root plus two odd subtrees), so the splits to try are `1` and
`n - 2`, `3` and `n - 4`, and so on. For each split, every left shape can be paired with every
right shape. Even `n` gives nothing at all.

## Builds on

This one introduces something new to the series: memoising a recursion. The same sizes come up
again and again. Building trees of 7 nodes asks for every tree of 5 nodes twice (once as a left
subtree, once as a right), and each of those asks for trees of 3. `memo[n]` stores every tree of
`n` nodes the first time they're built, so each size is enumerated once.

The number of trees grows fast: 1, 1, 2, 5, 14, 42... for 1, 3, 5, 7, 9, 11 nodes. Those are the
Catalan numbers, which count many kinds of "split into a left part and a right part" structures.
For `n = 19`, the largest odd input, it's 4862 trees.

When pairing shapes, the code copies each left and right subtree with `copyBinaryTree` rather than
pointing at the cached ones. LeetCode would accept shared subtrees, but then every tree in the
answer would share nodes with others, and changing one would silently change the rest. Copying
costs time and memory; it buys trees that are genuinely independent.

**Complexity:** proportional to the output: roughly the number of trees times their size, since
each tree is built (and copied) once.
