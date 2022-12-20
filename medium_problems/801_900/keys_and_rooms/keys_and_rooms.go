package keys_and_rooms

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	queue := NewQueue()
	queue.Enqueue(0)
	visited[0] = true
	for !queue.IsEmpty() {
		x := queue.Dequeue().Data
		val := rooms[x]
		for _, val := range val {
			if !visited[val] {
				queue.Enqueue(val)
				visited[val] = true
			}
		}
	}
	for _, val := range visited {
		if !val {
			return false
		}
	}

	return true
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
