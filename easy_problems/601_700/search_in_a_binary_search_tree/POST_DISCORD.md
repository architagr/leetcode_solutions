**365 Days of LeetCode Challenge — Day 19/365**
**Search in a Binary Search Tree** (Easy)
🔗 https://leetcode.com/problems/search-in-a-binary-search-tree/

**Intuition:** at each node, `root.Val` vs `val` tells you which single subtree could
still hold the answer, so you never have to check both children or backtrack. It's just
a straight walk from the root down to wherever `val` lives, or down to `nil` if it
doesn't.

**Full solution:**
```go
func searchBST(root *TreeNode, val int) *TreeNode {
	// Ran off the tree without finding val.
	if root == nil {
		return nil
	}
	// BST property: everything in root.Right is >= root.Val, so if root.Val
	// is already too big, val (if present) can only live in the left subtree.
	// The right subtree is ruled out without ever visiting it.
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	// Mirror image: root.Val is too small, so only the right subtree is
	// still worth checking.
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	// Neither comparison fired, so root.Val == val. Return root itself (not
	// just a bool) so its Left/Right pointers carry the whole subtree along,
	// which is what the problem asks for.
	return root
}
```

**Walkthrough** on `root = [4,2,7,1,3]`:

Example 1, `val = 2` (found):

!["Example 1"](tree1.jpg "Example 1")

- `searchBST(4, val=2)`: `4 > 2` → recurse left into `2`, right subtree (`7`) eliminated

![Step 1: at node 4, 4 > 2, recurse left, right subtree eliminated](images/walkthrough-1.svg)

- `searchBST(2, val=2)`: `2 == 2`, so it's a match, returns subtree `[2,1,3]` as expected

![Step 2: at node 2, match found, returns subtree [2,1,3]](images/walkthrough-2.svg)

Example 2, `val = 5` (not found):

!["Example 2"](tree2.jpg "Example 2")

- `searchBST(4, val=5)`: `4 < 5` → recurse right into `7`, left subtree (`2,1,3`) eliminated

![Step 3: at node 4, 4 < 5, recurse right, left subtree eliminated](images/walkthrough-3.svg)

- `searchBST(7, val=5)`: `7 > 5` → recurse left into `nil`, which hits the base case and
  returns `nil` straight away

![Step 4: at node 7, 7 > 5, recurse left into nil, base case returns nil](images/walkthrough-4.svg)

Same shape both times: one comparison per level, one subtree gone per comparison. That's
what makes this a nice easy one to trace by hand.

O(h) time, O(h) space, where h is the tree's height.
