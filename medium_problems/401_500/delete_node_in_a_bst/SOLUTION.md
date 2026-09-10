## Solution walkthrough

The implementation is `deleteNode(root *TreeNode, key int) *TreeNode` with `successor` and
`predecessor` helpers, in `main.go`.

![Example 1](images/1.jpg)

We'll trace `root = [5,3,6,2,4,null,7]`, `key = 3`. Root `5` with children `3` and `6`; `3`
has children `2` and `4`; `6` has a right child `7`.

1. **The search half is Day 26's walk.**

   ```go
   if key > root.Val {
       root.Right = deleteNode(root.Right, key)
   } else if key < root.Val {
       root.Left = deleteNode(root.Left, key)
   }
   ```

   One comparison picks the only direction the key could be in.

   ![Step 1: at 5, the key is smaller, so recurse left](images/walkthrough-1.png)

2. **The reassignment is not decoration.** `root.Left = deleteNode(...)` rather than calling
   and discarding the result. Deleting inside a subtree can change which node roots that
   subtree, so every caller has to take back what comes out and re-attach it. This is what
   makes the leaf case work at all — `root = nil` deep in the recursion only takes effect
   because the parent assigns the returned nil into its own child pointer.

3. **A leaf is the easy case.** `if root.Left == nil && root.Right == nil { root = nil }`.

4. **A node with two children can't be unhooked.** Something has to fill the hole, and only
   two values in the whole tree can: the in-order predecessor and the in-order successor —
   the values immediately before and after it in sorted order.

   ![Step 2: 3 has two children](images/walkthrough-2.png)

   That follows from the property Day 32 used. If the in-order sequence must stay sorted,
   whatever fills the hole has to sit between everything in the left subtree and everything
   in the right. The only candidates are the largest value on the left and the smallest on
   the right.

5. **Copy the successor's value up, then delete the successor.**

   ```go
   root.Val = successor(root)
   root.Right = deleteNode(root.Right, root.Val)
   ```

   `successor` walks one step right, then all the way left — the smallest value greater
   than the node. Here that's `4`.

   ![Step 3: the successor's value is copied into the hole](images/walkthrough-3.png)

   Note the second line deletes `root.Val`, which by then is the *new* value. The node
   itself is never removed; its value is overwritten and the duplicate below is deleted
   instead.

6. **Why the recursion terminates.** It looks circular — delete calls delete — but the
   successor is the leftmost node of the right subtree, so by definition it has no left
   child. A node with at most one child is one of the easy cases, so each step descends
   into a strictly simpler problem and bottoms out at a leaf.

   ![Step 4: deleting the successor is an easy case](images/walkthrough-4.png)

7. **The third branch exists for a concrete reason.** `else { root.Val = predecessor(root)
   ... }` handles a node with a left child and no right one. `successor` does
   `root = root.Right` unconditionally and would panic on a nil right child, so this case
   can't fall through to it. The predecessor — rightmost of the left subtree, so no right
   child of its own — is the mirror answer and equally valid. The problem's own second
   picture shows an alternative correct output for exactly this reason.

   ![An equally valid answer for the same input](images/2.jpg)

**Complexity:** O(h) time, where h is the tree's height — one path down to find the node,
one more to find a successor, and the recursive delete follows that same path. O(log n)
balanced, O(n) skewed. Space is O(h) for the recursion stack.
