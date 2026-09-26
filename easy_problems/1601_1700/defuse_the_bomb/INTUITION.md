## Intuition

Each output value is the sum of a window of `|k|` numbers next to position `i`: the `k`
after it when `k > 0`, the `|k|` before it when `k < 0`. Moving from `i` to `i + 1`
shifts that window by one, so it's a fixed sliding window again: subtract the number that
leaves, add the one that enters.

What's new is the circle. The window after the last element continues at the front. Index
arithmetic modulo `n` handles that without copying the array: `code[(i+k)%n]` is the
number `k` places after `i`, wrapping as needed. For the backwards case the code writes
`((i-k)+n)%n`, adding `n` before the modulo so the index never goes negative (Go's `%`
keeps the sign of the left operand, so `-1 % 4` is `-1`, not 3).

## Builds on

- [Day 92: Maximum Average Subarray I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/maximum_average_subarray_i/) — one running sum over a fixed-length window, adjusted by one subtraction and one addition per step

The window here excludes `i` itself, which changes the order of operations. Before the
update, `sum` holds `code[i..i+k-1]`. Subtracting `code[i]` and adding `code[i+k]` gives
`code[i+1..i+k]`, exactly "the next `k`". So each iteration updates first and records
second, where Day 92 recorded between the add and the subtract.

For `k < 0` the code runs the same idea mirrored: prime with the last `|k|` elements and
walk `i` from the end towards the front.

`k == 0` needs no branch: `ans` comes from `make`, already all zeros.

"All numbers are replaced simultaneously" is why the result goes into a new slice.
Writing into `code` in place would corrupt the sums for positions that still need the
original values.

**Complexity:**
- Time: O(n + |k|), which is O(n) since `|k| < n`.
- Space: O(n) for the output.
