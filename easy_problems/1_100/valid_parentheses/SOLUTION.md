# Valid Parentheses — solution walkthrough

From `valid_parentheses.go`:

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

## What has to be remembered

A counter is not enough, and it is worth seeing exactly where it breaks.

Counting openers against closers accepts `)(`. Tracking a depth that rises and falls accepts `([)]`, because depth records *how many* brackets are open and never *which*.

What the problem needs is the sequence of brackets opened and not yet closed. And when a closer arrives, only the most recently opened one can match it: in `([`, a `)` is wrong because `[` is still open and has to close first.

Most recent in, first out. That is a stack.

## The map is keyed by the closer

```go
pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
```

The instinct is to write `'(' : ')'`. Keying by the closing bracket is what makes the loop short:

```go
open, isCloser := pairs[A[i]]
```

One lookup answers both questions. `isCloser` says whether this character closes anything, and `open` says what must be on top if it does. Keyed the other way, you need a separate membership test before you can branch at all.

## The three cases

![Step 1](images/walkthrough-1.png)

An opener goes straight on the stack. No decision to make.

![Step 2](images/walkthrough-2.png)

The stack now records both unclosed brackets, in order.

```go
if len(stack) == 0 || stack[len(stack)-1] != open {
	return false
}
```

![Step 3](images/walkthrough-3.png)

A closer has two ways to be wrong, and both are in that one condition.

`len(stack) == 0` is a closer with nothing open at all, as in `)`. The order matters: Go short-circuits `||`, so the emptiness test has to come first or the index below it reads past the end of an empty slice.

`stack[len(stack)-1] != open` is a closer meeting the wrong bracket, which is the `([)]` case. `)` needs `(` and finds `[`.

Neither can be rescued by anything later in the string, so the function returns immediately rather than reading on.

## The final check

```go
return len(stack) == 0
```

Reaching the end proves every closer matched. It does not prove every opener was closed: `(((` never enters the closer branch at all.

So the stack has to be empty. Anything left on it was opened and never closed.

## The capacity hint

```go
stack := make([]byte, 0, len(A))
```

At most one opener per character, so the stack can never exceed the string length. Allocating that up front means `append` never has to grow and copy.

It is not needed for correctness, and on a string that is mostly closers it over-allocates. It is a reasonable trade when the bound is this obvious.

## Complexity

- **Time: O(n).** Each character is looked at once and does constant work.
- **Space: O(n).** The stack, whose worst case is a string of all openers.

## Test

`main_test.go` runs all five worked examples. `(]` and `([)]` are the interesting pair: the first is caught by the wrong-bracket test, and the second is the one a depth counter would wrongly accept.
