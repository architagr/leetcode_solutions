**365 Days of LeetCode Challenge — Day 9/365**
**Binary Tree Postorder Traversal** (Easy)
🔗 https://leetcode.com/problems/binary-tree-postorder-traversal/

**Intuition:** postorder means children before parent: recurse left, recurse
right, then append the current node. The recursion guarantees every
descendant is already recorded before a node appends its own value, so
there's no extra state to track. Nice part is you don't manage any of that
yourself, the recursion just does it right by construction.

![Example 1](images/1.png "Example1")

**Full solution:**
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

**Walkthrough** on `root = [1,null,2,3]` (expected `[3,2,1]`):

- `traversal(3, [])`: both children `nil` → leaf bottoms out immediately →
  appends `3` → `arr = [3]`

![Step 1: node 3 is a leaf — both children hit the nil base case, so it appends itself first](images/walkthrough-1.svg)

- `traversal(2, [])`: left (`3`) done, right (`nil`) done → appends `2` →
  `arr = [3, 2]`

![Step 2: node 2's left (3) and right (nil) are both resolved, so it appends itself next](images/walkthrough-2.svg)

- `traversal(1, [])`: left (`nil`) done, right (`2`) done → appends `1` →
  `arr = [3, 2, 1]` ✓

![Step 3: node 1's left (nil) and right (2) are both resolved, so it appends itself last](images/walkthrough-3.svg)

O(n) time, O(h) space for the recursion stack (h = tree height) plus O(n) for
the output slice.

**Follow up:** LeetCode also wants an iterative version. Typical fix is one
stack, pushing nodes and prepending their values instead of appending, which
builds the reverse of a root-right-left traversal and flips it into postorder
for free.
