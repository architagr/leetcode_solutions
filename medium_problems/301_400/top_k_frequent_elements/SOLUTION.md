# Top K Frequent Elements — solution walkthrough

From `top_k_frequent_elements.go`:

```go
type entry struct {
	ele, count int
}

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

	h := &minHeap{}
	for ele, count := range m {
		heap.Push(h, entry{ele: ele, count: count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := make([]int, 0, h.Len())
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(entry).ele)
	}
	return ans
}
```

## Counting is the boring half

```go
m := make(map[int]int)
for _, v := range nums {
	m[v]++
}
```

![Step 1](images/walkthrough-1.png)

O(n), no decisions. Everything interesting happens to the `d` distinct values that come out of it.

## Why not sort the counts

Building a slice of `(value, count)` and sorting by count is correct, and it is O(d log d).

The follow-up on the problem page asks for better than O(n log n), which rules that out when d approaches n. And the objection from day 77 applies again: sorting computes a total ordering when only the top k is wanted. The 7th most frequent element gets its position calculated and then thrown away.

## The heap, capped at k

```go
heap.Push(h, entry{ele: ele, count: count})
if h.Len() > k {
	heap.Pop(h)
}
```

Push, and if that tips the heap over k, evict. The heap never holds more than k+1 entries momentarily and k the rest of the time, so every operation is O(log k).

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

Pushing the third entry takes the heap to 3 with k = 2, so one has to go.

## `Less` compares counts, and it is a min-heap

```go
func (h minHeap) Less(i, j int) bool { return h[i].count < h[j].count }
```

Two things in one line.

**It compares `count`, not `ele`.** The heap is ordered by frequency; the value is cargo.

**It is `<`, not `>`.** Day 76 got a max-heap by writing `>` and lying to `container/heap`. Here the honest `<` is wanted, and this is the problem where that choice does the work.

## Why the min-heap is right when you want the most frequent

This reads backwards every time, and day 77 set it up.

Ask what the structure does on each arrival: decide whether the new entry belongs in the retained set, and if so, which incumbent it displaces.

The incumbent displaced is always **the least frequent of those currently kept** — the weakest member, the one on the boundary. That is the entry you need instant access to. So it has to be at the root, which makes it a min-heap.

A max-heap of size k would hold the most frequent at the root. That entry is in no danger whatsoever, and finding the one to evict would mean scanning the heap — O(k), which discards the reason for using one.

Getting this backwards does not crash. It silently returns the k **least** frequent elements.

![Step 4](images/walkthrough-4.png)

## The output order

```go
for h.Len() > 0 {
	ans = append(ans, heap.Pop(h).(entry).ele)
}
```

Draining a min-heap gives ascending frequency, so the least frequent of the top k comes out first.

The problem says any order is acceptable, so it is returned as it comes rather than reversed. Worth noticing: tomorrow's problem *does* specify an order, and that single difference changes which tool fits.

## Complexity

- **Time: O(n + d log k)** — O(n) to count, then d pushes and pops on a heap of size at most k.
- **Space: O(d)** for the count map, plus O(k) for the heap.

The sorting version is O(n + d log d). When k is small and d is large, the difference is the point of the question.

## Test

`main_test.go` has the worked examples. `heap_test.go` adds k equal to the distinct count, k of 1, all-identical input and negatives, plus 2000 random trials against counting and sorting — skipping trials where the kth and (k+1)th frequencies tie, since the problem guarantees a unique answer and several would otherwise be correct.
