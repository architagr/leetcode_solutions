## Intuition

The tricky part of this problem isn't finding leaves — it's knowing whether a leaf is
a **left** child or a **right** child, since only left leaves count. A node itself
can't tell you this about itself; only its **parent** knows which side it's on.

So the recursion needs to check "is my left child a leaf?" from the parent's
perspective, rather than each node trying to detect its own leaf-and-side status in
isolation. That's why the check `root.Left != nil && root.Left.Left == nil &&
root.Left.Right == nil` happens in the parent's call, looking one level down at its
left child — not inside a base case for the child itself.

The rest is a standard tree recursion: sum the left leaves found in the left subtree,
sum the left leaves found in the right subtree, add in the current node's left child's
value if that child qualifies as a left leaf, and return the total.

**Complexity:**
- Time: O(n) — every node is visited exactly once.
- Space: O(h) for the recursion stack, where h is the tree's height (O(log n) for a
  balanced tree, O(n) for a completely skewed one).
