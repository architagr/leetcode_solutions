## 365 Days of LeetCode Challenge — Day 4/365

# Sum of Left Leaves

🔗 https://leetcode.com/problems/sum-of-left-leaves/ · Difficulty: Easy

### The problem

Given the root of a binary tree, return the sum of all left leaves — leaves that are
the left child of their parent.

### The intuition

The tricky part isn't finding leaves — it's knowing whether a leaf is a **left** child
or a **right** child, since only left leaves count. A node itself can't tell you this
about itself; only its **parent** knows which side it's on.

So the recursion needs to check "is my left child a leaf?" from the parent's
perspective, rather than each node trying to detect its own leaf-and-side status in
isolation. The rest is a standard tree recursion: sum the left leaves found in each
subtree, add in the current node's left child's value if that child qualifies, and
return the total.

### The solution

```go
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := sumOfLeftLeaves(root.Left)
	r := sumOfLeftLeaves(root.Right)
	// Only a parent can tell whether its child is a "left" leaf, since a
	// node has no idea which side of its own parent it's on. So this check
	// happens here, from root's perspective, looking at root.Left.
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		l += root.Left.Val
	}
	return l + r
}
```

Walking it through `[3,9,20,null,null,15,7]` (expected `24`):
- `9` is a leaf and `3`'s left child → counts.
- `15` is a leaf and `20`'s left child → counts.
- `7` is a leaf but `20`'s *right* child → doesn't count.
- `sumOfLeftLeaves(20)` detects `15` as a left leaf → returns `15`.
- `sumOfLeftLeaves(3)` detects `9` as a left leaf → `9 + 15 = 24`. ✓

**Complexity:** O(n) time — every node visited once. O(h) space for the recursion
stack, where h is the tree height.

Full code: `easy_problems/401_500/sum_of_left_leaves/` in the repo.
