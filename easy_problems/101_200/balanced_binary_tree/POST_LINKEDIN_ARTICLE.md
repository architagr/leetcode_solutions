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

The naive way to solve this is to write a `height(node)` helper, then walk every node
comparing `height(node.Left)` against `height(node.Right)`. It's correct, but it's
wasteful: computing `height()` for a node already walks that node's whole subtree. Do
that separately at every ancestor above it and you're re-walking the same nodes over and
over — O(n) work repeated at O(n) levels, so **O(n²)** in the worst case (think of a
long, skewed tree).

The fix is to notice that height and balance are both naturally **bottom-up**
computations, and they can share a single pass. A node's height depends only on its
children's heights. A node's balance depends only on its children's heights too. So
instead of two separate traversals, do one post-order traversal that hands back *both*
pieces of information at once:

- how tall is this subtree, and
- is everything under it still balanced so far?

And if a subtree buried deep in the tree is already broken, there's no reason to keep
computing anything above it — the whole tree is unbalanced, full stop. The recursion can
bail out immediately the moment it finds a violation instead of grinding through the rest
of the tree.

This is the "augment the return value" pattern: rather than a function that only answers
one question, have it return everything the parent needs to make its own decision in one
shot. No subtree's height ever gets computed twice, and no node ever needs to see further
than its immediate children.

**Complexity:** O(n) time (every node visited once), O(h) space for the recursion stack
(the tree's height — O(log n) balanced, O(n) worst case).

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

**`isBalanced` is a thin wrapper.** All the real work happens in `validateTree`, which
returns two things for the subtree rooted at `node`: its `height`, and whether it (and
everything below it) `ok`-ays as balanced. `isBalanced` only cares about `ok` — the
height of the whole tree is thrown away with `_`.

**The base case.** Go's named return values are initialized up front, so `nil` resolves
to `height = 0, ok = true` with no special-casing. That's what makes every leaf's
children behave correctly without extra code.

**Right before left.** Balance doesn't care which child is checked first — the result is
identical either way — but this implementation happens to recurse into `node.Right`
before `node.Left`, checking `ok` immediately after each call returns. The moment one
side fails, the function returns early and the other side is never even visited. That's
the short-circuiting described above, in code.

**The balance check.** Once both children report back a real height (meaning both are
individually balanced), this node checks itself: `diff := leftHeight - rightHeight`.
If `|diff| > 1`, `ok` flips to `false` and the leftover `height = 0` is harmless — callers
never trust `height` when `ok` is `false`.

**Combining into this node's height.** If the node passes, its height is one more than
the taller of its two children — the standard recursive definition of tree height,
computed bottom-up as the recursion unwinds.

---

## Visual walkthrough

The trace below follows **Example 2**, `root = [1,2,2,3,3,null,null,4,4]`, because it
shows both the "everything's fine so far" path and the short-circuiting failure path in
the same tree:

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

**Step 1 — the root's right subtree resolves first.** The right child of the root (a leaf
"2") has two `nil` children, so both recursive calls immediately return `(0, true)`, and
this node computes `height = max(0,0)+1 = 1`, `ok = true`.

![Step 1](images/walkthrough-1.svg)

**Step 2 — descend into the root's left subtree, right child first.** Inside the call for
the left "2", the right child ("3", a leaf) resolves next, also to `(1, true)`.

![Step 2](images/walkthrough-2.svg)

**Step 3 — the left "3" and its two "4" leaves resolve.** Both "4" nodes return
`(1, true)`. Their parent computes `diff = 1-1 = 0`, passes, and reports
`height = max(1,1)+1 = 2`.

![Step 3](images/walkthrough-3.svg)

**Step 4 — the left "2" combines its two children.** `rightHeight = 1`, `leftHeight = 2`.
`diff = 1`, within `[-1, 1]`, so it's fine and reports `height = max(2,1)+1 = 3`.

![Step 4](images/walkthrough-4.svg)

**Step 5 — the root combines its children and fails.** `rightHeight = 1` (from the leaf
"2"), `leftHeight = 3` (from the "2" subtree just resolved). `diff = 2`, and `2 > 1`, so
`ok` flips to `false` right here. `isBalanced` returns `false`.

![Step 5](images/walkthrough-5.svg)

---

That's Day 5. See you tomorrow for the next one. 🌳
