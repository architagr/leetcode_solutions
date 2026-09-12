---
meta_title: "A heap keeps one promise, which is why it is fast"
meta_description: "Sorting answers a bigger question than this problem asks. A heap maintains only that the largest is at the root, and maintaining less is what makes it cheap."
---

![Day 76](HERO.png)

## 365 Days of LeetCode Challenge — Day 76/365

**[1046. Last Stone Weight](https://leetcode.com/problems/last-stone-weight/)** (Easy)

Repeatedly smash the two heaviest stones together until at most one remains.

Heaps start today, and yesterday's Kth Largest in a Stream was the bridge — the same structure, used there to hold a running answer.

## Why sorting is the wrong shape

Sort descending, take the first two, smash them, put the remainder back — in the right place, which is an insertion. O(n) per round, paid every round.

The deeper mismatch is that **sorting answers a question nobody asked**. This problem never needs the third-heaviest stone, or the ordering of the rest. It needs the maximum, repeatedly, from a collection that keeps changing.

## A heap keeps exactly one promise

The largest element is at the root. Nothing else.

It makes no claim about the order of anything below. Read a heap's backing array left to right and it looks close to random — and that is the point. **Maintaining less is what makes it cheaper.**

![Step 1](images/walkthrough-1.png)

The invariant is visible in the shape: every parent is at least as large as its children, and nothing more holds.

- Find the maximum: O(1).
- Remove it: O(log n).
- Add one: O(log n).

Exactly the operations this problem performs, and no others.

## One inverted comparison

```go
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
```

Go's `container/heap` always builds a **min**-heap — it puts whatever `Less` calls smallest at the root.

A max-heap is obtained by lying to it. That single inverted comparison is the whole difference, and it is the most common place to go wrong with this package.

## The loop

![Step 2](images/walkthrough-2.png)

Two pops give the two heaviest. `x >= y` is guaranteed by the heap, so `x - y` is never negative and no `abs` is needed.

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

Equal stones give `z == 0` and nothing is pushed — both destroyed, as specified.

![Step 5](images/walkthrough-5.png)

## The `Pop` that looks backwards

```go
func (h *MaxHeap) Pop() any {
	old := *h
	n := old.Len()
	x := old[n-1]      // the LAST element, not the root
	*h = old[:n-1]
	return x
}
```

That looks wrong for a pop, and it is correct.

`container/heap` splits the work: the package's `heap.Pop` swaps the root with the last element and sifts down, **then** calls your `Pop` to detach what is now at the end. Your method never touches the root.

Returning `old[0]` instead is a classic mistake, and it produces silently wrong answers rather than a crash.

## One thing this could do better

```go
maxHeap := &MaxHeap{}
heap.Init(maxHeap)          // Init on an empty heap does nothing
for _, v := range stones {
	heap.Push(maxHeap, v)   // O(log n) each, so O(n log n) total
}
```

Appending all the stones first and calling `Init` once heapifies in **O(n)** — Floyd's method, sifting down from the middle, genuinely linear.

It does not change the overall complexity, since the main loop is O(n log n) anyway. Worth knowing which one you are paying for.

## Complexity

- **Time: O(n log n)**. Up to n rounds, each a constant number of O(log n) operations.
- **Space: O(n)**, or O(1) extra if the input slice is heapified in place.

## Builds on

- [Day 75: Kth Largest Element in a Stream](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/701_800/kth_largest_element_in_a_stream/) — the same structure, used there to hold a running answer and here to repeatedly take the largest

Full code and the step-by-step walkthrough:
[last_stone_weight](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1001_1100/last_stone_weight/SOLUTION.md)

#DSA #LeetCode #Golang #Heap #PriorityQueue #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
