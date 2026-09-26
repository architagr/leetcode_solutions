## Intuition

Four days ago the window counted white blocks and kept the minimum. Here it counts vowels
and keeps the maximum. Structurally it's the same program, with `isVowel(b)` in place of
`b == 'W'`. The medium label is mostly about input size: `s` can be `10^5` long, so
recounting each window (O(n·k)) can time out, and the sliding count (O(n)) is the
expected answer.

## Builds on

- [Day 94: Minimum Recolors to Get K Consecutive Black Blocks](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/2301_2400/minimum_recolors_to_get_k_consecutive_black_blocks/) — the same fixed window counting characters that match a test; there it was 'W', here it's a vowel

That's the value of having done the easier versions first. Recognising "fixed length,
count of things matching a test, best over all windows" as one shape means this problem
takes minutes rather than a fresh derivation.

`isVowel` is a chain of five comparisons, which is plenty fast. A `[26]bool` lookup
table would also work; for five letters the comparisons are arguably easier to read.

One optional improvement: the maximum possible count in a window is `k` itself. Once
`max` reaches it, no later window can do better and the loop could return early. On
Example 1 that happens at `"iii"`, three windows before the end.

**Complexity:**
- Time: O(n).
- Space: O(1).
