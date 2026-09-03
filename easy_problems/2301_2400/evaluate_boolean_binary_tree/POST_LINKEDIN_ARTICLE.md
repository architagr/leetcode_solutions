## 365 Days of LeetCode Challenge — Day 28/365

# Evaluate Boolean Binary Tree

🔗 https://leetcode.com/problems/evaluate-boolean-binary-tree/ · Difficulty: Easy

### The problem

You get the root of a **full binary tree** (every node has 0 or 2 children, no
exceptions) that encodes a boolean expression. Leaves hold `0` or `1`, standing for
`False` or `True`. Every other node holds `2` or `3`, standing for `OR` or `AND`, and
its value comes from applying that operator to its two children's values.

Return the boolean result of evaluating the whole tree.

### The intuition

The tree itself *is* the expression. Every leaf is a literal, every internal node is an
operator sitting on top of two operands, and because the tree is guaranteed **full**,
there's no in-between case to handle: a node is either a bare value with no children, or
an operator with exactly two children to combine.

That's a post-order recursion. You can't know what a node evaluates to until you know
both of its children, so you go left, then right, and only combine after both come back.
Leaves are the base case since there's nothing left to wait on. What I like about this
one is how little the code has to decide beyond that: the recursion shape just mirrors
the tree shape.

The only other piece is decoding the integers: `2` means `OR`, `3` means `AND`, and on a
leaf, `1` is `True` while anything else (`0`) is `False`.

### The solution

![Example 1](images/1.png "Example1")

```go
const (
	FALSE = 0
	TRUE  = 1
	OR    = 2
	AND   = 3
)

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == TRUE
	}

	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)
	if root.Val == OR {
		return left || right
	}
	return left && right
}
```

Walking it through `root = [2,1,3,null,null,0,1]` (root is `OR`, its left child is leaf
`1`, its right child is `AND` whose own children are leaves `0` and `1`; expected
`true`):

- The three leaves settle first, since they hit the base case right away: root's left
  child (`1`) returns `True`; `AND`'s left child (`0`) returns `False`; `AND`'s right
  child (`1`) returns `True`.

![Step 1: leaves hit the base case, root.Val == TRUE decides each](images/walkthrough-1.svg)

- Back at the `AND` node: `left = False`, `right = True`. `root.Val` isn't `OR`, so it
  falls to `left && right`, which comes out `False && True = False`.

![Step 2: AND node combines left && right = False && True = False](images/walkthrough-2.svg)

- Back at the `root` (`OR`) node: `left = True` (its own left child), `right = False`
  (the just-resolved `AND` subtree). `root.Val == OR`, so `left || right` gives
  `True || False = True`. Checks out.

![Step 3: root OR combines left || right = True || False = True](images/walkthrough-3.svg)

**Complexity:** O(n) time, since every node gets visited and evaluated exactly once.
O(h) space for the recursion stack, where h is the tree's height: O(log n) if it's
balanced, O(n) if it's basically a straight line.

Full code: `easy_problems/2301_2400/evaluate_boolean_binary_tree/` in the repo.

#DSA #LeetCode #100DaysOfCode #BinaryTree #Recursion #Golang #CodingInterview #Programming

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
