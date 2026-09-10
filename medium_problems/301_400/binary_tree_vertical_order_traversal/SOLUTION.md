## Solution walkthrough

The implementation is `verticalOrder(root *TreeNode) [][]int` in `main.go`, a BFS over a
queue of `customNode` values that pair a node with its column index.

![Example 1](images/1.png)

We'll trace `[3,9,20,null,null,15,7]`: root `3` with children `9` and `20`; `20` has
children `15` and `7`. Expected output `[[9],[3,15],[20],[7]]`.

1. **Column index is a coordinate carried down.** The root is `0`, a left child is
   `parent - 1`, a right child is `parent + 1`. `customNode` embeds a `TreeNode` and adds
   `order`, so the queue carries the pairing.

2. **`push` does two jobs at once.**

   ```go
   push := func(n *TreeNode, o int) {
       l, ok := m[o]
       if !ok {
           l = make([]int, 0)
       }
       l = append(l, n.Val)
       m[o] = l
       q = append(q, &customNode{TreeNode: *n, order: o})
   }
   ```

   It records the value into its column's list *and* enqueues the node for later
   expansion. So the map is built during the push rather than during the pop, and a node's
   value lands in its column at the moment it is discovered.

   ![Step 1: the root enters at column 0](images/walkthrough-1.png)

3. **The BFS itself only expands children.** Popping a node pushes its left child at
   `order-1` and its right at `order+1`. Nothing is recorded on pop, because the push
   already did it.

   ![Step 2: BFS drains depth 1 before depth 2](images/walkthrough-2.png)

4. **This is the step that requires BFS.** `15` sits in column `0`, the same column as the
   root, and it must appear *after* the root because it's lower down. A breadth-first walk
   reaches `3` at depth 0 and `15` at depth 2, in that order, so appending as it goes
   produces `[3, 15]` correctly with no sorting.

   ![Step 3: 15 joins column 0 after the root](images/walkthrough-3.png)

   A depth-first walk would break this. It would drive one branch to the bottom before
   touching the other, so a deep node from the left subtree could be appended to a column
   before a shallower node from the right subtree that belongs above it. Same grouping,
   wrong order within it — which is exactly the difference from Day 15, where order within
   a group never depended on arrival.

5. **Assembly walks the columns in order.**

   ```go
   for i := -101; i <= 101; i++ {
       if l, ok := m[i]; ok {
           result = append(result, l)
       }
   }
   ```

   This is the one shortcut in the solution worth naming. It leans on the constraint that
   the tree holds at most 100 nodes, so no column index can fall outside `[-100, 100]`.
   Collecting the map's keys and sorting them would be the version that doesn't depend on
   the constraint holding.

   ![Step 4: reading the columns left to right](images/walkthrough-4.png)

6. **A small note on `customNode`.** It embeds `TreeNode` by value, so `&customNode{TreeNode:
   *n, ...}` copies the node. That's harmless here because only `Left`, `Right` and `Val`
   are read from the copy and the pointers still refer to the real children — but embedding
   a `*TreeNode` would avoid the copy.

**Complexity:** O(n) for the traversal — every node is pushed and popped once. The assembly
loop is a fixed 203 iterations, O(1) under the constraint; the sorted-keys version would be
O(c log c) in the number of columns. Space is O(n) for the map and the queue.
