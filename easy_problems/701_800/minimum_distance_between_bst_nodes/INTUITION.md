## Intuition

The key property to lean on is what makes a tree a **BST** in the first place: an
in-order traversal (left, node, right) visits every value in **strictly sorted order**.

That one fact turns "find the minimum difference between *any* two nodes" — which
sounds like it could require checking every pair, an O(n²) affair — into something much
simpler. Once the values are sorted, the smallest possible difference between any two of
them can only ever occur between two values that are **adjacent** in that sorted order.
Any pair that skips over a value in between can't beat the gap between neighbors,
because the skipped value sits strictly between them and splits that gap into two
smaller (or equal) pieces.

So the whole problem reduces to two steps:
1. Collect every node's value via an in-order traversal — this hands back an
   already-sorted list, no separate sort needed.
2. Walk that sorted list once, tracking the smallest gap between consecutive entries.

**Complexity:**
- Time: the in-order traversal itself visits each of the `n` nodes once, so that part is
  O(n). The implementation here builds the sorted list with nested `append` calls at
  every recursive step (`append(inOrder(root.Left), append([]int{root.Val},
  inOrder(root.Right)...)...)`), which repeatedly copies slice contents as it merges —
  on a heavily skewed tree this pattern degrades toward O(n²) rather than staying
  linear. The final scan over the sorted array is a clean O(n).
- Space: O(n) for the collected slice, plus O(h) for the recursion stack (h = tree
  height).
