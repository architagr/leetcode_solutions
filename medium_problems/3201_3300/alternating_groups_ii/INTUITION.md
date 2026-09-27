## Intuition

A group is `k` consecutive tiles, around a circle, where every neighbouring pair differs in
colour. Two ideas handle the two awkward parts.

**The circle.** A group can start near the end and wrap to the front, but it's never longer
than `k`, so it never reaches more than `k - 1` tiles past the end. Appending the first
`k - 1` tiles to the array turns every circular group into an ordinary window, and no group
is counted twice because only starting positions `0..n-1` can fit a full window.

**Alternation.** Instead of checking each group from scratch, keep a window of tiles that
alternate. Extend it one tile at a time. If the new tile repeats the previous colour, no
alternating stretch can contain both, so the window restarts at the new tile. Whenever the
window reaches `k` tiles, one group has been found; count it and move `left` forward by
one so the window stays at `k` and the next start position gets checked on the next tile.

## Builds on

- [Day 115: Minimum Swaps to Group All 1's Together II](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/2101_2200/minimum_swaps_to_group_all_1s_together_ii/) — a fixed-size window over a circular array; there with % n, here by appending the first k-1 elements

Where yesterday wrapped with `% n`, this code builds the extended slice directly with
`append(colors, colors[:k-1]...)`. One subtlety of `append` in Go: if `colors` had spare
capacity, it would write the copied tiles into the caller's backing array. A slice built
from a literal or read from input has no spare capacity, so a new array is allocated here,
but it's the kind of line worth a second look in production code.

**Complexity:**
- Time: O(n + k). Each tile of the extended array is visited once.
- Space: O(n + k) for the extended copy.
