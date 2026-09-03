**365 Days of LeetCode Challenge — Day 13/365**
**Binary Tree Tilt** (Easy)
🔗 https://leetcode.com/problems/binary-tree-tilt/

**Intuition:** Each node's tilt is the absolute difference between everything under
its left child and everything under its right child, not just the two direct kids.
Redoing that sum from scratch at every node wastes work, so compute each subtree's
sum once, bottom-up (postorder), and add every node's tilt into a shared running
total as you go. The fun bit is that one recursive call does both of those jobs at
the same time, accumulate and return, without them getting tangled up.

![Example 2](images/2.jpg "Example2")

**Full solution:**
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

**Walkthrough** on `[4,2,9,3,5,null,7]` (expected `15`):

- Leaves `3` and `5` are the base case: `l=0, r=0`, `tilt+=0`, each just returns its
  own value.

![Step 1: leaves 3, 5, 7 are base cases — each returns its own Val, tilt+=0](images/walkthrough-1.svg)

- `sum(2)`: `l=3, r=5` → `tilt += |3-5| = 2` (running total `2`) → returns `10`

![Step 2: node 2 resolves — l=3, r=5, tilt+=2, returns 10](images/walkthrough-2.svg)

- `sum(9)`: no left child → `l=0`; `r=sum(7)=7` → `tilt += |0-7| = 7` (running total
  `9`) → returns `16`

![Step 3: node 9 resolves — l=0 (no left child), r=7, tilt+=7, returns 16](images/walkthrough-3.svg)

- `sum(4)` (root): `l=10, r=16` → `tilt += |10-16| = 6` (running total `15`) →
  returns `30`. `findTilt` returns `15` ✓

![Step 4: node 4 resolves — l=10, r=16, tilt+=6, res=15](images/walkthrough-4.svg)

O(n) time, since each node is visited once, O(h) space for the recursion stack,
where h is the tree's height.
