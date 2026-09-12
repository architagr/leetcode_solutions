---
meta_title: "The two pops come back in the wrong order"
meta_description: "Postfix needs no brackets because evaluation order is already in the tokens. The only real trap is that popping reverses the operands."
---

![Day 62](HERO.png)

## 365 Days of LeetCode Challenge — Day 62/365

**[150. Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/)** (Medium)

Evaluate an expression written in postfix, where each operator follows its operands.

## Why postfix has no parentheses

`(2 + 1) * 3` needs brackets because `+` sits *between* its operands, and without them precedence rules must decide what binds first.

Postfix is `2 1 + 3 *`. Each operator applies to the two values immediately before it. Evaluation order is already encoded in token order — there is nothing to disambiguate, so there is nothing for parentheses to do.

That property is the whole basis of the algorithm: the expression arrives in the order it should be evaluated.

## Which makes the algorithm short

Read left to right. Numbers cannot be used yet, so they wait. An operator consumes the two most recent waiting values and leaves its result in their place.

Only the two most recent can be its arguments — the same reasoning as day 58's brackets, with operands instead of openers.

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

Two values in, one out.

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

## The actual trap

```go
a, b := poplast(), poplast()
push(b - a)
```

Popping returns the operands **in reverse**. The first pop is the *right* operand; the second is the *left*.

For `+` and `*` it makes no difference — they commute, which is exactly why the bug survives casual testing. For `-` and `/` it picks between the right answer and a plausible wrong one:

- `["5","3","-"]` means `5 - 3 = 2`. With `a - b` you get `-2`.
- `["6","2","/"]` means `6 / 2 = 3`. With `a / b` you get `0`.

Go's integer `/` truncates toward zero, which is what the problem specifies, so nothing extra is needed there.

## Two things this implementation could drop

**The `pointer` field.** It is always `len(stack) - 1`. Two values that must agree rather than one that cannot disagree — the same observation as day 60's `this.len`.

**The hand-written first token.** The first token is pushed before the loop, which then runs over `tokens[1:]`. It works, because valid postfix always starts with an operand. It is also unnecessary — the `default` branch already pushes numbers, so starting at `tokens[0]` would do the same with one special case fewer.

That is the shape day 22 fixed with a dummy head. Here the special case can just be deleted.

## The end

```go
return stack[pointer]
```

A well-formed expression leaves exactly one value. More means too many operands; an empty stack mid-operation means too few. The problem guarantees validity so neither is checked — in code parsing untrusted input, both would have to be.

## Complexity

- **Time: O(n)**. One pass, constant work per token.
- **Space: O(n)** for the stack.

## Builds on

- [Day 58: Valid Parentheses](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/valid_parentheses/) — the same "only the most recent entries can be resolved right now" reasoning, with operands in place of brackets

Full code and the step-by-step walkthrough:
[evaluate_reverse_polish_notation](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/evaluate_reverse_polish_notation/SOLUTION.md)

#DSA #LeetCode #Golang #Stack #Parsing #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
