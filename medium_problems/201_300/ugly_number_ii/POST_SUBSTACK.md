---
meta_title: "Closing the heap arc on a problem the heap loses"
meta_description: "The heap solution needs a set alongside it, because 6 arrives twice. Three pointers dedupe for free and turn O(n log n) into O(n)."
tags: [golang, dynamic-programming, heap, dsa]
---

![Day 81](HERO.png)

*365 Days of LeetCode Challenge — Day 81/365*

**[264. Ugly Number II](https://leetcode.com/problems/ugly-number-ii/)** (Medium)

An ugly number is a positive integer whose only prime factors are 2, 3 and 5. Return the nth ugly number.

This is the last day of the heap arc, and I chose it for that slot because the heap loses.

## Why the direct approach does not work

Start at 1 and walk upward. For each integer, divide out all the 2s, then all the 3s, then all the 5s, and see whether 1 is left.

That is correct. It is also hopeless, and the reason is worth quantifying rather than hand-waving: the 1690th ugly number is **2,123,366,400**. The problem's upper bound on n is 1690, so this approach asks you to test over two billion integers to find 1690 of them.

The density collapses as the numbers grow. Almost everything you test is a miss, and the misses are the entire cost.

So the move is to stop **testing** candidates and start **generating** them.

## Reading the definition forwards

The definition says: an ugly number's prime factors are limited to 2, 3 and 5.

Turn it around and it becomes a construction rule: the only way to produce an ugly number is to take a smaller ugly number and multiply it by 2, 3 or 5.

Nothing else produces one. Which means the sequence can be built entirely out of itself — given all the ugly numbers found so far, the next one is the smallest unused product of one of them with 2, 3 or 5.

![Step 1](images/walkthrough-1.png)

## The heap solution, and the thing it needs

"Repeatedly take the smallest from a changing collection" is precisely what this arc has spent a week on. The heap solution writes itself:

Push 1. Pop the smallest — that is the next ugly number. Push its three multiples. Repeat n times.

It is correct, and it is O(n log n).

It also needs a **set**, and that is the interesting part.

6 enters the heap twice: once as 2×3 when 3 is popped, and once as 3×2 when 2 is popped. Without a set to remember what has already been seen, 6 is emitted twice and every position after it is off by one.

So the working solution is a heap *plus* a set of every value ever pushed. And that extra structure is a tell. It means the algorithm is generating candidates it will have to identify and discard — doing work in order to undo it.

## One pointer per multiplier

```go
count_2, count_3, count_5 := 1, 1, 1
```

The alternative keeps the whole sequence in an array and maintains three pointers, one per multiplier. Each pointer marks the earliest position its multiplier has not yet been applied to.

```go
val_2 = 2 * dp[count_2]
val_3 = 3 * dp[count_3]
val_5 = 5 * dp[count_5]
min_val = min(val_2, min(val_3, val_5))
```

![Step 2](images/walkthrough-2.png)

Three candidates, and the smallest of them is the next ugly number.

It is worth being sure about why no fourth candidate is needed. Any other product of an earlier ugly number with 2, 3 or 5 is either already in the array, or it uses a position further along than one of the three pointers — and since the array is increasing, such a product is larger than the candidate that pointer currently offers. So the true next value is always one of these three.

![Step 3](images/walkthrough-3.png)

## The three ifs, which are not an else-if

```go
if min_val == val_2 {
	count_2++
}
if min_val == val_3 {
	count_3++
}
if min_val == val_5 {
	count_5++
}
```

Here is the whole solution, and the single most important thing about it is that these are three independent `if` statements and not an `if / else if` chain.

![Step 4](images/walkthrough-4.png)

Watch what happens when the sequence reaches 6.

At that point `p2` is pointing at 3 and `p3` is pointing at 2. So `val_2` is 2×3 = 6, and `val_3` is 3×2 = 6. Both candidates are 6. Both are correct. The value 6 is genuinely produced by two different multipliers.

Both pointers have to advance.

If only one advanced — which is exactly what `else if` would cause — then on the very next iteration the other pointer would offer 6 again, 6 would be appended a second time, and the whole sequence from there on would be wrong.

**That is the deduplication.** There is no set, no lookup, no membership test. Just the recognition that a value produced by two multipliers has consumed both of them, so both must move on.

This is the part of the solution I find genuinely satisfying. The heap version detects duplicates after creating them. This version makes them impossible to create.

![Step 5](images/walkthrough-5.png)

## A small thing about indexing

```go
dp := make([]int, n+1)
dp[1] = 1
```

The array is 1-indexed: `dp[0]` is allocated and never read.

It costs one integer and it means `dp[i]` is the ith ugly number rather than the (i+1)th, which removes every off-by-one from the loop bounds, the pointers and the return statement. The pointers start at 1 for the same reason.

## Why this beats the heap

No heap. No set. Three integers, an array, and per step: three multiplications, two comparisons, and up to three increments.

Every candidate is generated exactly once, because a pointer only ever moves forward and never revisits a position. The heap version generates roughly 3n candidates and throws away the duplicates.

**O(n) against O(n log n)**, and with a much smaller constant factor on top of the better complexity.

## Closing the arc on a problem the heap does not win

This is deliberate as the final day.

Days 76 through 78 built the case for the heap, and it is a strong case — Last Stone Weight, Kth Largest, Top K Frequent are all problems where the heap is exactly the right tool and the alternatives are worse.

Day 79 found a problem where sorting reads better, because one added ordering requirement made the heap's comparator awkward while leaving sorting's obvious.

Today is a problem where the heap is the *obvious* answer — "repeatedly take the smallest" could not signal it more clearly — and a plain array with three counters is strictly better on every axis.

Day 41 did the same thing to BFS, replacing it with a two-pass DP on a problem that looked exactly like the previous day's graph traversal.

Recognising the structure a problem suggests is most of the skill. Knowing when to put it down is the rest, and it is the harder half, because it requires noticing that the obvious tool is generating work it then has to undo.

## Complexity

- **Time: O(n).** Constant work per number generated.
- **Space: O(n)** for the sequence. This one is unavoidable — later terms are computed from earlier ones, so they have to be kept.

Full code and the step-by-step walkthrough:
[ugly_number_ii](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/ugly_number_ii/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
