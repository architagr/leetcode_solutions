## Intuition

Searching a BST already finds the insertion point. Walk down the way a search would, left
when the value is smaller, right when it's bigger, and since the value isn't in the tree
the walk has to fall off the bottom somewhere. That empty slot is the only place a new leaf
can go without breaking the ordering of anything above it.

So insertion is a search that doesn't stop at "not found". It builds the node there.

## Builds on

- [Day 45: Search in a Binary Search Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/search_in_a_binary_search_tree/) — the same one-path walk; one comparison per level decides the side and the other subtree is never looked at

The part worth slowing down on is how the new node gets attached. A search only has to
return something. An insert has to change a pointer in the parent, and the parent is one
call up the stack by the time the empty slot is found.

The code handles that by having every call return the root of its own subtree, and having
every caller store the return value back into the child it recursed into:

```go
root.Left = insertIntoBST(root.Left, val)
```

At the bottom, `root` is nil and the call returns a fresh `&TreeNode{Val: val}`. The parent
stores it into `Left` or `Right`, and that's the insertion. Every call above that stores
back the same pointer it already had, which is a harmless write.

I like this shape more than the alternative, which is to stop one level early and check
`if root.Left == nil { root.Left = &TreeNode{...} }` on each side. That version works but
needs the nil check duplicated for both children, plus a separate case for an empty tree.
Here the empty tree is the base case: `insertIntoBST(nil, 5)` returns the new node as the
whole tree, with no special handling.

The problem allows any valid BST as the answer, and a rebalancing insert would also pass.
This one never rotates anything. The new value always becomes a leaf, and the existing
nodes stay exactly where they were.

**Complexity:**
- Time: O(h), where h is the height. One node per level on a single path. That's O(log n)
  on a balanced tree and O(n) on a skewed one, like inserting sorted values one by one.
- Space: O(h) for the recursion stack.
