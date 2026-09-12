---
meta_title: "Finding a list's middle without counting it first"
meta_description: "Two pointers, one moving twice as fast as the other. When the fast one hits the end, the slow one is exactly halfway, and no length was ever computed."
tags: [golang, linked-list, two-pointers, dsa]
---

![Day 20](HERO.png)

*365 Days of LeetCode Challenge — Day 20/365*

**[876. Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/)** (Easy)

Yesterday's problem was about rewriting pointers. Today's is about reading them, and it introduces a walk that turns up in three more problems before this arc is finished.

The task: return the middle node of a singly linked list, and when the length is even, return the second of the two middle nodes.

## Start with the answer that works

Walk the list and count. Walk it again and stop at position `count/2`. Two passes, O(n) time, O(1) space.

I want to be clear that this is a good answer. It is easy to write, easy to get right, and nobody should feel clever for rejecting it. If it came up in an interview and you wrote this, you have solved the problem.

What is interesting is the constraint it quietly assumes: that you are allowed two passes, and that asking for the length is cheap. On a singly linked list the length is not a stored property. It is something you pay a full traversal to learn.

## Removing the first pass

Run two pointers from the head. Move one of them one node per step, the other two nodes per step.

When the fast pointer reaches the end of the list, the slow one has covered exactly half the distance, because over the same number of steps it moved at exactly half the speed.

There is nothing linked-list-specific in that argument. It is a statement about speed and time that would be just as true of two runners. What makes it useful here is that it converts "where is the halfway point" into something you can answer *while* walking, rather than after.

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

Note that both pointers start on the same node rather than one starting a step ahead. From a shared starting line, after k iterations `end` has covered 2k nodes and `middle` has covered k, and the halving is exact. Stagger the start and you shift the answer by one.

## The walk

The list is `[1,2,3,4,5]`.

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

After one iteration `middle` is on 2 and `end` is on 3.

![Step 3](images/walkthrough-3.png)

After two, `middle` is on 3 and `end` is on 5, which is the last node.

![Step 4](images/walkthrough-4.png)

`end.Next` is `nil`, the condition fails, and the loop stops with `middle` on node 3.

The function returns that node, not its value. In a linked list a node carries everything after it, so handing back node 3 hands back `[3,4,5]`, which is why the expected output in the problem statement looks like a list rather than a number.

## Two conditions, two different bugs

```go
for end != nil && end.Next != nil {
```

It is tempting to read that as one defensive check written twice. It is not. Each half catches a different length parity.

`end != nil` is the even-length case. Walk `[1,2,3,4]` and the fast pointer visits node 1, then node 3, then steps to `nil` exactly. Without this test the very next thing evaluated is `end.Next` on a nil pointer.

`end.Next != nil` is the odd-length case. Walk `[1,2,3,4,5]` and the fast pointer lands on node 5, where `end.Next.Next` reads through nothing.

Remove either one and the function panics on half of all possible inputs. That is an unpleasant shape of bug, because whichever example you happen to test first has even odds of passing and convincing you the code is fine.

The order is load-bearing as well. Go evaluates `&&` left to right and stops at the first false, so `end != nil` must be written first. Reverse them and the second test is the one that crashes, on exactly the input the first test exists to handle.

## The even case needs no code

The problem says to return the second middle when the length is even. Nothing in this function implements that rule. There is no `if`, no parity check, no adjustment.

Trace `[1,2,3,4,5,6]`. The fast pointer visits 1, 3, 5 and then steps to `nil`. The slow pointer visits 1, 2, 3, 4 and stops. Node 4 is the second of the two middles.

It falls out of where the loop condition chooses to stop. I find this genuinely satisfying, and it is also a warning: the behaviour is a consequence of the condition rather than an intention expressed in the code, so changing that condition changes the answer in a way nothing in the function will flag.

Which brings up the variant worth memorising:

```go
for end.Next != nil && end.Next.Next != nil {
```

That stops one node earlier and leaves the slow pointer on the *last node of the first half* rather than the first node of the second. When this walk appears inside a harder problem the goal is usually to cut the list in two, and cutting needs the node before the split, not the node after it. That is the form day 26 uses.

## Complexity

- **Time: O(n).** The loop runs about n/2 times and the fast pointer touches about half the nodes. One pass.
- **Space: O(1).** Two pointers. Nothing allocated, nothing that grows.

## Builds on

Nothing yet. This is the first appearance of the two-pointer walk in the series, and days 21, 23 and 26 all lean on it.

Full code and the step-by-step walkthrough:
[middle_of_the_linked_list](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/middle_of_the_linked_list/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
