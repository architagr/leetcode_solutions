## Solution walkthrough

The implementation is `LevelOrder(root *TreeNode) [][]int`, a thin wrapper that seeds an
empty accumulator and calls the recursive `traversal`, in
`binary_tree_level_order_traversal.go`.

![Example 1](images/1.jpg)

We'll trace `[3,9,20,null,null,15,7]`: root `3` has children `9` and `20`, and `20` has
children `15` and `7`. Expected output: `[[3],[9,20],[15,7]]`.

1. **The accumulator is the output.** `arr := make([][]int, 0)` and
   `traversal(root, arr, 0)`. There's no separate queue or per-level buffer — the thing
   being built up during the traversal is the answer, one inner slice per level, and the
   root starts at level `0`.

2. **The base case returns the accumulator unchanged.** `if head == nil { return arr }`.
   Nothing to append, nothing to grow, hand back what came in. The tests cover the empty
   tree, and this is the line that makes it return `[][]int{}` rather than panicking.

3. **Grow the outer slice on first arrival at a depth.**

   ```go
   if (len(arr) - 1) < level {
       a := make([]int, 0)
       arr = append(arr, a)
   }
   ```

   `len(arr) - 1` is the highest level index that currently exists. If the level being
   visited is beyond it, there's no slot to append into yet, so one gets created. This
   fires exactly once per level — the first node to reach that depth opens it, and
   everything after finds it already there.

   ![Step 1: level 0 opens and the root appends](images/walkthrough-1.png)

4. **Every node appends into the slot for its own depth.** `arr[level] = append(arr[level],
   head.Val)`. This is the line the whole approach rests on. The node doesn't care when it
   was visited or which branch it came from; it writes into `arr[level]`, and the level is
   a parameter it was handed.

   ![Step 2: the recursion goes left first, so 9 lands in arr[1]](images/walkthrough-2.png)

5. **Left before right, at every depth.**

   ```go
   arr = traversal(head.Left, arr, level+1)
   arr = traversal(head.Right, arr, level+1)
   ```

   Both calls pass `level+1`, and both capture the returned slice rather than relying on
   mutation — `append` can reallocate the outer slice when step 3 grows it, so the caller
   has to take back what the callee produced.

   Descending into `Left` first is what makes each level read left to right. When the
   recursion comes back up and takes the right branch, `20` appends into `arr[1]`, the
   same slot `9` used, and lands after it.

   ![Step 3: 20 joins the slot 9 already opened](images/walkthrough-3.png)

6. **Deeper levels open the same way.** `15` is the first node to reach depth 2, so step 3
   creates `arr[2]` for it.

   ![Step 4: level 2 opens for 15](images/walkthrough-4.png)

   Then `7` finds that slot already there and appends after it, completing
   `[[3],[9,20],[15,7]]`.

   ![Step 5: 7 lands in arr[2] as well](images/walkthrough-5.png)

The test's third case is a good check on the left-to-right claim: a tree yielding
`[[1],[2,3],[4,5,6,7],[8,9],[10]]` has four nodes on one level arriving from two different
branches, and they still come out in order.

**Complexity:** O(n) time — every node is visited once and does constant work. Space is
O(h) for the recursion stack plus O(n) for the output. Worth noting against BFS: the
recursion's peak memory is the tree's height, where a queue-based solution's would be the
width of the widest level. On a skewed tree this one is worse; on a bushy one it's better.
