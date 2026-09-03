**365 Days of LeetCode Challenge — Day 21/365**
**Minimum Distance Between BST Nodes** (Easy)
🔗 https://leetcode.com/problems/minimum-distance-between-bst-nodes/

An in-order traversal of a BST visits values in strictly sorted order, so the minimum
difference between any two nodes can only show up between two values that land next to
each other once sorted. Traverse once to get that sorted list for free, then scan for
the smallest gap between neighbors. Kind of a fun one honestly, the whole solve is just
trusting that a BST already sorts itself for you.

![Example 1](images/1.jpg "Example1")

**Full solution:**
```go
func minDiffInBST(root *TreeNode) int {
	// No pair of nodes to compare in an empty tree or a single-node tree
	// (constraints guarantee n >= 2, so this mainly guards root == nil).
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	// In-order traversal of a BST visits every value in strictly increasing
	// order, so the minimum difference between ANY two nodes can only occur
	// between two values that end up adjacent here.
	arr := inOrder(root)
	minVal := math.MaxInt
	// Only adjacent pairs in the sorted list need checking.
	for i := 1; i < len(arr); i++ {
		minVal = min(minVal, abs(arr[i]-arr[i-1]))
	}
	return minVal
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

// inOrder returns this subtree's values in sorted order (left, node, right).
func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// Concatenate: left subtree values, then this node's value, then right
	// subtree values — the classic in-order recipe.
	return append(inOrder(root.Left), append([]int{root.Val}, inOrder(root.Right)...)...)
}
```

Walkthrough on `root = [4,2,6,1,3]` (expected `1`):

`arr := inOrder(root)` flattens the BST into sorted order `[1, 2, 3, 4, 6]`:

![Step 1: in-order traversal collects [1, 2, 3, 4, 6] from the tree](images/walkthrough-1.svg)

The loop scans adjacent gaps `(1,2)=1`, `(2,3)=1`, `(3,4)=1`, `(4,6)=2`, and keeps the
smallest:

![Step 2: scanning adjacent gaps 1, 1, 1, 2 — minVal settles at 1](images/walkthrough-2.svg)

`minVal` gets returned:

![Step 3: minDiffInBST(root) returns 1](images/walkthrough-3.svg)

The traversal is O(n), though the nested `append`s inside `inOrder` mean it can degrade
toward O(n²) on a skewed tree instead of staying linear. The scan afterward is a clean
O(n). Space is O(n) for the array plus O(h) for the recursion stack, h being the tree's
height.
