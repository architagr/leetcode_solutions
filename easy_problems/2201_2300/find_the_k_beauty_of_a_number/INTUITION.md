## Intuition

The obvious version slices each length-`k` substring and parses it with `strconv.Atoi`.
With at most ten digits that's perfectly fast. This solution does something I enjoyed more:
it keeps the window as a number and updates it arithmetically, the way Day 92 updated a
sum.

Adding a digit on the right of a number is `x*10 + digit`. Dropping the leftmost digit of
a `k`-digit number is `x % 10^(k-1)`: the remainder after dividing by the place value of
that leading digit. So the window slides with one multiply-add and one modulo per step,
and never parses a string after the first one.

## Builds on

- [Day 92: Maximum Average Subarray I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/maximum_average_subarray_i/) — add the new element, drop the oldest; here adding is x*10 + digit and dropping is x % 10^(k-1)

The loop uses the same "prime one short" shape. `k--` first, then `d` becomes
`10^(k-1)` and `x` holds the first `k-1` digits. Each step appends one digit (the window
is now `k` digits), tests divisibility, and applies `x %= d` to drop the oldest digit,
leaving `k-1` digits again.

Two cases from the problem's notes are handled by the arithmetic or by one guard:

- **Leading zeros are allowed.** In `"430043"`, the window `"04"` is the number 4. The
  arithmetic produces exactly that: 0 then `0*10 + 4`.
- **0 is not a divisor.** The window `"00"` is the number 0, and `num % 0` would panic.
  The check is `x > 0 && num%x == 0`, and `&&` stops before the modulo.

For `k = 1`, `numsStr[:0]` is `""`, `strconv.Atoi("")` returns an error that the code
ignores, and `x` stays 0, which happens to be the right starting value. It works, but it
works by accident; I'd rather it didn't lean on an ignored error.

**Complexity:**
- Time: O(d), where d is the number of digits (at most 10).
- Space: O(d) for the string form of `num`.
