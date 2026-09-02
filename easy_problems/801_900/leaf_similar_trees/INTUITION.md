## Intuition

"Leaf-similar" is really just asking: if you read off each tree's leaves from left to
right, do you get the same sequence? So the problem reduces to two much simpler steps —
extract the leaf sequence of each tree, then compare the two sequences directly. There's
no need to compare the trees' shapes at all; only the leaves, in left-to-right order,
matter.

Extracting a tree's leaf sequence in left-to-right order is exactly what a depth-first
traversal that always visits the left subtree before the right subtree gives you for
free: recurse left, recurse right, and whenever you land on a node with no children,
that's a leaf — record its value. Because the left subtree is always fully explored
before the right subtree, the leaves naturally come out in left-to-right order without
any extra bookkeeping.

Once both leaf sequences are collected into two slices, comparing them for
leaf-similarity is just comparing two lists: same length, and same value at every
index.

**Complexity:**
- Time: O(n + m) — every node of both trees is visited exactly once, where n and m are
  the node counts of `root1` and `root2`.
- Space: O(n + m) for the two leaf slices, plus O(h1 + h2) for the recursion stacks
  (h1, h2 being each tree's height).
