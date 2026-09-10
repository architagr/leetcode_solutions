**365 Days of LeetCode Challenge — Day 20/365**
**Maximum Level Sum of a Binary Tree** (Medium)
🔗 https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/

Two things have to be true here and only one is about sums.

The sums are the easy half: nil-sentinel BFS from Day 16, add each level, keep the biggest total.

The other half is in the wording — return the **smallest** level whose sum is maximal. Ties go to the shallower level. There's no code for that anywhere. It's one character: the comparison is `sum > maxSum`, strictly greater, so a later level that merely matches doesn't replace the earlier one. Write `>=` and you still return a maximal level, just the wrong one on ties.

```go
if node == nil {
	if sum > maxSum {
		maxLevel = currentLevel
		maxSum = sum
	}
	currentLevel++
	sum = 0
	if len(queue) > 0 {
		push(nil)
	}
	continue
}
```

One more quiet decision: `maxSum` starts at `root.Val`, not 0. Values can be negative, and a tree where every level sums negative would never beat a zero start.

O(n) time, O(w) space. No level is ever stored — just a running sum and two counters.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1101_1200/maximum_level_sum_of_a_binary_tree/SOLUTION.md
