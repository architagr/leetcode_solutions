---
meta_title: "Two arcs, fifty-eight days apart, meet here"
meta_description: "Merge k Sorted Lists is day 22's merge with the if replaced by a heap. The heap holds k heads, never N nodes, and that is the whole complexity argument."
---

![Day 80](HERO.png)

## 365 Days of LeetCode Challenge — Day 80/365

**[23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)** (Hard)

Merge `k` sorted linked lists into one sorted list.

This is the problem the last two months have been pointing at. It needs the linked list arc from day 22 and the heap arc from this week, and it needs both at once.

## Day 22, with one thing changed

Day 22 merged **two** sorted lists: compare the heads, take the smaller, advance. The comparison was an `if` — with exactly two candidates, that is all a comparison needs to be.

With `k` lists, "take the smallest head" is no longer one comparison. It is a repeated query for the minimum of a collection that changes after every answer.

Which is the definition of what a heap is for.

So: **day 22's merge, with the `if` replaced by a heap.** The dummy head, the splicing, the shape of the loop — all unchanged.

## The heap holds heads, not nodes

```go
for _, l := range lists {
	if l != nil {
		*h = append(*h, l)
	}
}
heap.Init(h)
```

![Step 1](images/walkthrough-1.png)

Only one node per list.

This is the observation the complexity rests on. A list's second element **cannot** be the next smallest overall while its first is still unplaced — the list is sorted, so its own head beats it. It is not a candidate yet, and becomes one at exactly the moment its predecessor is placed.

So the heap is size `k`, not size `N`. That is O(N log k) rather than O(N log N).

## Two details in that setup

**Empty lists are filtered out.** A `nil` must never be pushed — `Less` would dereference it on the first sift. The problem's own examples include `[[]]` and `[]`.

**`Init` rather than k pushes.** Heapifying once is O(k) by Floyd's method; pushing one at a time is O(k log k). Day 76 flagged this and did not use it. Here it is used.

## The loop

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

Pop the smallest head, splice it on, push its successor — which has just become a candidate. One node out, at most one in, so the heap stays at k.

## Why not just scan the k heads

Correct, and O(k) per placed node — so O(N·k) overall.

The waste is precise: between two rounds only **one** head changed, the list that was just advanced. The scan re-examines all k anyway. A heap keeps the ordering between rounds and pays O(log k) to absorb that single change.

## Splicing, and the cycle it can create

The nodes already exist. Copying their values into fresh ones costs O(N) memory to produce what pointer rewriting gives for free — day 22's rule, and the one 206 and 21 were corrected for.

But splicing has a consequence worth naming:

```go
tail.Next = nil
```

Every spliced node arrives with a `Next` still pointing into its original list.

In practice the last node popped is the overall maximum, which is its own list's tail, so its `Next` is already nil. The explicit cut is belt and braces — and the tests walk the result with a step cap so a cycle **fails** rather than hangs. A merge that reuses nodes is exactly the shape of code where an unterminated list becomes an infinite one.

![Step 5](images/walkthrough-5.png)

## Complexity

- **Time: O(N log k)**. Each of N nodes pushed and popped once on a heap bounded by k, plus O(k) to heapify.
- **Space: O(k)**. Nothing allocated per node; the output reuses the input.

Scanning is O(N·k). Concatenating and sorting is O(N log N).

## Builds on

- [Day 22: Merge Two Sorted Lists](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/merge_two_sorted_lists/) — the dummy head and the splice-don't-copy rule, both used here unchanged
- [Day 75: Kth Largest Element in a Stream](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/701_800/kth_largest_element_in_a_stream/) — a heap held at a fixed size while values stream through it

Full code and the step-by-step walkthrough:
[merge_k_sorted_lists](https://github.com/architagr/leetcode_solutions/blob/main/hard_problems/1_100/merge_k_sorted_lists/SOLUTION.md)

#DSA #LeetCode #Golang #Heap #LinkedList #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
