---
meta_title: "Ask which earlier values this one answers"
meta_description: "The brute force scans right and re-reads the same descending runs. Flipping the question turns it into one pass with a monotonic stack."
tags: [golang, stack, monotonic-stack, dsa]
---

![Day 59](HERO.png)

*365 Days of LeetCode Challenge — Day 59/365*

**[496. Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/)** (Easy)

Given two arrays where `nums1` is a subset of `nums2`, find for each value in `nums1` the first value to its right in `nums2` that is larger. If there is none, report `-1`.

## The brute force is fine, and it is worth knowing what it wastes

For each element, walk rightwards until you find something bigger. That is O(n x m), it is easy to write, and it is correct.

The waste in it is specific rather than vague, and naming it is what points at the better solution.

Suppose the array contains a long descending run: `9, 8, 7, 6, 5`. Start at the 9 and you walk the entire run without success. Start at the 8 and you walk almost the same run. Start at the 7 and again. Every one of those walks re-reads values that a previous walk already established were too small.

The information exists. The algorithm just throws it away between iterations.

## Turning the question round

Here is the reframe, and it is the whole solution.

The brute force asks, standing at position `i`: *what is to the right of me?* That question can only be answered by looking forward, which means scanning.

Ask instead, as each value arrives: *which earlier values does this one answer?*

That question can be answered immediately, because the values it answers are the ones already behind you. A value arriving resolves every earlier value smaller than it that is still unanswered — and once resolved, a value never needs looking at again.

Same problem. One direction of asking requires a search; the other requires only a lookup at the right place.

## What has to be remembered

The values seen so far that have not yet found anything larger.

Now the observation that names the technique. That set is always **decreasing from bottom to top**.

Why: suppose it were not, and some earlier value `a` sat below a later value `b` with `a < b`. But `b` arrived after `a`, and when it arrived it would have resolved `a` — `b` is larger, and `a` was waiting for exactly that. So `a` would have been removed. It cannot still be sitting underneath.

The invariant is not imposed. It falls out of the rule for removing things.

A stack that maintains an ordering invariant this way is called a **monotonic stack**, and this is the first of two days in a row that use one.

## The loop

```go
for _, v := range nums2 {
	for len(stack) > 0 && stack[len(stack)-1] < v {
		next[stack[len(stack)-1]] = v
		stack = stack[:len(stack)-1]
	}
	stack = append(stack, v)
}
```

![Step 1](images/walkthrough-1.png)

Nothing to compare against yet, so the first value waits.

![Step 2](images/walkthrough-2.png)

3 arrives and beats the waiting 1. That is 1's answer, recorded, and 1 leaves the stack for good.

![Step 3](images/walkthrough-3.png)

The same thing happens to 3 when 4 arrives.

![Step 4](images/walkthrough-4.png)

2 beats nothing on the stack, so the inner loop does not run at all and 2 simply stacks on top of 4. The stack is `[4, 2]`, decreasing, exactly as the invariant promises.

![Step 5](images/walkthrough-5.png)

And 5 clears both of them in a single outer iteration: 2 first, then 4.

## The objection that step 5 raises

A nested loop that can run several times per outer iteration looks quadratic. It is reasonable to be suspicious of it.

It is not, and the argument fits in two sentences. Every value in `nums2` is pushed onto the stack exactly once. Every value is popped at most once, because a popped value is never pushed again.

So however unevenly the inner loop's iterations are distributed — nothing for four outer steps, then four in a single step — the total across the entire run is bounded by the number of values. The work is amortised, not avoided.

That is precisely the difference from the brute force. The brute force **re-reads** values it has already rejected, over and over. This one **consumes** them.

## The leftovers

```go
for _, v := range stack {
	next[v] = -1
}
```

Whatever is still waiting when `nums2` runs out never met anything larger. In the running example that is just the 5, which was the maximum.

## One assumption to check before reusing this

The answers live in a map keyed by **value**:

```go
next := make(map[int]int, len(nums2))
```

That is necessary here because `nums1` is a subset of `nums2` in a different order, so positions do not correspond and an array indexed by position would be useless.

It is only *safe* because the problem guarantees all values are unique. Two equal values would share a single map entry, and the second would overwrite the first's answer.

This is the kind of assumption that is easy to inherit without noticing. If you lift this shape into a problem where duplicates are possible — day 61's Daily Temperatures is one, since temperatures repeat — you need to key by index instead.

## Complexity

- **Time: O(n + m).** Each value in `nums2` is pushed once and popped at most once; then one pass over `nums1`.
- **Space: O(n)** for the stack and the map.

## Builds on

- [Day 58: Valid Parentheses](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/valid_parentheses/) — the same "the top of the stack is the only thing that can be resolved right now" reasoning, applied to values instead of brackets

Full code and the step-by-step walkthrough:
[next_greater_element_i](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/401_500/next_greater_element_i/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
