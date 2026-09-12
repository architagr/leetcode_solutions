## Solution walkthrough

The implementation is `findLeaves(root *TreeNode) [][]int`, a wrapper over the recursive
`parse`, plus a small `maxVal` helper, in `main.go`.

![Example 1](images/1.jpg)

We'll trace `[1,2,3,4,5]`: root `1` with children `2` and `3`; `2` has children `4` and
`5`. Expected output: `[[4,5,3],[2],[1]]`.

1. **The wrapper throws away the return value.** `_ = parse(root, &arr)`. `parse` returns
   a round number, which matters to a parent but not to the caller — the answer was built
   up inside `arr` as a side effect. The blank identifier says that on purpose.

2. **`parse` is post-order.** Both children are visited before the current node does
   anything:

   ```go
   leftIndex := parse(root.Left, arr)
   rightIndex := parse(root.Right, arr)
   index := maxVal(leftIndex, rightIndex) + 1
   ```

   That ordering isn't a preference. A node's round depends on its children's rounds, so
   the children have to have reported first.

3. **The nil case returns `-1`, and that's the load-bearing constant.** A leaf's two
   children are both nil, so `maxVal(-1, -1) + 1` is `0` — leaves land in round 0.

   Return `0` from the nil case instead and every node's round shifts up by one, leaving
   an empty bucket at the front. The `-1` is what anchors the numbering to the leaves.

   ![Step 1: leaf 4 computes round 0 from two nil children](images/walkthrough-1.png)

4. **`max` of the two children, not `min` or a sum.** A node is removed in the round after
   *both* its children have gone, so it has to wait for the slower side. `maxVal(0, 0) + 1`
   makes node `2` round 1.

   ![Step 2: leaf 5 returns 0 and joins the same bucket](images/walkthrough-2.png)

   ![Step 3: node 2 waits for both children, so max(0,0)+1 = 1](images/walkthrough-3.png)

5. **Buckets grow lazily.** `if len(*arr) <= index { *arr = append(*arr, []int{}) }` —
   the first node to reach a round creates its bucket, growing the outer slice on demand
   level slices.

6. **Round is height, not depth.** This is the point of the whole problem, and node `3`
   shows it: it sits at depth 1, shallower than `4` and `5` at depth 2, and still lands in
   round 0 with them, because it's a leaf.

   ![Step 4: leaf 3 is round 0 despite being shallower](images/walkthrough-4.png)

7. **The root waits for its deepest side.** `maxVal(1, 0) + 1` is `2`.

   ![Step 5: the root lands in round 2](images/walkthrough-5.png)

   Final answer `[[4,5,3],[2],[1]]`. The order within a bucket is whatever the traversal
   produced, and the problem explicitly allows that — `[[3,5,4],[2],[1]]` is equally
   correct — so nothing sorts.

8. **`arr` is a `*[][]int`.** Growing the outer slice inside a recursive call has to be
   visible to every other call, so the pointer is passed rather than the slice threaded
   through as a return value. Day 8 solved the same reallocation problem the other way,
   returning the slice; both work.

**Complexity:** O(n) time — every node is visited exactly once. Simulating the problem's
description literally, stripping leaves and re-walking, would be O(n * h), which is the
whole reason to compute heights instead. Space is O(h) for the recursion stack plus O(n)
for the output.
