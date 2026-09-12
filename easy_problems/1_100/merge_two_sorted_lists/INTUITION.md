# Merge Two Sorted Lists — intuition

## The problem in one line

Two sorted lists in, one sorted list out, built by splicing together the nodes that already exist.

## The idea

At every moment, the smallest node not yet placed is the head of one list or the head of the other. It cannot be anywhere else, because both lists are sorted, so nothing behind a head is smaller than that head.

So the algorithm writes itself: compare the two heads, take the smaller, advance that list, repeat. When one list runs out, everything left in the other is already sorted and already larger than everything placed, so it gets attached in one move rather than walked.

That last part is easy to miss. The obvious loop keeps going node by node until both lists are empty, which works but does n extra pointer assignments to achieve what one assignment achieves.

## The dummy head

The irritating part of this problem is not the comparison, it is the first node. Until something has been placed there is no "previous node" to attach to, so the naive version ends up with two copies of the pick-the-smaller logic: one to choose the head, and one inside the loop.

A throwaway node solves it:

```go
dummy := &ListNode{}
tail := dummy
```

`tail` now always refers to something real, from the very first iteration, so the loop never has to ask whether this is the first node. At the end the answer is `dummy.Next` and the dummy itself is discarded.

That is the trick worth taking from this problem. It is not specific to merging. Any time you are building a linked list front to back, a dummy head removes the special case for the first element.

## Splicing, not copying

The problem statement says the result "should be made by splicing together the nodes of the first two lists". That is a real constraint, not decoration.

Copying the values into new nodes gives the same output and costs O(n) extra space. Relinking the nodes that already exist costs nothing beyond the two pointers, and it is the same move as day 19: change what points where, and the structure changes without anything being allocated.

## Why `<=` and not `<`

```go
if list1.Val <= list2.Val {
```

With equal values, `<=` takes from `list1` first, which keeps nodes from the first list ahead of equal nodes from the second. That is what makes the merge stable, and stability is the property that matters when this merge is the inner step of a merge sort.

Using `<` still produces a correctly sorted list. It just quietly reverses the relative order of equal elements.

## Complexity

- **Time: O(n + m).** Every node is looked at once, and the leftover tail is attached without being walked.
- **Space: O(1).** Two pointers and one dummy node. No allocation proportional to the input.
