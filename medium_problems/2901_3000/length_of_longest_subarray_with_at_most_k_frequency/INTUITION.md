## Intuition

Day 102 asked for the longest substring where no letter appears more than twice. This asks
for the longest subarray where no number appears more than `k` times. It's the same
problem with the 2 turned into a parameter and letters turned into integers (up to `10^9`,
so a map rather than a 26-slot array).

Grow the right edge and count. Only the value that just arrived can have gone over `k`,
so that's the only count worth checking. If it's fine, record the window. If not, drop
elements from the left until that one value is back to `k`.

## Builds on

- [Day 102: Maximum Length Substring With Two Occurrences](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/3001_3100/maximum_length_substring_with_two_occurrences/) — this exact problem with k fixed at 2 and letters instead of numbers

This version records only in the "still valid" branch and `continue`s past the shrink.
It looks like it might miss a window, but it can't. Before the new element arrived, the
window `[start, i-1]` was valid and already recorded. The shrink moves `start` forward by
at least one, so the window after it, `[start', i]`, is no longer than the one just
recorded. Nothing is lost.

Unlike Fruit Into Baskets, this map never deletes zero counts. It doesn't need to: nothing
here asks how many keys there are, only what one key's count is.

**Complexity:**
- Time: O(n).
- Space: O(n) for the counts, since values are arbitrary integers.
