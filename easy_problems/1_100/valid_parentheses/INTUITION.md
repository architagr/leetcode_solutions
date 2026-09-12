# Valid Parentheses — intuition

## The problem in one line

Given a string of brackets, is every one closed by the right kind, in the right order?

## Why the obvious counting approach fails

Count openers and closers and compare? That accepts `)(`, which is balanced by count and wrong by order.

Track a depth counter that goes up on an opener and down on a closer? That accepts `([)]`, because depth says nothing about *which* bracket is open at each level.

Both fail for the same reason: a single number cannot remember what kind of bracket you are inside, and this problem is entirely about that.

## What you actually need to remember

At any point in the string, the thing that matters is the list of brackets opened and not yet closed, in the order they were opened.

And when a closer arrives, only one of them can possibly match: the most recent one. `([` followed by `)` is wrong not because `(` is unmatched but because `[` is still open and has to close first.

Most recent in, first one out. That is a stack, and noticing this problem *is* a stack is the whole exercise.

## The loop

Push openers. On a closer, look at the top:

- Nothing there: a closer with nothing open, so it is invalid.
- Wrong kind: this closer cannot match what is open, so it is invalid.
- Right kind: they cancel, so pop.

At the end, anything still on the stack was opened and never closed.

## Keying the map by the closing bracket

```go
pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
```

The natural instinct is to map opener to closer. Keying it the other way is better here, because the lookup then answers two questions at once: whether this character is a closer at all, and what it needs on top of the stack.

```go
open, isCloser := pairs[A[i]]
```

One lookup, one branch. Key it by the opener and you need a separate test for whether the character is an opener before you can decide anything.

## Failing immediately

Nothing can rescue a mismatch. Once a closer meets the wrong top, no later character makes the string valid, so the function returns there rather than reading the rest.

That matters less for correctness than for what it says: the invariant is checked continuously, not reconstructed at the end.

## Complexity

- **Time: O(n).** One pass, constant work per character.
- **Space: O(n).** The stack, worst case a string of all openers.
