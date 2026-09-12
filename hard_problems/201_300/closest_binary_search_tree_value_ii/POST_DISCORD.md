**365 Days of LeetCode Challenge — Day 57/365**
**Closest Binary Search Tree Value II** (Hard)
🔗 https://leetcode.com/problems/closest-binary-search-tree-value-ii/

Second hard of the challenge, and like the first it isn't the traversal that's hard. It's a claim you have to notice and trust:

**In a sorted array, the k values closest to a target are always contiguous.**

Suppose an answer skipped some value `v` and took a further one `w`, with `v` between the closest element and `w`. Sorted order puts `v` between them in value too, so `v` is at least as close as `w` — swap them and the answer is no worse. So the k closest can always be taken as a run.

That turns "pick k things" into "pick a window". Flatten the BST in-order (yesterday's iterator, verbatim), find the single closest value, then expand outward taking whichever neighbour is nearer:

```go
i, j := minIndex-1, minIndex+1
for i >= 0 && j < len(inOrderArr) && l < k {
	if absDiff(inOrderArr[i], target) < absDiff(inOrderArr[j], target) {
		i--
	} else {
		j++
	}
	l++
}
return inOrderArr[i+1 : j]
```

Greedy is safe because distance grows monotonically as you move away either way — it's a merge step run outward from a centre.

O(n + k) time, O(n) space. The known better answer is O(k + h): two stack iterators from the target, one backwards, one forwards.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/hard_problems/201_300/closest_binary_search_tree_value_ii/SOLUTION.md
