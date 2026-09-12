**365 Days of LeetCode Challenge — Day 60/365**
**Min Stack** (Medium)
🔗 https://leetcode.com/problems/min-stack/

A stack where push, pop, top and getMin are all O(1).

Two fixes that do not work. Scanning in getMin is O(n). Keeping a single `min` field works until the first pop - pop the element that IS the minimum and the field is stale, with nothing to recompute it from short of a scan.

That second failure is the insight: the minimum is not a property of the stack, it is a property of the stack AT A GIVEN HEIGHT. Popping returns you to a height whose minimum you already discarded.

So do not discard it.

```go
type dataNode struct {
	val, min int
}

func (this *MinStack) Push(val int) {
	currMin := this.GetMin()
	if currMin > val {
		currMin = val
	}
	this.dataArr = append(this.dataArr, dataNode{val: val, min: currMin})
}

func (this *MinStack) GetMin() int { return this.dataArr[this.len-1].min }
```

The payoff is that Pop contains no minimum-maintenance code at all. Removing the top exposes an entry whose min was computed when it was on top and is still correct, because nothing below it ever changed.

O(1) everything, O(n) space - two ints per element instead of one.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/min_stack/SOLUTION.md
