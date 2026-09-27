## Intuition

This is the first problem in a while that the repo doesn't solve with a window, and the
reason is worth understanding.

Windows count well when the condition is "at most" or "at least": pieces or extensions of a
valid window stay valid. "Exactly `goal`" is neither. Zeros make it worse: a window summing
to `goal` can shed or gain zeros at either end and still sum to `goal`, so there's no single
right place to stop, and with `goal = 0` it isn't clear the window should hold anything.

Prefix sums handle "exactly" directly. If `prefix[i]` is the running sum up to `i`, the
subarray after position `j` and up to `i` sums to `prefix[i] - prefix[j]`. It equals
`goal` exactly when `prefix[j] = prefix[i] - goal`. So walking left to right, the number
of valid subarrays ending at `i` is the number of earlier prefixes equal to
`sum - goal`. A map from prefix value to how many times it's appeared answers that in O(1).

## Builds on

- [Day 101: Grumpy Bookstore Owner](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1001_1100/grumpy_bookstore_owner/) — prefix sums, where the sum of any range is the difference of two prefix values

One detail: the empty prefix (sum 0, before the first element) is also a valid `j`; it's
what makes a subarray starting at index 0 countable. This code handles it with
`if sum == goal { count++ }` instead of putting `0: 1` in the map up front. Both work.
Tomorrow's problem uses the seeded-map version.

There's also a window-based answer: count subarrays with sum *at most* `goal`, subtract
those with sum at most `goal - 1`, and the difference is exactly `goal`. It uses O(1)
space. The prefix map is the more general tool, since it doesn't need the values to be
non-negative.

**Complexity:**
- Time: O(n).
- Space: O(n) for the map of prefix counts.
