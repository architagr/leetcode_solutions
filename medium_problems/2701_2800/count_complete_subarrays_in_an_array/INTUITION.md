## Intuition

Yesterday: substrings containing all of `a`, `b` and `c`. Today: subarrays containing
every distinct value that appears anywhere in the array. It's the same problem once you know
how many distinct values there are, so the solution starts by counting them (`k`) in one
pass over the array.

A subarray can never have *more* than `k` distinct values, so "equal to `k`" means "at
least `k`", and that's the kind of constraint that survives growing: add elements to a
complete subarray and it stays complete. So count by start, as yesterday. Grow the right
edge until the window has `k` distinct values; then every subarray starting at `start` and
ending at `end` or later is complete, which is `len(nums) - end` of them. Drop
`nums[start]`, check again, repeat.

## Builds on

- [Day 118: Number of Substrings Containing All Three Characters](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1301_1400/number_of_substrings_containing_all_three_characters/) — counting 'at least' windows by start: once valid, all len - end extensions count

Distinct values are tracked with a count map that deletes keys at zero, so `len(fre)` is
the number of distinct values in the window. The code does the decrement and delete as one
branch: if the count is 1, delete; otherwise decrement.

Example 2 is a nice sanity check: with only one distinct value, every subarray is complete,
and the count comes out as `4 + 3 + 2 + 1 = 10`.

**Complexity:**
- Time: O(n). One pass to count distinct values, one pass with the window.
- Space: O(n) for the maps.
