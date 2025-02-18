package laststoneweight

import "container/heap"

type MaxHeap []int

func (h *MaxHeap) Len() int { return len(*h) }
func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}
func (h *MaxHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for _, v := range stones {
		heap.Push(maxHeap, v)
	}
	for maxHeap.Len() > 1 {
		x := heap.Pop(maxHeap).(int)
		y := heap.Pop(maxHeap).(int)
		z := x - y
		if z != 0 {
			heap.Push(maxHeap, z)
		}
	}
	if maxHeap.Len() == 0 {
		return 0
	}
	return heap.Pop(maxHeap).(int)
}
