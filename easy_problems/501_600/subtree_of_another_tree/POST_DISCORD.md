**365 Days of LeetCode Challenge — Day 14/365**
**Subtree of Another Tree** (Easy)
🔗 https://leetcode.com/problems/subtree-of-another-tree/

**Intuition:** this is "are two trees identical" (the classic Same Tree check), run at
every node of `root` as a candidate anchor. Checking values before the full structural
comparison saves a lot of wasted recursion, and honestly that one comparison is doing
most of the work here.

![Example 1](images/1.jpg "Example1")

**Full solution:**
```go
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return subRoot == nil
	}

	if root.Val == subRoot.Val && equalBinaryTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func equalBinaryTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	if subRoot == nil {
		return root == nil
	}
	if root.Val != subRoot.Val {
		return false
	}
	return equalBinaryTree(root.Left, subRoot.Left) && equalBinaryTree(root.Right, subRoot.Right)
}
```

**Walkthrough** on `root = [3,4,5,1,2]`, `subRoot = [4,1,2]` (expected `true`):
- `isSubtree(3, subRoot)`: `3 != 4`, so it skips the equality check and recurses into
  `4` and `5`

![Step 1: node 3 vs subRoot's 4 — values differ, skip equalBinaryTree, recurse into 4 and 5](images/walkthrough-1.svg)

- `isSubtree(4, subRoot)`: `4 == 4`, so `equalBinaryTree(4, subRoot)` actually runs

![Step 2: node 4 vs subRoot's 4 — values match, call equalBinaryTree](images/walkthrough-2.svg)

- `equalBinaryTree` walks in lockstep: `4=4`, `1=1`, `2=2`, all matches, so it returns
  `true`, and `isSubtree` returns `true` too. Node `5` never gets checked, `||`
  short-circuits before it's reached.

![Step 3: equalBinaryTree walks 4/1/2 against 4/1/2 in lockstep — all match, isSubtree returns true, node 5 never checked](images/walkthrough-3.svg)

Runs in O(m·n) time (`m` = nodes in `root`, `n` = nodes in `subRoot`), O(h1 + h2) space
for the two recursion stacks.
