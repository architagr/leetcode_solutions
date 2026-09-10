## Intuition

The general version of this problem needs a search. The BST version doesn't, and the
reason is the same property Day 26 leaned on: in a BST, comparing a value against a node
tells you which subtree it must be in, if it's anywhere.

## Builds on

- [Day 26: Search in a Binary Search Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/search_in_a_binary_search_tree/) — a comparison in a BST is a direction, not a verdict. Here the same comparison is made twice, once per target

Run that comparison for both targets at once and there are exactly three cases at any node:

- Both targets are smaller than the node. Then both live in the left subtree, so the node
  they share can't be this one — it's further left.
- Both are larger. Symmetric: the answer is further right.
- Otherwise, the targets are on opposite sides, or one of them *is* this node.

That third case is the answer, and it deserves a moment because it covers two different
situations in one line. If `p` and `q` sit on opposite sides, this node is where the paths
to them diverge, so it's the lowest node with both below it. And if one target equals this
node, the definition of LCA allows a node to be a descendant of itself, so this node is
again the answer — its own path and the other target's path meet right here.

The neat part is that nothing has to detect which of those two it is. "Not both smaller,
not both larger" is exactly the union of them, so one `return root` handles both.

There's no search for `p` or `q` anywhere in this code. The problem guarantees both exist
in the tree, so the walk never has to confirm it — it only has to find the point where
they stop agreeing on a direction.

The nil guards at the top handle degenerate input rather than anything structural: both
targets nil means every node trivially qualifies, so the current root is returned, and one
nil means there's nothing sensible to answer with.

**Complexity:**
- Time: O(h), where h is the tree's height — one comparison per level, and the walk never
  branches or backtracks. That's O(log n) on a balanced tree and O(n) on a skewed one.
- Space: O(h) for the recursion stack. The recursion is a tail call in shape, so an
  iterative version with a `for` loop would make this O(1).
