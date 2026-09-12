**365 Days of LeetCode Challenge — Day 20/365**
**Middle of the Linked List** (Easy)
🔗 https://leetcode.com/problems/middle-of-the-linked-list/

The obvious answer counts the nodes, then walks to position count/2. Two passes, and honestly fine.

The one-pass version runs two pointers from the head, one moving a node per step and one moving two. When the fast one hits the end, the slow one is exactly halfway. The length never exists as a number anywhere.

```go
func middleNode(head *ListNode) *ListNode {
	middle, end := head, head

	for end != nil && end.Next != nil {
		middle = middle.Next
		end = end.Next.Next
	}
	return middle
}
```

Both loop conditions matter, and they catch different things. `end != nil` is the even-length case, where the fast pointer steps exactly onto nil. `end.Next != nil` is the odd-length case, where it lands on the last node. Drop either one and half of all inputs panic.

Returning the second middle on an even list is not implemented anywhere. It falls out of that condition.

O(n) time, O(1) space.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/middle_of_the_linked_list/SOLUTION.md
