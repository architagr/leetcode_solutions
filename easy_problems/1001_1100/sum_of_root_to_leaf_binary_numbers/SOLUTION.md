## Solution walkthrough

The implementation is `sumRootToLeaf(root *TreeNode) int` in `main.go`, with the actual
recursion delegated to a helper `sum(root *TreeNode, rootLevelSum int) int`.

![Example 1](images/1.png "Example1")

We'll trace it on the example above, `root = [1,0,1,0,1,0,1]`, expected output `22`.

1. **Entry point.** `sumRootToLeaf` just kicks off the recursion with an empty running
   value: `return sum(root, 0)`. `rootLevelSum` is the binary number built from the
   root down to (but not including) the current node.

2. **Base case.** `if root == nil { return 0 }` — an empty subtree contributes nothing,
   and also protects the top-level call in case `root` itself is `nil`.

3. **Extend the running value by one bit.** `rootLevelSum *= 2` shifts every bit already
   accumulated one place to the left, making room for the current node's bit;
   `currentSum := rootLevelSum + root.Val` drops that bit (`0` or `1`) into the newly
   opened units place. This is exactly how a binary number grows digit by digit as you
   read it left to right.

4. **Leaf check — stop and report.** `if isLeafNode(root) { return currentSum }` — once
   a leaf is reached, `currentSum` already *is* the complete binary number for this
   root-to-leaf path, so it's returned as-is. `isLeafNode` (`node.Left == nil &&
   node.Right == nil`) is a small named helper rather than an inline check, which keeps
   the leaf condition self-documenting at the call site.

5. **Otherwise, recurse into both children and add their results.**
   `leftSum := sum(root.Left, currentSum)` and `rightSum := sum(root.Right, currentSum)`
   pass the extended `currentSum` down as the new `rootLevelSum` for each child — every
   node below inherits the bits accumulated so far. `return leftSum + rightSum` combines
   whatever the two subtrees contribute; internal nodes never sum a value for themselves,
   only for their leaves.

Walking it through `root = [1,0,1,0,1,0,1]` (root=1, left=0, right=1, leftleft=0,
leftright=1, rightleft=0, rightright=1) — four leaves, four root-to-leaf binary numbers:

- `sum(root=1, 0)`: `rootLevelSum` becomes `1`, `currentSum = 1` ("bits=1"). Not a leaf,
  so recurse into both children.
  - `sum(left=0, 1)`: `rootLevelSum` becomes `10b = 2`, `currentSum = 2` ("bits=10").
    Not a leaf, recurse further.
    - `sum(leftleft=0, 2)`: `currentSum = 100b = 4`. `leftleft` is a leaf → returns `4`.

    ![Step 1: first path root → 0 → 0 reaches a leaf, currentSum = 100b = 4](images/walkthrough-1.png)

    - `sum(leftright=1, 2)`: `currentSum = 101b = 5`. `leftright` is a leaf → returns
      `5`. Back in the `left` call: `leftSum = 4 + 5 = 9`.

    ![Step 2: second path root → 0 → 1 reaches a leaf, currentSum = 101b = 5, left subtree total so far = 9](images/walkthrough-2.png)

  - `sum(right=1, 1)`: `rootLevelSum` becomes `11b = 3`, `currentSum = 3` ("bits=11").
    Not a leaf, recurse further.
    - `sum(rightleft=0, 3)`: `currentSum = 110b = 6`. `rightleft` is a leaf → returns
      `6`.

    ![Step 3: third path root → 1 → 0 reaches a leaf, currentSum = 110b = 6; left subtree already totals 9](images/walkthrough-3.png)

    - `sum(rightright=1, 3)`: `currentSum = 111b = 7`. `rightright` is a leaf → returns
      `7`. Back in the `right` call: `rightSum = 6 + 7 = 13`.

    ![Step 4: fourth path root → 1 → 1 reaches a leaf, currentSum = 111b = 7; all four leaves now computed](images/walkthrough-4.png)

  - Back in the top-level `sum(root, 0)` call: `leftSum = 9` and `rightSum = 13` have
    both returned, so `sum` returns `leftSum + rightSum = 9 + 13 = 22`. ✓

  ![Step 5: recursion unwinds, leftSum=9 and rightSum=13 bubble up to root, total = 22](images/walkthrough-5.png)

**Complexity:** O(n) time — every node visited once. O(h) space for the recursion
stack, where h is the tree height.
