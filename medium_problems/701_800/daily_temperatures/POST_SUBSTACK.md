---
meta_title: "One character decides whether duplicates break this"
meta_description: "Daily Temperatures is yesterday's monotonic stack over indices. Equal temperatures are not warmer, and that is what the <= is for."
tags: [golang, stack, monotonic-stack, dsa]
---

![Day 61](HERO.png)

*365 Days of LeetCode Challenge — Day 61/365*

**[739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)** (Medium)

Given daily temperatures, return for each day how many days you must wait for a warmer one. Zero if no warmer day ever comes.

## Yesterday's problem, asked for a distance

Day 59 asked *which* value comes next and is larger. Today asks *how far away* it is.

That sounds like a cosmetic difference and it changes two things, both worth stating explicitly because each one is a decision.

**The stack holds indices rather than values.** A distance is computed from positions, so what goes on the stack is `i`, and the answer is `stack.Top() - i`.

**Values repeat.** Day 59 stored its answers in a map keyed by value, and I flagged at the time that this was only safe because the problem guaranteed uniqueness. Temperatures are bounded between 30 and 100 with up to 100,000 days, so duplicates are not merely possible, they are certain.

Storing by index removes that problem rather than working around it. Indices are unique by construction; there is nothing to guard against.

It is a small illustration of something general: a lot of "handling edge cases" is really choosing a representation that does not have the edge case.

## Scanning the other way

Day 59 scanned left to right and asked, of each arriving value, *which earlier values does this one answer?* Answers were discovered out of order, so they were written into a map.

This scans right to left and asks, standing on each day, *what is the first thing to my right that beats me?*

The difference is when the answer becomes known. Here it is known while standing on the element it belongs to, so it goes straight into `results[i]` and the loop moves on. No map, no deferred bookkeeping, no second pass.

Both directions are monotonic stacks and both are linear. Which one to write depends on where the answer lands.

## The code

```go
func dailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))
	stack := new(Stack)

	for i := len(temperatures) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && temperatures[stack.Top()] <= temperatures[i] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			results[i] = 0
		} else {
			results[i] = stack.Top() - i
		}
		stack.Push(i)
	}
	return results
}
```

![Step 1](images/walkthrough-1.png)

The last day has nothing after it, so its answer is 0. It still goes on the stack, because it is a candidate for every day to its left.

![Step 2](images/walkthrough-2.png)

76 at index 4 beats 71 at index 3, so the wait is `4 - 3 = 1`.

## What the stack actually contains

It holds indices of days, all to the right of the current one, that are still candidates for being somebody's answer.

A day stops being a candidate the instant a warmer-or-equal day appears to its left. Once that happens, everybody further left would reach the nearer day first, so the further one can never be anyone's answer again — and nothing will change that, because we only ever move further left.

That is what the inner loop enforces:

```go
for !stack.IsEmpty() && temperatures[stack.Top()] <= temperatures[i] {
	stack.Pop()
}
```

![Step 3](images/walkthrough-3.png)

Index 2 holds 75. The 71 sitting at index 3 is cooler, so it is eliminated. The next candidate down is the 76 at index 4, giving `4 - 2 = 2`.

After the pops, the stack reads increasing in temperature from top to bottom, and its top is exactly the nearest day warmer than the current one.

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

## The character this problem turns on

```go
temperatures[stack.Top()] <= temperatures[i]
```

Not `<`. This is the single most likely thing to get wrong here, and it fails quietly.

The reasoning is one sentence: an equal temperature is not a *warmer* one. A day holding exactly the same temperature can never satisfy anybody's search, so leaving it on the stack is leaving a wrong answer waiting to be found.

With `<`, equal days survive the pops. The next day to the left then looks at the top of the stack, finds an equal temperature sitting there, and reports it as the warmer day it was looking for.

Concretely: `[70, 70]` should produce `[0, 0]`. Neither day is followed by anything warmer. With `<` it produces `[1, 0]`, claiming day 0 waits one day for warmth that never arrives.

And notice where this lands. Day 59 could ignore duplicates because the problem promised there were none. Today duplicates are guaranteed. The entire handling of that difference is one character, in a comparison that most people write without pausing.

## The zero case

```go
if stack.IsEmpty() {
	results[i] = 0
}
```

An empty stack after the pops means nothing to the right is warmer.

`results` was allocated with `make`, so every entry is already 0 and this assignment writes a value that is already there. It is redundant and I would keep it — relying on Go's zero value works, but saying "no warmer day means zero" out loud is the kind of thing a reader should not have to infer.

## Why the nested loop is still linear

Each index is pushed exactly once and popped at most once. The inner loop may run several times in a single outer iteration, but each of those iterations removes an index permanently.

Same argument as yesterday, and worth repeating precisely because it keeps looking wrong. A nested loop reads as quadratic until you count what it *consumes* rather than what it visits.

## A note on the stack

The file defines its own linked-list `Stack` with `Push`, `Pop`, `Top`, `IsEmpty` and `Count`. Go has no stack in its standard library, and the common idiom is a slice: `s[len(s)-1]` to peek, `s = s[:len(s)-1]` to pop, which days 58 and 62 use.

Both are fine. The named methods make the main loop read close to the English description of the invariant, which for this problem is worth something.

## Complexity

- **Time: O(n).** Each index is pushed once and popped at most once.
- **Space: O(n)** for the stack. The worst case is a strictly decreasing sequence, where nothing is ever popped and every index piles up.

## Builds on

- [Day 59: Next Greater Element I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/401_500/next_greater_element_i/) — the same monotonic stack, holding indices here because the answer is a distance and values can repeat

Full code and the step-by-step walkthrough:
[daily_temperatures](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/701_800/daily_temperatures/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
