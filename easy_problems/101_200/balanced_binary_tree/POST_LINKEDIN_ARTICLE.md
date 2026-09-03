# 365 Days of LeetCode Challenge — Day 5/365

## Balanced Binary Tree (Easy)

🔗 [leetcode.com/problems/balanced-binary-tree](https://leetcode.com/problems/balanced-binary-tree/)

Given a binary tree, determine if it is height-balanced: at every single node, the
heights of its left and right subtrees must differ by no more than 1.

**Example 1 — balanced:**

![Example 1](balance_1.jpg)

Input: `root = [3,9,20,null,null,15,7]` → Output: `true`

**Example 2 — not balanced:**

![Example 2](balance_2.jpg)

Input: `root = [1,2,2,3,3,null,null,4,4]` → Output: `false`

---

## The intuition

The obvious way to solve this is a `height(node)` helper, then walking every node and
comparing `height(node.Left)` against `height(node.Right)`. It works, and the waste in
it is easy to miss: computing `height()` for a node already walks that node's whole
subtree. Do that separately at every ancestor above it and you end up re-walking the
same nodes over and over. That's O(n) work repeated at O(n) levels, so O(n²) in the
worst case. A long, skewed tree makes it obvious.

The fix comes from noticing that height and balance are both bottom-up computations, and
they can share one traversal. A node's height depends only on its children's heights. A
node's balance depends on the same thing. So one post-order pass can hand back both
pieces at once:

- how tall this subtree is, and
- whether everything under it is still balanced.

If a subtree buried deep in the tree is already broken, there's no point computing
anything above it. The whole tree fails regardless. The recursion can bail out the
instant it finds a violation instead of grinding through the rest of the tree, and this
is the part of the solution I find most satisfying: a broken tree gets abandoned early
instead of fully walked.

This is the augment-the-return-value pattern: instead of a function that answers one
question, have it return everything the caller needs to decide in one shot. No subtree's
height is ever computed twice, and no node needs to look past its immediate children.

This runs in O(n) time, since every node is visited once, and O(h) space for the
recursion stack, where h is the tree's height: O(log n) for a balanced tree, O(n) worst
case for a skewed one.

---

## The code

```go
func isBalanced(root *TreeNode) bool {
	// Only the balance flag matters here; the overall tree height is discarded.
	_, ok := validateTree(root)
	return ok
}

// validateTree does a single post-order pass that returns, for the subtree rooted at
// node, both its height and whether it (and everything beneath it) is height-balanced.
// Computing both in one pass avoids recomputing subtree heights repeatedly, which is
// what makes a naive "check height() at every node" approach O(n^2) in the worst case.
func validateTree(node *TreeNode) (height int, ok bool) {
	// Base case: an empty subtree has height 0 and is trivially balanced. This is what
	// lets leaf nodes resolve cleanly with no special-casing above.
	height = 0
	ok = true
	if node == nil {
		return
	}
	rightHeight := 0
	rightHeight, ok = validateTree(node.Right)
	if !ok {
		// Short-circuit: the right subtree is already unbalanced, so this subtree
		// (and everything above it) can never be balanced either. Skip checking the
		// left subtree entirely.
		return
	}
	leftHeight := 0
	leftHeight, ok = validateTree(node.Left)
	if !ok {
		// Same short-circuit, triggered by the left subtree instead.
		return
	}
	// Both children are individually balanced (ok == true so far); now check that
	// their heights don't differ by more than 1 at this node.
	diff := leftHeight - rightHeight
	if diff > 1 || diff < -1 {
		ok = false
		return
	}
	// This node passes its own check: its height is one more than its taller child.
	height = maxValue(leftHeight, rightHeight) + 1
	return
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### Walking through it

`isBalanced` itself doesn't do much. All the real work is in `validateTree`, which
returns two things for the subtree rooted at `node`: its `height`, and whether it (and
everything below it) `ok`-ays as balanced. `isBalanced` only cares about `ok`. The
height of the whole tree just gets thrown away with `_`.

The base case is where named return values pay off. Go initializes them up front, so
`nil` resolves to `height = 0, ok = true` with no special-casing needed. That's what
makes every leaf's children behave correctly without extra code.

One detail worth calling out: this implementation recurses into `node.Right` before
`node.Left`. Balance doesn't care which child goes first, the result is identical either
way, but checking `ok` immediately after each call means that the moment one side fails,
the function returns early and the other side never even gets visited. That's the
short-circuit from the intuition section, written out in code.

Once both children report back a real height, meaning both are individually balanced,
this node checks itself: `diff := leftHeight - rightHeight`. If `|diff| > 1`, `ok` flips
to `false` and the leftover `height = 0` doesn't matter, since callers never trust
`height` when `ok` is `false`.

If the node passes, its height is one more than the taller of its two children. That's
the standard recursive definition of tree height, computed bottom-up as the recursion
unwinds.

---

## Visual walkthrough

The trace below follows Example 2, `root = [1,2,2,3,3,null,null,4,4]`, since it's the
one that shows both the "everything's fine so far" path and the short-circuiting
failure path in the same tree:

```
              1
            /   \
           2     2      <- right "2" is a leaf (both children null)
          / \
         3   3           <- right "3" is a leaf (both children null)
        / \
       4   4
```

Because the code calls `validateTree(node.Right)` before `validateTree(node.Left)` at
every level, the actual order of computation is:

**Step 1: the root's right subtree resolves first.** The right child of the root (a leaf
"2") has two `nil` children, so both recursive calls immediately return `(0, true)`, and
this node computes `height = max(0,0)+1 = 1`, `ok = true`.

![Step 1](images/walkthrough-1.png)

**Step 2: on to the left subtree, right child first.** Inside the call for the left "2",
the right child ("3", a leaf) resolves next, also to `(1, true)`.

![Step 2](images/walkthrough-2.png)

**Step 3: the left "3" and its two "4" leaves resolve.** Both "4" nodes return
`(1, true)`. Their parent computes `diff = 1-1 = 0`, passes, and reports
`height = max(1,1)+1 = 2`.

![Step 3](images/walkthrough-3.png)

**Step 4: the left "2" combines its two children.** `rightHeight = 1`, `leftHeight = 2`.
`diff = 1`, within `[-1, 1]`, so it's fine and reports `height = max(2,1)+1 = 3`.

![Step 4](images/walkthrough-4.png)

**Step 5: the root combines its children and fails.** `rightHeight = 1` (from the leaf
"2"), `leftHeight = 3` (from the "2" subtree just resolved). `diff = 2`, and `2 > 1`, so
`ok` flips to `false` right here. `isBalanced` returns `false` without ever looking at
anything else.

![Step 5](images/walkthrough-5.png)

---

That's Day 5 done. Back tomorrow with the next one.

#DSA #LeetCode #100DaysOfCode #CodingInterview #SoftwareEngineering #BinaryTree #Recursion #Golang

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
