---
meta_title: "This is a heap being used as a sort"
meta_description: "Pushing all n and popping k is O(n log n), the same as sorting. The version the question is really asking for is a min-heap capped at k."
tags: [golang, heap, priority-queue, dsa]
---

![Day 77](HERO.png)

*365 Days of LeetCode Challenge — Day 77/365*

**[215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)** (Medium)

Given an integer array and an integer `k`, return the kth largest element.

## The definition is the first trap

"Kth largest" means the element at **position k in sorted-descending order**. It does not mean the kth distinct value.

In `[5, 5, 4]`, the 2nd largest is **5**, not 4. There are two fives and they occupy two separate positions.

This matters because the natural mental shortcut — "the kth biggest *different* number" — is wrong, and it is wrong in a way that passes `[3,2,1,5,6,4]` and every other example without repeats. It is the first thing the tests for this solution check.

## Three ways to answer it

**Sort and index.** `sort.Sort(sort.Reverse(...))` then `nums[k-1]`. O(n log n), two lines, and completely correct. I want to be clear that this is a good answer. In production, on an array that fits in memory, it is probably what you should write.

**A heap of everything, popped k times.** What this solution does.

**A min-heap capped at k.** O(n log k), and the one the question is angling for.

## What this solution costs

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

![Step 1](images/walkthrough-1.png)

Every element goes in. The root is the maximum; nothing else about the arrangement means anything.

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

The loop keeps only the most recently popped value, so after k iterations `x` holds the kth largest. Nothing is accumulated, which is neat.

Now add it up. Pushing n elements one at a time costs O(log n) each, so O(n log n). Popping k times adds O(k log n).

![Step 4](images/walkthrough-4.png)

The total is **O(n log n)** — exactly what sorting and indexing costs.

What did the heap buy? It stopped after k pops instead of ordering the remaining n − k elements. That is a real saving in wall-clock terms when k is small, and it is a constant factor, not an asymptotic improvement.

It is worth being blunt about this rather than letting "I used a heap" stand in for "I made it faster": **this is a heap performing a sort, interrupted early.**

## The version the question is angling for

Keep a min-heap and never let it grow past k:

- Push each element as you scan.
- If the heap now holds more than k, pop the smallest.
- After one pass, the heap contains exactly the k largest elements, and its root is the kth largest.

Every operation is on a heap of size at most k, so the whole thing is **O(n log k)**.

The difference is not cosmetic. This question is usually asked because k is small and n is enormous — the top 10 of a million. Then log k is about 3 and log n is about 20, and you are also holding 10 elements instead of a million.

## Why that has to be a min-heap

Here is the part that trips people up, and tomorrow's problem is built on it.

You want the k **largest**. So the instinct is to reach for a max-heap.

But think about what the structure has to do on every arrival: decide whether this new element deserves to be in the retained set, and if so, which incumbent it displaces.

The incumbent it displaces is always **the smallest of the ones currently kept**. That is the weakest member, the one on the boundary. So that is the element you need instant access to — which means it must be at the root — which means a min-heap.

A max-heap of size k would keep the *largest* of your retained set at the root. That element is in no danger and tells you nothing useful. You would have to scan the heap to find the one to evict, and scanning a heap is O(k), which throws away the reason for using one.

**To keep the k largest, use a min-heap. To keep the k smallest, use a max-heap.** The inversion is the whole idea, and it looks wrong every time until you say out loud which element you need to look at.

## Quickselect, for completeness

Partition the array around a pivot, as in quicksort. After partitioning you know exactly which position the pivot landed in, so you recurse into whichever side contains position k — and only that side.

Average O(n), because the work halves each time. Worst case O(n²) if pivots are chosen badly, which is fixable with a random or median-of-medians pivot.

It is asymptotically the best of the three and the one I would be least confident writing correctly in an interview without a lot of care over the partition boundaries.

## Day 75 was this question with one constraint changed

Two days ago, Kth Largest Element in a **Stream**: values arrive one at a time and the answer is needed after each.

That constraint removes almost everything. You cannot sort a stream that has not ended. You cannot quickselect over data you have not seen. A size-k min-heap is not the clever answer there — it is close to the only answer, because it maintains the running result in O(log k) per arrival with O(k) memory.

Here, the whole array is available up front, and that permits everything. Same question, different constraints, and the constraints do the deciding. That is worth internalising more than any of the individual techniques.

## Complexity

- **Time: O(n log n)** as written. O(n log k) with a size-k min-heap. O(n) average with quickselect.
- **Space: O(n)** here, holding every element; O(k) for the capped version.

## Builds on

- [Day 75: Kth Largest Element in a Stream](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/701_800/kth_largest_element_in_a_stream/) — the same question asked of a fixed array rather than an arriving stream, which changes which heap is the right one

Full code and the step-by-step walkthrough:
[kth_largest_element_in_an_array](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/kth_largest_element_in_an_array/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
