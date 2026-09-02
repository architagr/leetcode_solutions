## 365 Days of LeetCode Challenge — Day 27/365

# Find a Corresponding Node of a Binary Tree in a Clone of That Tree

🔗 https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/ · Difficulty: Easy

### The problem

You're given two binary trees, `original` and `cloned` — `cloned` is guaranteed to be an
exact structural copy of `original`, same shape and same values, just different node
objects sitting in memory. You're also given `target`, a reference to a specific node
*inside* `original`. The task: return a reference to the corresponding node inside
`cloned` — without mutating either tree or `target`.

### The intuition

The detail that makes this problem tractable is easy to skim past: `cloned` isn't just
*a* copy, it's a **structurally identical** one. Same shape, same values, node for node.
That guarantee means you never have to search `cloned` on its own, trying to somehow
recognize `target`'s position independently. Instead, you can walk both trees **in
lockstep** — one step down `original`, the matching step down `cloned`, in the very same
recursive call — and the two pointers will always be looking at "the same" logical
position in the tree, just in two different objects.

At each paired position, the question is simply: *is this the node?* Compare values. If
they match, the `cloned`-side pointer is standing exactly where `target` stands in
`original`, so it's the answer — return it immediately, no further searching needed.

If it's not a match, recurse left in *both* trees together first. If that search turns
up something, propagate it straight back up the call stack. Otherwise, recurse right in
both trees together, and return whatever that finds instead. Because `cloned` mirrors
`original` node-for-node, the base case only needs to check one tree running out of
nodes — wherever `original` ends, `cloned` ends too, at exactly the same moment.

It's a plain DFS/pre-order traversal underneath. The only twist is that it's a **paired**
traversal over two trees at once — which is what lets the search work purely by
structure and value, with zero extra bookkeeping (no hash maps, no path recording,
nothing) needed to line the two trees up.

### The solution

![Example 1](images/1.png "Example1")

```go
// getTargetCopy walks original and clones in lockstep: every recursive call
// advances both pointers together, so wherever the traversal is standing in
// original, clones is standing at the structurally-matching node in the
// cloned tree. That's guaranteed because clones is an exact copy of
// original, which lets the search find target's counterpart purely by
// position/value instead of needing to compare node identities.
func getTargetCopy(original, clones, target *TreeNode) *TreeNode {
	// clones mirrors original node-for-node, so clones == nil alone is
	// enough to detect "ran out of tree" along this path; original doesn't
	// need its own nil check.
	if clones == nil {
		return nil
	}
	// Values are guaranteed unique, so a value match means clones is
	// standing exactly where target stands in original — return this
	// cloned-tree node as the answer.
	if clones.Val == target.Val {
		return clones
	}
	// Search the left subtrees of both trees together first.
	n := getTargetCopy(original.Left, clones.Left, target)
	if n != nil {
		return n
	}
	// Not found on the left; whatever the right subtree search finds (or
	// doesn't) is this call's result.
	return getTargetCopy(original.Right, clones.Right, target)
}
```

Walking it through `original = [7,4,3,null,null,6,19]`, `target = 3` (expected: the `3`
node from `cloned`):

- `getTargetCopy(7, 7, target)`: `clones.Val(7) != target.Val(3)` → recurse left first,
  into `(4, 4)`.

  ![Step 1: at the root pair (7,7), values don't match, recurse left into (4,4)](images/walkthrough-1.svg)

- `getTargetCopy(4, 4, target)`: `clones.Val(4) != target.Val(3)`. `4` is a leaf in both
  trees, so recursing left goes to `(nil, nil)` → returns `nil` immediately, and
  recursing right also goes to `(nil, nil)` → returns `nil`. This call returns `nil` —
  the target isn't anywhere under `4`.

  ![Step 2: at pair (4,4), values don't match and both children are nil, so this branch is a dead end — returns nil](images/walkthrough-2.svg)

- Back in `getTargetCopy(7, 7, target)`: the left recursion returned `nil`, so it falls
  through to search the right subtree instead: `(3, 3)`.

- `getTargetCopy(3, 3, target)`: `clones.Val(3) == target.Val(3)` → match! Returns
  `clones` directly — the `3` node from the *cloned* tree — and that result is
  propagated all the way back up as the final answer. ✓

  ![Step 3: at pair (3,3), values match — clones (the cloned-tree node) is returned as the answer](images/walkthrough-3.svg)

**Complexity:** O(n) time — in the worst case (target is the last node visited, or the
tree is degenerate) every node pair gets visited once. O(h) space for the recursion
stack, where h is the tree's height (O(log n) for a balanced tree, O(n) for a completely
skewed one).

Full code: `easy_problems/1301_1400/find_a_corresponding_node_of_a_binary_tree_in_a_clone_of_that_tree/` in the repo.
