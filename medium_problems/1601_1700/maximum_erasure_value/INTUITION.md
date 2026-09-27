## Intuition

"Erase a subarray of unique elements, score its sum" is Longest Substring Without
Repeating Characters with a different score. There the prize was the window's length; here
it's the window's sum. The window itself is identical: grow right, and on a repeat, remove
everything from the left edge through the earlier copy.

So the solution is Day 104's code with a running `sum` added to the window and removed
from it alongside the map entries.

## Builds on

- [Day 104: Longest Substring Without Repeating Characters](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1_100/longest_substring_without_repeating_characters/) — the same map of value to index, deleting from the left through the earlier copy on a repeat

Where the answer is recorded follows the pattern from earlier in the arc: right before a
repeat forces a shrink, and once more after the loop. That's valid because every number is
positive. Between two repeats the window only gains elements, so its sum only goes up, and
the moment before a repeat is when it's largest.

With negative numbers that would stop being true: a window could be worth more after
dropping a negative element from the left, and recording only at repeats could miss it. The
constraint `nums[i] >= 1` is what makes the lazy recording safe.

**Complexity:**
- Time: O(n). Each element is added once and removed at most once.
- Space: O(n) for the map.
