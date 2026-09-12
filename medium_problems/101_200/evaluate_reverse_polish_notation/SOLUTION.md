# Evaluate Reverse Polish Notation — solution walkthrough

From `main.go`:

```go
func evalRPN(tokens []string) int {
	pointer := 0
	stack := make([]int, 0, len(tokens))
	x, _ := strconv.Atoi(tokens[0])
	stack = append(stack, x)
	poplast := func() int {
		x := stack[pointer]
		stack = stack[:pointer]
		pointer--
		return x
	}
	push := func(x int) {
		stack = append(stack, x)
		pointer++
	}
	for _, val := range tokens[1:] {
		switch val {
		case "+":
			a, b := poplast(), poplast()
			push(a + b)
		case "-":
			a, b := poplast(), poplast()
			push(b - a)
		case "*":
			a, b := poplast(), poplast()
			push(a * b)
		case "/":
			a, b := poplast(), poplast()
			push(b / a)
		default:
			x, _ := strconv.Atoi(val)
			push(x)
		}
	}
	return stack[pointer]
}
```

## Why postfix needs no brackets

`(2 + 1) * 3` needs parentheses because `+` sits between its operands, and without them precedence rules have to decide what binds first.

Postfix is `2 1 + 3 *`. Every operator applies to the two values immediately preceding it, so evaluation order is already encoded in token order. Nothing to disambiguate.

That property is what the algorithm exploits: the expression arrives in the order it should be evaluated.

## The loop

Numbers cannot be used yet, so they wait. An operator consumes the two most recent waiting values and leaves its result where they were.

Only the two most recent can be an operator's arguments — the same reasoning as day 58's brackets, with operands instead.

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

Two operands waiting.

![Step 3](images/walkthrough-3.png)

`+` pops both, adds, pushes the result. The stack shrinks by one: two values in, one out.

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

## The order of the two pops

This is where this problem's bugs live.

```go
a, b := poplast(), poplast()
push(b - a)
```

Popping returns operands in reverse order. The first pop is the **right** operand, the second is the **left**.

For `+` and `*` it makes no difference, because they commute. For `-` and `/` it decides between a right answer and a wrong one that looks perfectly reasonable:

- `["5","3","-"]` means `5 - 3 = 2`. With `a - b` it returns `-2`.
- `["6","2","/"]` means `6 / 2 = 3`. With `a / b` it returns `0`, because Go truncates integer division.

Note also that Go's `/` truncates toward zero for integers, which is what the problem specifies. No adjustment is needed.

## The `pointer` field

The function tracks `pointer` alongside the slice:

```go
push := func(x int) {
	stack = append(stack, x)
	pointer++
}
```

`pointer` is always `len(stack) - 1`. Two values that must agree rather than one that cannot disagree — the same observation as day 60's `this.len`.

Using `len(stack)-1` directly would remove the possibility of them drifting apart. Nothing is wrong today; the closures keep them in step.

## The first token is handled outside the loop

```go
x, _ := strconv.Atoi(tokens[0])
stack = append(stack, x)
...
for _, val := range tokens[1:] {
```

The first token is pushed before the loop starts, and the loop runs over `tokens[1:]`.

It works, because a valid postfix expression always begins with an operand. It is also unnecessary: the `default` branch already pushes numbers, so starting the loop at `tokens[0]` would do exactly the same thing with one special case fewer.

This is the same shape as the merge in day 22, where a dummy head removed a hand-written first step. Here the first step could simply be deleted.

## The end

```go
return stack[pointer]
```

A well-formed postfix expression leaves exactly one value. More would mean too many operands; an empty stack during an operation would mean too few.

The problem guarantees validity, so neither is checked. In code parsing untrusted input, both would need to be.

## Complexity

- **Time: O(n).** One pass, constant work per token, with `Atoi` proportional to the digits of a number rather than the expression.
- **Space: O(n)** for the stack.

## Test

`main_test.go` covers the worked examples, including the nested one whose intermediate results feed later operators, which is the case that catches a wrong pop order.
