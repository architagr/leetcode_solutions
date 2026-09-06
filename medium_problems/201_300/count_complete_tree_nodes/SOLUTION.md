# Solution Walkthrough

The implementation is `countNodes(root *TreeNode) int` in `main.go`. We'll trace it on LeetCode's example:

![Example tree from LeetCode](images/example.jpg)

## Algorithm

1. **Base case.** If `root == nil`, there are no nodes, so return `0`.

2. **Recursively count right subtree.** `rightCount := countNodes(root.Right)` descends into the right child and counts all nodes in that subtree.

3. **Recursively count left subtree.** `leftCount := countNodes(root.Left)` descends into the left child and counts all nodes in that subtree.

4. **Combine.** Add the two subtree counts plus 1 for the current node: `return rightCount + leftCount + 1`.

## Execution Trace

Walking through `[1,2,3,4,5,6]`:

**Step 1:** `countNodes(1)` - start at root

![Step 1: Processing root node 1](images/walkthrough-1.png)

**Step 2:** Process right subtree `countNodes(3)` - node 3 has left child (6), returns `2`

![Step 2: Right subtree counted as 2 nodes](images/walkthrough-2.png)

**Step 3:** Process left subtree `countNodes(2)` - has two children (4, 5), returns `3`

**Step 4:** Combine at root: rightCount=2, leftCount=3, return `2 + 3 + 1 = 6` ✓

![Step 3: All nodes counted, total = 6](images/walkthrough-3.png)

## Complexity

- **Time:** O(n) — we visit every node in the worst case
- **Space:** O(h) where h is height — recursion call stack depth

Note: While this solution is correct and straightforward, a complete binary tree's structure allows for an O(log² n) solution by using binary search on the height and leveraging the tree's "completeness" property to skip entire subtrees.
