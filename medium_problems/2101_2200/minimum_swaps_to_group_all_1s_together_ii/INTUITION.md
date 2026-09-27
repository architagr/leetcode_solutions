## Intuition

When all the ones are grouped, they fill a block whose length is exactly the number of
ones. So pick where that block will be, and count what's wrong with it: every zero inside
the block must be swapped with a one outside it. There are exactly as many ones outside as
zeros inside, so the cost of a block is its number of zeros, and each swap fixes one.

The answer is the smallest number of zeros in any block of that length. That's Minimum
Recolors from Day 94 (fewest whites in a window of size `k`) with zeros instead of
whites and the window size fixed by counting ones first. The circular part is Day 99:
index with `(i + k) % n` and let the window wrap past the end.

## Builds on

- [Day 94: Minimum Recolors to Get K Consecutive Black Blocks](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/2301_2400/minimum_recolors_to_get_k_consecutive_black_blocks/) — the fewest 'wrong' elements in any window of a fixed size; there whites, here zeros
- [Day 99: Defuse the Bomb](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1601_1700/defuse_the_bomb/) — a fixed window that wraps around the end with (i + k) % n

The code counts the ones (`maxOne`) and returns 0 if there are none. Then it primes a
window one short, `k = maxOne - 1`, and for every start position `i` from 0 to `n - 1`
adds `nums[(i+k)%n]`, takes the minimum, and removes `nums[i]`. Starting positions near
the end wrap around to the front, which is how Example 3's `[1,1,0,0,1]` finds a block with
no zeros at all.

**Complexity:**
- Time: O(n). One pass to count, one to slide.
- Space: O(1).
