## Intuition

The key BST property to lean on: an **in-order traversal of a BST visits nodes in
strictly increasing sorted order**. That's not a coincidence — it's the definition of
a BST (left subtree < node < right subtree), applied recursively.

Once you know the values come out sorted, the rest is a classic fact about sorted
arrays: the minimum absolute difference between *any* two elements always occurs
between two **adjacent** elements in sorted order. You never need to compare a value
against every other value — comparing it only to its immediate predecessor is enough,
because any non-adjacent pair's gap is at least as large as the smallest adjacent gap
between them.

So the algorithm becomes: do an in-order traversal, and as each node is visited,
compare its value to the value of the *previous* node visited (the previous one in
sorted order, i.e. its in-order predecessor) — not the previous one pushed onto a
stack or the parent. Track the smallest such gap seen so far, and that's the answer.
This turns an O(n²) all-pairs comparison into a single O(n) sweep, at the cost of
remembering just one extra value (the previous node) as you go.

**Complexity:**
- Time: O(n) — every node is visited exactly once.
- Space: O(h) for the recursion stack, where h is the tree's height (O(log n) for a
  balanced tree, O(n) for a completely skewed one).
