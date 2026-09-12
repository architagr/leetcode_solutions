---
meta_title: "The minimum belongs to the height, not the stack"
meta_description: "A single min field breaks on the first pop. Storing the minimum per entry makes pop free, because the old answer was never overwritten."
---

![Day 60](HERO.png)

## 365 Days of LeetCode Challenge — Day 60/365

**[155. Min Stack](https://leetcode.com/problems/min-stack/)** (Medium)

A stack where `push`, `pop`, `top` and `getMin` are all O(1).

## Two fixes that do not work

**Scan the stack in `getMin`.** O(n) per call. The requirement says constant.

**Keep one `min` field, updated on push.** Push is fine. The first `pop` breaks it: pop the element that *is* the current minimum and the field is stale, with nothing to recompute it from short of a scan.

## That second failure says what the structure needs

It is worth sitting with, because it is the insight.

The minimum is not a property of *the stack*. It is a property of **the stack at a given height**. Popping returns you to a height whose minimum you have already thrown away.

So do not throw it away.

## Store it per entry

```go
type dataNode struct {
	val, min int
}
```

Every entry carries the value pushed **and** the minimum of everything at or below it.

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

0 does not beat -2, so the existing minimum is copied up. Every entry has a correct answer for its own height, even when it duplicates its neighbour's.

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

`getMin` is a field read. No scan, no comparison.

![Step 5](images/walkthrough-5.png)

## The step that makes the design

Look at what `Pop` does to maintain the minimum:

```go
func (this *MinStack) Pop() {
	defer func() { this.len-- }()
	this.dataArr = this.dataArr[:this.len-1]
}
```

Nothing. There is no minimum-maintenance code at all.

Removing the top exposes an entry whose `min` was computed when *it* was on top, and is still correct because nothing below it ever changed. The previous minimum is restored for free — because it was never overwritten in the first place.

## Two things in this implementation worth naming

**The `defer` in `Push`** is doing something subtle. `Push` calls `GetMin`, which reads `dataArr[this.len-1]`, and the deferred increment means `this.len` still describes the stack *before* this push at that moment. Correct — and a reader has to reconstruct the ordering to see why.

**`this.len` always equals `len(this.dataArr)`.** The struct keeps both. Nothing is wrong today, but it is two things that must agree rather than one thing that cannot disagree.

## The other standard solution

Two stacks: one for values, one holding only minima, pushed when a value is less than *or equal to* the current minimum.

Less space when new minima are rare. But push on strictly-less-than and duplicate minima get popped too early — a real edge case the pair-per-entry version simply does not have.

## Complexity

- **Time: O(1)** for all four operations, no amortisation.
- **Space: O(n)**, two integers per element.

## Builds on

- [Day 58: Valid Parentheses](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/valid_parentheses/) — the same structure, now asked to answer a question about its contents rather than just hold them

Full code and the step-by-step walkthrough:
[min_stack](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/min_stack/SOLUTION.md)

#DSA #LeetCode #Golang #Stack #SystemDesign #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
