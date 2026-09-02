## Intuition

The tree itself is the expression. Every leaf holds a boolean literal (`0`/`False` or
`1`/`True`), and every internal node holds an operator (`OR` or `AND`) that combines
whatever its two children evaluate to. Since the problem guarantees a **full** binary
tree — every node has either 0 or 2 children — there's never an ambiguous case: a node
is either a pure leaf value, or an operator with exactly two operands to evaluate.

That maps directly onto a **post-order recursion**: to know a node's own boolean value,
you first need the boolean values of both its children, then you combine them with the
node's operator. Leaves are the base case (their value *is* the answer, no children to
wait on); internal nodes recurse into `Left` and `Right` first, and only combine
afterward — which is exactly why the children's evaluations have to come back "up" the
call stack before the parent can produce its own result.

The only other piece is translating the encoded integers into meaning: `2` is `OR`, `3`
is `AND`, and for a leaf, `1` means `True` (anything else, i.e. `0`, means `False`).

**Complexity:**
- Time: O(n) — every node is visited and evaluated exactly once.
- Space: O(h) for the recursion stack, where h is the tree's height (O(log n) for a
  balanced tree, O(n) for a completely skewed one).
