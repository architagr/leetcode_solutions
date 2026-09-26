## Intuition

This one isn't about a window over the input as given. The students can be picked from
anywhere in the array, so position means nothing, and there are C(n, k) possible groups.
The window only appears after sorting.

Sort the scores. Now take any group of `k` scores, and look at its lowest and highest.
Between them in sorted order there are at least `k` scores (the group itself fits there).
So the `k` consecutive sorted scores starting at the group's lowest one have a highest
score no bigger than the group's highest. A run of neighbours is always at least as good as
any scattered choice. The only groups worth checking are the `n - k + 1` windows of
length `k` in the sorted array.

And within a sorted window, the lowest score is the first element and the highest is the
last. The window's difference is `nums[i] - nums[i-k]` after the usual `k--`, two array
reads, with no running state to maintain at all. It's the lightest sliding window in this
run.

## Builds on

- [Day 92: Maximum Average Subarray I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/maximum_average_subarray_i/) — the fixed window of length k and the k-- offset that makes nums[i-k] its first element

`ans` starts at the full range, `nums[n-1] - nums[0]`, which is an upper bound for every
window. When `k = 1` the offset becomes 0 and every window's difference is 0, so Example 1
works without a special case.

One side effect: `sort.Ints` sorts the caller's slice in place. LeetCode doesn't care. A
function in a real codebase called with someone's list of scores probably should copy
first.

**Complexity:**
- Time: O(n log n) for the sort; the window pass is O(n).
- Space: O(1) extra beyond the sort's own stack usage.
