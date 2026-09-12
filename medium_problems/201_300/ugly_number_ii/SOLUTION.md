# Ugly Number II — solution walkthrough

From `ugly_number_II.go`:

```go
func NthUglyNumber(n int) int {
	dp := make([]int, n+1)
	count_2, count_3, count_5 := 1, 1, 1
	min_val, val_2, val_3, val_5 := 0, 0, 0, 0

	dp[1] = 1 //ist ugly number is 1
	for i := 2; i <= n; i++ {
		val_2 = 2 * dp[count_2]
		val_3 = 3 * dp[count_3]
		val_5 = 5 * dp[count_5]
		min_val = min(val_2, min(val_3, val_5))
		if min_val == val_2 {
			count_2++
		}
		if min_val == val_3 {
			count_3++
		}
		if min_val == val_5 {
			count_5++
		}
		dp[i] = min_val
	}
	return dp[n]
}
```

## Why not test each integer in turn

Walk upward from 1. For each candidate, divide out all the 2s, 3s and 5s, and check whether 1 is left.

Correct, and hopeless. The 1690th ugly number is 2,123,366,400. Almost every integer tested is a miss, and the misses dominate completely.

The move is to stop **testing** candidates and start **generating** them.

## Reading the definition forwards

An ugly number's only prime factors are 2, 3 and 5. Read the other way: the only way to produce one is to multiply a smaller ugly number by 2, 3 or 5.

So the sequence can be built out of itself. Given every ugly number found so far, the next is the smallest unused product of one of them with 2, 3 or 5.

![Step 1](images/walkthrough-1.png)

## The heap solution this is not

The natural structure for "repeatedly take the smallest" is a heap — this arc has spent a week on exactly that.

Push 1. Pop the smallest, push its three multiples, repeat n times. Correct, and O(n log n).

It also needs a **set**. 6 arrives as 2×3 and again as 3×2, so without deduplication it is emitted twice and every count after it is wrong.

That set is the tell. It means the structure is generating work it then discards.

## One pointer per multiplier

```go
count_2, count_3, count_5 := 1, 1, 1
```

Each pointer marks the earliest position in `dp` that its multiplier has not yet been applied to.

```go
val_2 = 2 * dp[count_2]
val_3 = 3 * dp[count_3]
val_5 = 5 * dp[count_5]
min_val = min(val_2, min(val_3, val_5))
```

![Step 2](images/walkthrough-2.png)

Three candidates, one per multiplier, and the smallest is the next ugly number. Nothing else can be: any other product is either already in `dp` or larger than one of these three.

![Step 3](images/walkthrough-3.png)

## The three `if`s are not an `else if`

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

This is the whole solution, and the reason it is three separate `if`s rather than a chain matters.

![Step 4](images/walkthrough-4.png)

When `i` reaches 6, `val_2` is 2×3 and `val_3` is 3×2. Both equal 6. Both pointers must advance.

Advance only one — which is what `else if` would do — and the other produces 6 again on the next iteration, and 6 appears twice in the sequence.

**That is the deduplication.** No set, no lookup: just the observation that a value produced by two multipliers has consumed both of them.

![Step 5](images/walkthrough-5.png)

## The array is 1-indexed

```go
dp := make([]int, n+1)
dp[1] = 1
```

`dp[0]` is allocated and never used, so `dp[i]` is the ith ugly number rather than the (i+1)th. It costs one integer and removes every off-by-one from the loop and the return.

The pointers start at 1 for the same reason — position 1 is the first real entry.

## Why this beats the heap

No heap and no set. Three integers, three multiplications and two comparisons per step.

Every candidate is generated exactly once, because a pointer never revisits a position. The heap generates 3n candidates and discards the duplicates; this generates exactly what it needs.

**O(n) against O(n log n)**, with a smaller constant on top.

## Closing the arc on a problem the heap does not win

This is deliberate as the last day of the heap arc.

Day 76 to 78 built the case for the heap. Day 79 found a problem where sorting reads better. Today is a problem where the heap is the obvious answer and a plain array with three counters is strictly better.

Day 41 did the same thing to BFS, replacing it with a two-pass DP. Recognising the structure a problem suggests is most of the skill; knowing when to put it down is the rest.

## Complexity

- **Time: O(n).** Constant work per number: three multiplications, two comparisons, up to three increments.
- **Space: O(n)** for the sequence. Unavoidable — later terms are computed from earlier ones.

## Test

`main_test.go` covers `n = 10` giving 12 and `n = 1` giving 1. The interesting case to add is any `n` past 6, since 6 is the first value produced by two different multipliers and the first that a single `else if` would duplicate.
