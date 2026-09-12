# Next Greater Element I — intuition

## Builds on

- [Day 58: Valid Parentheses](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/valid_parentheses/) — the same "the top of the stack is the only thing that can be resolved right now" reasoning, applied to values instead of brackets

## The problem in one line

For each value in `nums1`, find the first value to its right in `nums2` that is larger.

## The brute force, and what it wastes

For each element, scan rightwards until you find something bigger. O(n²), and correct.

What it wastes is that the scans repeat each other. If you scan right from index 3 and pass over a long descending run, then scan right from index 4 and pass over almost the same run, you are re-reading values you already know are too small.

## Turning the question round

Instead of asking "what is to the right of me", ask, for each value as you arrive at it, "**which earlier values does this one answer?**"

That flips the work. A value arriving at position `i` resolves every earlier value smaller than it that is still unanswered. And once resolved, those values never need to be looked at again.

## What needs remembering

The values seen so far that have not yet found anything bigger.

They have a useful property: they are **decreasing**. If an earlier value were smaller than a later one, the later one would have resolved it on arrival, so it would not still be waiting.

So the pending set is always a decreasing sequence, and the incoming value resolves a prefix of it from the top down.

Most recent in, first out again. A stack, and one that maintains an ordering invariant — which is why it is called a monotonic stack.

## Why this is linear

Every value is pushed exactly once and popped at most once. The inner loop can run many times on one iteration, but every one of those iterations removes something permanently.

That is the difference from the brute force: the brute force re-reads values, this one consumes them.

## The leftovers

Whatever is still on the stack at the end never met anything larger, so those get `-1`.

## Why the map

`nums1` is a subset of `nums2` and in a different order, so the answers are computed against `nums2` once and then read off by value. The problem guarantees every value is unique, which is what makes a value-keyed map safe here.

## Complexity

- **Time: O(n + m).** Each value in `nums2` pushed and popped once, then one pass over `nums1`.
- **Space: O(n)** for the stack and the map.
