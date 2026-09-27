## Intuition

Yesterday: the longest run of ones if you may flip up to `k` zeros. Today: the longest run
of *either* letter if you may change up to `k` answers. The only difference is that the
run can be made of Ts or of Fs, and you don't know which in advance.

So don't choose. Solve it twice. The longest all-T run you can make is the longest window
with at most `k` Fs in it. The longest all-F run is the longest window with at most `k`
Ts. Take the bigger one.

That's exactly what this solution does: `countData(answerKey, k, c)` is Day 106's window
with `c` as the letter being counted, and `maxConsecutiveAnswers` calls it for `'T'` and
for `'F'` and returns the max.

## Builds on

- [Day 106: Max Consecutive Ones III](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1001_1100/max_consecutive_ones_iii/) — the longest window with at most k of one kind of element, which this runs once per letter

Inside `countData`, the window records its length when it breaks (the count of `c` goes
past `k`) and once after the loop, the same lazy recording as the last few days. It starts
by counting `answerKey[0]` separately and running `end` from 1, which is just a different
way of writing the first step.

A single-pass version also exists: track both counts in one window, and shrink whenever the
*smaller* of the two exceeds `k` (the smaller count is the one you'd change). It's a nice
trick, but two calls to a function you already trust is easier to get right, and it's
still O(n).

**Complexity:**
- Time: O(n). Two linear passes.
- Space: O(1).
