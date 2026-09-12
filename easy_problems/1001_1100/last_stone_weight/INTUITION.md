# Last Stone Weight — intuition

## Builds on

- [Day 75: Kth Largest Element in a Stream](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/701_800/kth_largest_element_in_a_stream/) — the same structure, used there to hold a running answer and here to repeatedly take the largest

## The problem in one line

Repeatedly smash the two heaviest stones together until at most one is left.

## Why sorting is the wrong shape

Sort descending, take the first two, smash them, and put the remainder back — in the right place, which means an insertion. That is O(n) per round, and you pay it every round.

The real mismatch is that sorting answers a question nobody asked. You never need the third-heaviest stone, or the order of the rest. You need the maximum, repeatedly, from a collection that keeps changing.

## What a heap is for

A heap keeps exactly one promise: the largest element is at the root. It says nothing about the order of anything else, and that is the point — maintaining less means maintaining it more cheaply.

- Find the maximum: O(1).
- Remove it: O(log n).
- Add a new element: O(log n).

That is precisely the set of operations this problem performs, and nothing more.

## The shape

A binary heap is a complete binary tree where every parent is at least as large as its children. "Complete" means it fills level by level, left to right, which is why it can live in a flat array with no pointers: the children of index `i` are at `2i+1` and `2i+2`.

Pushing appends and then swaps upward while the new value beats its parent. Popping moves the last element to the root and swaps downward. Both walk one root-to-leaf path, so both are O(log n).

## The loop

While more than one stone remains, pop twice and push the difference back if it is non-zero. Equal stones cancel and nothing is pushed.

At the end, one stone or none. Returning 0 for the empty case is not an edge case bolted on — it is the answer the problem specifies.

## Complexity

- **Time: O(n log n).** Heapifying is O(n), and each of up to n rounds does a constant number of O(log n) operations.
- **Space: O(n)** for the heap, or O(1) extra if you heapify the input array in place.
