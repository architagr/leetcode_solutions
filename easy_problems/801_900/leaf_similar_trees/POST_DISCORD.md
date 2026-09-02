**365 Days of LeetCode Challenge — Day 22/365**
**Leaf-Similar Trees** (Easy)
🔗 https://leetcode.com/problems/leaf-similar-trees/

**Intuition:** Stop comparing tree shapes — extract each tree's leaves left-to-right
into a slice (left subtree fully explored before right, so order comes for free),
then it's just a list-equality check: same length, same value at every index.

![Example tree](images/1.png "Example tree")

**Full solution:**
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

**Walkthrough** on Example 1 — `root1 = [3,5,1,6,2,9,8,null,null,7,4]`,
`root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]` (expected `true`):

![Example 1](images/2.jpg "Example1")

- `leafs(root1)` → `6, 7, 4` (left subtree), then `9, 8` (right subtree)

![Step 1: root1's leaves collected left-to-right → (6, 7, 4, 9, 8)](images/walkthrough-1.svg "Step 1")

- `leafs(root2)` (different shape) → same order: `6, 7, 4, 9, 8`

![Step 2: root2's leaves collected left-to-right → (6, 7, 4, 9, 8)](images/walkthrough-2.svg "Step 2")

- Same length, same values at every index → `true`

![Step 3: comparing the two leaf sequences element by element → true](images/walkthrough-3.svg "Step 3")

O(n + m) time, O(n + m) space (plus O(h1 + h2) recursion stack).
