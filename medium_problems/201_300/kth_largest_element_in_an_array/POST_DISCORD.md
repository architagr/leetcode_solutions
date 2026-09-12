**365 Days of LeetCode Challenge — Day 77/365**
**Kth Largest Element in an Array** (Medium)
🔗 https://leetcode.com/problems/kth-largest-element-in-an-array/

First the definition: "kth largest" is position k in sorted-descending order, NOT the kth distinct value. In `[5,5,4]` the 2nd largest is 5. Deduplicating passes every simple example and fails there.

```go
for _, v := range nums {
	heap.Push(mHeap, v)
}
var x int
for i := 0; i < k; i++ {
	x = (heap.Pop(mHeap)).(int)
}
return x
```

Pushing n one at a time is O(n log n); popping k adds O(k log n). So this is O(n log n) - the same as sorting and indexing. The heap only buys stopping after k pops instead of ordering everything. A constant factor, not an asymptotic one. This is a heap performing a sort, interrupted early.

The version the question really wants: a MIN-heap capped at k. Push each element, and if the heap exceeds k, pop the smallest. O(n log k).

Why a MIN-heap when you want the largest? To retain the k largest, the element you need instant access to is the SMALLEST of the ones you kept - that is what an arriving element has to beat. A max-heap of size k tells you nothing about what to evict.

Quickselect does it in O(n) average, and is the fiddliest of the three.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/kth_largest_element_in_an_array/SOLUTION.md
