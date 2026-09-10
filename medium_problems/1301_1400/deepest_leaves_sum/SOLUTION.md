## Solution walkthrough

The implementation is `deepestLeavesSum(root *TreeNode) int` in `main.go`, wrapping a
recursive closure `foo` that carries the current level.

![Example 1](images/1.png)

We'll trace `[1,2,3,4,5,null,6,7,null,null,null,null,8]`. The deepest leaves are `7` and
`8` at level 3, so the answer is `15`.

1. **Two ints hold the whole answer.** `deepestLevel` and `sumOfDeepestLevel`, captured by
   the closure. Nothing per-level is stored — no map, no slice of levels — regardless of
   how wide the tree gets.

2. **The base case is a leaf, not nil.**

   ```go
   if node.Left == nil && node.Right == nil {
   ```

   Almost every recursion in this batch bottoms out at `if node == nil`. This one doesn't,
   and it's consistent: only leaves contribute to the answer, so a leaf is the meaningful
   terminal case. The recursive calls are guarded, so a nil child is never passed down and
   never needs representing.

3. **Three cases at a leaf, and the first is the interesting one.**

   ```go
   if level > deepestLevel {
       sumOfDeepestLevel = node.Val
       deepestLevel = level
   } else if level == deepestLevel {
       sumOfDeepestLevel += node.Val
   }
   ```

   Deeper than anything seen: everything accumulated so far belonged to a shallower level
   and is now worthless, so the sum is **replaced**, not added to. Equal: accumulate.
   Shallower: fall through and do nothing.

   The replacement is why one pass suffices. You never need to know the final depth in
   advance, because discovering a deeper leaf invalidates the previous answer outright.

   The first leaf reached is `7` at level 3, which sets the baseline.

   ![Step 1: the first leaf sets the level and the sum](images/walkthrough-1.png)

4. **Shallower leaves are discarded.** `5` sits at level 2 and contributes nothing.

   ![Step 2: a shallower leaf is ignored](images/walkthrough-2.png)

5. **Equal-depth leaves accumulate.** `8` is also at level 3, so the sum becomes `15`.

   ![Step 3: an equal-depth leaf is added](images/walkthrough-3.png)

   ![Step 4: the final answer, 15](images/walkthrough-4.png)

   Note the order this happens in doesn't matter. Had `8` been reached before `7`, the same
   two additions would occur in the other order. And had a level-4 leaf existed anywhere,
   whichever branch found it first would have reset the sum and the level-3 work would have
   been thrown away.

6. **`level` starts at 0 and the root is never nil.** `foo(root, 0)` is called without a
   nil check, and `foo` dereferences `node.Left` on its first line. That's safe only
   because the constraints guarantee at least one node. Allow an empty tree and this panics
   immediately — worth knowing, since the same function with a nil base case would have
   been robust to it for free.

**Complexity:** O(n) time, every node visited once in a single pass rather than the two the
obvious solution needs. Space is O(h) for the recursion stack; the answer itself is two
ints.
