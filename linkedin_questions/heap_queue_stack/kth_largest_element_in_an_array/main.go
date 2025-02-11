package kthlargestelementinanarray

import "container/heap"

type MaxHeap []int

func (h *MaxHeap) Len() int { return len(*h) }
func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}
func (h *MaxHeap) Push(val any) {
	*h = append(*h, val.(int))
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	mHeap := &MaxHeap{}
	heap.Init(mHeap)
	for _, v := range nums {
		heap.Push(mHeap, v)
	}
	var x int
	for i := 0; i < k; i++ {
		x = (heap.Pop(mHeap)).(int)
	}
	return x
}
