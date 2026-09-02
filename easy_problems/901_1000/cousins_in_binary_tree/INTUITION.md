## Intuition

Two nodes are cousins if they're at the **same depth** but have **different
parents**. That's two separate facts to track per node — its depth, and some way to
identify its parent — so the natural approach is: find both facts for `x`, find both
facts for `y`, then compare.

The tricky bit is "identify its parent." We don't have parent pointers in a standard
binary tree, and we don't want to build a whole `map[*TreeNode]*TreeNode` just to
answer one comparison. But the problem guarantees every node has a **unique value**, so
a node's own `Val` works as a stand-in for its identity — instead of returning a
pointer to the parent, we can return the parent's `Val` and compare those.

That's exactly what `foo` does: given a node to start searching from and a target
value, it walks the tree (checking each node's children before descending into them,
so it can hand back the *child's value* as `parent` the moment it finds the target)
and returns `(depth, parent, found)` for that target. `isCousins` just calls `foo`
twice — once for `x`, once for `y` — and checks `depthX == depthY && parentX !=
parentY`.

Using `0` as the "no parent" sentinel works because node values are constrained to be
`>= 1`, so `0` can never collide with a real parent value — it only ever shows up if
the target turns out to be the root itself (which has no parent).

**Complexity:**
- Time: O(n) — `foo` is called twice, and each call visits at most every node once.
- Space: O(h) for the recursion stack per call, where h is the tree's height (O(log n)
  for a balanced tree, O(n) for a completely skewed one).
