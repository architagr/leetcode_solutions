**365 Days of LeetCode Challenge — Day 33/365**
**Binary Search Tree Iterator** (Medium)
🔗 https://leetcode.com/problems/binary-search-tree-iterator/

An iterator promises an ordering and hides *when* the work happens. That second half is the only real decision in this problem, and this solution takes the blunt end of it: do everything up front.

The constructor walks the tree once in-order — ascending, by the property Days 29, 30 and 32 all used — flattens it into a slice, and keeps an index. After that the tree is never touched again.

```go
func (this *BSTIterator) Next() int {
	val := this.inorder[this.index]
	this.index++
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.inorder)
}
```

Both O(1), no traversal state to restore. What it costs is O(n) memory for the iterator's whole lifetime and a full walk before the caller has asked for anything — want the first three values of a million nodes and you paid for all million.

The follow-up asks for O(h) memory: keep a stack of the leftmost spine, and on each next() pop a node then push the left spine of its right child. Each node is pushed and popped once, so it amortises to O(1).

Neither is wrong. Worth being able to say which you wrote and why.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_search_tree_iterator/SOLUTION.md
