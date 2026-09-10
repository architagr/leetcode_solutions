**365 Days of LeetCode Challenge — Day 51/365**
**Deepest Leaves Sum** (Medium)
🔗 https://leetcode.com/problems/deepest-leaves-sum/

The obvious solution is two passes: find the maximum depth, then walk again summing everything at that depth. Both halves already exist in this batch — Day 1 for the depth, Day 14 for the accumulation.

One pass is enough, and the trick is being willing to throw work away.

Carry the level down, keep two ints. At every leaf: deeper than anything seen? **Replace** the sum, because everything accumulated so far belonged to a shallower level. Equal? Add. Shallower? Ignore.

```go
if level > deepestLevel {
	sumOfDeepestLevel = node.Val
	deepestLevel = level
} else if level == deepestLevel {
	sumOfDeepestLevel += node.Val
}
```

The replacement is why you never need the final depth in advance — discovering a deeper leaf invalidates the previous answer outright, so whatever survives was accumulated at the true maximum.

One structural oddity worth noticing: the base case is a **leaf**, not nil. Only leaves contribute, and the recursive calls are guarded, so nil never gets passed down. The cost is that `node` is dereferenced unchecked — safe only because the constraints guarantee at least one node.

O(n) time, O(h) space, two ints of state.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1301_1400/deepest_leaves_sum/SOLUTION.md
