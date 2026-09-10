## Solution walkthrough

The implementation is `upsideDownBinaryTree(root *TreeNode) *TreeNode` in `main.go` — a
single recursive function, thirteen lines, doing two unrelated jobs at once.

![How the rotation works](images/1.jpg)

We'll trace `[1,2,3,4,5]`: root `1` with children `2` and `3`; `2` has children `4` and
`5`. Expected output `[4,5,2,null,null,3,1]`.

1. **The recursion only ever goes left.** `upsideDownBinaryTree(root.Left)` is the only
   recursive call. The right children are moved but never recursed into, which is why the
   cost is the length of the left spine rather than the size of the tree.

2. **The base case finds the new root.** `if root == nil || root.Left == nil { return root }`
   — keep taking `Left` until there isn't one. Here that's `4`.

   ![Step 1: descend the left spine to find the deepest node](images/walkthrough-1.png)

3. **The return value is a pass-through.**

   ```go
   x := upsideDownBinaryTree(root.Left)
   // ... rewiring, none of which touches x ...
   return x
   ```

   `x` is produced once, at the bottom, and handed upward untouched through every frame.
   Nothing above the base case ever computes a new return value.

   ![Step 2: the new root travels up unmodified](images/walkthrough-2.png)

   That's the key to reading this function: it's doing two separate jobs on one recursion.
   One is finding and propagating the new root. The other is local pointer surgery at each
   level. They share a traversal and otherwise have nothing to do with each other.

4. **The rewiring, on the way back up.**

   ```go
   newRight := root
   newNode := root.Left
   newLeft := root.Right
   newNode.Left = newLeft
   newNode.Right = newRight
   ```

   The three temporaries name what each node is about to become, which is the only reason
   this reads clearly. `root.Left` is promoted above `root`; the old right child becomes
   its left; the old parent becomes its right.

   Doing this after the recursive call matters: the child subtree has to be turned over
   before its old parent can be hung underneath it.

   ![Step 3: frame for node 2 rewires 4's pointers](images/walkthrough-3.png)

5. **The clearing that only matters once.**

   ```go
   root.Left = nil
   root.Right = nil
   ```

   This looks like it should destroy the work. It doesn't, because for every node except
   the original root, the parent's frame is about to overwrite both pointers as part of its
   own rewiring — the frame for `1` sets `2.Left = 3` and `2.Right = 1`, replacing the nils
   the frame for `2` just wrote.

   The clearing survives only for the topmost call. The original root becomes a leaf in the
   flipped tree and genuinely needs both pointers cleared. It's written uniformly and takes
   effect once.

   ![Step 4: the top frame overwrites the nils and finishes the tree](images/walkthrough-4.png)

6. **Why moving a right child wholesale is safe.** The problem guarantees every right child
   has a left sibling and no children of its own. So a right child is always a leaf, and
   reattaching it as someone's left child can't drag a subtree along with it. Without that
   guarantee this rotation wouldn't be well defined.

**Complexity:** O(h) time, where h is the length of the left spine — the recursion never
descends right. Space is O(h) for the recursion stack.
