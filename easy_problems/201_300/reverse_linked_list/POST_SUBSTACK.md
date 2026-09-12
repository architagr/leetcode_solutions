---
meta_title: "Reversing a linked list without moving a single node"
meta_description: "206 is four lines in a loop, and the order of those four lines is the entire problem. Here is why each one has to sit exactly where it does."
tags: [golang, linked-list, pointers, dsa]
---

![Day 19](HERO.png)

*365 Days of LeetCode Challenge — Day 19/365*

**[206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)** (Easy)

Eighteen days of binary trees end today, and I want to start the linked list stretch by pointing at something that has been sitting in the series already.

Day 10 was Flatten Binary Tree to Linked List. That problem sounds like it is about trees, and it is filed that way, but what the solution actually did was rewrite `Next` pointers until a tree was shaped like a line. No node was copied. No new structure was allocated. The shape changed because the arrows changed.

Today's problem is that idea with the tree taken away. The structure is already a line, and rewriting the same pointers puts it in the other order. If you have been following along, you have done this before without me calling it that.

## What a linked list actually is

It helps to be pedantic about this for a minute, because the pedantry is the solution.

A linked list is not a sequence of values. It is a pile of nodes scattered wherever the allocator put them, each holding one value and one pointer. The order is not a property of the nodes. It lives entirely in the pointers. Nothing about the first node marks it as first. It is first because nothing points at it.

Once that lands, reversing the list stops looking like a data-moving problem. Nothing needs to move. Every node stays exactly where it is in memory. You walk the list once and turn each arrow around, and the new order is a consequence of having done that.

## The loop that does not work

Here is the version most people write first:

```go
head.Next = prev
head = head.Next
```

Two lines, reads fine, and it is broken. Why it is broken is the whole of the problem.

The first line points the current node backwards. That overwrites `head.Next`, which was the only reference to the node you were about to visit. The second line then reads that same field and gets `prev`, so instead of walking forward you walk back into the part of the list you already reversed.

The rest of the list is not corrupted. It is just unreachable, which amounts to the same thing.

So you save it before you break it:

```go
next := head.Next
head.Next = prev
prev = head
head = next
```

Three variables now, four lines, and the order of those lines matters more than anything else in the problem. `next` has exactly one job: hold the rest of the list across the statement that destroys the link to it.

## The whole function

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

`prev` starting at `nil` is doing more work than it looks like. It is not an empty slot waiting to be filled. It is already the correct final `Next` for whatever node is currently at the front, because that node is going to end up last, and the last node in a list points at `nil`. Starting `prev` there is what lets the old head terminate the reversed list with no special case anywhere.

## Watching it run

The list is `1 -> 2 -> 3`.

![Step 1](images/walkthrough-1.png)

Nothing has happened yet. `prev` is `nil`, `head` is node 1.

![Step 2](images/walkthrough-2.png)

`next := head.Next`. Node 2 is now held by a local variable. This line produces no visible change to the list, which is exactly why it is easy to leave out.

![Step 3](images/walkthrough-3.png)

`head.Next = prev`. The reversal itself. Node 1 points at `nil`.

Stop on this picture for a second, because it is the most interesting moment in the problem. The list is currently cut in two. Node 1 is a complete, correct, one-element reversed list. Nodes 2 and 3 are still in their original order. And the only thing in the entire program referring to node 2 is a local variable declared one line earlier. Lose that local and the rest of the list is garbage.

![Step 4](images/walkthrough-4.png)

Two more passes through the same four lines.

![Step 5](images/walkthrough-5.png)

`head` is `nil`, so the loop ends. `prev` is on the original tail, which is the head of the reversed list.

Return `prev`, not `head`. `head` is `nil` by this point, and returning it is the most common way to fail this question, and it fails silently: you get an empty list, which reads as though the reversal did nothing at all.

Then look at where the nodes are in that final picture. Same three positions they occupied in the first one. Nothing moved. Only the arrows changed, which is what I meant at the top.

## There is no other order for those four lines

This is worth stating plainly because it is unusual:

- `next := head.Next` must precede `head.Next = prev`, or the rest of the list becomes unreachable.
- `prev = head` must precede `head = next`, or `prev` lands on the wrong node.

Between them, those two constraints pin all four statements. There is no alternative arrangement that reverses the list. A four-line loop body with exactly one valid ordering is a strange little object, and it is why a problem marked Easy still filters people.

## The edge cases handle themselves

An empty list: the loop body never runs, `prev` is still `nil`, and `nil` is returned. A single node: the body runs once, points it at `nil`, and returns it.

Neither needed an `if`. When the edge cases fall out of the loop condition instead of being guarded in front of it, that is usually a sign the loop is expressing the right invariant rather than the convenient one.

## Complexity

- **Time: O(n).** One visit per node, constant work at each.
- **Space: O(1).** Three pointers, whatever the length of the list. Nothing is allocated.

That O(1) is the actual question being asked. You can solve 206 by walking the list into a slice, reversing the slice, and building a new list from it. Same O(n) time, much easier to get right, and O(n) space. It is a perfectly good answer to "reverse this list" and a failing answer to "reverse this list" as asked in an interview, and the difference between those two is this problem's entire reason for existing.

## Builds on

- [Day 10: Flatten Binary Tree to Linked List](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/flatten_binary_tree_to_linked_list/) — the same move: rewriting `Next` pointers to change a structure's shape without moving a single node

Full code and the step-by-step walkthrough:
[reverse_linked_list](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/reverse_linked_list/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
