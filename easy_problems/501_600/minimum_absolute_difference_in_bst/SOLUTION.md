## Solution walkthrough

The implementation is `getMinimumDifference(root *TreeNode) int` in `main.go`.

![Example 1](images/1.jpg "Example1")

We'll trace it on the example above, `root = [4,2,6,1,3]` (expected output `1`).

1. **Set up state.** `res := math.MaxInt` starts the running minimum as large as
   possible, and `var prev *TreeNode` tracks the previously-visited node in in-order
   sequence — initially `nil`, since nothing has been visited yet.

2. **Define the in-order recursion.** `helper` is a closure so it can read and mutate
   `res` and `prev` directly, without threading them through return values.
   `if root == nil { return }` is the base case: an empty subtree contributes nothing.

3. **Recurse left first.** `helper(root.Left)` visits the entire left subtree before
   the current node does anything — this is what guarantees in-order (sorted) visit
   order for a BST.

4. **Compare against `prev`, then advance it.**
   `if prev != nil { res = min(res, root.Val-prev.Val) }` — only compare once there
   *is* a previous node (skipped on the very first node visited). Because traversal
   order is sorted, `root.Val - prev.Val` is always non-negative, so there's no need
   for `abs()`. Then `prev = root` makes the current node the predecessor for
   whichever node in-order visits next.

5. **Recurse right.** `helper(root.Right)` continues the in-order walk into the right
   subtree, using the now-updated `prev`.

6. **Kick it off and return.** `helper(root)` runs the traversal starting at the root,
   then `return res` reports the smallest adjacent gap found.

Walking it through `root = [4,2,6,1,3]`, whose in-order sequence is `1, 2, 3, 4, 6`:

- `helper(4)` recurses left into `helper(2)`, which recurses left into `helper(1)`,
  which recurses left into `helper(nil)` (returns immediately). Back at node `1`:
  `prev` is still `nil`, so the comparison is skipped, and `prev` becomes `1`.

  ![Step 1: descend to leftmost node 1, prev initialized](images/walkthrough-1.svg)

- `helper(1)` then recurses right into `helper(nil)` and returns. Back at node `2`:
  `prev` is `1`, so `res = min(∞, 2-1) = 1`, then `prev` becomes `2`.

  ![Step 2: at node 2, diff 2-1=1, res becomes 1](images/walkthrough-2.svg)

- `helper(2)` recurses right into `helper(3)`. At node `3`: `prev` is `2`, so
  `res = min(1, 3-2) = 1` (unchanged), then `prev` becomes `3`.

  ![Step 3: at node 3, diff 3-2=1, res stays 1](images/walkthrough-3.svg)

- Control returns up to the root, node `4`: `prev` is `3`, so
  `res = min(1, 4-3) = 1` (unchanged), then `prev` becomes `4`.

  ![Step 4: back at root 4, diff 4-3=1, res stays 1](images/walkthrough-4.svg)

- `helper(4)` recurses right into `helper(6)`. At node `6`: `prev` is `4`, so
  `res = min(1, 6-4) = min(1, 2) = 1` (unchanged), then `prev` becomes `6`.
  `helper(6)` recurses right into `helper(nil)` and the traversal ends.

  ![Step 5: at node 6, diff 6-4=2, res stays 1 — final answer](images/walkthrough-5.svg)

- `getMinimumDifference` returns `res = 1`. ✓

Note that the algorithm never explicitly sorts anything — the BST's in-order property
does that for free, so each node only ever needs to be compared against the single
node visited immediately before it.
