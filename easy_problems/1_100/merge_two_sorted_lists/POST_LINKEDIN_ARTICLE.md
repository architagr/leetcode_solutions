---
meta_title: "The dummy node trick, and why merging needs it"
meta_description: "Merging two sorted lists is easy until you ask which node goes first. A throwaway head removes that special case entirely, and the trick generalises."
---

![Day 22](HERO.png)

## 365 Days of LeetCode Challenge — Day 22/365

**[21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)** (Easy)

Two sorted lists in, one sorted list out, built by splicing together the nodes that already exist.

## The idea is the easy part

At any moment, the smallest node not yet placed is the head of one list or the head of the other. It cannot be anywhere else: both lists are sorted, so nothing sitting behind a head is smaller than that head.

So compare the two heads, take the smaller, advance that list, repeat.

## The annoying part is the first node

Until something has been placed there is no previous node to attach to. The straightforward version ends up writing the pick-the-smaller logic twice: once before the loop to choose the head, and once inside the loop for everything else.

A throwaway node removes that entirely.

```go
dummy := &ListNode{}
tail := dummy
```

`tail` now refers to something real from the very first iteration, so the loop never has to ask whether this is the first node. At the end the answer is `dummy.Next`, and the dummy is discarded.

This is the part of the problem worth keeping. It is not specific to merging: whenever you build a linked list front to back, a dummy head deletes the special case for the first element.

## The function

```go
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
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

## Watching it run

`list1 = [1,2,4]`, `list2 = [1,3,4]`.

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

`1 <= 1`, so list1's node is taken.

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

list1 runs out, and everything left in list2 is attached in a single assignment.

## Attaching the remainder, not walking it

```go
if list1 != nil {
	tail.Next = list1
} else {
	tail.Next = list2
}
```

One pointer assignment, not a loop. That is correct because two things are true at once by this point: the leftover list is still sorted, and every node in it is at least as large as everything already placed, because the only way its head survived the loop is by losing every comparison it entered.

The obvious version keeps looping node by node until both lists are empty. It produces the same answer with n extra assignments.

## Three nil checks that are not needed

An earlier version of this opened with guards for "both empty", "list1 empty" and "list2 empty". None of them earn their place.

Both empty: the loop never runs, the `else` sets `tail.Next = nil`, and `dummy.Next` is `nil`. One empty: the loop never runs and the other list is attached whole. The general path already handles every empty case, and the guards were three branches restating it.

## `<=` rather than `<`

On ties this takes from `list1`, which keeps equal values in their original relative order. Using `<` still produces a correctly sorted list, so no test that checks sortedness will catch the difference.

What it changes is stability, and stability is the property you care about when this merge is the combining step inside a merge sort, which is the usual reason to have written it at all.

## Complexity

- **Time: O(n + m)**. Every node inspected once, remainder attached without traversal.
- **Space: O(1)**. One dummy node and two pointers. The problem says "splicing", and splicing is what keeps this constant.

Full code and the step-by-step walkthrough:
[merge_two_sorted_lists](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/merge_two_sorted_lists/SOLUTION.md)

#DSA #LeetCode #Golang #LinkedList #MergeSort #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
