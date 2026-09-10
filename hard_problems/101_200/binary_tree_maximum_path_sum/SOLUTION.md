## Solution walkthrough

The implementation is `maxPathSum(root *TreeNode) int` over the recursive
`recurssivePathSum`, plus a `maxValue` helper, in `main.go`.

![Example 2](images/2.jpg)

We'll trace `[-10,9,20,null,null,15,7]`: root `-10` with children `9` and `20`; `20` has
children `15` and `7`. Expected answer: `42`, the path `15 -> 20 -> 7`, which never touches
the root.

1. **The answer lives outside the return value.** `max := -2000` in the wrapper, passed
   down as `*int`. The seed is below anything reachable — values are at least `-1000` and
   a path is non-empty — so the first real candidate always replaces it. A `0` seed would
   be wrong on an all-negative tree, where the answer is the least negative single node.

2. **Post-order, because both children must report first.**

   ```go
   left := recurssivePathSum(root.Left, max)
   right := recurssivePathSum(root.Right, max)
   ```

3. **Each node computes two different things.** This is the whole problem. As the top of a
   path, a node may use both children. As a link inside some ancestor's path, it may use
   only one, because the ancestor attaches above it and a path can't have a node with three
   neighbours.

   ![Step 1: at a leaf the two roles agree](images/walkthrough-1.png)

   At a leaf they coincide, which is why leaves are a bad place to notice the distinction.

   ![Step 2: leaf 15 returns 15 and raises max](images/walkthrough-2.png)

   ![Step 3: leaf 7 returns 7, not beating the best so far](images/walkthrough-3.png)

4. **The candidate goes into `max`; the one-sided value is returned.** Node `20` is where
   the two part company: it records `15 + 7 + 20 = 42` as a candidate, but returns only
   `20 + max(15, 7) = 35`.

   ![Step 4: node 20 records 42 but returns 35](images/walkthrough-4.png)

   Returning 42 here would be the classic bug. The parent would attach to it and build a
   "path" that forks at `20`.

5. **The root loses, and that's the point of this example.** `9 + 35 - 10 = 34`, which
   doesn't beat `42`. The best path never passes through the root, which is exactly the
   case a root-anchored solution gets wrong.

   ![Step 5: the root's candidate loses to a path it isn't on](images/walkthrough-5.png)

6. **Negatives are handled by comparing against not taking the branch.** Every use of a
   child is guarded against just using `root.Val` alone. `maxValue(left+root.Val,
   root.Val)` means "extend into the left child, or don't."

7. **The nested `maxValue` calls are equivalent to a much shorter expression.** The code
   builds the candidate by comparing `root.Val` against `right+val`, `left+val` and
   `left+right+val`. Written directly, the candidate is `val + max(0, left) + max(0, right)`
   and the return is `val + max(0, left, right)`. Those agree with the implementation on
   every input — worth knowing, because the short form is what you'd want to write from
   scratch and the long form is what makes the code look harder than the idea.

**Complexity:** O(n) time, one visit per node with constant work. Space is O(h) for the
recursion stack, where h is the tree's height.
