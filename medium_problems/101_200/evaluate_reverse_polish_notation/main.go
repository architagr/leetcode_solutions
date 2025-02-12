package evaluatereversepolishnotation

import "strconv"

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
