# Solution Walkthrough

The implementation is `countNodes(root *TreeNode) int` in `main.go`. We'll trace it on LeetCode's example:

![Example tree from LeetCode](images/example.jpg)

1. **Base case.** If `root == nil`, there are no nodes, so return `0`.

2. **Recursively count right subtree.** `rightCount := countNodes(root.Right)` descends into the right child and counts all nodes in that subtree.

3. **Recursively count left subtree.** `leftCount := countNodes(root.Left)` descends into the left child and counts all nodes in that subtree.

4. **Combine.** Add the two subtree counts plus 1 for the current node: `return rightCount + leftCount + 1`.

Walking through `[1,2,3,4,5,6]`:

- `countNodes(1)`: Call on root node 1.

  ![Step 1: Processing root node 1](images/walkthrough-1.png)

- Recursively process `root.Right` (node 3): node 3 has no children, so `countNodes(3)` returns `0 + 0 + 1 = 1`.

  ![Step 2: Right subtree (node 3) counted as 1 node](images/walkthrough-2.png)

- Recursively process `root.Left` (node 2):
  - `countNodes(2)` calls `countNodes(2.Right)` (node 5) → returns `1`
  - Then calls `countNodes(2.Left)` (node 4) → returns `1`
  - Returns `1 + 1 + 1 = 3`

- Back at root: combine all counts.
  - `rightCount = 1` (from node 3)
  - `leftCount = 3` (from subtree rooted at node 2)
  - return `1 + 3 + 1 = 5`

Wait, the tree has 6 nodes, not 5. Let me retrace: the tree is `[1,2,3,4,5,6]`:

```
      1
     / \
    2   3
   / \
  4   5
      /
     6
```

Actually, in level-order array representation, index 5 (value 6) is the left child of index 2 (value 5). So:

```
      1
     / \
    2   3
   / \
  4   5
```

That's 5 nodes. The sixth node would be at index 6. The array `[1,2,3,4,5,6]` gives:

```
      1
     / \
    2   3
   / \ /
  4  5 6
```

So the right subtree of 3 has one child (6):
- `countNodes(3)`: `countNodes(3.Left)` (node 6) → `1`, `countNodes(3.Right)` (nil) → `0`, return `1 + 0 + 1 = 2`.
- `countNodes(1)`: `rightCount = 2`, `leftCount = 3`, return `2 + 3 + 1 = 6` ✓

  ![Step 3: All nodes counted, total = 6](images/walkthrough-3.png)

## Complexity

- **Time:** O(n) — we visit every node in the worst case
- **Space:** O(h) where h is height — recursion call stack depth

Note: While this solution is correct, a complete binary tree's structure allows for an O(log² n) solution by using binary search on the height and leveraging the tree's "completeness" property.
