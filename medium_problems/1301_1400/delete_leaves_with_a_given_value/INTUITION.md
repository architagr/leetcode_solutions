## Intuition

The statement describes a loop: delete the target leaves, see who became a leaf, delete
again, repeat until nothing changes. Taken literally that's several passes over the tree.

Post-order does it in one. If a node deals with its children first, then by the time it
looks at itself, every deletion below it has already happened. Whether it's a leaf *now*
is exactly the question the repeated passes were trying to answer. So each node asks once,
after its children, and the cascade falls out of the order.

## Builds on

- [Day 73: Binary Tree Pruning](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/801_900/binary_tree_pruning/) — the same bottom-up deletion, where a subtree's fate is decided only after its children report back

The piece to notice in this code is that `isLeaf(root)` is called twice. The first call,
before the recursion, decides whether there are children to visit. The second, after, sees
the children as they are *after* deletion. For the `2` on the left of Example 1, the first
check says "not a leaf, it has a child"; the recursion deletes that child; the second check
says "leaf now", and the node goes too.

The other thing I like compared to Day 73: this returns `*TreeNode` and the parent stores
the result, `root.Left = removeLeafNodes(root.Left, target)`. Returning nil is how a node
deletes itself. That also covers the root with no extra code. If the whole tree collapses,
the top call returns nil and that's the answer. Day 73 returned a boolean and needed a
wrapper function just to handle the root; this shape, the same one Delete Node in a BST
used, doesn't.

The recursion only descends into non-nil children, and the function is never called with a
nil root thanks to the constraints. Called on nil directly, `isLeaf(nil)` is false and the
next line would dereference nil, so this relies on that guarantee.

**Complexity:**
- Time: O(n), each node visited once.
- Space: O(h) for the recursion stack.
