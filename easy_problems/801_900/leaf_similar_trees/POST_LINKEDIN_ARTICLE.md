## 365 Days of LeetCode Challenge — Day 22/365

# Leaf-Similar Trees

🔗 https://leetcode.com/problems/leaf-similar-trees/ · Difficulty: Easy

### The problem

Consider all the leaves of a binary tree, from left to right, and read off their
values — that's the tree's **leaf value sequence**. Two trees are *leaf-similar* if
their leaf value sequences are the same, even if the trees themselves are shaped
completely differently.

![Example tree](images/1.png "Example tree")

For the tree above, the leaf value sequence is `(6, 7, 4, 9, 8)`.

### The intuition

"Leaf-similar" is really just asking: if you read off each tree's leaves left to
right, do you get the same sequence? So the problem reduces to two much simpler
steps — extract the leaf sequence of each tree, then compare the two sequences
directly. There's no need to compare the trees' shapes at all; only the leaves, in
left-to-right order, matter.

Extracting a tree's leaf sequence in left-to-right order is exactly what a
depth-first traversal that always visits the left subtree before the right subtree
gives you for free: recurse left, recurse right, and whenever you land on a node with
no children, that's a leaf — record its value. Because the left subtree is always
fully explored before the right subtree, the leaves naturally come out in
left-to-right order without any extra bookkeeping.

Once both leaf sequences are collected into two slices, comparing them for
leaf-similarity is just comparing two lists: same length, and same value at every
index.

### The solution

```go
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// Extract each tree's leaf sequence independently; from here on we only
	// ever compare two int slices, not the trees themselves.
	root1Leafs := leafs(root1)
	root2Leafs := leafs(root2)
	// Cheap early exit: different leaf counts can never be leaf-similar.
	if len(root1Leafs) != len(root2Leafs) {
		return false
	}
	for i := 0; i < len(root1Leafs); i++ {
		if root1Leafs[i] != root2Leafs[i] {
			return false
		}
	}
	return true
}

// leafs returns node's leaves in left-to-right order.
func leafs(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	// A childless node is a leaf; contribute just its own value.
	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}
	// Left subtree's leaves always come before the right subtree's, so the
	// concatenation below naturally preserves left-to-right order.
	return append(leafs(node.Left), leafs(node.Right)...)
}
```

Walking it through Example 1 — `root1 = [3,5,1,6,2,9,8,null,null,7,4]`,
`root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]` (expected `true`):

![Example 1](images/2.jpg "Example1")

- `leafs(root1)` collects `6`, `7`, `4` from the left subtree, then `9`, `8` from the
  right subtree.

![Step 1: root1's leaves collected left-to-right → (6, 7, 4, 9, 8)](images/walkthrough-1.svg "Step 1")

- `leafs(root2)` — a differently shaped tree — collects the same values in the same
  order: `6`, `7`, `4`, `9`, `8`.

![Step 2: root2's leaves collected left-to-right → (6, 7, 4, 9, 8)](images/walkthrough-2.svg "Step 2")

- Same length, same value at every index, so `leafSimilar` returns `true`.

![Step 3: comparing the two leaf sequences element by element → true](images/walkthrough-3.svg "Step 3")

**Complexity:** O(n + m) time — every node of both trees is visited exactly once.
O(n + m) space for the two leaf slices, plus O(h1 + h2) for the recursion stacks
(each tree's height).

Full code: `easy_problems/801_900/leaf_similar_trees/` in the repo.
