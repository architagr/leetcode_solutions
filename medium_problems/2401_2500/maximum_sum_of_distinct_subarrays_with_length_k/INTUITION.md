## Intuition

Two fixed-window problems from the start of the arc, combined. Day 92 slid a running sum
over windows of length `k`. Day 93 slid a count map over windows of length 3 and called a
window good when the map had one key per element. Here the window is length `k`, the
score is the sum, and a window only counts if its elements are distinct.

So the window carries both: a running sum and a count map that deletes keys at zero. A
window of `k` elements is all-distinct exactly when the map has `k` keys. When it does,
compare its sum with the best so far.

## Builds on

- [Day 93: Substrings of Size Three with Distinct Characters](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1801_1900/substrings_of_size_three_with_distinct_characters/) — a fixed window whose map deletes keys at zero, so all-distinct is len(map) == window size
- [Day 92: Maximum Average Subarray I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/maximum_average_subarray_i/) — the running sum over a fixed window, with the k-- offset

The code reuses the `k--` offset, so the distinctness check reads `len(m) == k+1`. The
sum is kept as `int64` to match the function's return type; with up to `10^5` values of up
to `10^5`, a window's sum can reach `10^10`, which needs 64 bits anyway.

`result` starts at 0, which is also the required answer when no window qualifies
(Example 2).

**Complexity:**
- Time: O(n). A constant number of map operations per step.
- Space: O(k) for the map.
