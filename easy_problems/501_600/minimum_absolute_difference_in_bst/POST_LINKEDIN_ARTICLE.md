## 365 Days of LeetCode Challenge — Day 11/365

# Minimum Absolute Difference in BST

🔗 https://leetcode.com/problems/minimum-absolute-difference-in-bst/ · Difficulty: Easy

### The problem

Given the root of a Binary Search Tree (BST), return the minimum absolute difference
between the values of any two different nodes in the tree.

### The intuition

The key BST property to lean on: an **in-order traversal of a BST visits nodes in
strictly increasing sorted order**. That's not a coincidence — it's the definition of
a BST (left subtree < node < right subtree), applied recursively.

Once you know the values come out sorted, the rest is a classic fact about sorted
arrays: the minimum absolute difference between *any* two elements always occurs
between two **adjacent** elements in sorted order. You never need to compare a value
against every other value in the tree — comparing it only to its immediate
predecessor is enough, because any non-adjacent pair's gap is at least as large as the
smallest adjacent gap between them.

So the algorithm becomes: do an in-order traversal, and as each node is visited,
compare its value to the value of the *previous* node visited (its in-order
predecessor) — not the previous one pushed onto a stack, and not its parent. Track the
smallest such gap seen so far, and that's the answer. This turns an O(n²) all-pairs
comparison into a single O(n) sweep, at the cost of remembering just one extra value
(the previous node) as you go.

### The solution

![Example 1](images/1.jpg "Example1")

```go
func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt
	// prev is the previously-visited node in in-order sequence, i.e. the
	// current node's in-order predecessor. nil until the first node is visited.
	var prev *TreeNode
	// helper performs a standard in-order traversal (left, node, right), which
	// for a BST visits values in strictly increasing sorted order. That means
	// the minimum gap between ANY two nodes must occur between two in-order
	// adjacent nodes, so comparing each node only to its immediate
	// predecessor is enough — no need to compare every pair.
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		helper(root.Left)

		// Skip the comparison on the very first node visited (no predecessor
		// yet). Since traversal order is sorted, root.Val >= prev.Val always,
		// so no abs() is needed.
		if prev != nil {
			res = min(res, root.Val-prev.Val)
		}
		// This node becomes the predecessor for whichever node is visited next.
		prev = root

		helper(root.Right)
	}

	helper(root)

	return res
}
```

Walking it through `root = [4,2,6,1,3]`, whose in-order sequence is `1, 2, 3, 4, 6`
(expected output `1`):

- Descend all the way left to node `1`. `prev` is still `nil`, so the comparison is
  skipped, and `prev` becomes `1`.

  ![Step 1: descend to leftmost node 1, prev initialized](images/walkthrough-1.svg)

- Back at node `2`: `prev` is `1`, so `res = min(∞, 2-1) = 1`, then `prev` becomes `2`.

  ![Step 2: at node 2, diff 2-1=1, res becomes 1](images/walkthrough-2.svg)

- At node `3`: `prev` is `2`, so `res = min(1, 3-2) = 1` (unchanged), then `prev`
  becomes `3`.

  ![Step 3: at node 3, diff 3-2=1, res stays 1](images/walkthrough-3.svg)

- Back up at the root, node `4`: `prev` is `3`, so `res = min(1, 4-3) = 1`
  (unchanged), then `prev` becomes `4`.

  ![Step 4: back at root 4, diff 4-3=1, res stays 1](images/walkthrough-4.svg)

- At node `6`: `prev` is `4`, so `res = min(1, 6-4) = min(1, 2) = 1` (unchanged), then
  `prev` becomes `6`. Traversal ends.

  ![Step 5: at node 6, diff 6-4=2, res stays 1 — final answer](images/walkthrough-5.svg)

`getMinimumDifference` returns `res = 1`. ✓

Notice the algorithm never explicitly sorts anything — the BST's in-order property
does that for free, so each node only ever needs to be compared against the single
node visited immediately before it.

**Complexity:** O(n) time — every node visited once. O(h) space for the recursion
stack, where h is the tree height (O(log n) balanced, O(n) skewed).

Full code: `easy_problems/501_600/minimum_absolute_difference_in_bst/` in the repo.
