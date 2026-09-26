## Intuition

The problem is phrased as an operation count, but it's a counting question underneath. To
make a particular stretch of `k` blocks all black, you repaint exactly the white blocks
in it: no more, no fewer. So the cost of a window is its number of whites, and the answer
is the smallest number of whites in any window of length `k`.

That's Day 92's problem with a different thing being summed. Instead of adding values,
add 1 for each `W`. Minimise instead of maximise.

## Builds on

- [Day 92: Maximum Average Subarray I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/maximum_average_subarray_i/) — the same prime-one-short, add, check, remove loop, with a count of whites in place of the sum

The code is almost line for line the same shape: `k--` to turn `k` into the offset of
the oldest block, prime the first `k` blocks, then add the new block, record the minimum,
and remove the oldest.

`ans` starts at `k`, which is the real worst case (a window of all whites costs `k`
repaints) rather than a sentinel like `math.MaxInt`. Every window's count is at most `k`,
so the first check always sets it properly anyway.

A possible early exit: once `ans` hits 0 nothing can beat it, and the loop could stop.
With `n <= 100` it doesn't matter.

**Complexity:**
- Time: O(n).
- Space: O(1).
