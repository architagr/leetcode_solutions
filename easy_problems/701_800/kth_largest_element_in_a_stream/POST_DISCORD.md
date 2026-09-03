**365 Days of LeetCode Challenge — Day 20/365**

**Kth Largest Element in a Stream** (LeetCode 703, Easy)
<https://leetcode.com/problems/kth-largest-element-in-a-stream/>

We only ever need to know the kth largest value seen so far, not the full sorted order
of everything. So we only need to track the top k scores. Among those top k, the kth
largest is exactly the smallest one, which turns the problem into: maintain a bounded
set of k values, and read or replace its minimum fast. A **min-heap** capped at size
`k` does exactly that in `O(log k)` per operation. Once that clicks, the code for both
the constructor and `Add` basically writes itself, they're the same eviction check
applied in two places.

**The code (`min.go`):**

```go
package kthlargestelementinastream

import (
	"container/heap"
)

// minHeap is a min-heap of ints: Less makes the smallest element sort to the
// root ((*mHeap)[0]), which is what lets us read the kth largest score in O(1).
type minHeap []int

func (mHeap *minHeap) Len() int {
	return len(*mHeap)
}

// Less uses < (not >) so the smallest value bubbles to the root — this is what
// makes it a min-heap instead of the more common max-heap.
func (mHeap *minHeap) Less(i, j int) bool {
	return (*mHeap)[i] < (*mHeap)[j]
}

func (mHeap *minHeap) Swap(i, j int) {
	(*mHeap)[i], (*mHeap)[j] = (*mHeap)[j], (*mHeap)[i]
}

// Push appends the new value to the end; heap.Push then sifts it up to
// restore the heap invariant.
func (h *minHeap) Push(val any) {
	*h = append(*h, val.(int))
}

// Pop removes and returns the last element of the slice. heap.Pop swaps the
// element to be removed to the end (after sifting down) before calling this,
// so it always operates on the correct element.
func (h *minHeap) Pop() any {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// KthLargest keeps a min-heap of at most k elements — the k largest scores
// seen so far. The heap's root is then always the smallest of that top-k
// group, i.e. exactly the kth largest score overall.
type KthLargest struct {
	h *minHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	obj := &minHeap{}
	heap.Init(obj)

	// len(nums) can be less than k (constraints allow it), so cap the
	// unconditional fill at whichever is smaller.
	limit := k
	if len(nums) < k {
		limit = len(nums)
	}

	// First k (or fewer) elements are automatically part of the top-k
	// since the heap isn't full yet — push them with no comparison.
	for i := 0; i < limit; i++ {
		heap.Push(obj, nums[i])
	}
	// For every remaining element, run the same "does it beat the current
	// weakest member of the top-k?" check that Add uses below.
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

	// Heap not yet at capacity k: val is automatically part of the
	// current top-k, so just push it — no eviction needed.
	if this.h.Len() < this.k {
		heap.Push(this.h, val)
		return (*this.h)[0]
	}
	// Heap already holds k elements: val only earns a spot if it beats
	// the current weakest of the top-k (the root). Otherwise leave the
	// heap untouched.
	if (*this.h)[0] < val {
		_ = heap.Pop(this.h)
		heap.Push(this.h, val)
	}
	// Root of the min-heap is, by construction, the kth largest value
	// seen so far.
	return (*this.h)[0]
}
```

**Walkthrough**, tracing the example `KthLargest(3, [4, 5, 8, 2])` then
`add(3), add(5), add(10), add(9), add(4)`:

`Constructor` pushes the first `k=3` elements (`4, 5, 8`) with no checks, filling the
heap to `{4, 5, 8}`, root `4`. The remaining element `2` gets tested against the root:
`4 < 2`? No, rejected.

![Step 1: Constructor fills the heap with 4, 5, 8; rejects 2](images/walkthrough-1.svg)

`Add(3)`: root `4`. `4 < 3`? No, rejected, heap unchanged, returns `4`.

![Step 2: Add(3) is rejected, heap unchanged, returns 4](images/walkthrough-2.svg)

`Add(5)`: root `4`. `4 < 5`? Yes, pop `4`, push `5`. Heap `{5, 5, 8}`, returns `5`.

![Step 3: Add(5) pops 4, pushes 5, returns 5](images/walkthrough-3.svg)

`Add(10)`: root `5`. `5 < 10`? Yes, pop `5`, push `10`. Heap `{5, 8, 10}`, returns `5`.

![Step 4: Add(10) pops 5, pushes 10, returns 5](images/walkthrough-4.svg)

`Add(9)`: root `5`. `5 < 9`? Yes, pop `5`, push `9`. Heap `{8, 9, 10}`, returns `8`.

![Step 5: Add(9) pops 5, pushes 9, returns 8](images/walkthrough-5.svg)

`Add(4)`: root `8`. `8 < 4`? No, rejected, heap unchanged, returns `8`.

![Step 6: Add(4) is rejected, heap unchanged, returns 8](images/walkthrough-6.svg)

Final returns: `[4, 5, 5, 8, 8]`, matches the expected output.

**Complexity:** `O(n log k)` for the constructor, `O(log k)` per `Add` call, `O(k)`
space. The heap never grows past size `k`.
