# Kth Largest Element in an Array — intuition

## Builds on

- [Day 75: Kth Largest Element in a Stream](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/701_800/kth_largest_element_in_a_stream/) — the same question asked of a fixed array rather than an arriving stream, which changes which heap is the right one

## The problem in one line

Return the kth largest element, counting duplicates as separate positions.

## The definition worth being careful about

"Kth largest" means the element at position k in sorted-descending order — not the kth *distinct* value.

In `[5, 5, 4]`, the 2nd largest is 5, not 4. Two fives occupy two positions. Getting this wrong produces a solution that passes the simple examples and fails on duplicates.

## Three solutions, and what separates them

**Sort and index.** O(n log n), two lines, completely correct. Worth saying plainly that this is a fine answer.

**A heap of everything, popped k times.** What this solution does. Also O(n log n) — it is a heap being used to perform a sort, one element at a time, and stopping early.

**A min-heap capped at k.** O(n log k). Keep the k largest seen so far; anything smaller than the smallest of them cannot be the answer and is discarded immediately. When k is small and n is large — the usual reason to ask this question — that is a real difference.

## Why the capped version needs a min-heap

This is the counter-intuitive part, and day 78 leans on it.

To hold the k largest, the element you need instant access to is **the smallest of them**, because that is the one a new arrival has to beat. So the root must be the minimum of the retained set, which means a min-heap.

A max-heap of size k would keep the largest at the root, which tells you nothing about whether a new element deserves a place.

## Day 75 is the same question, differently shaped

There, values arrived one at a time and the answer was needed after each — so a size-k min-heap was the only sensible structure.

Here the whole array is available up front, which permits sorting, quickselect and heaps of every shape. Same question, different constraints, and the constraints decide.

## Complexity

- **Time: O(n log n)** as written: O(n log n) to build the heap by pushing, plus O(k log n) to pop.
- **Space: O(n)** for the heap holding every element.
