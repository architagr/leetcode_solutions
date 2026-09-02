## Solution walkthrough

The implementation is `isUnivalTree(root *TreeNode) bool` in `main.go`.

![Example 2](images/2.png "Example2")

We'll trace it on `root = [2,2,2,5,2]`, which LeetCode's own example shows returning
`false`:

```
        2
       / \
      2   2
     / \
    5   2
```

1. **Base case.** `if root == nil { return true }` — an empty subtree is trivially
   uni-valued; there's nothing in it to disagree with the parent's value.

2. **Assume both sides are fine, then try to disprove it.** `left, right := true, true`
   starts optimistic. Each side only gets checked — and can only become `false` — if
   the corresponding child actually exists.

3. **Check the left child, if it exists.**
   `if root.Left != nil { left = isUnivalTree(root.Left) && root.Left.Val == root.Val }`
   — this does two things at once: it recurses into the left subtree to make sure *it's*
   internally uni-valued, *and* it compares the left child's value to the current
   node's value (`root.Val`), from the parent's vantage point. Both must hold for
   `left` to stay `true`.

4. **Check the right child the same way.**
   `if root.Right != nil { right = isUnivalTree(root.Right) && root.Right.Val == root.Val }`
   — identical logic, mirrored for the right side.

5. **Combine.** `return left && right` — the subtree rooted at `root` is uni-valued
   only if both the left and right sides checked out.

Walking it through `root = [2,2,2,5,2]`:

- `isUnivalTree(2)` (root) recurses left into node `A = 2`, and right into node
  `B = 2` (a leaf).
  - `isUnivalTree(A=2)` recurses left into leaf `C = 5`, and right into leaf `D = 2`.
    - `isUnivalTree(C=5)`: no children → `left=true, right=true` → returns `true`
      (this call has no idea `5` doesn't match its parent — that's not its job).
    - `isUnivalTree(D=2)`: no children → returns `true`, same reasoning.
    - `isUnivalTree(B=2)`: no children → returns `true`.

    ![Step 1: 5, 2, and 2 bottom out as leaves, each returning true](images/walkthrough-1.svg)

  - Back in `isUnivalTree(A=2)`: `left = isUnivalTree(C) && C.Val == A.Val` →
    `true && (5 == 2)` → `true && false` → `left = false`. `right =
    isUnivalTree(D) && D.Val == A.Val` → `true && (2 == 2)` → `right = true`. Returns
    `left && right = false && true = false`.

    ![Step 2: at node A (value 2), left child 5 mismatches, left becomes false](images/walkthrough-2.svg)

  - Back at the root: `left = isUnivalTree(A) && A.Val == root.Val` →
    `false && (2 == 2)` — the left-hand side of `&&` is already `false`, so `left =
    false` regardless of what the value comparison would have said. `right =
    isUnivalTree(B) && B.Val == root.Val` → `true && (2 == 2)` → `right = true`.
    Returns `left && right = false && true = false`. ✓

    ![Step 3: at the root, the left subtree already resolved to false, so the whole tree is false](images/walkthrough-3.svg)

The mismatch (`5` under a `2`) is detected two levels down, at node `A`, and then just
propagates upward as `false` through every ancestor's `left`/`right` combination —
nothing further up needs to re-check values it already knows are wrong.
