package daily_temperature

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
