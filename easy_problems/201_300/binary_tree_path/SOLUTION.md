## Solution walkthrough

The implementation is split across `binaryTreePaths(root *TreeNode) []string` (the
entry point) and `foo(node *TreeNode, current string, result *[]string)` (the recursive
helper for everything below the root), both in `main.go`.

1. **`binaryTreePaths` seeds the path with the root.** `current := fmt.Sprint(root.Val)`
   starts the path string as just the root's value, with no `"->"` prefix — the root is
   the one node that never needs a separator before it.

2. **Root-as-leaf special case.** `if root.Left == nil && root.Right == nil` — if the
   root has no children, the whole tree is a single node, so the only path is `current`
   itself. Returned immediately.

3. **Otherwise, recurse into whichever children exist.** `result` collects every
   completed path found by the recursion. For each non-nil child, `foo` is called with
   the child, the path so far (`current`), and a pointer to `result` so the recursive
   calls can all append into the same slice.

4. **`foo` extends the path and checks for a leaf.**
   `current += fmt.Sprintf("->%d", node.Val)` appends this node's value with the `"->"`
   separator — every call to `foo` is for a node below the root, so this separator is
   always needed here (unlike step 1).

5. **Leaf check inside `foo`.** `if node.Left == nil && node.Right == nil` — if this
   node has no children, `current` is now a complete root-to-leaf path, so it's
   appended to `*result` and the recursion for this branch stops (`return`).

6. **Otherwise, keep descending.** `foo` recurses into whichever of `node.Left`/
   `node.Right` are non-nil, passing along the updated `current` string — each
   recursive call gets its own local copy of `current` (Go strings are immutable and
   passed by value), so sibling branches don't interfere with each other's path.

Walking it through `root = [1,2,3,null,5]` (node `2` has a right child `5`, node `3` has
no children):
- `binaryTreePaths`: `current = "1"`. Root has both children, so calls `foo(2, "1", &result)`
  and `foo(3, "1", &result)`.
- `foo(2, "1", ...)`: `current = "1->2"`. Node `2` has a right child (`5`), so it's not a
  leaf — recurse: `foo(5, "1->2", &result)`.
- `foo(5, "1->2", ...)`: `current = "1->2->5"`. Node `5` has no children — leaf reached,
  append `"1->2->5"` to `result`.
- `foo(3, "1", ...)`: `current = "1->3"`. Node `3` has no children — leaf reached,
  append `"1->3"` to `result`.
- Final `result = ["1->2->5", "1->3"]`, matching the example (order may vary, which the
  problem explicitly allows).
