## Solution walkthrough

The implementation is `evaluateTree(root *TreeNode) bool` in `main.go`, backed by four
named constants: `FALSE = 0`, `TRUE = 1`, `OR = 2`, `AND = 3` — these mirror exactly the
encoding LeetCode uses for node values, so the code reads as directly as the problem
statement.

![Example 1](images/1.png "Example1")

We'll trace it on the example above, `root = [2,1,3,null,null,0,1]`.

1. **Base case: leaf node.**
   `if root.Left == nil && root.Right == nil { return root.Val == TRUE }` — because the
   tree is guaranteed **full** (every node has 0 or 2 children), checking that *both*
   children are `nil` is enough to know this is a leaf; there's no partial-children case
   to special-case. A leaf's value is `0` (`FALSE`) or `1` (`TRUE`), so comparing it
   against the `TRUE` constant directly yields the boolean — no separate lookup table
   needed.

2. **Recurse into both children first.**
   `left := evaluateTree(root.Left)` then `right := evaluateTree(root.Right)` — this is
   post-order: neither child's boolean is known until its own subtree has been fully
   evaluated, so both recursive calls must return before this node can combine anything.

3. **Combine with the node's operator.**
   `if root.Val == OR { return left || right }` handles the `OR` case (`root.Val == 2`);
   any node reaching this point that isn't `OR` must be `AND` (`root.Val == 3`, since
   leaves already returned in step 1 and the problem guarantees no other non-leaf
   value), so the final line falls through to `return left && right`.

Now the trace on `root = [2,1,3,null,null,0,1]` — root is `OR`, its left child is leaf
`1`, its right child is `AND` whose own children are leaves `0` and `1`:

- The three leaves hit the base case immediately: `evaluateTree` on root's left child
  (`1`) returns `True`; on `AND`'s left child (`0`) returns `False`; on `AND`'s right
  child (`1`) returns `True`.

  ![Step 1: leaves hit the base case, root.Val == TRUE decides each](images/walkthrough-1.svg)

- Back at the `AND` node: `left = False`, `right = True`. `root.Val` is `AND`, not
  `OR`, so it falls to `return left && right` → `False && True = False`.

  ![Step 2: AND node combines left && right = False && True = False](images/walkthrough-2.svg)

- Back at the `root` (`OR`) node: `left = True` (its own left child, resolved in step
  1), `right = False` (the just-resolved `AND` subtree). `root.Val == OR`, so
  `return left || right` → `True || False = True`. ✓

  ![Step 3: root OR combines left || right = True || False = True](images/walkthrough-3.svg)

**Complexity:** O(n) time — every node is visited and evaluated exactly once. O(h)
space for the recursion stack, where h is the tree's height.
