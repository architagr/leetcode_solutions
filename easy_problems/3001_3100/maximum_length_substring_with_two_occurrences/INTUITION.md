## Intuition

Every window so far has had a fixed length. This one doesn't, and that's the second half
of the sliding-window idea: a window that grows while it's valid and shrinks only when it
has to.

Here "valid" means no letter appears more than twice. Push the right edge forward one
letter at a time. The window stays valid until some letter reaches a third copy. At that
moment the window is too big, and the only way to fix it is to move the left edge past
the *first* copy of that letter. Anything to the left of that first copy has to go with
it, because a substring can't skip letters.

Two things make this efficient. The right edge only moves forward, and so does the left
edge. Each index enters the window once and leaves at most once, so the whole thing is
O(n) even though the code has a loop inside a loop.

And the longest valid window can be recorded lazily. Between violations the window only
grows, so its size just *before* a violation is the biggest it got. The code records
`right - left` at each violation (before shrinking) and once more after the loop for the
final window. It never needs to check on every step.

## Builds on

- [Day 93: Substrings of Size Three with Distinct Characters](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1801_1900/substrings_of_size_three_with_distinct_characters/) — a window that carries a count per letter, updated as letters enter and leave

The inner loop decrements the count of each letter it drops, including the matching copy
of the offending letter, which brings that letter back from 3 to 2. Then `left++` and
`break` step past it.

**Complexity:**
- Time: O(n). Each index is added once and removed at most once.
- Space: O(1). At most 26 letters in the map.
