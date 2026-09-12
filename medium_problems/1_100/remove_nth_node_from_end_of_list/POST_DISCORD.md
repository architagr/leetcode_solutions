**365 Days of LeetCode Challenge — Day 24/365**
**Remove Nth Node From End of List** (Medium)
🔗 https://leetcode.com/problems/remove-nth-node-from-end-of-list/

"nth from the end" is measured from a place you cannot start from. A singly linked list gives you the head; the end only exists once you have walked there. So first translate it: count - n is the same position measured from the front.

The part that catches people is smaller. To remove a node you need the node BEFORE it - a node cannot remove itself, since it has no pointer to whatever refers to it. Removal is always done by the predecessor.

So the second walk stops one node early, deliberately:

```go
temp = head
x := 1
for x < count-n {
	x++
	temp = temp.Next
}
temp.Next = temp.Next.Next
```

`x` starting at 1 rather than 0 is what makes that land on the predecessor.

Removing the head is its own branch, because the head has no predecessor. Day 22's dummy node erases exactly that special case - worth noticing there was another option.

The follow-up asks for one pass: put two pointers n apart and move both until the leader falls off the end. Same node visits, one traversal.

O(n) time, O(1) space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1_100/remove_nth_node_from_end_of_list/SOLUTION.md
