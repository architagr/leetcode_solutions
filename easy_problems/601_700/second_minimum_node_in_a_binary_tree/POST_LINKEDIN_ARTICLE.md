# 365 Days of LeetCode Challenge — Day 18/365

## Second Minimum Node In a Binary Tree (Easy)

🔗 https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/

---

### The problem

You're given a non-empty binary tree with a very specific shape rule: every node has
either **zero** or **two** children, and whenever a node has two children, its own value
is the smaller of its two children's values:

```
root.val = min(root.left.val, root.right.val)
```

Given a tree like this, find the **second smallest distinct value** among all the node
values in the tree. If no such value exists, return `-1`.

Two examples from the problem statement:

!["Example 1"](smbt1.jpg "Example 1")

```
Input: root = [2,2,5,null,null,5,7]
Output: 5
Explanation: The smallest value is 2, the second smallest value is 5.
```

!["Example 2"](smbt2.jpg "Example 2")

```
Input: root = [2,2,2]
Output: -1
Explanation: The smallest value is 2, but there isn't any second smallest value.
```

---

### The intuition

The obvious approach — collect every value into a set, sort it, and grab the second
element — works. But that special structural rule the problem hands us,
`root.val = min(root.left.val, root.right.val)`, lets us do a lot better, and it's worth
sitting with *why* before jumping to code.

**Fact 1 — the root is always the global minimum.** Because every parent's value is the
smaller of its two children's values, the minimum can't hide deeper in the tree — it
always bubbles up to the root. So `root.Val` *is* the smallest value in the tree, with no
search required.

**Fact 2 — a subtree's minimum equals its own root's value.** Apply the same rule
recursively: whatever value sits at the root of *any* subtree is also the smallest value
anywhere inside that subtree. Values never decrease as you descend — they only stay the
same or grow.

That second fact is the whole trick. Suppose we're walking the tree and we land on a node
whose value is strictly greater than the known global minimum. Two things are now true at
once:

- this node's value is a **candidate** for the second-minimum answer, and
- **nothing below it can beat that candidate**, because the entire subtree underneath is
  bounded below by this node's own value.

So the moment we find such a node, we can record it as a candidate and stop descending
into that branch entirely — a free pruning opportunity that a plain "collect everything"
approach doesn't get.

The only reason to keep recursing into a branch is when the current node's value is
*still equal* to the global minimum — meaning the tree hasn't "branched away" from the
minimum yet, so any second-minimum value is still hiding further down.

**Complexity.** Time is O(n) in the worst case (imagine a tree where only one deep leaf
differs from the minimum — you still have to walk down to it), and space is O(h) for the
recursion stack, where h is the tree's height. No auxiliary collection of values is ever
stored, unlike the sort-based approach.

---

### The actual code

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

Two package-level variables, `min` and `ans`, carry state across recursive calls without
threading them through return values.

### Tracing it against Example 1

`root = [2,2,5,null,null,5,7]` looks like this:

```
        2
       / \
      2   5
         / \
        5   7
```

**Step 1 — initialize.** `min` is pinned to `root.Val` (2). `ans` starts at infinity —
"no candidate yet."

<img src="images/walkthrough-1.svg" alt="Step 1: initialize min and ans" width="360" />

**Step 2 — `dfs(root)`, then `dfs(left child)`.** At the root, `root.Val` (2) equals
`min`, so we skip the `if` and take the `else if`, recursing into both children. Down the
left branch, the left child also has value 2 — same story, `min == root.Val`, recurse
again. But this node is a leaf, so both recursive calls immediately hit the `nil` base
case and do nothing.

<img src="images/walkthrough-2.svg" alt="Step 2: recursing while value equals min" width="360" />

**Step 3 — `dfs(right child)`, value 5.** Back at the root, we now visit the right
child, value 5. Now `min < root.Val && root.Val < ans` — i.e. `2 < 5 < ∞` — holds, so
`ans = 5`. Note this branch does **not** recurse into the node's own children (the 5 and
7 further down) — that's the pruning described above, in action.

<img src="images/walkthrough-3.svg" alt="Step 3: candidate found, subtree pruned" width="360" />

**Step 4 — unwind and return.** With no more calls left on the stack, control returns to
`findSecondMinimumValue`. `ans` is `5`, less than `math.MaxInt64`, so the function
returns `5` — matching the expected output.

<img src="images/walkthrough-4.svg" alt="Step 4: final return value" width="360" />

### Why Example 2 returns -1

For `root = [2,2,2]`, every single node has value 2, which equals `min` at every step.
`dfs` only ever takes the `else if` branch, recursing all the way to the leaves without
ever satisfying `min < root.Val < ans`. `ans` is never touched, so it's still
`math.MaxInt64` by the time `findSecondMinimumValue` checks it, and the function
correctly returns `-1`.

---

### Takeaway

When a problem statement hands you a structural invariant like
`root.val = min(root.left.val, root.right.val)`, don't just use it to sanity-check the
input — use it to prune your search. Here it told us two free facts (root is the global
min, and a subtree's minimum equals its root) that turned "collect everything and sort"
into a single DFS pass with early termination.
