**365 Days of LeetCode Challenge — Day 38/365**
**Kth Largest Element in a Stream** (Easy)
🔗 https://leetcode.com/problems/kth-largest-element-in-a-stream/

We never need the full sorted order, only the top k scores, and among those the kth largest is exactly the smallest one. So the job becomes: keep a bounded set of k values and read or replace its minimum fast. A min-heap capped at size `k` does that in O(log k). After that, the constructor and `Add` are the same eviction check written twice.

```go
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
	return KthLargest{k: k, h: obj}
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
```

The `minHeap` type and its `sort.Interface` + `heap.Interface` methods are in the repo.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/701_800/kth_largest_element_in_a_stream/SOLUTION.md
