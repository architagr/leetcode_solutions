## 365 Days of LeetCode Challenge — Day 9/365

# Binary Tree Postorder Traversal

🔗 https://leetcode.com/problems/binary-tree-postorder-traversal/ · Difficulty: Easy

### The problem

Given the root of a binary tree, return the postorder traversal of its nodes'
values — left subtree, then right subtree, then the node itself.

### The intuition

"Postorder" just names the order in which a node gets recorded relative to its
children: **left subtree, then right subtree, then the node itself**. So a node's
value is only ever appended to the result *after* everything underneath it has
already been appended — a node is always the last thing written among itself and
its descendants.

That ordering falls out naturally from a simple recursive shape: recurse into the
left child, recurse into the right child, then append the current node's value.
There's no need to track any extra state (like "have I visited this node's
children yet") — the recursion itself guarantees the children are fully processed,
and their values are already sitting in the result slice, before the current
node's `append` ever runs.

The only base case is an empty subtree (`nil`), which simply contributes nothing
and returns the result slice unchanged — that's what stops the recursion from
following `nil` children and is also why leaves resolve immediately: both of a
leaf's recursive calls hit this base case right away, so the very next line
appends the leaf's own value.

### The solution

![Example 1](images/1.png "Example1")

```go
func PostorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	return traversal(root, arr)
}

func traversal(A *TreeNode, arr []int) []int {
	if A == nil {
		return arr
	}

	arr = traversal(A.Left, arr)
	arr = traversal(A.Right, arr)
	arr = append(arr, A.Val)
	return arr
}
```

`PostorderTraversal` just seeds an empty slice and hands it off to `traversal`,
which does the real work and threads the growing result slice through every
recursive call: recurse left, recurse right, then append the current node last.

Walking `traversal(1, [])` on `root = [1,null,2,3]` (`1`'s left child is `nil`,
its right child is `2`, and `2`'s left child is `3`), expected output `[3,2,1]`:

- `traversal(1, [])` recurses left into `nil` → returns `[]` unchanged, then
  recurses right into `2`.
  - `traversal(2, [])` recurses left into `3`.
    - `traversal(3, [])` has both children `nil`, so it bottoms out immediately
      and appends its own value → `arr = [3]`.

    ![Step 1: node 3 is a leaf — both children hit the nil base case, so it appends itself first](images/walkthrough-1.svg)

  - Back in `traversal(2, [])`: left came back `[3]`, right (`nil`) adds
    nothing. Both sides done → appends `2` → `arr = [3, 2]`.

    ![Step 2: node 2's left (3) and right (nil) are both resolved, so it appends itself next](images/walkthrough-2.svg)

- Back in `traversal(1, [])`: left (`nil`) contributed nothing, right came back
  `[3, 2]`. Both sides done → appends `1` → `arr = [3, 2, 1]`.

  ![Step 3: node 1's left (nil) and right (2) are both resolved, so it appends itself last](images/walkthrough-3.svg)

`PostorderTraversal` returns `[3, 2, 1]`. ✓

**Complexity:** O(n) time — every node is visited exactly once. O(h) space for the
recursion stack, where h is the tree's height (O(log n) balanced, O(n) skewed),
plus O(n) for the output slice itself.

**Follow up worth thinking about:** LeetCode asks whether you can do this
iteratively instead of recursively — a common approach is a single stack that
pushes nodes and prepends (rather than appends) their values, effectively building
the reverse of a "root, right, left" traversal.

Full code: `easy_problems/101_200/binary_tree_postorder_traversal/` in the repo.
