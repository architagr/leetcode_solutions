**365 Days of LeetCode Challenge — Day 64/365**
**Exclusive Time of Functions** (Medium)
🔗 https://leetcode.com/problems/exclusive-time-of-functions/

Closes the stack arc, and here the stack is not a modelling choice. Function calls nest, so the log IS a recording of a call stack and the algorithm replays it.

"Exclusive" is the word doing the work. If 0 runs 0..6 and 1 runs 2..5 inside it, 0's exclusive time is 2, not 6.

Two interval rules, in the same short function, and they differ:

```go
out[idx] += time - sTime        // pause: no +1
out[idx] += time - sTime + 1    // end:   +1
```

Pausing is half-open - the function ran up to but not including the instant its child started. Ending is inclusive - a function starting and ending at the same instant ran for one unit, not zero.

And this is the line that makes it exclusive:

```go
if len(st) > 0 {
	st[len(st)-1][1] = time + 1
}
```

When a child ends at time t, the parent resumes at t+1, because instant t was already charged to the child. Delete it and the parent is charged its whole span including everything the child did - that is inclusive time, a different question.

Note the entry is MUTATED here, unlike day 60 where a stack entry was written once and never touched.

O(n) time, O(depth) space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/601_700/exclusive_time_of_functions/SOLUTION.md
