365 Days of LeetCode Challenge — Day 59/365

496. Next Greater Element I (Easy)
https://leetcode.com/problems/next-greater-element-i/

The brute force scans right from each element until something bigger turns up. O(n squared), and what it wastes is specific: the scans repeat each other. Scan right across a long descending run, then start one position later and scan across almost the same run.

Flip the question. Instead of "what is to the right of me", ask for each value as you arrive at it: which earlier values does this one answer?

A value resolves every earlier value smaller than it that is still waiting, and a resolved value is never looked at again.

What you remember is the values still waiting - and they are always decreasing, because if an earlier one were smaller than a later one, the later one would have resolved it on arrival. A stack that maintains an ordering invariant like that is a monotonic stack.

The obvious objection is that the inner loop makes this quadratic. It does not: every value is pushed once and popped at most once, so across the whole run the inner loop executes at most n times in total.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Stack #MonotonicStack #CodingInterview #Algorithms
