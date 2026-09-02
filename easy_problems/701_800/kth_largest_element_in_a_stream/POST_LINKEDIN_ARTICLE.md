# 365 Days of LeetCode Challenge — Day 20/365

## Kth Largest Element in a Stream

[LeetCode 703 — Kth Largest Element in a Stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/) · Difficulty: Easy

Imagine a university admissions office watching test scores roll in from applicants in
real time. After every single score submitted, they need to instantly know: what's the
`kth` highest score so far? That's the shape of today's problem — design a class that
tracks a running stream of numbers and, after every insertion, reports the kth largest
value seen so far.

```
KthLargest(int k, int[] nums)  // initialize with k and a starting stream of scores
int add(int val)                // add a new score, return the kth largest so far
```

### Example

```
Input:
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]

Output: [null, 4, 5, 5, 8, 8]

Explanation:
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);  // return 4
kthLargest.add(5);  // return 5
kthLargest.add(10); // return 5
kthLargest.add(9);  // return 8
kthLargest.add(4);  // return 8
```

## The intuition

The naive approach — keep every score seen so far, re-sort the whole collection every
time `add` is called, and read off the kth element — works, but it's wasteful. With up
to `10^4` calls to `add`, doing `O(n log n)` work on every single one adds up fast.

Here's the reframe that unlocks an efficient solution: we never actually need to know
the full sorted order of every score we've ever seen. We only ever need to answer one
narrow question — **what is the kth largest value right now?** That means we only need
to track the **top k scores**, nothing else. Everything below the top k is irrelevant to
every future query, forever (a smaller score can never become the kth largest as long as
at least k values above it exist).

And here's the key trick: among just those top k scores, the kth largest one is exactly
the **smallest** of the group — the "weakest link" holding onto a spot in the top k. So
the whole problem reduces to: efficiently maintain a bounded set of (at most) k values,
and be able to read and replace its minimum quickly. That is *precisely* what a
**min-heap** is built for.

The strategy in plain terms:
- Keep a min-heap holding at most `k` elements — the k largest scores seen so far.
- When a new value shows up:
  - If the heap isn't full yet (fewer than `k` elements), the new value is automatically
    part of the current top-k — just push it in.
  - Otherwise, compare it against the heap's root (the smallest of the current top-k).
    If the new value beats the root, the root gets evicted and the new value takes its
    place. If not, the new value doesn't crack the top-k, and the heap is left alone.
- After every `add`, the min-heap's root — its smallest element — is by construction the
  kth largest value seen across the whole stream, because the heap holds exactly the
  top-k values, and the smallest among them sits in the kth position once you sort
  everything in descending order.

The constructor does the exact same thing, just seeded with a batch of values instead of
one at a time: fill the heap with the first `k` elements of `nums` unconditionally, then
run every remaining element through the same "does it beat the current top-k's weakest
member?" test.

**Complexity:** `O(log k)` per `add` call (a bounded heap push/pop), `O(n log k)` for the
constructor over an initial array of length `n`, and `O(k)` space — the heap never grows
past size `k` no matter how long the stream runs.

## The solution, walked through

The implementation lives in `min.go`, built on Go's `container/heap` interface applied
to a `minHeap []int`.

### The `minHeap` type

```go
type minHeap []int

func (mHeap *minHeap) Len() int { return len(*mHeap) }
func (mHeap *minHeap) Less(i, j int) bool { return (*mHeap)[i] < (*mHeap)[j] }
func (mHeap *minHeap) Swap(i, j int) { (*mHeap)[i], (*mHeap)[j] = (*mHeap)[j], (*mHeap)[i] }

func (h *minHeap) Push(val any) { *h = append(*h, val.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[:n-1]
	return x
}
```

`Less(i, j int) bool { return (*mHeap)[i] < (*mHeap)[j] }` is the line that makes this a
**min-heap** rather than a max-heap: the smallest value always sorts to the front, so
`(*obj)[0]` is always the current minimum. `Push` and `Pop` follow the standard
`container/heap` contract — `Push` appends to the end and `heap.Push` sifts it up to
restore the heap invariant; `Pop` returns whatever `heap.Pop` has already swapped to the
end of the slice (the sift-down happens first, then this just removes the element).

### `Constructor(k int, nums []int) KthLargest`

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
	return KthLargest{
		k: k,
		h: obj,
	}
}
```

`limit` caps the initial unconditional fill at `k` — or fewer, since the constraints
allow `len(nums) < k`. The first loop pushes those first `limit` elements without any
comparison, since the heap isn't full yet. The second loop runs every *remaining*
element in `nums` through the same eviction test `Add` uses: if it beats the heap's
current minimum, the minimum is popped out and the new value takes its place; otherwise
nothing happens.

Tracing it against the example — `Constructor(3, [4, 5, 8, 2])`: `limit = 3`, so `4`,
`5`, `8` get pushed unconditionally, filling the heap to `{4, 5, 8}` with root `4`. Then
`i = 3`, `nums[3] = 2`: is `root(4) < 2`? No — `2` is rejected, heap stays `{4, 5, 8}`.

![Step 1: Constructor fills the heap with 4, 5, 8; rejects 2](images/walkthrough-1.svg)

### `Add(val int) int`

```go
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

Two branches:
- **Heap not yet full** (`this.h.Len() < this.k`): fewer than `k` values have been seen
  overall, so anything new is automatically part of the top-k — push it in directly.
- **Heap already at capacity `k`**: compare `val` against the root, the current weakest
  member of the top-k. If `val` beats it, pop the root and push `val` in its place. If
  `val` doesn't beat the root, leave the heap untouched.

Either way, the function returns `(*this.h)[0]` — the root, which is always the kth
largest value seen so far by construction.

Since the heap is already full for the rest of the example (`Len() == k == 3`), every
remaining call takes the comparison branch:

**`Add(3)`**: root is `4`. Is `4 < 3`? No — rejected, heap stays `{4, 5, 8}`, returns
`4`.

![Step 2: Add(3) is rejected, heap unchanged, returns 4](images/walkthrough-2.svg)

**`Add(5)`**: root is `4`. Is `4 < 5`? Yes — pop `4`, push `5`. Heap becomes
`{5, 5, 8}` with root `5`, returns `5`.

![Step 3: Add(5) pops 4, pushes 5, returns 5](images/walkthrough-3.svg)

**`Add(10)`**: root is `5`. Is `5 < 10`? Yes — pop that `5`, push `10`. Heap becomes
`{5, 8, 10}` with root `5` (the *other* `5` already in the heap), returns `5`.

![Step 4: Add(10) pops 5, pushes 10, returns 5](images/walkthrough-4.svg)

**`Add(9)`**: root is `5`. Is `5 < 9`? Yes — pop `5`, push `9`. Heap becomes
`{8, 9, 10}` with root `8`, returns `8`.

![Step 5: Add(9) pops 5, pushes 9, returns 8](images/walkthrough-5.svg)

**`Add(4)`**: root is `8`. Is `8 < 4`? No — rejected, heap stays `{8, 9, 10}`, returns
`8`.

![Step 6: Add(4) is rejected, heap unchanged, returns 8](images/walkthrough-6.svg)

Final sequence of returns: `[4, 5, 5, 8, 8]` — exactly the expected output.

### Complexity

- **Time:** `O(n log k)` for the constructor, `O(log k)` per `Add` call.
- **Space:** `O(k)` — the heap never holds more than `k` elements, regardless of how
  long the stream runs.

---

*Part of the 365 Days of LeetCode Challenge. Follow along for daily problem
breakdowns.*
