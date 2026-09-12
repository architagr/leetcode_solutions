## Intuition

The problem splits itself for you: find the node, then delete it. The first half is Day
26's walk — compare the key against the node and go the one direction it could be in.

## Builds on

- [Day 45: Search in a Binary Search Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/search_in_a_binary_search_tree/) — the comparison that picks a direction, which is the whole search half of this problem
- [Day 52: Validate Binary Search Tree](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1_100/validate_binary_search_tree/) — in-order is sorted, which is the property the deletion has to preserve and the reason the successor is the right value to move

The second half is where it gets interesting, because a node with two children can't just
be unhooked. Something has to take its place, and only two values in the entire tree can:
its in-order predecessor and its in-order successor — the values immediately before and
after it in sorted order.

That's the crux, and it follows directly from the property Day 52 used. If in-order must
stay sorted after the deletion, then whatever fills the hole must sit between the deleted
node's left subtree (all smaller) and its right subtree (all larger). The only two
candidates are the largest value on the left and the smallest on the right. Anything else
breaks the ordering somewhere.

So the implementation never removes an internal node at all. It overwrites the node's value
with a neighbour's, then recursively deletes that neighbour from the subtree it came from:

```go
root.Val = successor(root)
root.Right = deleteNode(root.Right, root.Val)
```

That looks circular on first read — delete calls delete — but it terminates, and for a
specific reason. The successor is the leftmost node of the right subtree, so by definition
it has no left child. A node with at most one child is one of the easy cases, so the
recursion descends into a strictly simpler problem every time and bottoms out at a leaf.

Three cases, then. A leaf is set to nil. A node with a right child takes its successor. A
node with only a left child takes its predecessor. That third case exists because
`successor` walks into `root.Right` unconditionally and would panic on a nil one — the
predecessor is the mirror answer, and it's equally valid.

The reassignment pattern matters too. `root.Left = deleteNode(root.Left, key)` rather than
just calling and discarding: deletion can change which node is the root of a subtree, so
every caller has to take back whatever comes out and re-attach it.

**Complexity:**
- Time: O(h), where h is the tree's height. The search is one path down; finding a
  successor is one more path; the recursive delete follows that same path. That's O(log n)
  on a balanced tree and O(n) on a skewed one.
- Space: O(h) for the recursion stack.
