## Solution walkthrough

The implementation is `isCousins(root *TreeNode, x int, y int) bool`, backed by a helper
`foo(node *TreeNode, currentDepth, searchVal int) (depth int, parent int, found bool)`,
both in `main.go`.

![Example 2](images/2.png "Example2")

We'll trace it on `root = [1,2,3,null,4,null,5]`, `x = 5`, `y = 4` — LeetCode's own
example shows this returns `true`:

1. **`isCousins` just delegates to `foo`, twice.**
   `depthX, parentX, _ := foo(root, 0, x)` and `depthY, parentY, _ := foo(root, 0, y)` —
   each call walks the tree from the root looking for one target value, and comes back
   with that node's depth and its parent's value. The `found` flag from each call is
   discarded (`_`) here — the problem guarantees `x` and `y` both exist in the tree, so
   it's never useful at this call site.

2. **`foo`'s defaults and base cases.** `depth = currentDepth`, `parent = 0`, `found =
   false` set "not found down this path" before anything else runs. `if node == nil {
   return currentDepth, 0, false }` handles walking off the tree. `if node.Val ==
   searchVal { return currentDepth, 0, true }` covers the node matching the search value
   directly on entry — its `parent` is reported as `0`, which doubles as the sentinel for
   "no parent" (the root case) because node values are constrained to be `>= 1`, so `0`
   can never collide with a real value.

3. **Checking a child before descending into it.** For both `node.Left` and
   `node.Right`, `foo` first checks `node.Left.Val == searchVal` (or `.Right`) directly,
   and if so returns `(currentDepth + 1, node.Val, true)` immediately. This is the only
   place `parent` ever becomes non-zero for a match below the root — it's the one
   vantage point where the target is being looked at *from its parent's own call*, so
   `node.Val` (the parent's value) is sitting right there to hand back.

4. **Recursing when a child isn't a direct match.** `depth, parent, found =
   foo(node.Left, currentDepth+1, searchVal)` descends into the left subtree; the
   following `if found { return }` short-circuits so the right subtree is skipped once
   the target has already turned up on the left. The right-side block mirrors this, but
   has no such early return after it — it's the last thing `foo` does, so it simply
   falls through to the final `return depth, parent, found`.

5. **The comparison, back in `isCousins`.** `return depthX == depthY && parentX !=
   parentY` — same depth, different parent, which is exactly the definition of cousins.

Tracing `foo` for `x = 5` first:

- `foo(1, 0, 5)`: node `1` doesn't match. Left child `2` doesn't match either, so it
  recurses into `foo(2, 1, 5)`.
  - `foo(2, 1, 5)`: node `2` doesn't match. `node.Left` is `nil`, so that branch is
    skipped entirely. `node.Right` is `4`, which also doesn't match `5`, so it recurses
    into `foo(4, 2, 5)`.
    - `foo(4, 2, 5)`: node `4` doesn't match and has no children — falls through to
      `return depth, parent, found` with the untouched defaults `(2, 0, false)`.
  - Back in `foo(2, 1, 5)`: the right branch's `found` came back `false`, so it also
    falls through, returning `(2, 0, false)`.
- Back in `foo(1, 0, 5)`: the left branch's `found` is `false`, so `if found { return }`
  doesn't fire — the search continues to `node.Right`, which is `3`. `node.Right.Val`
  (`3`) doesn't match `5`, so it recurses: `foo(3, 1, 5)`.
  - `foo(3, 1, 5)`: node `3` doesn't match. `node.Left` is `nil`. `node.Right` is `5` —
    a **direct child match** — so it returns immediately: `(1 + 1, 3, true)` = `(2, 3,
    true)`.
- Back in `foo(1, 0, 5)`: the right branch's result `(2, 3, true)` overwrites `depth,
  parent, found`, and the final `return depth, parent, found` sends back `(2, 3, true)`.

So `depthX = 2`, `parentX = 3`.

![Step 1: searching for x=5 — the miss down through node 2 and leaf 4 backtracks, then 5 is found as node 3's right child at depth 2](images/walkthrough-1.png)

Now `y = 4`, same starting point:

- `foo(1, 0, 4)`: node `1` doesn't match. Left child `2` doesn't match, so it recurses
  into `foo(2, 1, 4)`.
  - `foo(2, 1, 4)`: node `2` doesn't match. `node.Left` is `nil`, skipped. `node.Right`
    is `4` — a **direct child match** — so it returns immediately: `(1 + 1, 2, true)` =
    `(2, 2, true)`.
- Back in `foo(1, 0, 4)`: the left branch's `found` is `true`, so `if found { return }`
  fires right away — `node.Right` (`3`) is never even looked at. Returns `(2, 2, true)`.

So `depthY = 2`, `parentY = 2`.

![Step 2: searching for y=4 — found directly as node 2's right child at depth 2, so the right subtree is never visited](images/walkthrough-2.png)

6. **Back in `isCousins`:** `depthX == depthY` → `2 == 2` → `true`. `parentX !=
   parentY` → `3 != 2` → `true`. Both hold, so `isCousins` returns `true` — `5` and `4`
   are cousins: same depth, different parents. ✓

![Step 3: comparing the two results — same depth (2 and 2), different parents (3 and 2) — cousins](images/walkthrough-3.png)

**Complexity:**
- Time: O(n) — `foo` is called twice from `isCousins`, and each call visits at most
  every node once (less, when the target is found early and the sibling subtree gets
  pruned by `if found { return }`).
- Space: O(h) for the recursion stack per call, where h is the tree's height (O(log n)
  for a balanced tree, O(n) for a completely skewed one).
