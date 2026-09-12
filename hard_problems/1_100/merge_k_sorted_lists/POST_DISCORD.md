**365 Days of LeetCode Challenge — Day 80/365**
**Merge k Sorted Lists** (Hard)
🔗 https://leetcode.com/problems/merge-k-sorted-lists/

The problem the last two months have been pointing at. It needs the linked list arc from day 22 and the heap arc from this week, at the same time.

Day 22 merged TWO lists: compare heads, take the smaller, advance. That was an `if`, because with two candidates that is all a comparison needs to be.

With k lists, "take the smallest head" becomes a repeated query for the minimum of a collection that changes after every answer. Which is what a heap is for.

```go
for h.Len() > 0 {
	node := heap.Pop(h).(*ListNode)
	tail.Next = node
	tail = node
	if node.Next != nil {
		heap.Push(h, node.Next)
	}
}
tail.Next = nil
```

The detail the complexity rests on: the heap holds one node PER LIST, never every node. A list's second element cannot be the next smallest while its first is unplaced - the list is sorted. So the heap is size k, not N: O(N log k) instead of O(N log N).

Two more things. Filter nil lists before Init or `Less` dereferences one. And splicing means every reused node still points into its original list, so the merge must be terminated explicitly - the tests walk the result with a step cap so a cycle fails instead of hanging.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/hard_problems/1_100/merge_k_sorted_lists/SOLUTION.md
