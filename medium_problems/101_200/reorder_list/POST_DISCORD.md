**365 Days of LeetCode Challenge — Day 26/365**
**Reorder List** (Medium)
🔗 https://leetcode.com/problems/reorder-list/

Turn `L0 -> L1 -> ... -> Ln` into `L0 -> Ln -> L1 -> Ln-1 -> ...`

That looks like an arbitrary shuffle until you write it as two sequences:

L0, L1, L2, ... forward from the start
Ln, Ln-1, Ln-2, ... backward from the end

Interleaved. So it is a merge of the list with its own reverse, and the only obstacle is that a singly linked list has no backward walk. Reverse the back half, and walking it forward walks the original backward.

```go
func reorderList(head *ListNode) {
	mid := getMid(head)
	head2 := mid.Next

	mid.Next = nil
	head2 = reverseList(head2)

	head = mergeList(head, head2)
}
```

getMid is day 20. reverseList is day 19. mergeList is day 22 with the comparison removed, since the alternation is unconditional.

`mid.Next = nil` looks like tidying and is load-bearing: without it the first half runs into the second, which the merge is walking at the same time.

This closes the linked list arc, on a problem with nothing new in it. That is the point - most hard-feeling problems are old ideas standing next to each other.

O(n) time, O(1) space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/reorder_list/SOLUTION.md
