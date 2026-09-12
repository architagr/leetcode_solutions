# Top K Frequent Elements — intuition

## Builds on

- [Day 77: Kth Largest Element in an Array](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/kth_largest_element_in_an_array/) — the size-k min-heap that day named but did not use, applied here to frequencies rather than values

## The problem in one line

Return the `k` most frequent values in an array, in any order.

## Two steps, and only the second is interesting

**Count.** A map from value to occurrences. Unavoidable, O(n), nothing to decide.

**Select the top k by count.** This is where the choice is.

## Why not sort

Sorting the distinct values by frequency and taking the first k is correct and is O(d log d) in the number of distinct values.

The follow-up on the problem page explicitly asks for better than O(n log n), which rules it out when d is close to n.

And the same objection as yesterday applies: sorting produces a total ordering when only the top k is wanted. The 7th most frequent element is computed and then discarded.

## The heap, capped at k

Yesterday named this and did not use it. Here it is the answer:

- Push each `(value, count)` pair.
- If the heap exceeds k, pop.
- What survives is the top k.

Every operation runs on a heap of at most k entries, so this is **O(d log k)**.

## It has to be a min-heap

The part that reads backwards. You want the *most* frequent, and the heap is a *min*-heap.

The reason is what the structure has to do on each arrival: decide which entry to evict. The one to evict is always the least frequent of those currently held — the weakest member. So that entry has to be at the root, which makes it a min-heap.

A max-heap of size k would keep the most frequent at the root, which is never in danger and tells you nothing.

Day 77 stated this rule; today is the problem where getting it wrong silently returns the k *least* frequent elements.

## The order of the output

Draining the heap gives ascending frequency. The problem says any order is acceptable, so it is returned as it comes.

Worth noticing, because the next day's problem does specify an order, and that changes the tool.

## Complexity

- **Time: O(n + d log k)** — counting, then a heap of at most k over d distinct values.
- **Space: O(d)** for the counts, plus O(k) for the heap.
