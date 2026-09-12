365 Days of LeetCode Challenge — Day 60/365

155. Min Stack (Medium)
https://leetcode.com/problems/min-stack/

A stack where push, pop, top and getMin are all O(1).

Two fixes that do not work. Scanning in getMin is O(n) per call. Keeping a single min field updated on push works right up until the first pop - pop the element that IS the minimum and the field is stale, with nothing to recompute it from short of a scan.

That second failure is the insight. The minimum is not a property of the stack. It is a property of the stack AT A GIVEN HEIGHT, and popping returns you to a height whose minimum you already threw away.

So do not throw it away. Every entry stores the value pushed and the minimum of everything at or below it.

The result is that Pop contains no minimum-maintenance code at all. Removing the top exposes an entry whose min was computed when it was on top and is still correct, because nothing below it ever changed. The old minimum is restored for free - it was never overwritten.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Stack #SystemDesign #CodingInterview #Algorithms
