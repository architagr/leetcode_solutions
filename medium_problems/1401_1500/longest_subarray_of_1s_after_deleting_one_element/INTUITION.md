## Intuition

This closes the sliding-window run with a problem that the repo solves *without* a sliding
window, which I think makes a useful ending: the same question can be looked at in more
than one way, and it's worth seeing both.

Deleting a 0 glues together the run of ones before it and the run after it. Deleting a 1
can only shorten a run, so if there's any 0 at all, deleting a 0 is the better move. The
answer is: over every zero, the length of the run to its left plus the run to its right.

The solution compresses the array to make that direct. Runs of ones become their lengths,
and each zero stays as a `0` entry: `[0,1,1,1,0,1,1,0,1]` becomes `[0, 3, 0, 2, 0, 1]`.
Now every zero sits between exactly the two runs it would join, and the answer is the best
`arr[i-1] + arr[i+1]` over the zeros. Two zeros next to each other show up as neighbouring
`0` entries, contributing nothing, which is correct.

The one special case is an array of all ones (or a single element): `arr` has one entry,
there's no zero to delete, and one of the ones has to go, so the answer is `len(nums) - 1`.

## Builds on

- [Day 106: Max Consecutive Ones III](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1001_1100/max_consecutive_ones_iii/) — the longest window with at most k zeros; with k = 1 and one element removed, it's this problem

The window view gives the same answer in one pass and O(1) space. A window with at most one
0 is a stretch that becomes all ones after deleting that 0 (or any one element), so the
answer is the longest such window, minus the one deleted element:

```go
l, zeros, best := 0, 0, 0
for r, v := range nums {
	if v == 0 {
		zeros++
	}
	for zeros > 1 {
		if nums[l] == 0 {
			zeros--
		}
		l++
	}
	best = max(best, r-l) // window length minus the deleted element
}
```

It even covers the all-ones case without a branch, since the window is the whole array and
`r - l` is already one less than its length. I checked it against the run-length version on
a few thousand random arrays.

**Complexity (as written):**
- Time: O(n): one pass to compress, one over the runs.
- Space: O(n) for the compressed array.
