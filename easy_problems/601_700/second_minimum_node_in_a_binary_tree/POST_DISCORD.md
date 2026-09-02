**365 Days of LeetCode Challenge — Day 18/365**

## Second Minimum Node In a Binary Tree (Easy)
🔗 https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/

Every node in this tree has 0 or 2 children, and when it has 2, its value is the smaller
of the two children's values. That guarantees the root always holds the global minimum,
and it also means a subtree's minimum equals its own root's value — so the moment we find
a node above the minimum, we can record it as a candidate **and stop descending into it**,
because nothing underneath can be smaller.

**Examples:**

!["Example 1"](smbt1.jpg "Example 1")
```
Input: root = [2,2,5,null,null,5,7]
Output: 5
```

!["Example 2"](smbt2.jpg "Example 2")
```
Input: root = [2,2,2]
Output: -1
```

**The code:**

```go
var (
	ans int // best candidate found so far for the second-minimum value; math.MaxInt64 means "none yet"
	min int // global minimum value in the tree; by the tree's invariant this always equals root.Val
)

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	if min < root.Val && root.Val < ans {
		// root.Val is strictly between the known min and our best candidate so far,
		// so it's a new best candidate. Because root.val = min(left.val, right.val)
		// holds throughout the tree, this node's value is also the minimum of its
		// entire subtree, so nothing further down can beat this candidate — prune
		// by not recursing into root.Left/root.Right.
		ans = root.Val
	} else if min == root.Val {
		// Haven't branched away from the global minimum yet: the second-minimum
		// value, if any, must be further down, so keep searching both children.
		dfs(root.Left)
		dfs(root.Right)
	}
	// else: root.Val >= ans already, so this subtree can't improve ans either — dead end.
}

func findSecondMinimumValue(root *TreeNode) int {
	min = root.Val   // the tree invariant guarantees the root always holds the global minimum
	ans = math.MaxInt64 // sentinel: no second-minimum candidate found yet
	dfs(root)
	if ans < math.MaxInt64 {
		return ans
	}
	return -1 // every value in the tree equals min, so there is no second minimum
}
```

**Walkthrough on `root = [2,2,5,null,null,5,7]`:**

```
        2
       / \
      2   5
         / \
        5   7
```

**Step 1 — initialize.** `min = root.Val = 2`, `ans = ∞`.

<img src="images/walkthrough-1.svg" alt="Step 1: initialize min and ans" width="360" />

**Step 2 — `dfs(root)` then `dfs(left child)`.** Root's value (2) equals `min`, so we
recurse into both children. The left child is also 2, so we recurse into *its* children
too — but it's a leaf, so both calls immediately hit the `nil` base case. No candidate
found down this branch.

<img src="images/walkthrough-2.svg" alt="Step 2: recursing while value equals min" width="360" />

**Step 3 — `dfs(right child)`, value 5.** Now `min(2) < 5 < ans(∞)`, so `ans = 5`. This
branch does **not** recurse further — the subtree below (5 and 7) can't beat 5, so it's
pruned.

<img src="images/walkthrough-3.svg" alt="Step 3: candidate found, subtree pruned" width="360" />

**Step 4 — return.** `ans = 5 < math.MaxInt64`, so `findSecondMinimumValue` returns `5`.

<img src="images/walkthrough-4.svg" alt="Step 4: final return value" width="360" />

**Why `[2,2,2]` gives `-1`:** every node equals `min`, so `dfs` only ever takes the
`else if` branch all the way to the leaves. `ans` is never updated, stays at
`math.MaxInt64`, and the function falls through to `return -1`.

**Complexity:** O(n) time worst case, O(h) space for the recursion stack (h = tree
height).
