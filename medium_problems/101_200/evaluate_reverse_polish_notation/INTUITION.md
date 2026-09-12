# Evaluate Reverse Polish Notation — intuition

## Builds on

- [Day 58: Valid Parentheses](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/valid_parentheses/) — the same "only the most recent entries can be resolved right now" reasoning, with operands in place of brackets

## The problem in one line

Evaluate an expression written in postfix, where the operator follows its operands.

## Why postfix needs no parentheses

`(2 + 1) * 3` needs brackets because `+` sits between its operands and something has to say which operation binds first.

Written postfix it is `2 1 + 3 *`, and there is nothing to disambiguate. Each operator applies to the two values immediately before it, so the order of evaluation is already encoded in the order of the tokens. No precedence rules, no parentheses, no ambiguity.

That is the property the whole problem rests on: the expression is already in evaluation order.

## What that means for the algorithm

Read left to right. Numbers have nothing to do yet, so they wait. An operator consumes the two most recent waiting values and leaves its result in their place.

Most recent in, first out. A stack, for the same reason as day 58: when an operator arrives, only the two most recent operands can possibly be its arguments.

## The order of the two pops

This is where the bugs live.

Popping gives you the operands in reverse, so the first pop is the *right* operand and the second is the *left*.

For `+` and `*` that does not matter, because they commute. For `-` and `/` it decides between the right answer and a plausible-looking wrong one:

```go
a, b := poplast(), poplast()
push(b - a)
```

`a` is the right operand, `b` the left, and `b - a` is the subtraction in the order the expression meant.

Write `a - b` and `["5","3","-"]` returns -2 instead of 2. Both look like answers.

## The invariant at the end

A well-formed postfix expression leaves exactly one value on the stack. That is the result.

More than one means there were too many operands; an empty stack mid-operation means too few. The problem guarantees a valid expression, so neither is checked.

## Complexity

- **Time: O(n).** One pass, constant work per token.
- **Space: O(n)** for the stack.
