# Reverse Linked List — intuition

## Builds on

- [Day 10: Flatten Binary Tree to Linked List](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/flatten_binary_tree_to_linked_list/) — the same move: rewriting `Next` pointers to change a structure's shape without moving a single node

## The problem in one line

You get the head of a singly linked list. Return the same list, pointing the other way.

## The thing worth noticing

A linked list is not a sequence of values. It is a set of nodes, each holding one value and one pointer, and the *order* lives entirely in the pointers. Nothing about node 1 says "I am first". It is first because nothing points at it.

That means reversing the list is not a data-moving problem at all. Every node stays exactly where it is in memory. You walk the list once and turn each arrow around, and the order falls out of that.

Day 10 was the same move on a different structure. Flattening a binary tree into a linked list did not build a new list; it rewrote `Next` pointers until the tree's shape was a line. Here the structure is already a line, and rewriting the same pointers puts it in the other order. If you are following the series in order, you have done this once already.

## Why the loop needs three variables

The obvious version is two lines:

```go
head.Next = prev
head = head.Next
```

That does not work, and the reason it does not work is the whole problem. The first line destroys the link the second line needs. Once `head.Next` points at `prev`, the node you wanted to visit next is unreachable. You have overwritten the only pointer to it.

So you save it first:

```go
next := head.Next
head.Next = prev
prev = head
head = next
```

Four lines, three variables, and the order of them matters more than anything else here. `next` exists for exactly one reason: to hold onto the rest of the list across the line that breaks the link to it.

I like this problem because the fix is so small and so unforgiving. Swap any two of those lines and you either lose the tail or loop forever.

## Edge cases that come for free

An empty list and a single-node list both work without a special case. The loop body never runs for `nil`, so `prev` is still `nil` and that gets returned. For one node, the body runs once, points that node at `nil`, and returns it. No `if` needed for either, which is a good sign the loop is doing the right thing rather than the convenient thing.

## Complexity

- **Time: O(n).** One pass, one visit per node, constant work each.
- **Space: O(1).** Three pointers, regardless of list length. Nothing is allocated.

That O(1) is the point of the question. Copying the values into a slice, reversing the slice, and rebuilding a list is also O(n) time and it is much easier to write, but it is O(n) space, and an interviewer asking 206 is asking whether you can do it without the copy.
