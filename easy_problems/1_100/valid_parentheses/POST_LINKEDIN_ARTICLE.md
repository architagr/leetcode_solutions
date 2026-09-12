---
meta_title: "Why a counter cannot check brackets"
meta_description: "Counting accepts )(. A depth counter accepts ([)]. Both fail for the same reason, and that reason is exactly what a stack fixes."
---

![Day 58](HERO.png)

## 365 Days of LeetCode Challenge — Day 58/365

**[20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)** (Easy)

Given a string of brackets, is every one closed by the right kind, in the right order?

New topic today. Graphs ended on Course Schedule; stacks start here.

## Start with what does not work

**Count openers against closers.** Accepts `)(` — balanced by count, wrong by order.

**Track a depth that rises on an opener and falls on a closer.** Accepts `([)]` — depth never goes negative and ends at zero.

Both fail for the same reason, and it is worth naming precisely: a single number records **how many** brackets are open and never **which**. This problem is entirely about which.

## What actually has to be remembered

At any point, what matters is the list of brackets opened and not yet closed, in order.

And when a closer arrives, only one candidate can match it: the most recent opener. In `([`, a `)` is wrong not because `(` is unmatched, but because `[` is still open and must close first.

Most recent in, first out. That is a stack, and seeing that this problem *is* a stack is the whole exercise.

## The map is keyed by the closing bracket

```go
pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
```

The instinct is `'(' : ')'`. Keying it the other way is what makes the loop short:

```go
open, isCloser := pairs[A[i]]
```

One lookup answers both questions at once — whether this character closes anything, and what must be on top if it does. Keyed by the opener, you need a separate membership test before you can branch.

## The whole function

```go
func IsValid(A string) bool {
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0, len(A))

	for i := 0; i < len(A); i++ {
		open, isCloser := pairs[A[i]]
		if !isCloser {
			stack = append(stack, A[i])
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != open {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
```

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

## Two failures, one condition

```go
if len(stack) == 0 || stack[len(stack)-1] != open {
```

`len(stack) == 0` is a closer with nothing open, as in `)`.

`stack[len(stack)-1] != open` is a closer meeting the wrong bracket — the `([)]` case.

The order is load-bearing. Go short-circuits `||`, so the emptiness test must come first, or the index below it reads past the end of an empty slice.

## The last line is not a formality

```go
return len(stack) == 0
```

Reaching the end proves every closer matched. It does **not** prove every opener was closed — `(((` never enters the closer branch at all.

## Complexity

- **Time: O(n)**. One pass, constant work per character.
- **Space: O(n)**. The stack, worst case a string of all openers.

Full code and the step-by-step walkthrough:
[valid_parentheses](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/valid_parentheses/SOLUTION.md)

#DSA #LeetCode #Golang #Stack #CodingInterview #SoftwareEngineering #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
