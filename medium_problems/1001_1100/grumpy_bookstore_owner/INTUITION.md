## Intuition

The technique is a window of `minutes` consecutive minutes. Wherever it goes, the day
splits into three parts: before the window, inside it, and after it. Outside the window
things are as normal, so only the customers who arrive while the owner is in a good mood
count. Inside, everyone counts.

This solution computes that three-part total directly for every window position, using
two prefix-sum arrays built up front:

- `allGoodCount[i]`: customers in minutes `0..i` as if the owner were never grumpy.
- `grumpyCount[i]`: customers in minutes `0..i` who are actually satisfied, i.e. the
  ones who came in non-grumpy minutes. (The name reads the opposite way round; it counts
  the customers the grumpiness *doesn't* lose.)

With those, any window's total is three O(1) lookups: satisfied-as-usual before it, all
customers inside it, satisfied-as-usual after it. Slide the window and keep the best.

## Builds on

- [Day 92: Maximum Average Subarray I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/maximum_average_subarray_i/) — sliding a fixed-length window and scoring each position in constant time
- [Day 100: Maximum Points You Can Obtain from Cards](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1401_1500/maximum_points_you_can_obtain_from_cards/) — an answer made of what's inside a window plus what's outside it

There's a leaner version worth knowing. The customers who are satisfied no matter what
(non-grumpy minutes) are a fixed base; the window only changes how many grumpy-minute
customers are recovered. So: add up the base once, then run an ordinary fixed window
over "customers in grumpy minutes" and keep its best sum. The answer is base plus best.
That's the Day 92 loop with O(1) space instead of two O(n) arrays. Both are O(n) time, and
on Example 1 both give `10 + 6 = 16`.

If `minutes >= n`, the whole day is covered and the answer is every customer; the code
returns `allGoodCount[n-1]` for that.

**Complexity (as written):**
- Time: O(n): one pass to build the prefix sums, one to slide.
- Space: O(n) for the two prefix arrays.
