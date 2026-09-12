---
meta_title: "Reversing part of a list, and the two seams it leaves"
meta_description: "The reversal is day 19's, unchanged. What makes this a Medium is sewing the reversed run back in, and the two seams are not symmetric."
tags: [golang, linked-list, pointers, dsa]
---

![Day 25](HERO.png)

*365 Days of LeetCode Challenge — Day 25/365*

**[92. Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/)** (Medium)

Reverse only the nodes between positions `left` and `right`. Everything outside that range stays exactly as it was.

## The reversal is not the hard part

I want to get this out of the way first, because the problem's name suggests otherwise.

The reversal in this solution is the function from day 19. Not a variant of it, not adapted for a range. The same function, character for character, called once.

If reversing a whole list is an Easy and reversing part of one is a Medium, and the reversal itself is identical, then the difficulty is entirely in the difference between those two problems. That difference is the seams.

## Two seams, two different problems

A reversed run sitting in the middle of a list has to be joined back on at both ends, and the two ends are not symmetric.

**The front seam** is a node you pass on the way in. You see it before you need it, so the only requirement is to remember it. That is a matter of keeping a trailing pointer.

**The back seam** is a node you cannot see before you need it. It is the node the reversed run ends on, and which node that is is precisely what the reversal decides. You cannot save a pointer to it in advance, because in advance it does not exist yet.

Noticing that asymmetry is most of understanding this problem.

## The approach

Make the run look like an ordinary list, reverse it with the tool you have, reattach it.

1. Walk to `left`, keeping a pointer one step behind.
2. Walk on to `right`, save what comes after it, and cut the list there.
3. Reverse the detached run.
4. Reattach the front, then find the new tail and reattach the back.

## Walking in, and keeping the trailing pointer

```go
for tempHead != nil {
	count++
	if count == left {
		break
	}
	leftNode = tempHead
	tempHead = tempHead.Next
}
leftSide := tempHead
```

![Step 1](images/walkthrough-1.png)

Two pointers come out of this. `tempHead` sits on position `left`, the run's first node. `leftNode` sits on the node before it, kept by being assigned *before* the step forward rather than after.

`leftSide` names the run's first node, which is worth flagging now because it will not stay the first node for long.

## The cut, which is the good idea

```go
for tempHead != nil {
	if count == right {
		break
	}
	count++
	tempHead = tempHead.Next
}

rightNode = tempHead.Next
tempHead.Next = nil
```

![Step 2](images/walkthrough-2.png)

`rightNode` saves whatever followed the run, and then `tempHead.Next = nil` severs it.

This is the decision the whole solution rests on. After that single assignment, the stretch from `leftSide` to `tempHead` is a complete, ordinary, nil-terminated linked list. Nothing about it says "I am the middle of something".

Which means the next line is allowed to be:

```go
leftSide = reverse(leftSide)
```

![Step 3](images/walkthrough-3.png)

Day 19's function. No bounds parameters, no counter, no awareness that it is operating on a fragment. It walks until it hits `nil` and `nil` is exactly where we put it.

The alternative is a reversal that takes `left` and `right` and stops in the right places, and that function is harder to write, harder to test, and does not already exist. Cutting first converts a new problem into one that is already solved, which is the same move day 23 made.

Note what happens to `leftSide` across that line. It went in naming the run's first node and comes out naming the run's new first node, which is a different node. The node it used to name is now the run's tail, and it no longer has a name. That is the back seam problem, arriving.

## Reattaching the front

```go
if leftNode == nil {
	head = leftSide
} else {
	leftNode.Next = leftSide
}
```

![Step 4](images/walkthrough-4.png)

`leftNode` is `nil` exactly when `left` is 1, because the walk broke on its first iteration before assigning anything. That nil is carrying information: it means the run started at the head, so there is no predecessor, so the thing to update is `head` itself.

This is the third appearance of "the head has no predecessor" in this arc. Day 22 met it while merging and erased it with a dummy node. Day 24 met it while removing and took an explicit branch. Today takes an explicit branch as well.

A dummy node in front of `head` would collapse this `if` into its `else` here too, and the fact that the same fix keeps applying is the thing to take away. Three problems that look unrelated share a structural quirk, and one technique addresses all three.

## Reattaching the back

```go
for leftSide.Next != nil {
	leftSide = leftSide.Next
}
leftSide.Next = rightNode
```

![Step 5](images/walkthrough-5.png)

Since the tail could not be saved in advance, it gets found by walking.

This walk is the price of reusing `reverse` unchanged. A bespoke reversal could have returned the tail alongside the head and saved this pass. It would also have been a second implementation of logic that already exists and is already known to work, and the pass is O(run length) on a run that was just traversed twice anyway.

I would pay the walk every time. The cost is a constant factor on one pass; the saving is not having a second reversal in the codebase that can drift away from the first.

## Complexity

- **Time: O(n).** A walk to `right`, a reversal of the run, a walk back across the reversed run. Three partial passes, linear.
- **Space: O(1).** Five pointers and a counter.

## Builds on

- [Day 19: Reverse Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/reverse_linked_list/) — the reversal itself, reused here without a single change

Full code and the step-by-step walkthrough:
[reverse_linked_list_ii](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1_100/reverse_linked_list_ii/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
