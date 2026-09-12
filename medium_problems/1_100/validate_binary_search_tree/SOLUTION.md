## Solution walkthrough

The implementation is `IsValidBst(root *TreeNode) bool` in
`validate_binary_search_tree.go`, over the `inorderTraversal` helper. The file also carries
a second, slower approach kept for contrast, covered at the end.

![Example 2](images/2.jpg)

We'll trace `[5,1,4,null,null,3,6]` — the invalid example. Root `5` with children `1` and
`4`; `4` has children `3` and `6`. Expected result: `false`.

1. **Collect the values in-order.**

   ```go
   arr = inorderTraversal(A.Left, arr)
   arr = append(arr, A.Val)
   arr = inorderTraversal(A.Right, arr)
   ```

   Left subtree, then the node, then the right subtree. The slice is threaded through as a
   parameter and a return value — the same handling of `append` reallocating as Day 8.

   The walk descends fully left before emitting anything, so `1` comes out first.

   ![Step 1: in-order emits the leftmost node first](images/walkthrough-1.png)

   ![Step 2: then the root](images/walkthrough-2.png)

2. **The violation shows up as an ordering break, not as a comparison against an ancestor.**
   Node `3` sits in the root's right subtree, so it must be greater than `5`. It isn't. In
   the in-order sequence that means `3` is emitted straight after `5`.

   ![Step 3: 3 is emitted after 5](images/walkthrough-3.png)

3. **Then a single linear scan finds it.**

   ```go
   for i := 0; i < len(arr)-1; i++ {
       if arr[i] >= arr[i+1] {
           return false
       }
   }
   return true
   ```

   `arr` is `[1,5,3,4,6]`, and the pair `(5,3)` fails. No bounds were threaded down the
   recursion and no ancestor was consulted — the out-of-place node simply landed in the
   wrong slot.

   ![Step 4: the adjacent pair 5,3 fails the scan](images/walkthrough-4.png)

4. **`>=`, not `>`.** The problem says strictly less and strictly greater, so equal
   neighbours are invalid too. Using `>` would accept a tree containing duplicates.

5. **The loop bound is `len(arr)-1`.** It compares pairs, so it stops one short. On a
   single-node tree the loop body never runs and the function returns `true`, which is
   correct — and worth noting because the constraints guarantee at least one node, so the
   empty case never arises here.

6. **The second approach, and why it isn't the one used.** `IsValidBstApproch1` implements
   the local definition honestly: at each node, scan the entire left subtree for its
   maximum and the entire right subtree for its minimum, then require
   `max < root.Val < min`. That is correct — it compares against whole subtrees, not just
   immediate children, so it catches the deep-violation case too. But it re-scans subtrees
   at every node, which makes it O(n^2) on a skewed tree. It's in the file as the contrast.

   The tempting version that *is* wrong is neither of these: checking each node against its
   two immediate children only. That accepts trees where a node deep in a left subtree
   exceeds an ancestor several levels up.

**Complexity:** O(n) time — one traversal plus one linear scan. Space is O(n) for the
collected slice plus O(h) for the recursion stack. The standard improvement is to compare
each value against the previously visited one during the traversal rather than
materialising the slice, dropping space to O(h); Day 47 already does exactly that.
