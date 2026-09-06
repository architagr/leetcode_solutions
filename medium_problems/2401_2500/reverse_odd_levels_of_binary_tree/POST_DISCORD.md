**365 Days of LeetCode Challenge — Day 32/365**
**Reverse Odd Levels of Binary Tree** (Medium)
🔗 https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/

BFS with a slice per level works. But reversing a level is the same as swapping every mirror pair on it, and mirror pairs are something recursion hands you directly. Descend two nodes at a time, swap when the depth is odd. No buffers, O(log n) stack.

The catch: the two recursive calls have to cross. Pair them the tidy-looking way (left with left) and you permute the level instead of reversing it.

```go
func rev(l, r *TreeNode, d int) {
	if l == nil {
		return
	}
	if d%2 == 1 {
		l.Val, r.Val = r.Val, l.Val
	}
	rev(l.Left, r.Right, d+1)
	rev(l.Right, r.Left, d+1)
}
```

Called as `rev(root.Left, root.Right, 1)`, since the root has no mirror partner and never needs one.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/2401_2500/reverse_odd_levels_of_binary_tree/SOLUTION.md
