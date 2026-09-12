365 Days of LeetCode Challenge — Day 61/365

739. Daily Temperatures (Medium)
https://leetcode.com/problems/daily-temperatures/

Yesterday's monotonic stack with two changes.

The stack holds indices rather than values, because the answer is a distance and a distance needs positions.

And values repeat. Day 59 stored answers in a map keyed by value, which I flagged at the time was only safe because uniqueness was guaranteed. Temperatures repeat constantly. Indices are unique by construction, so the problem disappears rather than needing a workaround.

Then there is one character that decides whether duplicates break it:

temperatures[stack.Top()] <= temperatures[i]

Equal temperatures are not WARMER. A day holding the same temperature can never be anybody's answer, so it has to be popped. Write < instead and equal days stay on the stack, and the first one found reports a warmer day that is merely an equal one. On [70, 70] the correct answer is [0, 0] and < gives [1, 0].

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Stack #MonotonicStack #CodingInterview #Algorithms
