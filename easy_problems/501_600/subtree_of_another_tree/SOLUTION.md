## Solution walkthrough

The implementation is two functions in `main.go`: `isSubtree(root, subRoot *TreeNode) bool`,
which searches for a matching anchor node, and its helper `equalBinaryTree(root, subRoot
*TreeNode) bool`, which checks whether two trees are structurally identical.

![Example 1](images/1.jpg "Example1")

We'll trace it on the example above, `root = [3,4,5,1,2]`, `subRoot = [4,1,2]` (expected `true`).

### `isSubtree` — find a candidate anchor

1. **Base case: an empty `subRoot` always matches.** `if subRoot == nil { return true }`
   — trivially, every tree contains the empty tree as a subtree.

2. **Base case: an empty `root` can't contain a non-empty `subRoot`.**
   `if root == nil { return subRoot == nil }` — since we already know `subRoot != nil`
   at this point, this returns `false`.

3. **Try `root` itself as the anchor.**
   `if root.Val == subRoot.Val && equalBinaryTree(root, subRoot) { return true }` — the
   value check first is a cheap filter: there's no point running the full structural
   comparison (`equalBinaryTree`) unless the roots even agree on value.

4. **Otherwise, recurse into both children.**
   `return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)` — if `root`
   itself isn't a match, maybe some node deeper in its left or right subtree is. Because
   `||` short-circuits, the right subtree is only searched if the left subtree search
   comes back empty-handed.

Walking it through the example, `isSubtree(root=3, subRoot=4)`:
- `subRoot` isn't nil, `root` isn't nil, so we compare values: `3 != 4` → skip
  `equalBinaryTree`, fall through to the recursive case.

![Step 1: node 3 vs subRoot's 4 — values differ, skip equalBinaryTree, recurse into 4 and 5](images/walkthrough-1.png)

- Recurse left: `isSubtree(root=4, subRoot=4)`. Now `4 == 4`, so this time
  `equalBinaryTree(4-subtree, subRoot)` actually runs.

![Step 2: node 4 vs subRoot's 4 — values match, call equalBinaryTree](images/walkthrough-2.png)

### `equalBinaryTree` — confirm the anchor matches exactly

5. **Both nil → equal.** `if root == nil { return subRoot == nil }` — two empty trees
   match.

6. **One nil, one not → not equal.** `if subRoot == nil { return root == nil }` — since
   `root != nil` was just ruled out above, this returns `false` whenever only one side
   is empty.

7. **Values must match.** `if root.Val != subRoot.Val { return false }` — an immediate
   mismatch stops the comparison right there.

8. **Recurse into both children, and require both sides to agree.**
   `return equalBinaryTree(root.Left, subRoot.Left) && equalBinaryTree(root.Right,
   subRoot.Right)` — the two trees are walked in lockstep, left-to-left and
   right-to-right; every corresponding pair of nodes has to match for the whole thing to
   be equal.

Continuing the trace, `equalBinaryTree(root=[4,1,2], subRoot=[4,1,2])`:
- `4 == 4` at the top, then `equalBinaryTree(1, 1)` (both leaves, values match, both
  children nil on both sides → `true`) and `equalBinaryTree(2, 2)` (same reasoning →
  `true`). Both sides agree, so the whole call returns `true`.
- Back in `isSubtree(root=4, subRoot=4)`, since `4 == 4 && equalBinaryTree(...) ==
  true`, it returns `true` immediately.
- Because of `||` short-circuiting, `isSubtree(root=5, subRoot)` is **never called** —
  node `5` is never even examined, since the left branch already found a match.
- That `true` propagates all the way back up through `isSubtree(root=3, subRoot)`. ✓

![Step 3: equalBinaryTree walks 4/1/2 against 4/1/2 in lockstep — all match, isSubtree returns true, node 5 never checked](images/walkthrough-3.png)

**Complexity:** O(m·n) time in the worst case (`m` = nodes in `root`, `n` = nodes in
`subRoot`) — up to `m` candidate anchors, each costing up to `n` work in
`equalBinaryTree`. O(h1 + h2) space for the two recursion stacks, where `h1`/`h2` are
the heights of `root`/`subRoot`.
