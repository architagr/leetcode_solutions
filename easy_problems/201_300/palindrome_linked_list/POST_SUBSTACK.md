---
meta_title: "A palindrome check with no way to walk backwards"
meta_description: "A singly linked list cannot be read right to left, so reverse half of it. Both halves of this solution are problems from earlier in the week."
tags: [golang, linked-list, two-pointers, dsa]
---

![Day 23](HERO.png)

*365 Days of LeetCode Challenge — Day 23/365*

**[234. Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)** (Easy)

This is the day the arc pays off, so I want to say up front what is about to happen: today's solution is day 19 and day 20, run one after the other, with nothing new added.

That is the actual lesson. Most hard-feeling problems are not a new idea, they are two old ideas standing next to each other.

## Why the array version does not port

On an array, checking for a palindrome takes about four lines. Put an index at each end, walk them toward each other, compare as you go, stop when they cross.

Try to write that for a singly linked list and you get stuck immediately, at the part where the right-hand index moves left. There is no index. There is no pointer to the previous node. A node in a singly linked list knows what comes after it and has no information whatsoever about what came before.

The array algorithm is not merely slower here. It cannot run.

## The solution that dodges the question

Walk the list, push every value into a slice, then run the array algorithm on the slice.

This is a correct solution. O(n) time, O(n) space, easy to write, easy to get right. It is also the solution that made me think this was a boring problem the first time I saw it.

The follow-up on the problem page is what makes it interesting: can you do it in O(n) time and O(1) space?

## The reframe

Read the obstacle again. The problem is that you cannot walk the second half backwards.

You know how to make a list run the other way. That was day 19.

So: reverse the second half. Then walking it *forwards* is walking the original *backwards*, and the array algorithm becomes available after all.

Three steps, two of them already written:

1. Find the middle, with day 20's fast/slow walk.
2. Reverse everything after it, with day 19's in-place reversal.
3. Walk both halves inward, comparing.

## The function

```go
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := reverse(slow.Next)

	left, right := head, second
	ok := true
	for right != nil {
		if left.Val != right.Val {
			ok = false
			break
		}
		left = left.Next
		right = right.Next
	}

	slow.Next = reverse(second)
	return ok
}
```

## That loop condition is not the one from day 20

```go
for fast.Next != nil && fast.Next.Next != nil {
```

Day 20's version tested `fast` and `fast.Next`, and left `slow` on the first node of the second half. This one tests a node earlier and leaves `slow` on the last node of the first half.

Day 20's write-up flagged this variant and said it was the one worth memorising because harder problems usually want it. Here is the harder problem, one arc later, wanting it.

The reason is specific: the node this solution needs is the one whose `Next` pointer gets detached and later repaired. That is the boundary node on the *first* half's side. Use day 20's exact condition and `slow` sits one node too far along, and the split is wrong on even-length lists.

## The walk

`[1,2,2,1]`.

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

`slow` stops on the first `2`, the last node of the first half.

![Step 3](images/walkthrough-3.png)

One call to `reverse` and the back half runs the other way.

Look at what `slow.Next` points to now: the node that used to start the second half and is now its end. That looks like a stale pointer and it is not. It is the handle the restore step needs, and leaving it alone is deliberate.

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

Both pairs match, `right` runs out, and the answer is true.

## Why the shorter half drives the loop

```go
for right != nil {
```

This looks like an arbitrary choice between two pointers. It is not.

On an odd-length list the two halves are different lengths. The middle node belongs to the first half and has no partner in the second, because there is nothing for it to pair with.

Because `right` is the half that runs out first, that middle node is never compared. Which is correct: in a palindrome the middle element has only itself to match, and it matches itself trivially.

There is no parity check in this function. No `if length % 2`. The asymmetry in the data is handled by choosing which pointer ends the loop, and if you drive it with `left` instead you dereference `nil` on every odd-length input.

## Putting the list back

![Step 6](images/walkthrough-6.png)

```go
slow.Next = reverse(second)
```

Reversing the second half again restores its direction, and reattaching it to `slow` makes the caller's list exactly what it was when it arrived.

I want to spend a moment on this, because it is the part that separates a solution that passes from a solution that is right.

The function is called `IsPalindrome`. It takes a list and returns a bool. Nothing in that signature, and nothing in that name, tells a caller their data will be rearranged. But without the restore line, half of their list is left pointing the wrong way, and the bug does not show up in the function that caused it. It shows up somewhere else entirely, in whatever code touches that list next, which is the worst place for a bug to appear.

It costs one more O(n/2) pass and changes no complexity class. An interviewer asking this question is, I would guess, at least as interested in whether you notice the mutation as in whether you find the reversal trick, because the reversal trick is memorable and noticing side effects is a habit.

One more detail in the same spirit: the result is stored in `ok` and returned after the repair, rather than returned from inside the comparison loop. An early `return false` on the first mismatch would leave the list broken on exactly the inputs that fail, which is a genuinely nasty way to be wrong.

## Complexity

- **Time: O(n).** Four passes over roughly half the list each: find the middle, reverse, compare, restore. That is 2n operations at worst, which is linear.
- **Space: O(1).** Five pointers and a bool, regardless of list length.

The slice version is O(n) time and O(n) space. This one is what the follow-up was asking for, and it is assembled entirely from parts this series already had lying around.

## Builds on

- [Day 19: Reverse Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/reverse_linked_list/) — the in-place reversal, used here on half the list
- [Day 20: Middle of the Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/middle_of_the_linked_list/) — the fast/slow walk that finds where to split

Full code and the step-by-step walkthrough:
[palindrome_linked_list](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/palindrome_linked_list/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
