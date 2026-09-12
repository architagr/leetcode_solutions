# Ugly Number II — intuition

## The problem in one line

Return the nth positive integer whose only prime factors are 2, 3 and 5.

## Why not test each integer

Walk upward from 1, dividing out every 2, 3 and 5, and check whether 1 remains. Correct, and hopeless: the 1690th ugly number is over two billion, so almost everything tested is a miss.

The fix is to stop testing candidates and start **generating** them.

## Every ugly number is an earlier one, multiplied

Nothing produces an ugly number except multiplying a smaller ugly number by 2, 3 or 5. That is what the definition says, read forwards.

So the sequence can be built from itself: keep the ugly numbers found so far, and the next one is the smallest unused product of an earlier one with 2, 3 or 5.

## The heap solution, which this is not

Push 1. Repeatedly pop the smallest, push its three multiples, repeat n times.

Correct, O(n log n), and it needs a **set** alongside it — because 6 arrives as both 2×3 and 3×2, and without deduplication it is emitted twice.

That extra set is the tell. It means the structure is generating work it then has to throw away.

## The three-pointer version

Keep every ugly number in order, and one pointer per multiplier. Each pointer marks the earliest value that multiplier has not yet been applied to.

The next ugly number is the smallest of `2×dp[p2]`, `3×dp[p3]`, `5×dp[p5]`. Advance whichever pointer produced it.

**Advance every pointer that produced it.** When 6 comes up as both 2×3 and 3×2, both `p2` and `p3` move. That is the deduplication — not a set, just the observation that a value produced twice must consume both producers.

## Why this is better than the heap

No heap and no set. Three integers and an array, three multiplications and two comparisons per step, and every candidate is generated exactly once.

**O(n) against the heap's O(n log n)**, with a smaller constant.

This closes the heap arc on a problem where the heap is the obvious answer and the right answer is something else. That is the same shape as day 41, where a two-pass DP replaced a BFS — and the same shape as day 79, where sorting beat the heap.

## Complexity

- **Time: O(n).** Constant work per number.
- **Space: O(n)** for the sequence, which is required to compute later terms.
