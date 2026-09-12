**365 Days of LeetCode Challenge — Day 21/365**
**Linked List Cycle** (Easy)
🔗 https://leetcode.com/problems/linked-list-cycle/

The set-of-visited-nodes answer is correct and O(n) space. The follow-up asks for O(1), which is where yesterday's walk comes back.

Same two pointers, one moving a node per step and one moving two. What changes is the question: if the list has an end, the fast pointer finds it. If it does not, neither pointer can ever leave.

```go
func HasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
```

Why they must meet, in one sentence: inside the loop the forward gap between them shrinks by exactly one per iteration, and a non-negative integer dropping by one cannot skip zero.

That is also why the steps are 1 and 2, not 1 and 3. A gap shrinking by two can step over zero and the pointers pass without ever landing together.

`slow == fast` compares nodes, not values. Comparing values reports a cycle on `[1,1]`.

O(n) time, O(1) space.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/linked_list_cycle/SOLUTION.md
