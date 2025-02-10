package alloonedatastructure

type Node struct {
	Val   string
	Count int
	Prev  *Node
	Next  *Node
}

type AllOne struct {
	CountMap map[string]*Node
	Head     *Node
	Tail     *Node
}

func newNode(val string) *Node {
	return &Node{Val: val, Count: 1}
}

func Constructor() AllOne {
	cm := make(map[string]*Node)
	head := newNode("")
	tail := newNode("")

	head.Next = tail
	tail.Prev = head

	return AllOne{
		CountMap: cm,
		Head:     head,
		Tail:     tail,
	}
}

func (this *AllOne) insert(newNode *Node) {
	last := this.Tail.Prev

	last.Next = newNode
	newNode.Prev = last

	newNode.Next = this.Tail
	this.Tail.Prev = newNode
}

func (this *AllOne) swap(lhs, rhs *Node) {
	prev := lhs.Prev
	next := rhs.Next

	prev.Next = rhs
	rhs.Prev = prev

	rhs.Next = lhs
	lhs.Prev = rhs

	lhs.Next = next
	next.Prev = lhs
}

func (this *AllOne) sortBackward(node *Node) {
	for node.Prev != this.Head && node.Count > node.Prev.Count {
		this.swap(node.Prev, node)
	}
}

func (this *AllOne) Inc(key string) {
	if node, ok := this.CountMap[key]; ok {
		node.Count++
		this.sortBackward(node)
	} else {
		newNode := newNode(key)
		this.CountMap[key] = newNode
		this.insert(newNode)
	}
	// this.printAll("Inc")
}

func (this *AllOne) delete(node *Node) {
	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev

	node.Next = nil
	node.Prev = nil
}

func (this *AllOne) sortForward(node *Node) {
	for node.Next != this.Tail && node.Count < node.Next.Count {
		this.swap(node, node.Next)
	}
}

func (this *AllOne) Dec(key string) {
	node := this.CountMap[key]
	node.Count--
	if node.Count == 0 {
		delete(this.CountMap, key)
		this.delete(node)
	} else {
		this.sortForward(node)
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.Head.Next == this.Tail {
		return ""
	}
	return this.Head.Next.Val
}

func (this *AllOne) GetMinKey() string {
	if this.Head.Next == this.Tail {
		return ""
	}
	return this.Tail.Prev.Val
}
