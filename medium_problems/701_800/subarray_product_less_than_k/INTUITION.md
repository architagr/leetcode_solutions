## Intuition

Day 103 counted substrings with a window, using one fact: if a stretch is valid, every
piece of it is too. Then, for each right edge, the number of valid subarrays ending there
equals the length of the longest valid window ending there.

The same fact holds for products of positive integers. Every element is at least 1, so
dropping elements can only lower the product (or keep it). If a stretch has product below
`k`, so does every piece. So: grow the right edge multiplying in each value, divide
values out from the left while the product is `k` or more, and add the window's length.

## Builds on

- [Day 103: Count Substrings That Satisfy K-Constraint I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/3201_3300/count_substrings_that_satisfy_k_constraint_i/) — counting every valid subarray by adding the window's length at each right edge, because pieces of a valid window are valid

The product is kept by multiplication and exact division, which is fine because every
value divided out was multiplied in earlier. It never gets large: before each
multiplication it's below `k` (at most `10^6`), and the new value is at most 1000, so it
stays under `10^9`. The code still uses `int64` for it.

`k <= 1` is the edge case Example 2 hints at. No product of positive integers is below 1,
so the shrink loop empties the window every time (`left` passes `right`), the length is 0,
and the count stays 0 without any special branch.

**Complexity:**
- Time: O(n).
- Space: O(1).
