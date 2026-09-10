Day 50/365 · Populating Next Right Pointers (Medium)

Push right before left, so BFS walks each level backwards.

Now "the node to my right" is just "the one I visited before me": current.Next = prev. No lookahead at all.

https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/populating_next_right_pointers_in_each_node/SOLUTION.md

#golang
