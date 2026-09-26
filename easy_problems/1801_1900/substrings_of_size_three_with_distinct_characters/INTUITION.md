## Intuition

Same fixed window as yesterday, size 3 this time, but what the window has to remember
changes. A sum was enough to answer "how big is this window's total". To answer "are these
three letters all different", the window keeps a count per letter.

A map from letter to count does it, with one rule that makes the check trivial: when a
count drops to zero, delete the key. Then the number of keys in the map is exactly the
number of different letters in the window, and "good" is just `len(m) == 3`.

## Builds on

- [Day 92: Maximum Average Subarray I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/maximum_average_subarray_i/) — the fixed window that drops one element and adds one per step; here the running sum becomes a map of counts

Skip the delete and the check breaks quietly: a letter that left the window would still
be a key with count 0, and `len(m)` would overcount. On `"xyzzaz"` the window `"yzz"`
would still see `x` in the map and report three keys.

To be fair to the simpler approach: for a window of exactly 3 you could skip the map and
compare `s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2]`. The map version is
worth learning because it doesn't care about the window size. Change 3 to 10 and it still
works, and the next several problems use the same structure.

The loop order differs slightly from Day 92: this primes a full window and checks it
before the loop, then removes the old letter before adding the new one. Both orders work;
Day 92's avoids the separate first check.

**Complexity:**
- Time: O(n). Each step is a constant number of map operations.
- Space: O(1). The map never holds more than 3 keys.
