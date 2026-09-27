## Intuition

The last few days showed what a traversal on its own can't do. Preorder alone didn't pin down a
general tree; preorder plus postorder still couldn't tell a lone left child from a lone right one.
A traversal of *values* loses the shape.

The fix is to record the shape directly: write down the empty spots too. Do a preorder walk, and
whenever a child is missing, write an empty token for it. Now the string contains every decision
the tree makes, and reading it back is the same walk in reverse. Take the next token: if it's
empty, this position is nil; otherwise make a node with that value, then build its left subtree
from the following tokens, then its right.

## Builds on

- [Day 8: Binary Tree Preorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/) — the node-left-right order used both to write the tree and to read it back
- [Day 17: Construct String from Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/601_700/construct_string_from_binary_tree/) — turning a tree into a string where missing children have to be represented so the shape survives

In this code the nil marker is the empty string, so the tokens joined by commas look like
`"1,2,,,3,4,,,5,,"`. The empty tree serialises to `""`, which splits back into one empty token,
which deserialises to nil, so Example 2 needs no special case.

`deserialize`'s inner function returns the tokens it didn't consume along with the node it built.
That's how the right subtree knows where to start: it receives whatever the left subtree left
behind. A shared index would work equally well; returning the remainder keeps the function free of
outside state.

A tree with `n` nodes produces `n + 1` nil tokens (every node has two child slots, and all but
`n - 1` of them are empty), so the string is O(n). Both directions visit each token once.

**Complexity:**
- Time: O(n) each way.
- Space: O(n) for the string and tokens, O(h) recursion.
