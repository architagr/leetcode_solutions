# Middle of the Linked List — intuition

## The problem in one line

Return the middle node of a singly linked list. If the length is even, return the second of the two middles.

## The obvious solution, and what is wrong with it

Walk the list counting nodes, then walk it again to position `count/2`. Two passes, O(n) time, O(1) space, and it is a completely acceptable answer.

But it needs the length before it can do anything, and there is a way to avoid ever computing it.

## The trick

Run two pointers from the head. Move one of them one node per step and the other two nodes per step. When the fast one reaches the end, the slow one has travelled exactly half as far, because it moved at exactly half the speed over the same elapsed time.

That is the whole idea. It is not a clever fact about linked lists, it is just arithmetic. The reason it is worth knowing is that it gets the middle in one pass without ever knowing the length, and a singly linked list is exactly the structure where you cannot cheaply ask for the length.

I like this one because the code is shorter than the explanation:

```go
for end != nil && end.Next != nil {
    middle = middle.Next
    end = end.Next.Next
}
```

## The two conditions in the loop

Both are load-bearing, and they are guarding different things.

`end != nil` catches an even-length list. The fast pointer steps onto `nil` exactly, and without this check the next `end.Next` dereferences nothing.

`end.Next != nil` catches an odd-length list. The fast pointer lands on the final node, where `end.Next.Next` would walk off the end.

Drop either and you get a nil dereference on half of all inputs, which is the kind of bug that passes the first test case you try.

## Why the even case lands on the second middle for free

For `[1,2,3,4,5,6]` the fast pointer visits nodes 1, 3, 5, then steps to `nil`. The slow pointer visits 1, 2, 3, 4 and stops on 4. That is the second of the two middles, which is what the problem asked for, and no special case was written to make it happen. It falls out of the loop condition.

If the problem had wanted the first middle, the loop would have to be `end.Next != nil && end.Next.Next != nil` instead. That is worth remembering, because the version of this walk that appears inside harder problems usually wants the first middle, not the second.

## Complexity

- **Time: O(n).** The fast pointer touches about half the nodes and the slow one about half again, so it is a single pass in total.
- **Space: O(1).** Two pointers.
