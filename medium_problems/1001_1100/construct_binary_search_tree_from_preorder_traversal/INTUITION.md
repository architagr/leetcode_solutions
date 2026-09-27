## Intuition

This starts a run on rebuilding trees from traversals. Preorder writes a node, then its
entire left subtree, then its entire right subtree. So the first value is the root, and the
rest of the list is two blocks side by side: left subtree, then right subtree.

What tells you where one block ends? The BST ordering. Everything in the left subtree is
smaller than the root, everything in the right is bigger. So the right block starts at the
first value bigger than the root. Split there and recurse on both blocks.

## Builds on

- [Day 8: Binary Tree Preorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/) — the order being undone here: node, then all of the left subtree, then all of the right
- [Day 53: Validate Binary Search Tree](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1_100/validate_binary_search_tree/) — every node lives inside a range set by its ancestors, the idea behind the linear-time version

Each split scans its block for that first bigger value. On a balanced tree that totals
O(n log n), but on a sorted input (a chain) every block is almost the whole list and it
becomes O(n²). With 100 values that's irrelevant; at scale it isn't.

The linear version flips the question. Instead of finding where a block ends, walk the
preorder list once with an index and give each recursive call an upper bound: "take values
while they're below this". A left child's bound is its parent's value; a right child inherits
the parent's bound. Every value is consumed exactly once:

```go
func bstFromPreorderLinear(preorder []int) *TreeNode {
	i := 0
	var build func(bound int) *TreeNode
	build = func(bound int) *TreeNode {
		if i == len(preorder) || preorder[i] > bound {
			return nil
		}
		node := &TreeNode{Val: preorder[i]}
		i++
		node.Left = build(node.Val)
		node.Right = build(bound)
		return node
	}
	return build(math.MaxInt)
}
```

I checked it against the repo version on a thousand random BSTs. It's the Validate BST idea
from Day 53 run in reverse: there the bounds checked a tree, here they build one.

**Complexity (as written):**
- Time: O(n²) worst case (sorted input), O(n log n) on a balanced tree.
- Space: O(h) for recursion; the sub-slices share the original array.
