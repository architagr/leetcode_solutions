**365 Days of LeetCode Challenge — Day 63/365**
**Flatten Nested List Iterator** (Medium)
🔗 https://leetcode.com/problems/flatten-nested-list-iterator/

The obvious solution flattens everything in the constructor into a slice and indexes it. It returns the right values and is NOT an iterator - it does all the work up front whether the caller wants one element or all of them, and keeps a copy of every integer alive.

An iterator does nothing until asked and keeps only enough state to resume. Here that state is the path from the outer list to the cursor: which list, and how far into it, per level.

```go
func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		top := this.stack[len(this.stack)-1]
		if top.i == len(top.list) {
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}
		item := top.list[top.i]
		if item.IsInteger() {
			return true
		}
		top.i++
		this.stack = append(this.stack, &frame{list: item.GetList()})
	}
	return false
}
```

HasNext is the method that does the work - it pops finished lists, descends into nested ones, and stops on an integer. Next only reads and steps past.

The `top.i++` before descending is what stops an infinite loop: without it, the parent resumes pointing at the list it just finished and descends again.

Same design as day 55's BST iterator. O(1) construction, O(depth) space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/301_400/flatten_nested_list_iterator/SOLUTION.md
