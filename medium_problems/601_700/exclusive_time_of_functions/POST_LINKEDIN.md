365 Days of LeetCode Challenge — Day 64/365

636. Exclusive Time of Functions (Medium)
https://leetcode.com/problems/exclusive-time-of-functions/

This closes the stack arc, and it is the one where the stack is not a modelling choice at all. Function calls nest - a function that starts while another is running must finish before that one resumes. That is what a call stack is, and the log is a recording of one. The algorithm replays it.

The word doing the work is "exclusive". If function 0 runs from 0 to 6 and function 1 runs from 2 to 5 inside it, then 0's exclusive time is 2, not 6. The time spent waiting on 1 belongs to 1.

Two interval rules live in the same short function and they differ. Pausing uses time - start with no +1, because the function ran up to but not including the instant its child started. Ending uses time - start + 1, because an end timestamp is inclusive - a function starting and ending at the same instant ran for one unit, not zero.

And one line makes the whole thing exclusive: when a function ends, the one underneath resumes at time + 1, not time, because that instant was already charged to the child. Delete it and you have computed inclusive time instead.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Stack #CodingInterview #SoftwareEngineering #Algorithms
