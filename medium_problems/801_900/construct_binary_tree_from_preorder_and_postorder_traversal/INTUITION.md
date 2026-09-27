## Intuition

Two days ago the BST ordering told us where the left subtree's block ended in a preorder list.
Without a BST there's no ordering to use, so a second traversal has to supply that
information. Here it's postorder.

Preorder is root, left block, right block. Postorder is left block, right block, root. The
value right after the root in preorder, `preorder[1]`, is the root of the left subtree. In
postorder a subtree's root comes *last* in its block. So find `preorder[1]` in postorder:
everything up to and including it is the left subtree. That tells you the left subtree's size,
and the same size cuts both lists into left and right blocks. Recurse on each pair.

## Builds on

- [Day 130: Construct Binary Search Tree from Preorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1001_1100/construct_binary_search_tree_from_preorder_traversal/) — splitting a preorder list into root, left block and right block
- [Day 9: Binary Tree Postorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_postorder_traversal/) — postorder writes a subtree's root last, which is how the left block's end is found here

The answer isn't always unique, which is why the problem accepts any valid tree. If a node has
exactly one child, preorder and postorder look the same whether that child is on the left or
the right. With `preorder = [1,2]` and `postorder = [2,1]`, 2 could be either. This code
always treats `preorder[1]` as a left child, which is one of the valid answers.

The search for `preorder[1]` in postorder is a linear scan per call, so the whole thing is
O(n²) in the worst case. With at most 30 values it doesn't matter; a map from value to
postorder index, built once, would make it O(n).

**Complexity:**
- Time: O(n²) worst case (O(n) with an index map).
- Space: O(h) recursion; sub-slices share the arrays.
