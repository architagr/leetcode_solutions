## Intuition

A binary search tree already tells you which way to go — that's the entire point of the
"search" property. At every node, comparing the target `val` against `root.Val` doesn't
just say "match or no match," it says "the value you want, if it exists, can only live
in one specific subtree."

- If `root.Val > val`, everything in the right subtree is even bigger than `root.Val`,
  so it's automatically too big too. The only place `val` could still be hiding is the
  **left** subtree.
- If `root.Val < val`, symmetrically, `val` can only be hiding in the **right** subtree.
- If neither is true, they're equal — you've found the node, and since the problem asks
  for the *subtree* rooted at that node (not just a boolean), returning the node itself
  is enough; its `Left`/`Right` pointers already carry the rest of that subtree along
  with it.

This is exactly what makes a BST search different from searching a plain binary tree:
there's no need to check both children and no need to backtrack. Each comparison
eliminates an entire half of the remaining tree, so the recursion always walks a single
root-to-node path — never branches, never revisits.

The one edge case is running out of tree: if `root` becomes `nil` before a match is
found, the value simply isn't there, and `nil` is the correct answer.

**Complexity:**
- Time: O(h), where h is the tree's height — each call moves one level down a single
  path, so at most one node per level is visited. That's O(log n) for a balanced BST,
  degrading to O(n) for a completely skewed one.
- Space: O(h) for the recursion stack, for the same reason.
