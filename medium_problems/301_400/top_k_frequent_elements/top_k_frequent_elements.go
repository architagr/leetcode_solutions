package top_k_frequent_elements

import "container/heap"

// entry is one distinct value and how often it occurred.
type entry struct {
	ele, count int
}

// minHeap orders by count ascending, so the LEAST frequent entry currently
// held is at index 0. That is the one to evict when the heap is over size k,
// which is why a top-k problem wants a min-heap rather than a max-heap.
type minHeap []entry

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(entry)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	// The heap never exceeds k entries: push, and if that tips it over,
	// evict the smallest. Whatever survives to the end is the top k.
	h := &minHeap{}
	for ele, count := range m {
		heap.Push(h, entry{ele: ele, count: count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Draining gives ascending frequency. The problem does not specify an
	// order, so this is left as it comes out rather than reversed.
	ans := make([]int, 0, h.Len())
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(entry).ele)
	}
	return ans
}
