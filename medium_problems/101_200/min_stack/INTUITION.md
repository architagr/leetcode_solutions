# Min Stack — intuition

## Builds on

- [Day 58: Valid Parentheses](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/valid_parentheses/) — the same structure, now asked to answer a question about its contents rather than just hold them

## The problem in one line

A stack where `push`, `pop`, `top` and `getMin` are all O(1).

## Why the obvious fixes do not work

**Scan the stack in `getMin`.** O(n) per call, and the requirement says constant.

**Keep one `min` field, updated on push.** Push is fine. Pop is the problem: pop the current minimum and the field is now wrong, and there is nothing left to recompute it from without scanning.

That second failure is the interesting one, and it points at what the structure actually needs. The minimum is not a property of the stack. It is a property of *the stack at a particular height*, and popping moves you back to a height whose minimum you have already thrown away.

## Store it per entry

So do not throw it away. Every entry records the value pushed **and** the minimum of everything at or below it:

```go
type dataNode struct {
	val, min int
}
```

Push computes the new entry's `min` as the smaller of the incoming value and the current top's `min`. `getMin` reads the top's `min`. Pop needs to do nothing at all: removing the top exposes an entry whose `min` was correct when it was written and is still correct now.

That last point is the whole trick. Popping restores the previous minimum for free, because the previous minimum was never overwritten — it was stored alongside the entry that was on top at the time.

## The cost

O(n) extra space, one integer per entry.

The usual alternative is a second stack holding only the minima, pushed to when a new minimum arrives. It uses less space when minima are rare and needs care about ties. Storing a pair per entry is simpler and has no edge cases, which is why it is the version worth knowing first.

## Complexity

- **Time: O(1)** for every operation, with no amortisation.
- **Space: O(n)**, two integers per element.
