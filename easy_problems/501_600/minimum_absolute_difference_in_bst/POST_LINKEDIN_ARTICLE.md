## 365 Days of LeetCode Challenge — Day 11/365

# Minimum Absolute Difference in BST

🔗 https://leetcode.com/problems/minimum-absolute-difference-in-bst/ · Difficulty: Easy

### The problem

Given the root of a Binary Search Tree (BST), return the minimum absolute difference
between the values of any two different nodes in the tree.

### The intuition

The whole trick rides on one BST fact: walk the tree in-order (left, node, right) and
the values come out in sorted order. That's not a clever discovery so much as just
what the BST rule means once you apply it recursively at every node: left subtree
smaller, right subtree bigger, all the way down.

Once you know you're dealing with sorted values, an old fact about sorted arrays does
the rest of the work for you: the smallest gap between any two values always sits
between two neighbors. You don't need to check a value against every other value in
the tree. Checking it against whatever came right before it is enough, because any gap
between two non-neighbors can only be as large or larger than the smallest neighboring
gap between them.

So the algorithm mostly falls out on its own. Traverse in-order, and at each node
compare its value to the value of whichever node was visited right before it, its
in-order predecessor, not the parent, not whatever happens to be sitting on a stack.
Keep the smallest gap seen so far and that's your answer. What I like about this one
is how it turns an O(n²) all-pairs comparison into a single O(n) pass, and the only
cost is remembering one extra pointer as you go.

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

- Descend all the way left to node `1`. `prev` is still `nil`, so the comparison gets
  skipped, and `prev` becomes `1`.

  ![Step 1: descend to leftmost node 1, prev initialized](images/walkthrough-1.svg)

- Back at node `2`: `prev` is `1`, so `res = min(∞, 2-1) = 1`, then `prev` becomes `2`.

  ![Step 2: at node 2, diff 2-1=1, res becomes 1](images/walkthrough-2.svg)

- At node `3`: `prev` is `2`, so `res = min(1, 3-2) = 1` (no change), then `prev`
  becomes `3`.

  ![Step 3: at node 3, diff 3-2=1, res stays 1](images/walkthrough-3.svg)

- Back up at the root, node `4`: `prev` is `3`, so `res = min(1, 4-3) = 1` (no change),
  then `prev` becomes `4`.

  ![Step 4: back at root 4, diff 4-3=1, res stays 1](images/walkthrough-4.svg)

- At node `6`: `prev` is `4`, so `res = min(1, 6-4) = min(1, 2) = 1` (no change), then
  `prev` becomes `6`. Traversal ends.

  ![Step 5: at node 6, diff 6-4=2, res stays 1 — final answer](images/walkthrough-5.svg)

`getMinimumDifference` returns `res = 1`. ✓

Nothing in here explicitly sorts anything. The BST's in-order property does that job
for free, so each node only ever has to be checked against the one node that came
right before it. That's the part I find satisfying about this problem, the sorting is
basically hiding in plain sight in the traversal order and you don't have to go
looking for it.

This runs in O(n) time since every node gets visited exactly once, and O(h) space for
the recursion stack, where h is the tree's height. That's O(log n) for a balanced
tree, O(n) if the tree is a straight line.

Full code: `easy_problems/501_600/minimum_absolute_difference_in_bst/` in the repo.

#LeetCode #100DaysOfCode #DSA #CodingInterview #BinarySearchTree #InOrderTraversal #Golang

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
