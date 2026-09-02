## 365 Days of LeetCode Challenge — Day 15/365

# Merge Two Binary Trees

🔗 https://leetcode.com/problems/merge-two-binary-trees/ · Difficulty: Easy

### The problem

You're given two binary trees, `root1` and `root2`. Imagine overlaying one on top of
the other: wherever both trees have a node at the same position, the merged tree's
node value is the sum of the two. Wherever only one tree has a node, that node is used
as-is. Return the merged tree.

### The intuition

This is a tree recursion where the "combine" step is trivial — just add two numbers.
The real work is deciding what to do when one side of the merge is missing.

At every pair of positions `(root1, root2)` in the two trees, there are exactly three
cases:
- Both nodes exist → the merged node's value is their sum, and both children still
  need to be merged the same way, one level down.
- Only `root1` exists (`root2` is `nil` here) → nothing to add, so the merged subtree
  from here on is just whatever `root1` already is.
- Only `root2` exists → symmetric case, the merged subtree is just `root2`.

Because a `nil` subtree simply "loses" to whichever tree still has content, you never
need to build new nodes for the parts where only one tree is present — you can hand
back that existing subtree as-is and let it be reused inside the merged result. New
nodes only get allocated where *both* trees are actually present and their values need
summing.

The recursion runs in **preorder**: settle the current node's value first (so the
merged parent exists before its children are attached), then recurse into the left
pair and the right pair to build the merged subtrees, and finally wire those results
in as `Left`/`Right`.

### The solution

![Example 1](images/1.jpg "Example1")

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

Walking it through `root1 = [1,3,2,5]`, `root2 = [2,1,3,null,4,null,7]` (expected
`[3,4,5,5,4,null,7]`):

- `mergeTrees(1, 2)`: both non-nil → new node `1+2 = 3`. Its children aren't merged
  yet.

  ![Step 1: merged root created as 1+2=3, children still pending](images/walkthrough-1.svg)

  - `root.Left = mergeTrees(3, 1)`: both non-nil → new node `3+1 = 4`. Its own
    children come from base cases: `mergeTrees(5, nil)` returns `root1`'s `5`
    untouched, and `mergeTrees(nil, 4)` returns `root2`'s `4` untouched.

    ![Step 2: left subtree resolved — 3+1=4, reusing 5 from root1 and 4 from root2 as leaves](images/walkthrough-2.svg)

  - `root.Right = mergeTrees(2, 3)`: both non-nil → new node `2+3 = 5`. Left child
    comes from `mergeTrees(nil, nil)` → `nil` (no left child at all); right child
    comes from `mergeTrees(nil, 7)` → `root2`'s `7` untouched.

    ![Step 3: right subtree resolved — 2+3=5, right child reuses 7 from root2, no left child — merge complete](images/walkthrough-3.svg)

- Back at the root, `root.Left = 4` and `root.Right = 5` are wired in, giving the
  final merged tree `[3,4,5,5,4,null,7]`. ✓

**Complexity:** O(min(m, n)) time and space, where `m` and `n` are the node counts of
the two input trees — recursion only descends as far as *both* trees still have a node
at that position; once either side is `nil`, that branch returns immediately without
visiting further into whatever remains on the other side.

Full code: `easy_problems/601_700/merge_two_binary_trees/` in the repo.
