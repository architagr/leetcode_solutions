## Solution walkthrough

The implementation is split across `leafSimilar(root1, root2 *TreeNode) bool` (the
entry point) and `leafs(node *TreeNode) []int` (a recursive helper that extracts a
tree's leaf sequence), both in `main.go`.

We'll trace it on Example 1 from the problem statement:
`root1 = [3,5,1,6,2,9,8,null,null,7,4]`, `root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]`.

![Example 1](images/2.jpg "Example1")

1. **Collect both leaf sequences up front.**
   `root1Leafs := leafs(root1)` and `root2Leafs := leafs(root2)` — each call walks its
   tree independently and returns the leaf values it finds, in left-to-right order.

2. **Inside `leafs`, the base case handles an empty subtree.**
   `if node == nil { return []int{} }` — nothing to contribute.

3. **The leaf case.** `if node.Left == nil && node.Right == nil { return []int{node.Val} }`
   — a node with no children is a leaf, so it contributes exactly its own value as a
   one-element slice.

4. **The recursive case.** `return append(leafs(node.Left), leafs(node.Right)...)` —
   for an internal node, the leaf sequence is whatever `leafs` finds in the left
   subtree, followed by whatever it finds in the right subtree. Because the left call
   always happens (and is always fully evaluated, since Go evaluates arguments before
   `append` runs) before the right call's results are appended, the leaves come back in
   left-to-right order without any extra bookkeeping.

   Running `leafs` on `root1` visits `6`, `7`, `4` (from the left subtree rooted at
   `5`), then `9`, `8` (from the right subtree rooted at `1`):

   ![Step 1: root1's leaves collected left-to-right &#8594; (6, 7, 4, 9, 8)](images/walkthrough-1.png "Step 1")

   Running `leafs` on `root2` visits the same values, `6`, `7` (left subtree rooted at
   `5`), then `4`, `9`, `8` (right subtree rooted at `1`) — a differently shaped tree,
   but the same left-to-right leaf order:

   ![Step 2: root2's leaves collected left-to-right &#8594; (6, 7, 4, 9, 8)](images/walkthrough-2.png "Step 2")

5. **Back in `leafSimilar`, compare lengths first.**
   `if len(root1Leafs) != len(root2Leafs) { return false }` — a cheap early exit: two
   trees can't have the same leaf sequence if they don't even have the same number of
   leaves.

6. **Then compare values index by index.**
   ```go
   for i := 0; i < len(root1Leafs); i++ {
       if root1Leafs[i] != root2Leafs[i] {
           return false
       }
   }
   ```
   Any mismatch at any position means the sequences differ, so the trees aren't
   leaf-similar.

7. **If every index matched, the trees are leaf-similar.** `return true` falls through
   once the loop completes without finding a mismatch.

   `root1Leafs = [6, 7, 4, 9, 8]` and `root2Leafs = [6, 7, 4, 9, 8]` — same length, same
   values at every index, so `leafSimilar` returns `true`, matching the example:

   ![Step 3: comparing the two leaf sequences element by element &#8594; true](images/walkthrough-3.png "Step 3")

**Complexity:** O(n + m) time — every node of both trees is visited exactly once
(n and m are each tree's node count). O(n + m) space for the two leaf slices, plus
O(h1 + h2) for the recursion stacks.
