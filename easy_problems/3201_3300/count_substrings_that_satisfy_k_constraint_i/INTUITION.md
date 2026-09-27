## Intuition

Yesterday's window found the *longest* valid substring. This one has to *count* all of
them, and the trick that makes that possible is worth more than the problem itself.

The constraint has a useful property: if a substring satisfies it, so does every piece of
it. Removing characters can only lower the count of 0s and the count of 1s, so whichever
one was at most `k` stays at most `k`.

So fix the right end at `right` and let the window be the longest valid substring ending
there. Every shorter substring ending at `right` sits inside that window, so every one of
them is valid too. The number of valid substrings ending at `right` is exactly the
window's length. Add that up for every `right` and you've counted every valid substring
once, by its end position.

## Builds on

- [Day 102: Maximum Length Substring With Two Occurrences](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/3001_3100/maximum_length_substring_with_two_occurrences/) — the grow-right, shrink-left window that stays valid; here it counts substrings instead of measuring the longest

The window is kept the same way as yesterday: push `right`, and while both counts exceed
`k` (the constraint needs just one of them to be within `k`, so it breaks only when
both aren't), drop characters from the left.

This code makes `left` exclusive: it starts at `-1` and the window is `s[left+1..right]`,
so its length is simply `right - left`. The shrink loop does `left++` *before* removing
`s[left]`, which is the matching half of that convention.

**Complexity:**
- Time: O(n). Both edges only move forward.
- Space: O(1).
