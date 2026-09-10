## Intuition

The rotation is described per node: the left child becomes the new parent, the old parent
becomes its right child, and the old right child becomes its left child. Applied all the
way down the left spine, that turns the tree over.

Two questions fall out of that. Where does the new root come from, and when does the
rewiring happen?

## Builds on

- [Day 35: Increasing Order Search Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/increasing_order_search_tree/) — relinking the existing nodes into a different shape rather than building new ones, and the care that takes when you are mutating the pointers you navigate by

The new root is the deepest node on the left spine — keep taking `Left` until there isn't
one. That's the base case, `if root == nil || root.Left == nil { return root }`, and it's
the only place a value is ever produced.

Everything above it just passes that value along:

```go
x := upsideDownBinaryTree(root.Left)
// ... rewiring that never touches x ...
return x
```

So the return value is a pass-through. It's discovered once at the bottom and handed
upward untouched through every frame, while the actual work each frame does is local
pointer surgery on `root` and its two children. Those are two completely separate jobs
sharing one recursion, and separating them in your head is most of understanding this
solution.

The rewiring happens on the way back up, after the recursive call. That ordering matters:
the child subtree has to be turned over before its old parent can be hung underneath it.

There's one detail that looks like a bug and isn't. Each frame ends with:

```go
root.Left = nil
root.Right = nil
```

For every node except the original root, those nils are immediately overwritten by the
parent's frame, which is about to set this same node's `Left` and `Right` as part of its
own rewiring. The clearing only survives for the topmost call — the original root, which
becomes a leaf in the flipped tree and genuinely needs both pointers cleared. It's written
uniformly and only matters once.

The problem's guarantee is what makes any of this safe: every right child has a left
sibling and has no children of its own. So a right child is always a leaf, and moving it
wholesale to become someone's left child can't drag a subtree along with it.

**Complexity:**
- Time: O(h), where h is the length of the left spine — the recursion only ever descends
  left, so it never touches the right children except to move them.
- Space: O(h) for the recursion stack.
