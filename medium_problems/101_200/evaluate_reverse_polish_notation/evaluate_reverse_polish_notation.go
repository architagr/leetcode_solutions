package evaluate_reverse_polish_notation

import "strconv"

type Node struct {
	Data int
	Prev *Node
}
type Stack struct {
	count int
	Head  *Node
}

func NewNode(value int) *Node {
	newNode := new(Node)
	newNode.Data = value
	return newNode
}

func (stack *Stack) Push(value int) {
	newNode := NewNode(value)
	stack.count++
	if stack.Head == nil {
		stack.Head = newNode
		return
	}
	newNode.Prev = stack.Head
	stack.Head = newNode

}

func (stack *Stack) Pop() *Node {
	if stack.Head == nil {
		return nil
	}
	node := stack.Head
	stack.Head = stack.Head.Prev
	stack.count--
	return node
}

func (stack *Stack) Top() (value int) {
	if stack.Head == nil {
		return
	}
	value = stack.Head.Data
	return
}

func (stack *Stack) IsEmpty() bool {
	return stack.count == 0
}

func (stack *Stack) Count() int {
	return stack.count
}
func evalRPN(tokens []string) int {
	stack := new(Stack)
	for _, val := range tokens {
		switch val {
		case "+":
			a := stack.Pop().Data
			b := stack.Pop().Data
			stack.Push(b + a)
		case "-":
			a := stack.Pop().Data
			b := stack.Pop().Data
			stack.Push(b - a)
		case "*":
			a := stack.Pop().Data
			b := stack.Pop().Data
			stack.Push(b * a)
		case "/":
			a := stack.Pop().Data
			b := stack.Pop().Data
			stack.Push(b / a)
		default:
			a, _ := strconv.Atoi(val)
			stack.Push(a)
		}
	}
	return stack.Pop().Data
}
