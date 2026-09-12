# Daily Temperatures — solution walkthrough

From `main.go`:

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

## The two differences from day 59

Day 59 asked *which* value comes next and is larger. This asks *how far away* it is, and two things follow.

**The stack holds indices, not values.** A distance needs positions. `i` goes on the stack, and the answer is `stack.Top() - i`.

**Values repeat.** Day 59 stored answers in a map keyed by value, which was safe only because uniqueness was guaranteed. Temperatures repeat constantly, so that approach would collide. Indices are unique by construction — the problem disappears rather than needing a workaround.

## Scanning backwards

Day 59 went left to right, asking "which earlier values does this one answer?" and writing results into a map as it resolved them.

This goes right to left, asking "what is the first thing to my right that beats me?" The answer for index `i` is known while standing on `i`, so it is written straight into `results[i]` and the loop moves on. No map, no deferred bookkeeping.

![Step 1](images/walkthrough-1.png)

The rightmost day has nothing after it, so its answer is 0. It still goes on the stack — it is a candidate for the days to its left.

![Step 2](images/walkthrough-2.png)

Index 4 holds 76, which beats 71, so the wait is `4 - 3 = 1`.

## The invariant

The stack holds indices of days, all to the right of `i`, that are still candidates for being somebody's answer.

A day stops being a candidate as soon as a day warmer-or-equal appears to its left, because everyone further left would hit that nearer day first. That is what the inner loop removes:

```go
for !stack.IsEmpty() && temperatures[stack.Top()] <= temperatures[i] {
	stack.Pop()
}
```

![Step 3](images/walkthrough-3.png)

Index 2 holds 75. The 71 at index 3 is cooler, so it can never be the answer for index 2 or anything left of it — popped. The next candidate is the 76 at index 4, giving `4 - 2 = 2`.

After the pops, the stack reads increasing in temperature from top to bottom, and its top is the nearest warmer day.

![Step 4](images/walkthrough-4.png)

75 at index 2 already beats 74, so nothing is popped and the answer is `2 - 1 = 1`.

![Step 5](images/walkthrough-5.png)

## Why `<=` rather than `<`

```go
temperatures[stack.Top()] <= temperatures[i]
```

Equal temperatures are not *warmer*. A day holding the same temperature can never be anyone's answer, so it must be popped.

Write `<` and equal days stay on the stack. The first one found then produces an answer that claims a warmer day exists when it is merely an equal one. On `[70, 70]` the correct answer is `[0, 0]`, and `<` gives `[1, 0]`.

This is exactly the duplicate-value hazard day 59 avoided by having unique values. Here duplicates are guaranteed to occur, and this one character is where they are handled.

## The empty case

```go
if stack.IsEmpty() {
	results[i] = 0
}
```

An empty stack after the pops means nothing to the right is warmer. `results` was allocated with `make`, so its entries are already zero and the assignment is redundant — harmless, and explicit about intent rather than relying on the zero value.

## Why linear

Each index is pushed exactly once and popped at most once. The inner loop can run several times in one iteration, but every one of those iterations permanently removes an index.

Same argument as day 59, and worth repeating because a nested loop looks quadratic until you count what it consumes rather than what it visits.

## The hand-rolled stack

The file defines a linked-list `Stack` with `Push`, `Pop`, `Top`, `IsEmpty` and `Count`. Go has no stack type, and the usual idiom is a slice with `s[len(s)-1]` and `s = s[:len(s)-1]`, which days 58 and 62 use.

Both are fine. The named methods make the main loop read closer to the explanation.

## Complexity

- **Time: O(n).** Each index pushed once, popped at most once.
- **Space: O(n)** for the stack, worst case a strictly decreasing sequence where nothing is ever popped until the end.

## Test

`main_test.go` covers the worked examples, including `[30,40,50,60]` where every day is answered by the next, and the case ending in the maximum where the tail answers are all 0.
