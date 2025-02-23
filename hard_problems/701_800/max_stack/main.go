package maxstack

import "container/heap"

type Node struct {
	val, id int
}

type MaxHeap []Node

func (h *MaxHeap) Len() int {
	return len(*h)
}
func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MaxHeap) Less(i, j int) bool {
	if (*h)[i].val == (*h)[j].val {
		return (*h)[i].id > (*h)[j].id
	}
	return (*h)[i].val > (*h)[j].val
}

func (h *MaxHeap) Push(val any) {
	*h = append(*h, val.(Node))
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MaxStack struct {
	stack   []Node
	maxHeap *MaxHeap
	removed map[int]bool
	cnt     int
}

func Constructor() MaxStack {
	mHeap := &MaxHeap{}
	heap.Init(mHeap)
	return MaxStack{
		stack:   make([]Node, 0),
		maxHeap: mHeap,
		removed: make(map[int]bool),
		cnt:     1,
	}
}

func (this *MaxStack) popInlineStack() Node {
	n := len(this.stack)
	node := this.stack[n-1]
	this.stack = this.stack[:n-1]
	return node
}
func (this *MaxStack) peekInlineStack() Node {
	n := len(this.stack)
	return this.stack[n-1]
}

func (this *MaxStack) pushInlineStack(n Node) {
	this.stack = append(this.stack, n)
}

func (this *MaxStack) popInlineHeap() Node {
	x := heap.Pop(this.maxHeap)
	return x.(Node)
}
func (this *MaxStack) peekInlineHeap() Node {
	x := this.popInlineHeap()
	this.pushInlineHeap(x)
	return x
}

func (this *MaxStack) pushInlineHeap(n Node) {
	heap.Push(this.maxHeap, n)
}

func (this *MaxStack) Push(x int) {
	node := Node{
		val: x,
		id:  this.cnt,
	}
	this.pushInlineStack(node)
	this.pushInlineHeap(node)
	this.cnt++
}

func (this *MaxStack) popAllRemovedFromStack() {
	_, exist := this.removed[this.peekInlineStack().id]
	for exist {
		this.popInlineStack()
		_, exist = this.removed[this.peekInlineStack().id]
	}
}
func (this *MaxStack) Pop() int {
	this.popAllRemovedFromStack()
	top := this.popInlineStack()
	this.removed[top.id] = true
	return top.val
}

func (this *MaxStack) Top() int {
	this.popAllRemovedFromStack()
	return this.peekInlineStack().val
}

func (this *MaxStack) popAllRemovedFromHeap() {
	_, exist := this.removed[this.peekInlineHeap().id]
	for exist {
		this.popInlineHeap()
		_, exist = this.removed[this.peekInlineHeap().id]
	}
}

func (this *MaxStack) PeekMax() int {
	this.popAllRemovedFromHeap()
	return this.peekInlineHeap().val
}

func (this *MaxStack) PopMax() int {
	this.popAllRemovedFromHeap()
	x := this.popInlineHeap()
	this.removed[x.id] = true
	return x.val
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */
