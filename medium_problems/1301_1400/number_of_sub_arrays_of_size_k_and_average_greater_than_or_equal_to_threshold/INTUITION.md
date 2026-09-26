## Intuition

This is labelled medium, and it's the same loop as Day 92 with a count in place of a max.
The only new piece is one line at the top, and it's a line worth having in the toolbox:

```go
threshold *= k
```

"The average of this window is at least `threshold`" means `sum / k >= threshold`. Since
`k` is positive, multiply both sides by it: `sum >= threshold * k`. Do that
multiplication once, before the loop, and every window check becomes an integer
comparison. No division, no floats, nothing to round.

The problem's note that "averages are not integers" is there to catch float
comparisons that go wrong, and multiplying up sidesteps the issue completely. (Integer
division `sum/k >= threshold` would also happen to be right here, because the threshold is
an integer, but that takes a moment's thought to justify. The multiplied version needs
none.)

## Builds on

- [Day 92: Maximum Average Subarray I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/maximum_average_subarray_i/) — the running-sum window and the idea of comparing sums instead of averages when every window has the same length

The largest product is `10^4 * 10^5 = 10^9`, and the largest window sum is the same, so
Go's 64-bit `int` has no overflow risk here.

After that one line it's the familiar shape: `k--`, prime `k` elements, then add the new
element, count the window if `sum >= threshold`, subtract the oldest.

**Complexity:**
- Time: O(n).
- Space: O(1).
