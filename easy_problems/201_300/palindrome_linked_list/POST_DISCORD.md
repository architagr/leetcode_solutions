**365 Days of LeetCode Challenge — Day 23/365**
**Palindrome Linked List** (Easy)
🔗 https://leetcode.com/problems/palindrome-linked-list/

On an array this is nothing: an index at each end, walk inward, compare. A singly linked list has no index and no way back, so that algorithm cannot run.

Copying into a slice works and costs O(n) space. The follow-up asks for O(1), and this is where the week pays off: if you cannot walk the second half backwards, reverse it. Then walking it forwards walks the original backwards.

Day 20 finds the middle. Day 19 reverses the back half. Compare inward. No new technique.

```go
slow, fast := head, head
for fast.Next != nil && fast.Next.Next != nil {
	slow = slow.Next
	fast = fast.Next.Next
}

second := reverse(slow.Next)

left, right := head, second
for right != nil {
	if left.Val != right.Val {
		ok = false
		break
	}
	left, right = left.Next, right.Next
}

slow.Next = reverse(second)
```

Two details worth stealing. The loop is driven by `right`, so on odd lengths the unpaired middle node is skipped with no parity check. And the last line puts the caller's list back - nothing in the name `IsPalindrome` says it rearranges your data.

O(n) time, O(1) space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/palindrome_linked_list/SOLUTION.md
