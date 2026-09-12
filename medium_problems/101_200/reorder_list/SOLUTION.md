# Reorder List — solution walkthrough

From `main.go`:

```go
func reorderList(head *ListNode) {
	mid := getMid(head)
	head2 := mid.Next

	mid.Next = nil
	head2 = reverseList(head2)

	head = mergeList(head, head2)
}
```

Six lines, and three of them are calls to functions that are earlier problems in this arc. The whole solution is the observation that the target order is a merge.

## Reading the target order

The problem asks for `L0 -> Ln -> L1 -> Ln-1 -> L2 -> ...`.

Written as two sequences it is:

- `L0, L1, L2, ...` from the front
- `Ln, Ln-1, Ln-2, ...` from the back

Interleaved. So this is a merge of the list with its own reverse, and the only difficulty is that a singly linked list has no backward walk. Day 23 already answered that: reverse the back half.

## `getMid`

```go
func getMid(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
```

![Step 1](images/walkthrough-1.png)

This is day 20's walk in the variant that stops a node early, testing `fast.Next` and `fast.Next.Next` rather than `fast` and `fast.Next`.

That is required here. The node being returned is the one whose `Next` is about to be set to `nil`, so it has to be the *last node of the first half*. Day 20's own condition returns the first node of the second half, which is one node too far.

It also settles where the extra node goes on an odd-length list: it stays in the first half. For `[1,2,3,4,5]` the expected output is `[1,5,2,4,3]`, with 3 last, and that only works if the first half is the longer one.

## The cut

```go
head2 := mid.Next
mid.Next = nil
```

![Step 2](images/walkthrough-2.png)

Save the second half's head, then terminate the first half.

That `nil` is not tidiness. Without it the first half runs straight into the second, and since the merge walks both at once, the two walks trample each other and the result is a cycle or a mess. One assignment, and the solution is wrong without it.

## `reverseList`

```go
head2 = reverseList(head2)
```

![Step 3](images/walkthrough-3.png)

Day 19's walk, with `next` hoisted out of the loop rather than declared inside it. Point each node at its predecessor, return the node that ends up first.

After this, walking `head2` forward visits the original list's tail end backwards, which is the sequence the target order needs.

## `mergeList`

```go
func mergeList(a1, a2 *ListNode) *ListNode {
	if a2 == nil {
		return a1
	}
	t, t1, t2 := a1, a1.Next, a2.Next

	for a2 != nil {
		t.Next = a2
		a2.Next = t1
		t = t1
		a2 = t2
		if t1 != nil {
			t1 = t1.Next
		}
		if t2 != nil {
			t2 = t2.Next
		}
	}
	return a1
}
```

![Step 4](images/walkthrough-4.png)

This is the same alternating relink as day 22, with one difference that matters: day 22 chose which list to take from by comparing values, and here the alternation is unconditional. Front, back, front, back.

`t1` and `t2` hold each list's next node across the relinking. That is the day 19 problem again — save the pointer before the assignment that destroys it — except now there are two lists to keep alive at once, which is why there are two saved pointers rather than one.

The `if t1 != nil` and `if t2 != nil` guards are there because the halves can differ in length by one. When the shorter one runs out, advancing its saved pointer would dereference `nil`.

![Step 5](images/walkthrough-5.png)

## Why the function returns nothing

`reorderList` has no return value, and it does not need one.

Every operation relinks nodes that already exist. No node is allocated, no node moves, and in particular the node the caller's `head` points at is still the first node of the reordered list. The reassignment `head = mergeList(...)` on the last line is assigning to a local parameter and has no effect outside the function, which is fine precisely because the caller's pointer is already correct.

## Complexity

- **Time: O(n).** A pass to find the middle, a pass over half the list to reverse it, a pass to merge.
- **Space: O(1).** A handful of pointers across three small functions. Nothing allocated.

## Test

`main_test.go` covers both worked examples: `[1,2,3,4]` becoming `[1,4,2,3]` for the even case, and `[1,2,3,4,5]` becoming `[1,5,2,4,3]` for the odd one. The odd case is the one that fails if `getMid` uses day 20's loop condition instead of this one.
