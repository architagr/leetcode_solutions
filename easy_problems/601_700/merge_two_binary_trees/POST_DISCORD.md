**365 Days of LeetCode Challenge — Day 15/365**
**Merge Two Binary Trees** (Easy)
🔗 https://leetcode.com/problems/merge-two-binary-trees/

**Intuition:** At every pair of positions, sum the values if both trees have a node
there; otherwise just reuse whichever tree's subtree is non-nil as-is — no need to
build new nodes for the non-overlapping parts. New nodes only get allocated where both
sides actually overlap.

![Example 1](images/1.jpg "Example1")

**Full solution:**
```go
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// If one side is missing at this position, the merged subtree is just
	// whatever the other tree already has here — reused as-is, no new nodes.
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// Both nodes exist: this is the only case where a new node is allocated,
	// and its value is simply the sum of the two overlapping values.
	root := &TreeNode{Val: root1.Val + root2.Val}
	// Recurse in preorder so the merged parent exists before its children
	// are attached; each recursive call resolves to either a brand new
	// summed node or a reused subtree from whichever side is non-nil.
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}
```

**Walkthrough** on `root1 = [1,3,2,5]`, `root2 = [2,1,3,null,4,null,7]` (expected
`[3,4,5,5,4,null,7]`):

- `mergeTrees(1, 2)` → both non-nil → new node `1+2 = 3`, children pending

![Step 1: merged root created as 1+2=3, children still pending](images/walkthrough-1.svg)

- `mergeTrees(3, 1)` → new node `3+1 = 4`, reusing `5` (root1) and `4` (root2) as
  children

![Step 2: left subtree resolved — 3+1=4, reusing 5 from root1 and 4 from root2 as leaves](images/walkthrough-2.svg)

- `mergeTrees(2, 3)` → new node `2+3 = 5`, no left child (both nil), reusing `7`
  (root2) as right child

![Step 3: right subtree resolved — 2+3=5, right child reuses 7 from root2, no left child — merge complete](images/walkthrough-3.svg)

Final merged tree: `[3,4,5,5,4,null,7]` ✓

O(min(m, n)) time and space — recursion stops descending as soon as either side goes
`nil`.
