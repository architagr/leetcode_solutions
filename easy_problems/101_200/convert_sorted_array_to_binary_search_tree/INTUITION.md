## Intuition

A height-balanced BST needs roughly equal numbers of elements on the left and right of
every node. Since the input array is already **sorted**, that balance falls out almost
for free: if you always pick the **middle element** of a sorted slice as a node's value,
everything to its left is smaller (goes in the left subtree) and everything to its
right is larger (goes in the right subtree) — and both halves are roughly the same
size.

So the approach is a straightforward divide-and-conquer:
- Pick the middle element of the current slice as the root of this subtree.
- Recurse on the left half of the slice to build the left subtree.
- Recurse on the right half of the slice to build the right subtree.
- An empty slice means "no subtree here" (`nil`).

Because the recursion always splits the array roughly in half at each step, the
resulting tree is automatically balanced — you never have to rebalance anything after
the fact.

**Complexity:**
- Time: O(n) — each element of `nums` becomes exactly one tree node, created once.
- Space: O(log n) for the recursion stack (the tree has depth ~log₂n), plus O(n) for
  the output tree itself.
