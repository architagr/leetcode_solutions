## Solution walkthrough

The implementation is `PreorderTraversal(root *TreeNode) []int` plus its recursive
helper `traversal` in `binary_tree_preorder_traversal.go`.

![Example 1](images/1.png "Example1")

We'll trace it on the example above, `root = [1,null,2,3]` — node `1` has no left
child and a right child `2`, and `2` has a left child `3`.

1. **Entry point.** `PreorderTraversal` allocates an empty output slice,
   `arr := make([]int, 0)`, and hands it plus `root` to `traversal`. This keeps the
   public signature clean (`PreorderTraversal(root) []int`) while the recursive
   helper carries the extra `arr` parameter it needs to accumulate into.

2. **Base case.** `if A == nil { return arr }` — an empty subtree contributes
   nothing, so the accumulated slice is returned unchanged. This is what stops the
   recursion at the bottom of every branch.

3. **Visit the node itself, first.** `arr = append(arr, A.Val)` — this is the "pre"
   in preorder: the current node's value is recorded *before* either subtree is
   touched.

   ![Step 1: visit root 1, arr becomes [1]](images/walkthrough-1.svg)

4. **Recurse left.** `arr = traversal(A.Left, arr)` — the (possibly grown) slice is
   passed in and the (possibly grown again) slice is captured back out. Threading
   `arr` through the return value, rather than relying on the callee mutating a
   shared slice header, matters in Go because `append` can reallocate the backing
   array; capturing the return value is what guarantees the caller sees every
   element the callee added. Here, `1.Left` is `nil`, so this call immediately hits
   the base case and hands `arr` straight back unchanged.

5. **Recurse right.** `arr = traversal(A.Right, arr)` — same pattern, now on the
   right subtree. `1.Right` is `2`, so this call visits `2` (`arr = [1,2]`), then
   recurses into *its* children.

   ![Step 2: 1.Left is nil, visit 2, arr becomes [1,2]](images/walkthrough-2.svg)

   `2.Left` is `3`, so that recursive call visits `3` next (`arr = [1,2,3]`).

   ![Step 3: visit 3 via 2.Left, arr becomes [1,2,3]](images/walkthrough-3.svg)

6. **Unwind.** `3` has no children, so both of its recursive calls hit the base
   case and return `arr` untouched. That `[1,2,3]` bubbles back up through `3`'s
   call, then through `2`'s call (whose own `Right` is `nil`, another base case),
   then through `1`'s call, and finally out of `PreorderTraversal`. `return arr` at
   the end of `traversal` is what carries the final slice back through every frame
   on the stack.

   ![Step 4: 3's children are nil, unwind to return [1,2,3]](images/walkthrough-4.svg)

The result, `[1, 2, 3]`, matches the expected output — root before left before
right, all the way down.

**Complexity:** O(n) time — every node is visited exactly once, each doing O(1)
work. O(h) space for the recursion stack (h = tree height), plus O(n) for the
output slice.
