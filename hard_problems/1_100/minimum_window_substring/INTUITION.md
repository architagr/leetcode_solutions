## Intuition

This is the hard problem the whole arc has been building toward, and by now every piece of
it has appeared before.

"Contains every character of `t`, with duplicates" is an at-least condition, like Day 118:
once a window covers `t`, making it bigger keeps it covered. And the thing to minimise is
length, so the shape is the one from Minimum Size Subarray Sum: grow the right edge until
the window is valid, then shrink from the left while it stays valid, recording the shortest
valid window seen.

What's new is the bookkeeping for "covers `t`". Day 114 compared a window's counts to a
target's. This solution keeps a single map of what's still *needed*: it starts as the counts
of `t`, a character entering the window decrements its entry, and one leaving increments it.
The window covers `t` exactly when every entry is at most 0. Negative values are fine; they
mean the window has more of that character than it needs, so it can afford to lose some.

## Builds on

- [Day 114: Permutation in String](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/501_600/permutation_in_string/) — comparing a window's letter counts against a target string's counts
- [Day 118: Number of Substrings Containing All Three Characters](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1301_1400/number_of_substrings_containing_all_three_characters/) — an 'at least' window: once it contains everything, shrink from the left as far as it stays valid

Characters that aren't in `t` never touch the map. The code skips them on the right with
`continue`, and on the left the `continue` inside the shrink loop still runs `l++`
(it's the loop's post statement), so they're simply stepped over.

`found` checks every entry of the map, at most 52 letters, each time the shrink loop tests
its condition. That makes the whole thing O(52·n): linear, with a constant. The standard
refinement keeps one integer, the number of characters still missing, and updates it only
when an entry crosses 0 in either direction; the coverage check becomes a comparison with
zero. Same big-O, smaller constant.

**Complexity:**
- Time: O(|t| + 52·|s|), linear in the input.
- Space: O(1): at most 52 entries in the map.
