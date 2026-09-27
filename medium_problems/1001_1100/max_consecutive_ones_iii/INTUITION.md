## Intuition

The problem talks about flipping zeros, but nothing actually needs flipping. A stretch of
the array can be turned into all ones with at most `k` flips exactly when it contains at
most `k` zeros. So the question is really: what's the longest window with at most `k`
zeros?

Put like that, it's the variable window from the last few days with a very simple rule.
Push the right edge. Ones are free. Each zero adds to a count. When the count goes past
`k`, the window has to shrink from the left until one zero has left it.

## Builds on

- [Day 103: Count Substrings That Satisfy K-Constraint I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/3201_3300/count_substrings_that_satisfy_k_constraint_i/) — a binary-string window that counts one kind of character and shrinks once that count goes past k

The recording works like [Day 102](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/3001_3100/maximum_length_substring_with_two_occurrences/):
the window only grows between violations, so its length just before a violation (`r - l`,
not counting the zero that just arrived) is the best that stretch reached. The code records
it there, inside the shrink loop, and once more after the main loop for a final window
that never hit a violation.

Recording inside the shrink loop means `ans` is compared on every step of the shrink,
with `r - l` getting smaller each time. Only the first comparison can matter. It's
harmless, just a little more work than needed.

The `continue` for ones is a small nicety: a one can never cause a violation, so the rest
of the loop body doesn't need to run.

**Complexity:**
- Time: O(n). `l` and `r` each move forward at most n times.
- Space: O(1).
