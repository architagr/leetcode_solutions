# Kth Largest Element in an Array — solution walkthrough

From `main.go`:

```go
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
```

The `MaxHeap` type is the same as day 76's: `Less` inverted to greater-than, and `Pop` detaching the last element for `container/heap` to have swapped the root into.

## What "kth largest" means

Position k in sorted-descending order — **not** the kth distinct value.

In `[5, 5, 4]` the 2nd largest is 5. The two fives occupy two separate positions. A solution that deduplicates passes every simple example and fails here, which is why the tests cover it explicitly.

## The walk

![Step 1](images/walkthrough-1.png)

Every element goes into the heap. The root is the maximum; nothing else about the arrangement is meaningful.

![Step 2](images/walkthrough-2.png)

The first pop yields the largest — the 1st largest, not the answer unless `k` is 1.

![Step 3](images/walkthrough-3.png)

The second pop yields the 2nd largest, which for `k = 2` is the answer.

The loop keeps only the most recent popped value, so after k iterations `x` holds the kth largest. Nothing is accumulated.

## What this actually costs

![Step 4](images/walkthrough-4.png)

Pushing n elements one at a time is O(n log n). Popping k times adds O(k log n).

So this is **O(n log n)**, which is the same as sorting the array and indexing it. What the heap buys is stopping after k pops instead of ordering everything — a constant-factor saving, not an asymptotic one.

That is worth saying plainly: this is a heap performing a sort, interrupted early.

## The version the problem is really asking for

Keep a **min**-heap capped at k:

- Push each element.
- If the heap exceeds k, pop the smallest.
- After one pass, the heap holds the k largest and its root is the kth largest.

That is **O(n log k)**. When k is 5 and n is a million — the situation that makes this question interesting — log k is 2 and log n is 20.

## Why a min-heap for that

The counter-intuitive part, and day 78 depends on it.

To retain the k largest, the element you need at your fingertips is **the smallest of the ones you kept**, because that is what an arriving element must beat to earn a place. So the root has to be the minimum of the retained set.

A max-heap of size k puts the largest at the root, which tells you nothing about which element to evict.

## Quickselect, briefly

Partition around a pivot, as in quicksort, and recurse only into the side containing position k. Average O(n), worst case O(n²) unless the pivot is chosen carefully.

Asymptotically the best of the three, and the fiddliest to write correctly under time pressure.

## Day 75 was the same question with different constraints

There, values arrived one at a time and the answer was required after each, so a size-k min-heap was the only workable structure — you cannot sort a stream that has not finished.

Here the whole array is present up front, which admits sorting, quickselect, and heaps of any shape. Same question; the constraints decide the answer.

## Complexity

- **Time: O(n log n)** as written. O(n log k) with a size-k min-heap; O(n) average with quickselect.
- **Space: O(n)** here, O(k) for the capped version.

## Test

`main_test.go` covers the worked examples, both ends of k, all-equal input, negatives, and the duplicate case — plus 2000 random trials against sorting descending and indexing, which is the definition of the answer.
