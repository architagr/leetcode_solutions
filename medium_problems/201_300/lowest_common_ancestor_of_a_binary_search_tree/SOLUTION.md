## Solution walkthrough

The implementation is `lowestCommonAncestor(root, p, q *TreeNode) *TreeNode` in `main.go`.
There is no helper and no accumulator — the function calls itself on one side or returns.

![Example 1](images/1.png)

We'll trace `root = [6,2,8,0,4,7,9,null,null,3,5]` with `p = 2`, `q = 4`. Expected answer:
`2`.

1. **The nil guards are for degenerate input, not structure.**

   ```go
   if p == nil && q == nil {
       return root
   } else if p == nil || q == nil {
       return nil
   }
   ```

   Both targets nil means every node trivially has both as descendants, so the current root
   is returned. One nil means there's nothing sensible to answer with. Neither case arises
   from a valid call — the problem guarantees both nodes exist in the tree.

2. **Both targets smaller means go left.**

   ```go
   if root.Val > p.Val && root.Val > q.Val {
       return lowestCommonAncestor(root.Left, p, q)
   }
   ```

   At `6`, both `2` and `4` are smaller. In a BST that means both must be in the left
   subtree, so whatever node they share is also in there — it can't be `6`.

   ![Step 1: both targets are smaller than 6](images/walkthrough-1.png)

   One comparison per target eliminates the entire right subtree. Nothing is searched and
   nothing is revisited.

   ![Step 2: the right subtree is gone](images/walkthrough-2.png)

3. **Both larger means go right.** The mirror case, not exercised by this example.

4. **Anything else is the answer.** `return root`, and this single line covers two
   different situations.

   The first is targets on opposite sides: this node is where the paths to them diverge, so
   it's the lowest node with both beneath it.

   The second is one target being this node. The LCA definition allows a node to be a
   descendant of itself, so the node is its own answer — its path and the other target's
   path meet right here. That's this trace: at `2`, `p` is `2` itself, so neither "both
   smaller" nor "both larger" holds.

   ![Step 3: at 2, neither directional case holds](images/walkthrough-3.png)

   Nothing has to work out which of the two it was. "Not both smaller, not both larger" is
   exactly the union of them.

   ![Step 4: 2 is returned](images/walkthrough-4.png)

5. **The comparisons use `Val`, not pointer identity.** Fine here because BST values are
   unique. In the general-tree version tomorrow the same choice shows up again and matters
   more.

**Complexity:** O(h) time, where h is the tree's height — one comparison per level, never
branching, never backtracking. That's O(log n) balanced, O(n) skewed. Space is O(h) for
the recursion stack, and since every recursive call here is in tail position, rewriting it
as a `for` loop would make it O(1).
