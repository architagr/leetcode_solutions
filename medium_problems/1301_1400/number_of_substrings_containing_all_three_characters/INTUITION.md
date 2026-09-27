## Intuition

Yesterday's constraint ("product below `k`") survived *shrinking*: every piece of a valid
window was valid. This one ("contains a, b and c") survives *growing*: if a substring has
all three letters, adding more letters on either side can't take any away.

That flips how the counting works. Fix a start `left`, and find the first `right` where
`s[left..right]` contains all three. Every longer substring from the same start is valid
too, so that start contributes `len(s) - right` substrings in one go.

The window finds those first-valid positions efficiently. Grow the right edge. As soon as
the window holds all three letters, it's the shortest valid window for the current `left`,
so add `len(s) - right`, drop `s[left]`, and check again: the next start might be valid
at the same `right` too. Keep going until the window is missing a letter, then grow again.

## Builds on

- [Day 117: Subarray Product Less Than K](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/701_800/subarray_product_less_than_k/) — counting all valid subarrays in one pass; there pieces of a valid window stay valid, here bigger windows do, so the count flips from 'by end' to 'by start'

The counts live in a three-slot array, `counts[s[i]-'a']`, rather than a map. With a
fixed alphabet of three, that's the simplest possible structure.

**Complexity:**
- Time: O(n). Each index enters and leaves the window once.
- Space: O(1).
