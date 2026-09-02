## 365 Days of LeetCode Challenge — Day 21/365

# Minimum Distance Between BST Nodes

🔗 https://leetcode.com/problems/minimum-distance-between-bst-nodes/ · Difficulty: Easy

### The problem

Given the root of a Binary Search Tree (BST), return the minimum difference between the
values of any two different nodes in the tree.

### The intuition

The key property to lean on is what makes a tree a **BST** in the first place: an
in-order traversal (left, node, right) visits every value in **strictly sorted order**.

That one fact turns "find the minimum difference between *any* two nodes" — which sounds
like it could require checking every pair, an O(n²) affair — into something much
simpler. Once the values are sorted, the smallest possible difference between any two of
them can only ever occur between two values that are **adjacent** in that sorted order.
Any pair that skips over a value in between can't beat the gap between neighbors,
because the skipped value sits strictly between them and splits that gap into two
smaller (or equal) pieces.

So the whole problem reduces to two steps:
1. Collect every node's value via an in-order traversal — this hands back an
   already-sorted list, no separate sort needed.
2. Walk that sorted list once, tracking the smallest gap between consecutive entries.

### The solution

![Example 1](images/1.jpg "Example1")

We'll trace it on `root = [4,2,6,1,3]` (expected `1`).

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

**Step 1 — flatten the tree into a sorted list.** `arr := inOrder(root)` walks the tree
left, node, right, so on `[4,2,6,1,3]` it comes back as `[1, 2, 3, 4, 6]` — already
sorted, for free, because the tree is a BST.

![Step 1: in-order traversal collects [1, 2, 3, 4, 6] from the tree](images/walkthrough-1.svg)

**Step 2 — scan for the smallest neighboring gap.** `minVal` starts at `math.MaxInt` so
the first comparison always wins. The loop then walks the sorted array one adjacent pair
at a time — `(1,2)`, `(2,3)`, `(3,4)`, `(4,6)` — taking `abs` of each difference and
keeping the smallest. Because the minimum difference between any two BST values can only
occur between values adjacent once sorted, this single linear pass is enough.

![Step 2: scanning adjacent gaps 1, 1, 1, 2 — minVal settles at 1](images/walkthrough-2.svg)

**Step 3 — return the answer.** After the loop, `minVal` holds the smallest gap found
across the whole sorted sequence.

![Step 3: minDiffInBST(root) returns 1](images/walkthrough-3.svg)

**Complexity:** the traversal visits each of the `n` nodes once, but because `inOrder`
rebuilds and copies a slice at every recursive call via nested `append`s, the total
copying work can degrade toward O(n²) on a heavily skewed tree rather than staying
linear (a straightforward accumulator-based traversal would avoid this). The final scan
over the sorted array is a clean O(n). Space is O(n) for the collected slice plus O(h)
for the recursion stack, where `h` is the tree's height.

Full code: `easy_problems/701_800/minimum_distance_between_bst_nodes/` in the repo.
