# Last Stone Weight — solution walkthrough

From `main.go`:

```go
type MaxHeap []int

func (h *MaxHeap) Len() int { return len(*h) }
func (h *MaxHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Push(val interface{}) { *h = append(*h, val.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for _, v := range stones {
		heap.Push(maxHeap, v)
	}
	for maxHeap.Len() > 1 {
		x := heap.Pop(maxHeap).(int)
		y := heap.Pop(maxHeap).(int)
		z := x - y
		if z != 0 {
			heap.Push(maxHeap, z)
		}
	}
	if maxHeap.Len() == 0 {
		return 0
	}
	return heap.Pop(maxHeap).(int)
}
```

## Why not sort

Sort descending, take the first two, smash, insert the remainder back in position. The insert is O(n) and you pay it every round.

But the deeper mismatch is that sorting answers a bigger question than the one asked. This problem never needs the third-heaviest stone or the ordering of the rest. It needs *the maximum, repeatedly, from a changing collection* — and nothing else.

A heap maintains exactly that and no more, which is why it is cheaper.

## What the heap actually promises

Only one thing: the largest element is at the root.

It makes no claim about the order of anything else. Reading a heap's backing array from left to right gives you something that looks almost random, and that is fine — the ordering it does *not* maintain is the ordering nobody needs.

![Step 1](images/walkthrough-1.png)

The tree is drawn so the invariant is visible: every parent is at least as large as its children. Nothing else holds.

## `Less` is what makes it a max-heap

```go
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
```

Go's `container/heap` always builds a **min**-heap: it puts whatever `Less` reports as smallest at the root.

So a max-heap is obtained by lying to it — defining `Less` as greater-than. That one inverted comparison is the entire difference, and it is the single most common place to go wrong when reaching for this package.

## The loop

```go
for maxHeap.Len() > 1 {
	x := heap.Pop(maxHeap).(int)
	y := heap.Pop(maxHeap).(int)
	z := x - y
	if z != 0 {
		heap.Push(maxHeap, z)
	}
}
```

![Step 2](images/walkthrough-2.png)

Two pops give the two heaviest. `x >= y` is guaranteed by the heap, so `x - y` is never negative and no `abs` is needed.

![Step 3](images/walkthrough-3.png)

Pushing the remainder re-settles the heap in O(log n). The next round's maximum is found with no searching.

![Step 4](images/walkthrough-4.png)

Equal stones produce `z == 0` and nothing is pushed — both are destroyed, which is what the problem says happens.

![Step 5](images/walkthrough-5.png)

## `heap.Init` on an empty heap

```go
maxHeap := &MaxHeap{}
heap.Init(maxHeap)
for _, v := range stones {
	heap.Push(maxHeap, v)
}
```

`Init` is called on a heap with nothing in it, which does nothing. The values are then pushed one at a time, at O(log n) each, for O(n log n) overall.

Appending all the stones first and calling `Init` once would heapify in **O(n)** — Floyd's method, which sifts down from the middle of the array and is genuinely linear.

It does not change the overall complexity here, because the main loop is O(n log n) regardless. It is worth knowing which of the two you are paying for.

## The interface, and why `Pop` looks backwards

```go
func (h *MaxHeap) Pop() any {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[:n-1]
	return x
}
```

This removes the **last** element, not the root, which looks wrong for a pop.

It is correct because of the division of labour in `container/heap`: the package's `heap.Pop` first swaps the root with the last element and sifts down, *then* calls your `Pop` to detach what is now at the end. Your method never sees the root.

Implementing `Pop` to return `old[0]` is a classic mistake and produces silently wrong answers rather than a crash.

## The zero case

```go
if maxHeap.Len() == 0 {
	return 0
}
```

Stones can cancel exactly, leaving nothing. The problem specifies 0 for that, so this is the stated answer rather than defensive code.

## Complexity

- **Time: O(n log n).** Up to n rounds, each a constant number of O(log n) heap operations.
- **Space: O(n)** for the heap. Heapifying the input slice in place would make it O(1) extra.

## Test

`main_test.go` covers the worked example `[2,7,4,1,8,1]` giving 1, and `[1]` giving 1. The mutual-destruction path — everything cancelling to 0 — is the case worth adding if you extend it.
