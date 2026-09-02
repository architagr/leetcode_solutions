## 365 Days of LeetCode Challenge — Day 12/365

**Diameter of Binary Tree** — [LeetCode 543](https://leetcode.com/problems/diameter-of-binary-tree/) · Difficulty: **Easy**

The longest path between two nodes doesn't have to pass through the root — it can
sit entirely inside a subtree. Key trick: for **any** node, the longest path
*through* that node is `height(left subtree) + height(right subtree)`. Do one
post-order DFS that computes heights normally, and update a running max diameter
at every node along the way — one pass, no extra work.

**Example:**

!["Example 1"](diamtree.jpg "Example 1")

```
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
```

**Code:**

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

`dia` is a package-level running best, reset to `0` at the start of each call so
a previous call's result can't leak in. `calc(root)`'s own return value at the
top level is discarded — we only care about the side effect: every recursive
call updates `dia` whenever `left+right` (the path through the current node)
beats the current best, then returns `max(left, right)+1` — this node's height —
up to its parent so the same check can repeat one level higher.

**Walkthrough** on `root = [1,2,3,4,5]` (node 2's children are 4 and 5). Post-order
visits nodes in the order **4, 5, 2, 3, 1**:

**Step 1 — `calc(4)`:** leaf. `left=0 right=0`, `dia` stays `0`, height `1`.

![step 1](images/walkthrough-1.svg)

**Step 2 — `calc(5)`:** leaf. `dia` stays `0`, height `1`.

![step 2](images/walkthrough-2.svg)

**Step 3 — `calc(2)`:** `left=1 right=1` → `dia=max(0,2)=2`, height `2`.

![step 3](images/walkthrough-3.svg)

**Step 4 — `calc(3)`:** leaf. `dia` stays `2` (unchanged), height `1`.

![step 4](images/walkthrough-4.svg)

**Step 5 — `calc(1)`:** `left=2 right=1` → `dia=max(2,3)=3` (final), height `3`.
That's the path `4 → 2 → 1 → 3` (or `5 → 2 → 1 → 3`).

![step 5](images/walkthrough-5.svg)

`diameterOfBinaryTree` returns `dia = 3` — matches the expected output.

**Complexity:** `O(n)` time (every node visited once), `O(h)` space for the
recursion stack (`h` = tree height — worst case `O(n)`, `O(log n)` if balanced).
