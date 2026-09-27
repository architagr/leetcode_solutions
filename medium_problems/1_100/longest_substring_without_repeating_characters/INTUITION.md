## Intuition

This is probably the most famous sliding-window problem, and after Day 102 it's a smaller
step than its reputation suggests. There the window could hold each letter twice; here,
once. Grow the right edge, and when the new letter is already in the window, move the left
edge past its earlier copy.

What changes is what the window remembers. Instead of a count per letter, the map stores
*where* each letter sits: letter to index. When `s[i]` is already in the map at index
`r`, the left edge needs to land at `r + 1`, and the map says so directly.

## Builds on

- [Day 102: Maximum Length Substring With Two Occurrences](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/3001_3100/maximum_length_substring_with_two_occurrences/) — the same grow-and-shrink window with the limit at two copies; here the limit is one

This solution keeps the map equal to the window at all times. When a repeat shows up, it
deletes every letter from `start` through `r` before adding the new one. That makes two
things simple: "is this letter in the window?" is just "is it in the map?", and the window
length is `len(uniqueChar)`.

The deletions look like they could cost O(n) per step, but each letter is deleted at most
once after being added, so the total is O(n).

A common variant skips the deleting. Keep every letter's last position, and only move
`start` when that position is inside the current window:

```go
if j, ok := last[s[i]]; ok && j >= start {
	start = j + 1
}
last[s[i]] = i
best = max(best, i-start+1)
```

Here the map holds stale entries from before `start`, and the `j >= start` check is what
ignores them. Both are O(n). I find the deleting version easier to trust, because the map
never holds anything that isn't in the window.

**Complexity:**
- Time: O(n). Each character is added once and deleted at most once.
- Space: O(m) for the map, where m is the size of the character set.
