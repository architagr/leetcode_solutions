Day 39/365 · Delete Node in a BST (Medium)

An internal node is never actually removed.

Overwrite its value with the in-order successor, then delete THAT. The successor has no left child, so it is always easier — which is why this terminates.

https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/401_500/delete_node_in_a_bst/SOLUTION.md

#golang
