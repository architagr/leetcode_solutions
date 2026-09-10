---
meta_title: "Why preorder traversal in Go needs a return value"
meta_description: "The recursion is three lines. The catch is append — it can reallocate, so the slice must be threaded through as parameter and return."
tags: [golang, binary-tree, recursion, dsa, leetcode]
---

# Binary Tree Preorder Traversal

*365 Days of LeetCode Challenge — Day 8/365*

🔗 [LeetCode #144](https://leetcode.com/problems/binary-tree-preorder-traversal/) · Difficulty: Easy

Most of the tree problems in this series are interesting because of the tree. This one
isn't. The traversal is the easiest thing you'll write all week — three lines that read
like the definition of the word "preorder."

What makes it worth a newsletter edition is the second half, which has nothing to do
with trees at all. It's a Go problem, and it's the kind that doesn't announce itself:
your code compiles, runs, returns a slice, and quietly drops nodes.

### The problem

Given the root of a binary tree, return the values in preorder — each node before either
of its subtrees.

### The traversal, which is the easy part

"Preorder" tells you the whole algorithm up front: visit the node, then the left subtree,
then the right one. Root, left, right. Nothing is being searched for or compared. There's
one ordering rule and it gets applied at every node.

Recursion mirrors that almost word for word. A helper visits a node by appending its
value, then calls itself on `Left`, then on `Right`. You don't need an explicit stack —
the call stack already knows where you are in the tree and what's left to visit. That's
the part I genuinely like about tree recursion: the bookkeeping is free, and the job is
mostly trusting that it's happening.

The base case is the empty subtree. Visit nothing, contribute nothing, hand the slice
back exactly as it came in.

### The part that actually bites

In Go, `append` can reallocate the underlying array.

That sentence is the whole bug. If the backing array is full, `append` allocates a bigger
one, copies into it, and returns a slice header pointing somewhere new. The caller's
slice header still points at the old array — same length it always had, none of the new
values.

So you can't hand a slice down the recursion and assume the appends come back. You have
to thread it through as both a parameter and a return value, and every recursive call has
to capture what comes back:

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

Write `traversal(A.Left, arr)` without the `arr =` and the code still compiles. It still
runs. On a small enough tree it even returns the right answer, because nothing
reallocated. That's what makes it worth knowing about rather than worth looking up.

### Tracing it

`[1,null,2,3]` — node `1` has no left child, its right child is `2`, and `2` has a left
child `3`. Expected output is `[1,2,3]`.

![Example 1](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/101_200/binary_tree_preorder_traversal/images/1.png)

`traversal(1)` appends `1`, so `arr = [1]`.

![Step 1: visit root 1, arr becomes [1]](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/101_200/binary_tree_preorder_traversal/images/walkthrough-1.png)

`1.Left` is `nil`, so that call hits the base case immediately and does nothing.
`1.Right` is `2`, so `traversal(2)` appends `2` and `arr` becomes `[1,2]`.

![Step 2: 1.Left is nil, visit 2, arr becomes [1,2]](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/101_200/binary_tree_preorder_traversal/images/walkthrough-2.png)

`2.Left` is `3`, so `traversal(3)` appends `3`. `arr` is now `[1,2,3]`.

![Step 3: visit 3 via 2.Left, arr becomes [1,2,3]](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/101_200/binary_tree_preorder_traversal/images/walkthrough-3.png)

`3` has no children, so both of its calls hit the base case and return without touching
anything. The `[1,2,3]` slice unwinds back up through `3`, then `2` (whose `Right` is
also `nil`), then `1`, and out.

![Step 4: 3's children are nil, unwind to return [1,2,3]](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/101_200/binary_tree_preorder_traversal/images/walkthrough-4.png)

### Complexity

O(n) time — every node is visited once and does O(1) work. O(h) space for the recursion
stack, where h is the tree's height, plus O(n) for the output slice itself.

---

Tomorrow is postorder, which is the same three lines in a different order and has its own
small surprise waiting. If you've been following along since Day 1, the binary tree batch
is about half done.

Full code and the step-by-step walkthrough:
[binary_tree_preorder_traversal](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
