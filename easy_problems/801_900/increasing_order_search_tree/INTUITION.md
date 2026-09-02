## Intuition

The output shape the problem wants — every node with no left child and exactly one
right child — is just a sorted singly linked list built out of `TreeNode`s. The
question is really: "give me the nodes of this BST in ascending order, then chain them
together with `.Right`."

And "the nodes of a BST in ascending order" is exactly what an **in-order traversal**
produces, by definition of a binary search tree (left subtree < node < right subtree).
So the problem splits cleanly into two separate, easy pieces instead of one fused
recursive rewire:

1. Walk the tree in-order and collect every node (not its value — the node itself,
   since we're going to reuse and relink these same nodes) into a flat list.
2. Walk that flat list left to right and point each node's `.Right` at the next one,
   turning the list order into the vine-shaped tree the problem wants.

The one subtlety is that a node arrives at step 1 still wearing its *old* `Left` and
`Right` pointers from the original tree shape. If those aren't cleared, stray leftover
links could survive into the final structure once relinking starts touching `.Right`.
So each node's old children are detached the moment it's collected, before it's placed
in the list — that way step 2 is rewiring nodes that start out completely disconnected,
and the only links in the final tree are the ones step 2 explicitly creates.

**Complexity:**
- Time: O(n) — every node is visited once during the traversal and once during
  relinking.
- Space: O(n) for the list of collected nodes, plus O(h) for the recursion stack of the
  in-order traversal itself (h = tree height, O(log n) balanced / O(n) skewed). No new
  `TreeNode`s are allocated — the existing nodes are reused and relinked in place.
