**365 Days of LeetCode Challenge — Day 62/365**
**Evaluate Reverse Polish Notation** (Medium)
🔗 https://leetcode.com/problems/evaluate-reverse-polish-notation/

Why postfix needs no parentheses: `(2 + 1) * 3` needs brackets because `+` sits BETWEEN its operands. Written postfix it is `2 1 + 3 *`, and each operator applies to the two values immediately before it. Evaluation order is already in the token order.

So numbers wait on a stack, and an operator consumes the two most recent and leaves its result in their place.

```go
case "-":
	a, b := poplast(), poplast()
	push(b - a)
```

That is the trap. Popping returns the operands in REVERSE - first pop is the right operand, second is the left.

For `+` and `*` it makes no difference because they commute, which is exactly why the bug survives casual testing. For `-` and `/` it picks between right and plausibly wrong:

`["5","3","-"]` means 5 - 3 = 2. Backwards gives -2.
`["6","2","/"]` means 6 / 2 = 3. Backwards gives 0.

Go's integer `/` truncates toward zero, which is what the problem wants, so nothing extra is needed there.

Two things this version could drop: the `pointer` field is always len(stack)-1, and the first token is pushed by hand before the loop when the default branch already does it.

O(n) time and space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/evaluate_reverse_polish_notation/SOLUTION.md
