## 365 Days of LeetCode Challenge — Day 13/365

# Binary Tree Tilt

🔗 https://leetcode.com/problems/binary-tree-tilt/ · Difficulty: Easy

### The problem

Given the root of a binary tree, return the sum of every node's **tilt** — the
absolute difference between the sum of all values in its left subtree and the sum of
all values in its right subtree (a missing child counts as a subtree sum of `0`).

### The intuition

Computing one node's tilt needs the sum of *every* value in its left subtree and
*every* value in its right subtree — not just its immediate children. Re-summing each
subtree from scratch at every node would be wasteful: the same lower subtree's sum
would get recomputed over and over as you climb back up the tree.

The fix is to compute each subtree's sum exactly once, bottom-up, and hand it back to
the caller. That's a **postorder** traversal: visit the left subtree, visit the right
subtree, *then* do work at the current node — because the current node's tilt (and its
own contribution to its parent's subtree sum) depends on both children's totals
already being known.

So the recursive helper does two jobs on every call: it accumulates into a running
tilt total shared by every call (via a pointer to an int), and it returns the sum of
the subtree rooted at the current node so the parent can use it. Each node's tilt is
simply the absolute difference between what its left recursive call returned and what
its right recursive call returned.

### The solution

![Example 2](images/2.jpg "Example2")

```go
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// res is shared across every recursive call via pointer, so each node
	// can add its own tilt directly into one running total instead of
	// returning a (subtreeSum, tiltSum) pair that would need merging.
	res := 0
	sum(root, &res)
	return res
}

// sum does double duty: it accumulates every node's tilt into *res as a
// side effect, and it returns the sum of the subtree rooted at node so the
// caller (the parent) can use it when computing its own tilt.
func sum(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	// Postorder: both children must be fully resolved before this node's
	// tilt (which depends on both subtree sums) can be computed.
	l := sum(node.Left, res)
	r := sum(node.Right, res)
	// This node's tilt is the absolute difference between its left and
	// right subtree sums; add it straight into the shared accumulator.
	*res += absDiff(l, r)
	// Hand the parent this whole subtree's total as a single number.
	return l + r + node.Val
}

func absDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff
}
```

Walking it through `root = [4,2,9,3,5,null,7]` (expected `15`):

- `sum(2)` recurses into leaves `3` and `5`: `l=0, r=0` for each, `*res += |0-0| = 0`,
  each returns its own value.

  ![Step 1: leaves 3, 5, 7 are base cases — each returns its own Val, tilt+=0](images/walkthrough-1.svg)

- Back in `sum(2)`: `l=3, r=5`. `*res += |3-5| = 2` (running total `2`). Returns
  `3 + 5 + 2 = 10`.

  ![Step 2: node 2 resolves — l=3, r=5, tilt+=2, returns 10](images/walkthrough-2.svg)

- `sum(9)`: no left child so `l=0`; `sum(7)` (a leaf) gives `r=7`.
  `*res += |0-7| = 7` (running total `9`). Returns `0 + 7 + 9 = 16`.

  ![Step 3: node 9 resolves — l=0 (no left child), r=7, tilt+=7, returns 16](images/walkthrough-3.svg)

- Back in `sum(4)` (the root): `l=10, r=16`. `*res += |10-16| = 6` (running total
  `15`). Returns `10 + 16 + 4 = 30`. `findTilt` returns `res = 15`. ✓

  ![Step 4: node 4 resolves — l=10, r=16, tilt+=6, res=15](images/walkthrough-4.svg)

**Complexity:** O(n) time — every node is visited exactly once, doing O(1) work.
O(h) space for the recursion stack, where h is the tree's height (O(log n) balanced,
O(n) fully skewed).

Full code: `easy_problems/501_600/binary_tree_tilt/` in the repo.
