## Solution walkthrough

The implementation is `BinaryTreeRightSideView(root *TreeNode) []int`, a wrapper over the
recursive `findRight`, in `binary_tree_right_side_view.go`.

![Example 1](images/1.png)

We'll trace `[1,2,3,null,5,null,4]`: root `1` with children `2` and `3`; `2` has a right
child `5`; `3` has a right child `4`. Expected output: `[1,3,4]`.

1. **The wrapper allocates and guards.** `arr := make([]int, 0, 100)` — capacity 100
   because the constraints cap the tree at 100 nodes, so the slice never reallocates. An
   empty tree returns that empty slice rather than nil, which is what the tests compare
   against.

2. **`arr` doubles as the output and the bookkeeping.** There's no visited set and no
   per-level buffer. `arr` holds one value per level filled so far, in order, which means
   `len(arr)` is always the index of the next level that hasn't been recorded.

3. **The guard records first arrivals only.**

   ```go
   if len(arr) == level {
       arr = append(arr, head.Val)
   }
   ```

   `len(arr) == level` is true exactly when this node is the first one reached at its
   depth. Any later node at the same depth sees `len(arr)` already past `level` and
   appends nothing.

   ![Step 1: the root is the first node at level 0](images/walkthrough-1.png)

4. **Right before left is what makes "first" mean "rightmost".**

   ```go
   arr = findRight(head.Right, arr, level+1)
   arr = findRight(head.Left, arr, level+1)
   ```

   This is the whole solution in one line-order decision. The guard stores whichever node
   arrives first at a depth; reversing the usual child order makes that node the rightmost
   one. Swap these two lines and the same function returns the left side view.

   ![Step 2: 3 reaches level 1 before 2 does, so 3 is kept](images/walkthrough-2.png)

5. **The right branch runs to the bottom first.** `4` is the first node to reach level 2,
   so it's recorded.

   ![Step 3: 4 is the first arrival at level 2](images/walkthrough-3.png)

6. **Then the left branch is walked, and mostly ignored.** `2` sits at level 1, but
   `len(arr)` is already 3, so nothing is appended.

   ![Step 4: 2 arrives at a level that is already filled](images/walkthrough-4.png)

   Same for `5` at level 2. The final answer is `[1,3,4]`.

   ![Step 5: 5 is skipped too, leaving [1,3,4]](images/walkthrough-5.png)

7. **Why the left subtree is still walked at all.** It looks wasteful on this example,
   since nothing from it is used. It isn't, because the rightmost node at a level need not
   live in the right subtree — the right branch can simply run out of depth first. The
   test case expecting `[1,3,6,7]` and the problem's own second example both have visible
   nodes coming off a left branch. The traversal handles that with no special case: the
   right subtree gets first refusal at every depth, and where it has nothing, the left
   subtree's node is the first arrival.

**Complexity:** O(n) time, every node visited once with constant work. Space is O(h) for
the recursion stack plus O(h) for the output, which holds one value per level. Against a
BFS solution, this trades peak memory proportional to the widest level for peak memory
proportional to the height.
