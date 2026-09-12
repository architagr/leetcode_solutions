# Daily Temperatures — intuition

## Builds on

- [Day 59: Next Greater Element I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/401_500/next_greater_element_i/) — the same monotonic stack, holding indices here because the answer is a distance and values can repeat

## The problem in one line

For each day, how many days until a warmer one? Zero if there is no warmer day.

## It is day 59 with two changes

Day 59 asked *which* value comes next and is larger. This asks *how far away* it is.

Two consequences follow from that, and they are the whole difference:

**The stack holds indices, not values.** A distance needs positions, so `i` goes on the stack and the answer is `stack.Top() - i`.

**Values can repeat.** Day 59 could key a map by value because uniqueness was guaranteed. Temperatures repeat constantly, so a value-keyed map would collide. Indices are unique by construction, which removes the problem rather than working around it.

## This version scans backwards

Day 59 scanned left to right, asking "which earlier values does this one answer?" This scans right to left, asking "what is the first thing to my right that beats me?"

Both are monotonic stacks and both are linear. The backward version computes each answer at the moment it visits the element, so there is no bookkeeping to write into a map — the result slot is filled and the loop moves on.

## The invariant, going backwards

The stack holds indices of days to the right that are still candidates for being someone's answer.

A day is only a candidate if nothing between it and you is at least as warm. So when day `i` is processed, every index on top of the stack holding a temperature `<= temps[i]` is useless to everyone further left: they would hit day `i` first. Those get popped.

What is left is increasing in temperature from top to bottom, and the top is the nearest day warmer than `temps[i]`.

## Why `<=` and not `<`

Equal temperatures are not warmer. A day with the same temperature can never be anybody's answer, so it is popped rather than kept.

Using `<` leaves equal days on the stack, and the first one found becomes a wrong answer of "0 days until warmer" when it is not warmer at all.

## Complexity

- **Time: O(n).** Each index is pushed once and popped at most once.
- **Space: O(n)** for the stack, worst case a strictly decreasing temperature sequence.
