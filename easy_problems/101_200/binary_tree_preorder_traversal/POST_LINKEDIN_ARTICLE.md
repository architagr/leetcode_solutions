## 365 Days of LeetCode Challenge — Day 8/365

# Binary Tree Preorder Traversal

🔗 https://leetcode.com/problems/binary-tree-preorder-traversal/ · Difficulty: Easy

### The problem

Given the root of a binary tree, return the values in preorder: each node
before either of its subtrees.

### The intuition

"Preorder" tells you everything up front: visit the node first, then the left
subtree, then the right one. Root, left, right. There's no searching or
comparing hidden anywhere in this problem, just that one ordering rule applied
at every node.

Recursion is the obvious tool here because it mirrors the definition almost
word for word. A helper function visits a node by appending its value, then
calls itself on `Left`, then calls itself on `Right`. You don't need an
explicit stack for this — the call stack already tracks where you are in the
tree and what's left to visit. That's the part I actually like about tree
recursion: the bookkeeping is free, you just have to trust it's happening.

The thing that trips people up in Go specifically is `append`. It can
reallocate the underlying array, so you can't just mutate a slice and expect
the caller to see the change. You have to thread it through as both a
parameter and a return value, and every recursive call has to capture what
comes back rather than assume it happened by reference.

The base case is the empty subtree. Visit nothing, contribute nothing, hand
the slice back exactly as it came in.

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

Here's the trace on `[1,null,2,3]`: node `1` has no left child, its right
child is `2`, and `2` has a left child `3`. Expected output is `[1,2,3]`.

`traversal(1)` appends `1`, so `arr = [1]`.

![Step 1: visit root 1, arr becomes [1]](images/walkthrough-1.png)

`1.Left` is `nil`, so that call hits the base case immediately and does
nothing. `1.Right` is `2`, so `traversal(2)` appends `2` and `arr` becomes
`[1,2]`.

![Step 2: 1.Left is nil, visit 2, arr becomes [1,2]](images/walkthrough-2.png)

`2.Left` is `3`, so `traversal(3)` appends `3`. `arr` is now `[1,2,3]`.

![Step 3: visit 3 via 2.Left, arr becomes [1,2,3]](images/walkthrough-3.png)

`3` has no children, so both of its calls hit the base case and return
without touching anything. The `[1,2,3]` slice unwinds back up through `3`,
then `2` (whose `Right` is also `nil`), then `1`, and out. That's the whole
trace.

![Step 4: 3's children are nil, unwind to return [1,2,3]](images/walkthrough-4.png)

O(n) time, since every node gets visited once and does O(1) work. O(h) space
for the recursion stack, where h is the tree's height, plus O(n) for the
output slice itself.

The code lives at `easy_problems/101_200/binary_tree_preorder_traversal/` in
the repo.

#LeetCode #100DaysOfCode #CodingInterview #Algorithms #BinaryTree #Recursion #Golang

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
