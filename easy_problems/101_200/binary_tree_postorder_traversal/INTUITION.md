## Intuition

"Postorder" just names the order in which a node gets recorded relative to its
children: **left subtree, then right subtree, then the node itself**. So a node's
value is only ever appended to the result *after* everything underneath it has
already been appended — a node is always the last thing written among itself and its
descendants.

That ordering falls out naturally from a simple recursive shape: recurse into the
left child, recurse into the right child, then append the current node's value.
There's no need to track any extra state (like "have I visited this node's children
yet") — the recursion itself guarantees the children are fully processed, and their
values are already sitting in the result slice, before the current node's `append`
ever runs.

The only base case is an empty subtree (`nil`), which simply contributes nothing and
returns the result slice unchanged — that's what stops the recursion from following
`nil` children and is also why leaves resolve immediately: both of a leaf's
recursive calls hit this base case right away, so the very next line appends the
leaf's own value.

**Complexity:**
- Time: O(n) — every node is visited exactly once.
- Space: O(h) for the recursion stack, where h is the tree's height (O(log n) for a
  balanced tree, O(n) for a completely skewed one) — plus O(n) for the output slice
  itself, which is unavoidable since the result contains every node's value.
