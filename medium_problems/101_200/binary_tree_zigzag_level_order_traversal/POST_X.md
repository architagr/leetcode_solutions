Day 16/365 · Binary Tree Zigzag Level Order (Medium)

Yesterday's DFS never learned where a level ends. Zigzag needs it: you can't reverse a level before it's done.

So, back to BFS with a nil sentinel, flag flipped at each boundary.

https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_zigzag_level_order_traversal/SOLUTION.md

#golang #leetcode
