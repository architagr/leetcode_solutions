package find_if_path_exists_in_graph

import "fmt"

func validPath(n int, edges [][]int, source int, destination int) bool {
	visited := make([]bool, n)
	queue := NewQueue()
	at := make(map[int]map[int]bool)
	for i := 0; i < len(edges); i++ {
		val, ok := at[edges[i][0]]
		if !ok {
			val = make(map[int]bool)
		}
		val[edges[i][1]] = true
		at[edges[i][0]] = val

		val, ok = at[edges[i][1]]
		if !ok {
			val = make(map[int]bool)
		}
		val[edges[i][0]] = true
		at[edges[i][1]] = val

	}
	queue.Enqueue(source)
	visited[source] = true
	for !queue.IsEmpty() {
		x := queue.Dequeue().Data
		val := at[x]
		for key := range val {
			if !visited[key] {
				queue.Enqueue(key)
				visited[key] = true
			}
		}
	}
	return visited[destination]
}

type Node struct {
	Data int
	next *Node
}

func NewNode(value int) *Node {
	newNode := new(Node)
	newNode.Data = value
	return newNode
}

type Queue struct {
	count      int
	head, tail *Node
}

type IQueue interface {
	Enqueue(data int)
	Dequeue() *Node
	IsEmpty() bool
	Peep() (result int, err error)
}

func NewQueue() IQueue {
	return new(Queue)
}
func (queue *Queue) Enqueue(data int) {
	node := NewNode(data)
	if queue.IsEmpty() {
		queue.head = node
		queue.tail = node
	} else {
		queue.tail.next = node
		queue.tail = node
	}
	queue.count++
}

func (queue *Queue) Dequeue() *Node {
	if queue.IsEmpty() {
		return nil
	}
	node := queue.head
	queue.head = queue.head.next
	queue.count--
	if queue.IsEmpty() {
		queue.head = nil
		queue.tail = nil
	}
	return node
}
func (queue *Queue) Peep() (result int, err error) {
	if queue.IsEmpty() {
		err = fmt.Errorf("no data")
		return
	}
	result = queue.head.Data
	return
}

func (queue *Queue) IsEmpty() bool {
	return queue.count == 0
}
