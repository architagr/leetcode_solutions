## Intuition

The traversal here is the easy part. It's pre-order — node, then left, then right — and
the string is built in exactly that order.

What makes this a medium is a single clause buried in the formatting rules: empty
parentheses get omitted, *except* when a node has a right child and no left child. Then
you have to emit `()` anyway.

That exception isn't decoration. Without it the string stops being reversible. `1(2)`
would describe both "1 with a left child 2" and "1 with a right child 2", and the whole
point of the representation is that it maps one-to-one onto the tree. The empty pair is a
positional marker saying "the left slot is empty, what follows is the right child."

## Builds on

- [Day 8: Binary Tree Preorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/) — the node-then-left-then-right order this string is defined by
- [Day 3: Binary Tree Paths](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/binary_tree_path/) — building a string during a traversal, where each node contributes a piece and the recursion assembles them

So the question becomes: how do you write the branching so the exception falls out rather
than being special-cased?

The answer in this implementation is to make the two children asymmetric in the code, the
way they're asymmetric in the rules. Once a node is known to have at least one child, the
left parentheses are emitted unconditionally, and the right ones only if a right child
exists.

That single decision covers every case. When the left child exists, `parse` returns its
subtree and you get `(2(4))`. When it doesn't, `parse(nil)` returns the empty string and
the same line emits `()` — which is precisely the placeholder the exception asks for. No
branch tests for "right but no left" anywhere in the code, because the shape of the
recursion already produces it.

The mirror case never needs a placeholder: a node with a left child and no right child
just emits `(left)` and stops, and nothing about that is ambiguous.

**Complexity:**
- Time: O(n) for the traversal, and every node's value is converted once. The string
  concatenation is the caveat — each `+=` allocates a new string and copies what came
  before, so on a skewed tree the copying makes this O(n^2) in the worst case. A
  `strings.Builder` threaded through the recursion would keep it at O(n).
- Space: O(h) for the recursion stack, where h is the tree's height, plus O(n) for the
  string being built.
