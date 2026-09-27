## Intuition

Two strings and a budget sound like a string problem. They aren't, really. The only thing
that matters about position `i` is how much it costs to change: `|s[i] - t[i]|`. Replace
the two strings with that one row of numbers, and the question is the longest stretch of
the row whose sum fits in `maxCost`.

That's a variable window over non-negative numbers, the same family as Minimum Size
Subarray Sum, pointed the other way. Grow the right edge, adding each position's cost.
While the sum is over budget, drop costs from the left. After each step the window is the
longest affordable stretch ending at `r`; record its length.

## Builds on

- [Day 105: Minimum Size Subarray Sum](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/minimum_size_subarray_sum/) — a window over non-negative numbers, where adding raises the sum and removing lowers it; that one wanted the shortest window, this wants the longest

Why "non-negative" matters: a cost is an absolute difference, never below zero. So adding
a position never lowers the sum and removing one never raises it, which is what lets the
left edge move only forward. Some costs can be 0 (matching letters), and they're free to
include.

The code never builds the cost array; `absDiff(s[r], t[r])` computes each cost when it's
needed, and again for `s[l]`, `t[l]` when it leaves. If a single position costs more than
the whole budget, the shrink loop removes it too and the window is briefly empty, length 0.

**Complexity:**
- Time: O(n).
- Space: O(1).
