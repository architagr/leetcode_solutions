## Intuition

Strip the story away: two baskets, each holding one fruit type, picking from consecutive
trees until one doesn't fit. That's "the longest contiguous stretch containing at most two
distinct values".

Both pieces are already in the series. The window that grows right and shrinks left to
keep a stretch valid, from Day 102. And the count map that deletes a key when its count
reaches zero, so that `len(map)` is the number of distinct values in the window, from
Day 93. Put them together: grow the right edge, and while the map has more than two keys,
drop trees from the left.

## Builds on

- [Day 102: Maximum Length Substring With Two Occurrences](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/3001_3100/maximum_length_substring_with_two_occurrences/) — the grow-right, shrink-left window for a longest valid stretch
- [Day 93: Substrings of Size Three with Distinct Characters](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1801_1900/substrings_of_size_three_with_distinct_characters/) — a count map that deletes keys at zero, so its length is the number of distinct values

The delete is doing real work here. Without it, a fruit type that left the window would
still be a key with count 0, `len(fruitBaskets)` would stay at 3, and the shrink loop
would keep removing trees long after the window was valid again.

Recording the length after every step (rather than only when the window breaks, like some
earlier days) is simpler and just as correct: after the shrink loop, the window is the
longest valid one ending at `i`.

Swap 2 for `k` and this is the general "longest subarray with at most k distinct values",
a pattern that turns up a lot.

**Complexity:**
- Time: O(n). Each tree enters and leaves once.
- Space: O(1). The map never holds more than three keys.
