## Intuition

The problem describes a sequence of choices (front or back, `k` times), which suggests a
search or a greedy rule. Neither is needed, because the order doesn't matter. Whatever
sequence you follow, you end up holding some number `a` of cards from the front and
`k - a` from the back. The score depends only on `a`, and `a` can be anything from 0 to
`k`. That's `k + 1` possibilities.

Greedy (take whichever end is bigger) fails, and it's worth seeing how. On
`[1,1000,1,1,100]` with `k = 2`, greedy takes the 100, then a 1: 101. Taking both front
cards gives 1001. Greedy can't see the 1000 hiding behind a 1.

Now picture the array as a circle. The last `k - a` cards followed by the first `a` cards
are one contiguous stretch around the join. Every valid pick is a window of length `k`
that straddles the end of the array, starting with all `k` from the back and sliding
forward one card at a time until it's all `k` from the front. That's yesterday's circular
window, and each slide swaps one back card for one front card.

## Builds on

- [Day 99: Defuse the Bomb](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1601_1700/defuse_the_bomb/) — a fixed window that wraps from the end of the array to the start, slid with one subtraction and one addition

There's a mirror-image way to see it that some people find cleaner: the cards you *don't*
take are always a contiguous block of `n - k` in the middle. Maximising what you take is
minimising what you leave, so the answer is the total minus the smallest window sum of
length `n - k`. Same complexity; this code slides the taken window instead, which only
touches `2k` cards and never sums the middle.

In the code, `(start+i)%n` never actually wraps: `start + i` is at most `n - 1`. The
modulo is harmless and makes the circular reading explicit.

**Complexity:**
- Time: O(k). k cards to prime, then k swaps.
- Space: O(1).
