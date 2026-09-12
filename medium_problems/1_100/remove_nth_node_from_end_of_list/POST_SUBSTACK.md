---
meta_title: "Removing a node you are not allowed to be standing on"
meta_description: "A node in a singly linked list cannot remove itself. The whole difficulty of this problem is stopping one node early, and the off-by-one that hides there."
tags: [golang, linked-list, two-pointers, dsa]
---

![Day 24](HERO.png)

*365 Days of LeetCode Challenge — Day 24/365*

**[19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)** (Medium)

Remove the nth node counting from the end of the list, and return the head.

This is marked Medium and the algorithm is barely harder than yesterday's. What makes it worth its rating is that there are two separate places to be off by one, and both of them produce code that looks right.

## The position is stated from the wrong end

"nth from the end" describes a location relative to somewhere you have no access to.

A singly linked list hands you one thing: the head. The end of the list is not a place you can address. It is something that happens after enough steps, and finding out how many steps takes a full traversal.

So before any removal can happen, the position has to be translated. If the list has `count` nodes, the nth from the end is the `(count - n)`th from the front, and the front is the one place you can actually start.

```go
temp := head
count := 0
for temp != nil {
	count++
	temp = temp.Next
}
```

![Step 1](images/walkthrough-1.png)

Nothing interesting happens in that loop, which is appropriate. It exists to turn an unusable coordinate into a usable one.

Two passes over the list, O(n) time, O(1) space. The follow-up asks whether one pass is possible, and it is, and I will get to it. I want to note first that the one-pass version is not faster in any meaningful sense. It visits the same nodes. The improvement is in elegance, not in work done, and it is worth being honest about which of those you are optimising.

## The constraint that makes this a real problem

To remove a node from a singly linked list, you must be holding the node before it.

This sounds like a detail and it is the whole problem. A node has exactly one pointer and it points forward. It has no idea what refers to it, so it cannot ask that thing to point somewhere else. A node genuinely cannot remove itself from a singly linked list, and that is not a limitation of any particular implementation, it is what the data structure is.

Removal is an operation the predecessor performs.

So the second walk has to deliberately stop short of its target:

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

With `[1,2,3,4,5]` and `n = 2`, the arithmetic gives `count - n = 3`, and the loop leaves `temp` on the third node, holding the node whose `Next` gets rerouted.

`x` starting at 1 rather than 0 is the detail that makes that land. It counts which node `temp` is currently on, one-based, so the loop exits with `temp` on node `count - n`. Start it at 0 and you overshoot by one and delete the wrong node, and on most test inputs you will still get a list back that looks plausible.

![Step 3](images/walkthrough-3.png)

```go
temp.Next = temp.Next.Next
```

One assignment. Node 4 has not been destroyed; it is still sitting in memory with its own `Next` pointing at node 5. It is simply that nothing in the list refers to it any more, and in a structure defined entirely by its pointers, being unreferenced *is* being removed. Go's garbage collector handles the rest without being asked.

![Step 4](images/walkthrough-4.png)

## Removing the head

```go
} else if count == n {
	head = head.Next
}
```

When `n` equals the length, the node to remove is the head, and the rule above bites: there is no node before the head to perform the removal.

So this case gets expressed differently. Rather than unlinking, the head simply moves.

I want to connect this back to day 22, because the two problems are having the same argument and reaching different conclusions. Merging two sorted lists had the same difficulty, in the form of "there is no previous node to attach the first element to", and it solved it by allocating a dummy node in front so that the first element stopped being special.

The same fix works here. Put a dummy in front of the head, and the head acquires a predecessor like every other node, and this entire branch folds into the general path.

This solution takes the explicit branch instead. That is a defensible choice, and the point of raising it is not that the branch is wrong. It is that after day 22 you should recognise the shape of the problem and know that a tidier option exists, so that taking the branch is a decision rather than the only thing that occurred to you.

## The one-pass version

Put two pointers on the list with a gap of exactly `n` nodes between them. Advance both together. When the leading pointer runs off the end, the trailing pointer is `n` nodes from the end, and with the gap set one wider it is sitting on the predecessor.

That is day 20's idea with one change: the gap is created deliberately at the start rather than produced by moving at different speeds. Both are the same underlying trick, which is that the relationship between two pointers can encode information that neither pointer holds alone.

It touches the same nodes. It is one traversal rather than two. It is the better answer and it is not the faster one.

## Complexity

- **Time: O(n).** A full pass to count, then a partial pass to reach the predecessor.
- **Space: O(1).** One pointer and two integers.

## Builds on

- [Day 20: Middle of the Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/middle_of_the_linked_list/) — the same problem shape: a position defined relative to the end of a list you can only walk forward

Full code and the step-by-step walkthrough:
[remove_nth_node_from_end_of_list](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1_100/remove_nth_node_from_end_of_list/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
