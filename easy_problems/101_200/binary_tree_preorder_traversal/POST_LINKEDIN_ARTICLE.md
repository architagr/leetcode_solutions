## 365 Days of LeetCode Challenge — Day 8/365

# Binary Tree Preorder Traversal

🔗 https://leetcode.com/problems/binary-tree-preorder-traversal/ · Difficulty: Easy

### The problem

Given the root of a binary tree, return the preorder traversal of its nodes'
values — visit each node before either of its subtrees.

### The intuition

"Preorder" just names the order in which the three things at each node get
visited: the node **itself** first, then its **left** subtree, then its **right**
subtree (root → left → right). That ordering is the entire problem — there's no
searching, comparing, or bookkeeping beyond "visit this node, then recurse left,
then recurse right."

The natural way to express that is recursion that mirrors the definition
directly: a helper function visits a node by appending its value, then calls
itself on `Left`, then calls itself on `Right`. The recursion's call stack does
the traversal's bookkeeping for free — there's no need for an explicit stack,
since the position in the recursive calls already encodes "where am I in the
tree, and what's left to visit."

The one wrinkle in Go specifically is that `append` can reallocate the
underlying array, so the slice being built has to be threaded through as both an
argument and a return value — each recursive call returns the (possibly grown)
slice so the caller picks up wherever the callee left off, rather than each call
silently mutating a slice the caller no longer has a valid reference to.

The base case is the empty subtree (`nil`): visiting nothing contributes
nothing, so the accumulated slice is simply handed back unchanged.

### The solution

![Example 1](images/1.png "Example1")

```go
func PreorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	return traversal(root, arr)
}

// traversal visits nodes in root -> left -> right order, threading arr through
// as both a parameter and a return value. That threading matters in Go: append
// can reallocate arr's backing array, so a caller must capture the return value
// to see every element a callee appended, rather than relying on shared mutation.
func traversal(A *TreeNode, arr []int) []int {
	// Base case: an empty subtree contributes nothing, so hand arr back untouched.
	if A == nil {
		return arr
	}
	// "Pre"-order: record the current node's value before touching either subtree.
	arr = append(arr, A.Val)
	arr = traversal(A.Left, arr)

	arr = traversal(A.Right, arr)
	return arr
}
```

Walking it through `[1,null,2,3]` (node `1` has no left child, right child `2`;
`2` has left child `3`; expected output `[1,2,3]`):

- `traversal(1)` appends `1` → `arr = [1]`.

![Step 1: visit root 1, arr becomes [1]](images/walkthrough-1.svg)

- `1.Left` is `nil` — immediate base case, no change. `1.Right` is `2`, so
  `traversal(2)` appends `2` → `arr = [1,2]`.

![Step 2: 1.Left is nil, visit 2, arr becomes [1,2]](images/walkthrough-2.svg)

- `2.Left` is `3`, so `traversal(3)` appends `3` → `arr = [1,2,3]`.

![Step 3: visit 3 via 2.Left, arr becomes [1,2,3]](images/walkthrough-3.svg)

- `3` has no children — both its calls hit the base case. `[1,2,3]` unwinds back
  through `3`, `2` (whose `Right` is also `nil`), and `1`, all the way out. ✓

![Step 4: 3's children are nil, unwind to return [1,2,3]](images/walkthrough-4.svg)

**Complexity:** O(n) time — every node visited once, O(1) work each. O(h) space
for the recursion stack (h = tree height), plus O(n) for the output slice.

Full code: `easy_problems/101_200/binary_tree_preorder_traversal/` in the repo.
