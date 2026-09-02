## Solution

The implementation does a single post-order DFS over the tree, using a package-level
variable to track the best diameter seen so far while the recursion returns subtree
heights.

```go
var dia = 0

func diameterOfBinaryTree(root *TreeNode) int {
	dia = 0
	calc(root)
	return dia
}

func calc(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := calc(root.Left)
	right := calc(root.Right)
	dia = maxVal(left+right, dia)
	return maxVal(left, right) + 1
}
func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### Walking through it

- **`dia = 0`** — `diameterOfBinaryTree` resets the running-best diameter before
  starting a fresh traversal (important since `dia` is a package-level variable, not
  a local one — without the reset, a second call in the same process would carry
  over the previous call's answer).
- **`calc(root)`** kicks off the recursion and its return value (the tree's overall
  height) is discarded at the top level — all we actually want out of `calc` is its
  side effect of updating `dia` along the way.
- Inside `calc`, the **base case** `root == nil` returns `0`: an empty subtree
  contributes no height.
- Otherwise `calc` **recurses left, then right**, getting back the height of each
  child subtree.
- `dia = maxVal(left+right, dia)` is the heart of the algorithm: `left+right` is the
  length of the longest path that passes *through the current node* (down the left
  side, then down the right side). Comparing that against the running maximum lets
  every node take a turn being the path's center; the true diameter is whichever
  candidate turns out biggest.
- `return maxVal(left, right)+1` reports this node's **height** back to its parent
  — the taller of its two subtrees, plus one for the edge down to this node — so
  the parent can compute its own `left+right` candidate the same way.
- `maxVal` is just a small helper since Go's standard `max` wasn't relied on here —
  it returns the larger of two ints.

### Visual walkthrough

Using the example from the problem statement:

!["Example 1"](diamtree.jpg "Example 1")

```
Input: root = [1,2,3,4,5]   (2's children are 4 and 5)
```

`calc` is post-order, so it fully resolves the left subtree before touching the
right subtree, and it resolves a node's children before the node itself. For this
tree that visits nodes in the order **4, 5, 2, 3, 1**.

**Step 1 — `calc(4)`:** node 4 is a leaf, so both children return height `0`.
`dia` stays `0`, and `calc` returns `height = 1` up to node 2.

![step 1](images/walkthrough-1.svg)

**Step 2 — `calc(5)`:** same story — node 5 is also a leaf. `dia` is still `0`,
and its height is `1`.

![step 2](images/walkthrough-2.svg)

**Step 3 — `calc(2)`:** now both children of node 2 are known (`left=1`,
`right=1`). The path through node 2 has length `1+1=2`, which beats the current
`dia=0`, so `dia` updates to `2`. Node 2 reports `height = max(1,1)+1 = 2` up to
the root.

![step 3](images/walkthrough-3.svg)

**Step 4 — `calc(3)`:** node 3 is a leaf on the other side of the root. `dia`
stays `2` (a leaf can't beat that), and its height is `1`.

![step 4](images/walkthrough-4.svg)

**Step 5 — `calc(1)`:** the root now has `left=2` (from node 2) and `right=1`
(from node 3). The path through the root has length `2+1=3`, which beats the
current `dia=2`, so `dia` updates to its final value, `3`. That's the path
`4 → 2 → 1 → 3` (or `5 → 2 → 1 → 3`) from the problem's explanation.

![step 5](images/walkthrough-5.svg)

`diameterOfBinaryTree` returns `dia = 3`, matching the expected output.

### Complexity

- **Time:** `O(n)` — `calc` visits each node exactly once.
- **Space:** `O(h)` for the recursion stack, where `h` is the tree's height
  (`O(n)` worst case for a skewed tree, `O(log n)` for a balanced one).
