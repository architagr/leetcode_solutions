---
meta_title: "To keep the k most frequent, use a min-heap"
meta_description: "The inversion reads backwards every time. The element you need at the root is the weakest of the ones you kept, because that is the one an arrival displaces."
tags: [golang, heap, priority-queue, dsa]
---

![Day 78](HERO.png)

*365 Days of LeetCode Challenge — Day 78/365*

**[347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/)** (Medium)

Given an integer array and an integer `k`, return the `k` most frequent elements. Any order is acceptable.

## Two steps, and only one of them is a decision

**Count the occurrences.** A map from value to count, one pass, O(n). There is nothing to choose here and nothing that can go wrong.

**Select the top k by count.** Everything interesting lives in this step, and it operates on `d` distinct values rather than `n` elements — which matters, because `d` can be far smaller than `n`.

## Why not sort the counts

Take the map, flatten it into a slice of `(value, count)` pairs, sort by count descending, and read off the first k. Correct, and about six lines.

It is O(d log d). And the follow-up note on the problem page asks specifically for something better than O(n log n), which rules this out whenever `d` is close to `n` — an array of all-distinct values, say.

There is also the objection from day 77, which applies unchanged. Sorting produces a **total ordering** of every distinct value. This problem wants the top k. The relative position of the 7th and 8th most frequent elements is computed, carefully, and then thrown away.

## The heap, capped at k

```go
h := &minHeap{}
for ele, count := range m {
	heap.Push(h, entry{ele: ele, count: count})
	if h.Len() > k {
		heap.Pop(h)
	}
}
```

Push each pair. If that tips the heap past k, evict one immediately.

The heap therefore holds k entries almost always, and k+1 for the instant between a push and its matching pop. Every operation is O(log k) rather than O(log d), and there are d of them: **O(d log k)**.

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

## The line that does all the work

```go
func (h minHeap) Less(i, j int) bool { return h[i].count < h[j].count }
```

Two decisions are packed into that.

**It compares `count`, not `ele`.** The heap is ordered by frequency. The value itself is cargo — carried along so it can be returned at the end, never looked at for ordering.

**It uses `<`, not `>`.** Day 76 obtained a max-heap by writing `>` and lying to `container/heap` about which element was smallest. Here the honest `<` is what is wanted.

And this is the problem where that choice stops being a formality.

## Why a min-heap, when the goal is the most frequent

I said on day 77 that this reads backwards, and it does. Here is the reasoning that makes it stick.

Do not think about what you want to *end up* with. Think about what the structure has to *do* on every single arrival.

On each arrival it must answer: does this new entry belong in the set I am keeping, and if so, which current member does it push out?

The member it pushes out is always the **least frequent of the ones currently kept**. It is the weakest one, the one sitting on the boundary, the only one whose place is in question. Every other member is strictly safer than it.

So that is the entry you need constant-time access to. Which means it has to be at the root. Which makes it a min-heap.

Now consider the alternative properly. A max-heap of size k puts the **most** frequent entry at the root. That entry is in no danger at all — it is the last thing that would ever be evicted. To find the one you actually need, you would have to scan the heap, which is O(k), and scanning defeats the entire purpose of using a heap.

The rule, worth memorising in both directions:

- To keep the k **largest**, use a **min**-heap.
- To keep the k **smallest**, use a **max**-heap.

And the failure mode deserves naming. Getting this backwards does not crash, does not panic, and does not produce anything obviously malformed. It returns the k **least** frequent elements — a well-formed answer of the right length to a different question.

![Step 4](images/walkthrough-4.png)

## The order things come out in

```go
for h.Len() > 0 {
	ans = append(ans, heap.Pop(h).(entry).ele)
}
```

Draining a min-heap yields ascending order, so the *least* frequent member of the top k emerges first and the most frequent last.

The problem states that any order is acceptable, so this is returned as it comes rather than being reversed.

It is worth flagging rather than skipping past, because tomorrow's problem is the same question with an ordering requirement attached — and that one requirement changes which tool is the right one.

## Complexity

- **Time: O(n + d log k).** O(n) to count. Then d pushes and up to d pops, each on a heap bounded by k.
- **Space: O(d)** for the count map, which is unavoidable, plus O(k) for the heap.

The sorting version is O(n + d log d). When k is small and d is large — the case the question is built around — the gap between `log k` and `log d` is the whole point.

## Builds on

- [Day 77: Kth Largest Element in an Array](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/kth_largest_element_in_an_array/) — the size-k min-heap that day named but did not use, applied here to frequencies rather than values

Full code and the step-by-step walkthrough:
[top_k_frequent_elements](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/301_400/top_k_frequent_elements/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
