**365 Days of LeetCode Challenge — Day 22/365**
**Merge Two Sorted Lists** (Easy)
🔗 https://leetcode.com/problems/merge-two-sorted-lists/

The smallest node not yet placed is always the head of one list or the head of the other, because both are sorted. Compare heads, take the smaller, repeat.

The annoying part is the first node: until something is placed there is no previous node to attach to. A dummy head deletes that special case.

```go
func MergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}
	return dummy.Next
}
```

The leftover list is attached in one assignment, not walked - it is already sorted and already larger than everything placed.

`<=` rather than `<` keeps equal values in their original order, which is what makes this stable enough to be a merge sort's combining step.

O(n + m) time, O(1) space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/merge_two_sorted_lists/SOLUTION.md
