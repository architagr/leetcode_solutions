# Intuition

## Related Easy Problems

- [Day 8: Binary Tree Preorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/) — tree traversal patterns
- [Day 17: Two Sum IV - Input is a BST](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/two_sum_iv_input_is_a_bst/) — comparing nodes in a tree

---

In a perfect binary tree, nodes at mirrored positions (left of left subtree vs right of right subtree) are at the same level. Compare nodes at mirrored positions recursively. When the level number is odd, swap their values. Then recurse into the children, but swap the child pointers to maintain the mirror relationship.

Key: don't rearrange the tree structure, just swap node values.
