## Intuition

Up to now the variable window wanted to be as long as possible. This one wants to be as
short as possible, which flips when each edge moves: grow the right edge while the sum is
too small, and once it's big enough, record the length and shrink from the left to see if
a shorter window still qualifies.

That only works because every number is positive. Adding an element always increases the
sum and removing one always decreases it, so "too small" and "big enough" change in a
predictable direction as the edges move. With negative numbers, shrinking could make the
sum go up, and the whole approach falls apart.

## Builds on

- [Day 102: Maximum Length Substring With Two Occurrences](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/3001_3100/maximum_length_substring_with_two_occurrences/) — the variable window whose two edges only move forward; here it looks for the shortest valid window instead of the longest

This solution measures window sums through a prefix-sum array rather than a running total.
`prefixSum[i]` is the sum of `nums[0..i]`, so the window `(left, right]` sums to
`prefixSum[right] - prefixSum[left]`, with `getValue(-1)` returning 0 for the empty prefix.
`left` is exclusive and starts at `-1`, so the window's length is `right - left`, the same
convention as Day 103.

Each step moves exactly one edge forward, so the loop runs at most `2n` times. The early
return when the total is below `target` handles Example 3 and guarantees that `result`
gets set at least once.

A running sum does the same job in O(1) extra space:

```go
best, sum, left := len(nums)+1, 0, 0
for right, v := range nums {
	sum += v
	for sum >= target {
		best = min(best, right-left+1)
		sum -= nums[left]
		left++
	}
}
```

I checked it against the repo's version on a few thousand random arrays.

The prefix sums do earn their place for the follow-up, though. Because every number is
positive, `prefixSum` is strictly increasing, so for each `right` you could binary search
for the furthest `left` that still leaves a big enough sum. That's the O(n log n) version
the problem asks for.

**Complexity (as written):**
- Time: O(n). Each edge moves forward at most n times.
- Space: O(n) for the prefix sums.
