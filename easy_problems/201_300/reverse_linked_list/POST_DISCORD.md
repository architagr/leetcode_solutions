**365 Days of LeetCode Challenge — Day 19/365**
**Reverse Linked List** (Easy)
🔗 https://leetcode.com/problems/reverse-linked-list/

New topic. Day 10 flattened a binary tree into a linked list by rewriting `Next` pointers; today the structure is already a line and the same rewrite puts it in the other order.

A list moves no data when you reverse it. The nodes stay exactly where they are and you turn each arrow around. The catch: pointing a node backwards destroys the only reference to whatever came after it, so you save that first. There is no other order these four lines can go in.

```go
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
```

Return `prev`, not `head`. `head` is nil by then. That mistake fails silently and hands back an empty list.

O(n) time, O(1) space. Empty and single-node lists need no special case.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/reverse_linked_list/SOLUTION.md
