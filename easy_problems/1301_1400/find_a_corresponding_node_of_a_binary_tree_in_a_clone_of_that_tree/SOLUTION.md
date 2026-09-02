## Solution walkthrough

The implementation is `getTargetCopy(original, clones, target *TreeNode) *TreeNode` in
`main.go`. It walks `original` and `clones` **in lockstep** — one recursive call
advances both pointers together — so the position that matches `target` in `original`
is, by construction, the same position `clones` is standing on in the cloned tree.

![Example 1](images/1.png "Example1")

We'll trace it on the example above, `original = [7,4,3,null,null,6,19]`,
`target = 3`.

1. **Base case.** `if clones == nil { return nil }` — since `clones` mirrors `original`
   node-for-node, `clones` running out of nodes is enough to know the search has bottomed
   out along this path; there's no need to check `original == nil` separately.

2. **Check if this is the target.** `if clones.Val == target.Val { return clones }` — the
   pair currently being visited is compared by value. Because node values are unique
   (per the constraints) and the two trees are structurally identical, `clones` standing
   at this position means it corresponds exactly to `target`'s position in `original`.

3. **Recurse left, first.** `n := getTargetCopy(original.Left, clones.Left, target)` —
   both pointers step to their respective left children in a single call. If that
   returns a non-nil node, it's propagated straight up: `if n != nil { return n }`.

4. **Otherwise recurse right and return whatever it finds.**
   `return getTargetCopy(original.Right, clones.Right, target)` — the right subtree is
   only searched once the left subtree is confirmed empty of the target, and its result
   (found or `nil`) becomes this call's result.

Walking it through `original = [7,4,3,null,null,6,19]`, `target = 3` (expected: the `3`
node from `clones`):

- `getTargetCopy(7, 7, target)`: `clones.Val(7) != target.Val(3)` → recurse left first,
  into `(4, 4)`.

  ![Step 1: at the root pair (7,7), values don't match, recurse left into (4,4)](images/walkthrough-1.svg)

- `getTargetCopy(4, 4, target)`: `clones.Val(4) != target.Val(3)`. `4` is a leaf in both
  trees, so recursing left goes to `(nil, nil)` → returns `nil` immediately, and
  recursing right also goes to `(nil, nil)` → returns `nil`. This call returns `nil`
  — the target isn't anywhere under `4`.

  ![Step 2: at pair (4,4), values don't match and both children are nil, so this branch is a dead end — returns nil](images/walkthrough-2.svg)

- Back in `getTargetCopy(7, 7, target)`: the left recursion returned `nil`, so it falls
  through to `return getTargetCopy(original.Right, clones.Right, target)`, i.e.
  `(3, 3)`.

- `getTargetCopy(3, 3, target)`: `clones.Val(3) == target.Val(3)` → match! Returns
  `clones` directly — the `3` node from the *cloned* tree. This result is propagated all
  the way back up as the final answer. ✓

  ![Step 3: at pair (3,3), values match — clones (the cloned-tree node) is returned as the answer](images/walkthrough-3.svg)

**Complexity:** O(n) time — in the worst case every node pair is visited once (e.g. when
`target` is the last node reached, or the tree is degenerate). O(h) space for the
recursion stack, where h is the tree's height.
