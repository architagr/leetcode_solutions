## Intuition

"Some permutation of `s1` appears in `s2`" sounds like it involves generating
permutations, of which there are up to `n!`. It doesn't. Two strings are permutations of
each other exactly when they have the same letter counts. So the question is: does any
substring of `s2` with length `len(s1)` have the same counts as `s1`?

That's a fixed window of length `len(s1)` sliding over `s2`, carrying a count map, the
Day 93 structure. At each position, compare the window's counts with `s1`'s.

## Builds on

- [Day 93: Substrings of Size Three with Distinct Characters](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1801_1900/substrings_of_size_three_with_distinct_characters/) — a fixed window carrying letter counts in a map, deleting keys at zero

The comparison is `check`: the maps must have the same number of keys, and every letter's
count must match. Because the window map deletes letters at zero, a letter that has left
the window can't linger as a key with count 0 and make two maps look different.

`check` is O(26) (at most 26 keys), so the whole thing is O(26·n), linear. A tighter
version keeps a single number, "how many letters currently have matching counts", and
updates it in O(1) as each letter enters and leaves; when it reaches 26 the window matches.
It saves the constant, and it's a nice trick, but the map comparison is easier to read and
the same complexity.

**Complexity:**
- Time: O(26·n), which is O(n).
- Space: O(1): at most 26 keys per map.
