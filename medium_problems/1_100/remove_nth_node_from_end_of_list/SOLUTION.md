# Remove Nth Node From End of List — solution walkthrough

From `remove_nth_node_from_end_of_list.go`:

```go
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	if count == 1 {
		return nil
	} else if count == n {
		head = head.Next
	} else {
		temp = head
		x := 1
		for x < count-n {
			x++
			temp = temp.Next
		}
		temp.Next = temp.Next.Next
	}
	return head
}
```

## Pass one: the length

```go
temp := head
count := 0
for temp != nil {
	count++
	temp = temp.Next
}
```

![Step 1](images/walkthrough-1.png)

This is the translation step. "nth from the end" is a position measured from somewhere you cannot start; `count` converts it into `count - n`, a position measured from the head, which is somewhere you can.

Nothing clever happens here, and that is the point. The clever part of this problem is entirely in the second pass.

## The two cases that skip the walk

```go
if count == 1 {
	return nil
}
```

One node, and the constraints guarantee `1 <= n <= sz`, so `n` must be 1 and that node is the one being removed. Nothing is left.

```go
} else if count == n {
	head = head.Next
}
```

`n` equal to the length means the node to remove is the head itself. There is no node before the head to relink, so the removal is expressed as moving the head instead.

This is the case day 22's dummy node would have erased. Put a dummy in front and the head has a predecessor like every other node, and this branch disappears into the general path. Handling it explicitly is a fair choice, and it is worth recognising that it is a choice.

## Pass two: stop one short

```go
temp = head
x := 1
for x < count-n {
	x++
	temp = temp.Next
}
temp.Next = temp.Next.Next
```

![Step 2](images/walkthrough-2.png)

With `[1,2,3,4,5]` and `n = 2`, the target is `count - n = 3`, and the loop leaves `temp` on the third node.

The loop is doing something specific and easy to get wrong: it stops *before* the node being removed, not on it. In a singly linked list a node cannot remove itself. Removal means the previous node points past it, so the previous node is the one you have to be holding.

`x` starting at 1 rather than 0 is what makes `temp` land on the predecessor. `x` counts which node `temp` is currently on, one-based, so the loop runs until `temp` is on node `count - n`, whose `Next` is the node to drop.

![Step 3](images/walkthrough-3.png)

```go
temp.Next = temp.Next.Next
```

One assignment. Node 4 is still sitting in memory with its `Next` intact, but nothing in the list refers to it any more, so it is gone as far as the list is concerned and Go's garbage collector will deal with the rest.

![Step 4](images/walkthrough-4.png)

## Why `head` is returned rather than `temp`

`temp` is somewhere in the middle of the list by the end. The function returns `head`, which is unchanged except in the `count == n` branch that explicitly moved it.

That branch is the whole reason this function has to return anything at all. If the head could never change, the caller's pointer would still be valid and the removal could be done without a return value.

## The one-pass version the follow-up asks for

Advance one pointer `n` nodes ahead of another, then move both until the leading one reaches the end. The trailing pointer is then `n` from the end, and if you offset it by one more it lands on the predecessor.

It is the same fixed-gap idea as day 20, with the gap chosen as `n` instead of being created by moving at different speeds. It is not fewer node visits, and it is not asymptotically faster. It is one traversal instead of two, and it is the more elegant answer.

## Complexity

- **Time: O(n).** One full pass to count, then at most one partial pass to the predecessor.
- **Space: O(1).** A pointer and two integers.

## Test

`main_test.go` covers the three worked examples: `[1,2,3,4,5]` with `n = 2` giving `[1,2,3,5]`, the single node `[1]` with `n = 1` giving an empty list, and `[1,2]` with `n = 1` giving `[1]`. The second exercises the `count == 1` branch and the third exercises the general path with the shortest possible walk.
