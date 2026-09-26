## Intuition

This is the first problem in a run on sliding windows, and it's the cleanest example of
the idea, so it's worth being slow about it.

Every subarray of length `k` overlaps the next one in all but two places: the next window
loses its first element and gains one new element at the end. Recomputing each window's
sum from scratch costs `k` additions per window, O(n·k) in total. Keeping one running sum
and adjusting it (add the element coming in, subtract the element going out) costs two
operations per window, O(n) in total, regardless of `k`.

The other observation is that the average is just the sum divided by `k`, and `k` never
changes. The window with the largest average is the window with the largest sum. So the
loop only ever compares integer sums and divides once, at the very end. No floating point
inside the loop at all.

This solution uses a shape I'll keep coming back to over the next few days:

1. `k--`, so that `k` now means "how far back the oldest element of the window is".
2. Add up the first `k` elements, which is one short of a full window.
3. In the loop, add `nums[i]` (now the window is full), check it, then subtract
   `nums[i-k]` (the oldest element), leaving it one short again for the next step.

Priming one short means the "add, check, remove" order inside the loop is the same on
every iteration, including the first. There's no separate "check the first window" line
before the loop.

`max` starts at `math.MinInt` rather than 0 because every value can be negative, and a
best sum of 0 that no window ever had would be a wrong answer.

**Complexity:**
- Time: O(n). One pass, two arithmetic operations per step.
- Space: O(1).
