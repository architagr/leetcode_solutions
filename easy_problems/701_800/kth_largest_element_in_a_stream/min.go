package kthlargestelementinastream

import (
	"container/heap"
)

type minHeap []int

func (mHeap *minHeap) Len() int {
	return len(*mHeap)
}

func (mHeap *minHeap) Less(i, j int) bool {
	return (*mHeap)[i] < (*mHeap)[j]
}

func (mHeap *minHeap) Swap(i, j int) {
	(*mHeap)[i], (*mHeap)[j] = (*mHeap)[j], (*mHeap)[i]
}

func (h *minHeap) Push(val any) {
	*h = append(*h, val.(int))
}
func (h *minHeap) Pop() any {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
	h *minHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	obj := &minHeap{}
	heap.Init(obj)

	limit := k
	if len(nums) < k {
		limit = len(nums)
	}

	for i := 0; i < limit; i++ {
		heap.Push(obj, nums[i])
	}
	for i := limit; i < len(nums); i++ {
		if (*obj)[0] < nums[i] {
			_ = heap.Pop(obj)
			heap.Push(obj, nums[i])
		}
	}
	return KthLargest{
		k: k,
		h: obj,
	}
}

func (this *KthLargest) Add(val int) int {

	if this.h.Len() < this.k {
		heap.Push(this.h, val)
		return (*this.h)[0]
	}
	if (*this.h)[0] < val {
		_ = heap.Pop(this.h)
		heap.Push(this.h, val)
	}
	return (*this.h)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
