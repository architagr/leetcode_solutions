# Merge k Sorted Lists — intuition

## Builds on

- [Day 22: Merge Two Sorted Lists](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/merge_two_sorted_lists/) — the dummy head and the splice-don't-copy rule, both used here unchanged
- [Day 75: Kth Largest Element in a Stream](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/701_800/kth_largest_element_in_a_stream/) — a heap held at a fixed size while values stream through it

## The problem in one line

Merge `k` sorted linked lists into one sorted list.

## Two arcs meet here

Day 22 merged **two** sorted lists: compare the two heads, take the smaller, advance. The comparison was an `if`, because with two candidates that is all a comparison needs to be.

With `k` lists, "take the smallest head" stops being an `if` and becomes a query against a changing collection — which is what a heap is for.

That is the whole solution. Day 22's merge with the `if` replaced by a heap.

## What goes in the heap

Only the **head of each list** — never every node.

A list's second element cannot possibly be the next smallest while its first is still unplaced, so it is not a candidate yet. It becomes one exactly when its predecessor is placed.

So the heap holds at most `k` nodes regardless of how many nodes the lists contain between them. That is the difference between O(N log k) and O(N log N).

## Why not scan the k heads each round

Looking at all `k` heads and taking the minimum is correct and costs O(k) per node placed, so O(N·k) overall.

The scan re-examines the same heads over and over. A heap remembers the ordering between rounds, so it pays O(log k) instead — and only for the heads that actually changed.

## Splicing, not copying

The nodes already exist and are already in the right order within their own lists. Building new ones and copying values across costs O(N) memory to produce a list that could have been made by rewriting pointers.

This is day 22's rule, and the same one 206 and 21 were corrected for earlier in the series.

There is one consequence to watch: a spliced node's `Next` still points into its original list. The merged list must be terminated explicitly, or what you hand back can contain a cycle.

## Complexity

- **Time: O(N log k)** for N total nodes — each is pushed and popped once on a heap of size at most k.
- **Space: O(k)** for the heap. Nothing is allocated per node.
