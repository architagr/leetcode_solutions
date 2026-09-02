## Solution walkthrough

The implementation is `minDepth(root *TreeNode) int` in `main.go`, with a small
`minVal(a, b int) int` helper alongside it.

![Example 1](ex_depth.jpg)

We'll trace it on the example above, `root = [3,9,20,null,null,15,7]` (expected `2`).

```go
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left != 0 && right != 0 {
		return minVal(left, right) + 1
	} else if left != 0 {
		return left + 1
	}
	return right + 1
}
```

1. **Base case.** `if root == nil { return 0 }` — an empty subtree has no depth. This
   `0` sentinel is what lets a parent later distinguish "no subtree here" from "a real
   subtree that happens to be shallow."

2. **Recurse into both subtrees first.** `left := minDepth(root.Left)` and
   `right := minDepth(root.Right)` collect whatever minimum depth already exists below
   the current node, independent of which (if any) of those children is nil.

3. **Both children real (`left != 0 && right != 0`).** This is a genuine branching
   node — both subtrees exist, so the current node is not on the shortest path unless
   we pick the shallower branch: `return minVal(left, right) + 1`.

4. **Only the left child is real (`else if left != 0`).** The right subtree came back
   `0` because `root.Right` was nil, not because it's a zero-length path — so it must
   be excluded from the comparison entirely. `return left + 1` follows the one path
   that actually exists.

5. **Otherwise, the right child is the only real one (or both are nil).** Falls
   through to `return right + 1`. If `root` is a genuine leaf, `right` is `0` here and
   this correctly returns `1`. If only the right child exists, this follows it.

Walking it through `root = [3,9,20,null,null,15,7]`:

- `minDepth(3)`: recurse left into `9`, recurse right into `20`.
  - `minDepth(9)`: both children nil, `left=0, right=0`. `left != 0` is false, so it
    falls to step 5: `return right + 1 = 1`. (`9` is a leaf, depth `1`.)
  - `minDepth(20)`: recurse left into `15`, recurse right into `7`.
    - `minDepth(15)` → both children nil → falls to step 5 → returns `1`.
    - `minDepth(7)` → both children nil → falls to step 5 → returns `1`.

    ![Step 1: 9, 15, and 7 bottom out as leaves, each returning 1](images/walkthrough-1.svg)

    - Back in `minDepth(20)`: `left=1` (from `15`), `right=1` (from `7`). Both
      non-zero → step 3 → `minVal(1, 1) + 1 = 2`. Returns `2`.

    ![Step 2: at node 20, left=1 and right=1, both non-zero, min(1,1)+1 = 2](images/walkthrough-2.svg)

  - Back in `minDepth(3)`: `left=1` (from `9`), `right=2` (from `20`). Both non-zero
    → step 3 → `minVal(1, 2) + 1 = 2`. Returns `2`. ✓

  ![Step 3: at node 3, left=1 and right=2, both non-zero, min(1,2)+1 = 2 (final answer)](images/walkthrough-3.svg)

### Why the `left != 0` / `right != 0` guard is the whole problem

Example 2 from the problem statement, `root = [2,null,3,null,4,null,5,null,6]`, is a
tree with **no left children at all** — a straight right-leaning chain down to leaf
`6`, with the correct answer `5`. If the code ever computed
`minVal(left, right) + 1` at a node whose left child is nil, `left` would be `0`, and
`minVal(0, right) + 1` would collapse to `1` at the very first node — reporting the
tree bottoms out immediately, which is wrong, since a leaf is only a node with *no*
children, not a node with one missing child.

That's exactly what the `left != 0` / `right != 0` checks in steps 3–5 prevent: a
`0` coming back from a nil child is never allowed to compete in the `min` — the
code instead follows whichever single side actually exists.

![Step 4: at node 2, left is nil (0), so the min is skipped and the code follows the existing right child instead](images/walkthrough-4.svg)

**Complexity:** O(n) time — every node is visited exactly once. O(h) space for the
recursion stack, where h is the tree's height (O(log n) balanced, O(n) skewed).

Full code: `easy_problems/101_200/minimum_depth_of_binary_tree/` in the repo.
