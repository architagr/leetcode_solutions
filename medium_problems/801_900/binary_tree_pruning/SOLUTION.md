## Solution walkthrough

The implementation is `PruneTree(root *TreeNode) *TreeNode` over the recursive `parse`, in
`binary_tree_pruning.go`. `parse` returns a boolean — "does my subtree contain a 1?" — and
does the pruning as a side effect.

![Example 1](images/1.png)

We'll trace `[1,null,0,0,1]`: root `1` with no left child and a right child `0`, whose
children are `0` and `1`. Expected output `[1,null,0,null,1]`.

1. **The nil case reports false.** `if node == nil { return false }`. An absent subtree
   contains no `1`, so it contributes nothing to the `||` below. It also makes the parent's
   `node.Right = nil` a harmless no-op on a pointer that was already nil.

2. **Post-order, because a node can't answer on the way down.** It doesn't yet know what's
   beneath it, so both children report first and the node combines their answers with its
   own value.

   ![Step 1: a leaf 0 reports false](images/walkthrough-1.png)

3. **The parent does the pruning.**

   ```go
   right := parse(node.Right)
   if !right {
       node.Right = nil
   }
   ```

   This is the part that's easy to get wrong. A node cannot remove itself — it has no
   reference to its parent, and nulling a local variable changes nothing the caller can
   see. The child reports "nothing worth keeping down here", and the parent, which holds
   the pointer, is the one that clears it.

   ![Step 2: the parent clears the pointer to the pruned child](images/walkthrough-2.png)

4. **The combination is a plain or.** `return curr || right || left`, where `curr` is
   `node.Val == 1`. A node stays if it is itself a `1`, or if either subtree kept anything.

   Node `0` here is not a `1`, but its right subtree kept the leaf `1`, so it survives —
   which is the case that shows why "prune every 0" would be the wrong rule.

   ![Step 3: a 0 survives because its subtree kept a 1](images/walkthrough-3.png)

5. **The wrapper exists for the root.** The root has no parent, so nobody would ever clear
   it:

   ```go
   curr := parse(temp)
   if !curr {
       root = nil
   }
   ```

   Without this, a tree of all zeroes would come back intact instead of empty.

   ![Step 4: the root reports true and is kept](images/walkthrough-4.png)

6. **Right before left, and it doesn't matter.** `parse` recurses right first. For an
   ordering-sensitive problem that would be a decision; here each subtree's answer is
   independent of the other's, so the order is arbitrary.

7. **A note on the shape.** This returns a boolean and mutates the tree as a side effect.
   The more common formulation returns `*TreeNode` and has the caller reassign —
   `root.Left = prune(root.Left)` — which is what Day 39 used for deletion. Both work, and
   the reassigning version needs no wrapper, because returning nil for the root handles the
   whole-tree case naturally.

**Complexity:** O(n) time, every node visited once. Space is O(h) for the recursion stack,
where h is the tree's height.
