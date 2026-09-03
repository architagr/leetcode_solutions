## Solution walkthrough

The implementation is `increasingBST(root *TreeNode) *TreeNode` plus its helper
`inOrder(root *TreeNode) []*TreeNode` in `main.go`.

![Example 1](images/1.jpg "Example1")

`increasingBST` itself is short — it just calls `inOrder` and then rewires the result:

1. **Collect nodes in ascending order.** `inorder := inOrder(root)` walks the BST
   in-order and hands back a flat `[]*TreeNode` of every node, already in ascending
   value order (that's what in-order traversal of a BST guarantees).

2. **Relink them into a right-only chain.**
   ```go
   for i := 1; i < len(inorder); i++ {
       inorder[i-1].Right = inorder[i]
   }
   ```
   walks the slice and points each node's `.Right` at the next one in ascending order.

3. **Return the new root.** `return inorder[0]` — the smallest node, now the head of
   the chain.

The real work — and the part worth tracing carefully — happens inside `inOrder`:

```go
func inOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	res := make([]*TreeNode, 0)
	left := inOrder(root.Left)
	right := inOrder(root.Right)

	root.Left = nil
	root.Right = nil
	res = append(left, root)
	res = append(res, right...)
	return res
}
```

- **Base case.** `if root == nil { return []*TreeNode{} }` — an empty subtree
  contributes no nodes.
- **Recurse first, on the *old* tree shape.** `left := inOrder(root.Left)` and
  `right := inOrder(root.Right)` are computed *before* `root`'s own children are
  touched, so each recursive call still sees the original left/right children it needs
  to descend into.
- **Detach `root` from its old children.** `root.Left = nil; root.Right = nil` — once
  the subtrees have already been walked and their nodes collected, `root` no longer
  needs to point at its old children. Clearing both pointers now means the node that
  comes back out of this call is a clean, unlinked node — nothing left over from the
  original tree shape to interfere with the relinking pass in `increasingBST`.
- **Assemble in-order: left, then root, then right.** `res = append(left, root)` puts
  `root` right after everything from its left subtree, and
  `res = append(res, right...)` appends everything from its right subtree after that —
  the standard left → node → right in-order sequence, just built as a slice instead of
  printed or summed.

We'll trace `inOrder` on the smaller example from the statement, `root = [5,1,7]`
(a 3-node BST: `5` with left child `1` and right child `7`), since it's small enough to
draw every node explicitly.

![Example 2](images/2.jpg "Example2")

![Step 1: original tree shape, 5 with children 1 and 7, before inOrder runs](images/walkthrough-1.png)

- `inOrder(5)` recurses left and right before touching `5` itself:
  - `inOrder(1)`: both children nil, so `left=[]`, `right=[]`. `1.Left` and `1.Right`
    are already nil, so clearing them is a no-op. `res = append([], 1) = [1]`. Returns
    `[1]`.
  - `inOrder(7)`: same reasoning. Returns `[7]`.
- Back in `inOrder(5)`: `left=[1]`, `right=[7]`. Now `5.Left = nil` and `5.Right = nil`
  detach `5` from its old children. `res = append([1], 5) = [1, 5]`, then
  `res = append([1,5], 7...) = [1, 5, 7]`. Returns `[1, 5, 7]`.

![Step 2: inOrder returns [1, 5, 7], all three nodes already detached from old children](images/walkthrough-2.png)

Back in `increasingBST`, `inorder = [1, 5, 7]`. The relinking loop runs for
`i = 1` and `i = 2`:
- `i=1`: `inorder[0].Right = inorder[1]` → `1.Right = 5`.
- `i=2`: `inorder[1].Right = inorder[2]` → `5.Right = 7`.

`increasingBST` returns `inorder[0]`, node `1` — now the root of a chain
`1 → 5 → 7`, each linked purely through `.Right`, matching the expected output
`[1,null,5,null,7]`.

![Step 3: relinking sets 1.Right=5 and 5.Right=7, increasingBST returns node 1](images/walkthrough-3.png)

The same two-phase idea scales directly to the bigger example above
(`root = [5,3,6,2,4,null,8,1,null,null,null,7,9]`): `inOrder` collects all nine nodes
as `[1,2,3,4,5,6,7,8,9]` (already detached from their old children along the way), and
the relinking loop chains them `1 → 2 → 3 → 4 → 5 → 6 → 7 → 8 → 9`, matching the
expected output exactly.

**Complexity:** O(n) time — every node is visited once by `inOrder` and once by the
relinking loop. O(n) space for the `inorder` slice, plus O(h) for the recursion stack
(h = tree height).
