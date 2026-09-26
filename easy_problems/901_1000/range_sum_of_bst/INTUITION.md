## Intuition

Visit every node, add the ones whose value falls in `[low, high]`. That's the solution in
this folder, and it's correct for any binary tree at all. It sums the left subtree, sums
the right subtree, and adds the node's own value if it's in range.

## Builds on

- [Day 4: Sum of Left Leaves](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/401_500/sum_of_left_leaves/) — a whole-tree sum where each node adds itself only when a condition holds, and the subtree totals are added on the way back up

What this version doesn't use is the one thing the title promises: the tree is a BST. That
ordering tells you which subtrees can't possibly contribute, the same way it told
[Day 45: Search in a BST](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/search_in_a_binary_search_tree/)
which side to walk.

Everything in a node's left subtree is smaller than the node. So if the node's value is
already at or below `low`, nothing on its left can reach the range, and that whole subtree
can be skipped. Symmetrically, if the node is at or above `high`, its right subtree can be
skipped. The fix is two guards around the existing recursive calls:

```go
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Val > low { // only then can the left subtree reach low
		sum += rangeSumBST(root.Left, low, high)
	}
	if root.Val < high { // only then can the right subtree reach high
		sum += rangeSumBST(root.Right, low, high)
	}
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}
	return sum
}
```

On Example 1 that visits 4 of the 6 nodes; on Example 2, 4 of 10. The worst case is still
O(n), when the range covers the whole tree, but a narrow range on a large balanced tree
only walks the paths near the two boundaries plus the nodes actually in range.

I'd still write the plain version first in an interview, because it's obviously correct,
then add the two guards and say why. The guards are where the easy problem stops being a
tree-sum exercise and starts being a BST one.

**Complexity (as written):**
- Time: O(n), every node visited.
- Space: O(h) for the recursion stack.
