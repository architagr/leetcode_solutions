## Intuition

Only one thing about each number matters: whether it's odd. Replace every odd number with 1
and every even number with 0, and "a subarray with exactly `k` odd numbers" becomes "a
subarray whose sum is exactly `k`". That's yesterday's problem.

So the solution is yesterday's, with the running sum counting odd numbers. At each index,
the nice subarrays ending there correspond to earlier positions where the odd count was
`prefixSum - k`, and a map of how often each count has appeared gives that number in one
lookup.

## Builds on

- [Day 120: Binary Subarrays With Sum](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/901_1000/binary_subarrays_with_sum/) — counting subarrays with an exact sum by looking up earlier prefix sums in a map

The difference from yesterday's code is the empty prefix. Yesterday handled it with an
`if sum == goal` check. Here the map is seeded with `m[0] = 1` before the loop: "zero odd
numbers seen, once, before the array starts". Then `m[prefixSum-k]` counts subarrays that
start at index 0 without any special case. This is the more common way to write it.

Example 3 shows why a map is the right tool. `[2,2,2,1,2,2,1,2,2,2]` with `k = 2` has only
one pair of odd numbers, but 16 nice subarrays: four choices of where to start (0 to 3 evens
before the first odd) times four of where to end (0 to 3 evens after the second). The map
counts all of those without enumerating them.

**Complexity:**
- Time: O(n).
- Space: O(n) for the map.
