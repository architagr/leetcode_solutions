**365 Days of LeetCode Challenge — Day 24/365**
**Univalued Binary Tree** (Easy)
🔗 https://leetcode.com/problems/univalued-binary-tree/

**Intuition:** A tree is uni-valued exactly when every parent-child edge connects two
equal values. A node never needs the root's value directly. It just compares each
child to its own value as the recursion walks down. An empty subtree is trivially
uni-valued, so `nil` returns `true`.

![Example 2](images/2.png "Example2")

**Full solution:**
```go
func isUnivalTree(root *TreeNode) bool {
	// An empty subtree has nothing that could disagree with its parent's
	// value, so it's trivially uni-valued.
	if root == nil {
		return true
	}
	// Assume each side is fine unless a present child proves otherwise.
	// A side with no child simply stays true.
	left, right := true, true
	if root.Left != nil {
		// A child can only be compared to its parent's value from the
		// parent's own call, since the child has no way to know which
		// node called it. Also recurse so the left child's own subtree
		// is checked for internal uni-valuedness.
		left = isUnivalTree(root.Left) && root.Left.Val == root.Val
	}
	if root.Right != nil {
		// Mirror of the left-child check above.
		right = isUnivalTree(root.Right) && root.Right.Val == root.Val
	}
	// The whole subtree is uni-valued only if both sides checked out.
	return left && right
}
```

**Walkthrough** on `root = [2,2,2,5,2]` (expected `false`):
```
        2
       / \
      2   2
     / \
    5   2
```
- `isUnivalTree(A=2)`, `isUnivalTree(D=2)`, and `isUnivalTree(B=2)` all bottom out on
  leaves, so each just returns `true`

![Step 1: 5, 2, and 2 bottom out as leaves, each returning true](images/walkthrough-1.svg)

- At node `A=2`: left child `C=5` doesn't match (`5 != 2`), so `left = false`. Right
  child `D=2` matches, so `right = true`. `isUnivalTree(A)` returns `false`

![Step 2: at node A (value 2), left child 5 mismatches, left becomes false](images/walkthrough-2.svg)

- At the root: `left = isUnivalTree(A) && A.Val == root.Val` is `false`, since
  `isUnivalTree(A)` already came back `false` and short-circuits the rest. `right =
  true`. Root returns `false`, which matches

![Step 3: at the root, the left subtree already resolved to false, so the whole tree is false](images/walkthrough-3.svg)

Nice bit here: the mismatch (`5` under a `2`) gets caught two levels down and just
rides upward through each ancestor's `left`/`right` check, no re-checking needed.

O(n) time, O(h) space (tree height).
