# Reverse Linked List II — intuition

## Builds on

- [Day 19: Reverse Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/reverse_linked_list/) — the reversal itself, reused here without a single change

## The problem in one line

Reverse only the nodes between positions `left` and `right`, and leave everything else alone.

## Why this is a Medium and day 19 was an Easy

The reversal is identical. What makes this harder is that the reversed run has to be sewn back into a list on both sides, and the two seams are not symmetric.

The front seam is a node you walked past and have to remember. The back seam is a node you have to find *after* reversing, because reversing moved it.

## The approach

The neat way to think about it: make the middle run look like an ordinary standalone list, reverse it with the function you already have, and then reattach it.

1. Walk to `left`, keeping a pointer one step behind. That trailing pointer is the front seam.
2. Walk on to `right`. Save what comes after it, then cut the list there.
3. Reverse the detached run with day 19's function, unchanged.
4. Reattach: front seam points at the run's new head, and the run's new tail points at the saved remainder.

Step 3 is the reason for steps 1 and 2. By detaching the run first, the reversal needs no notion of "stop at position right" built into it, so day 19's function is reused literally rather than adapted.

## The two seams

**The front.** `leftNode` trails one step behind while walking to `left`. If `left` is 1 there is no node before the run, so `leftNode` stays `nil`, and that is the signal to update `head` instead of a `Next` pointer.

This is the same "the head has no predecessor" case as days 22 and 24. A dummy node in front would erase it here too.

**The back.** After reversing, the node that was the run's first is now its last. You cannot save a pointer to the tail beforehand, because which node ends up at the tail is exactly what the reversal changes. So the code walks to the end of the reversed run to find it.

That walk is the price of reusing the reversal as-is. Worth paying, I think, for a function that is already known to be correct.

## Complexity

- **Time: O(n).** Walk to `right`, reverse a run shorter than the list, walk the run once more to find its tail.
- **Space: O(1).** A handful of pointers.
