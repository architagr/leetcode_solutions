**365 Days of LeetCode Challenge — Day 8/365**
**Binary Tree Preorder Traversal** (Easy)
🔗 https://leetcode.com/problems/binary-tree-preorder-traversal/

**Intuition:** "Preorder" names the order directly — visit the node itself,
then its left subtree, then its right subtree. Recursion mirrors that
definition: visit, recurse left, recurse right. The Go-specific wrinkle is
threading the output slice through as both argument and return value, since
`append` can reallocate its backing array.

![Example 1](images/1.png "Example1")

**Full solution:**
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

**Walkthrough** on `[1,null,2,3]` (expected `[1,2,3]`):
- `traversal(1)` appends `1` → `arr = [1]`

![Step 1: visit root 1, arr becomes [1]](images/walkthrough-1.svg)

- `1.Left` is nil (base case); `1.Right` is `2` → `traversal(2)` appends `2` →
  `arr = [1,2]`

![Step 2: 1.Left is nil, visit 2, arr becomes [1,2]](images/walkthrough-2.svg)

- `2.Left` is `3` → `traversal(3)` appends `3` → `arr = [1,2,3]`

![Step 3: visit 3 via 2.Left, arr becomes [1,2,3]](images/walkthrough-3.svg)

- `3` has no children → unwind all the way back to the top → return `[1,2,3]` ✓

![Step 4: 3's children are nil, unwind to return [1,2,3]](images/walkthrough-4.svg)

O(n) time, O(h) space for recursion (plus O(n) for the output slice).
