365 Days of LeetCode Challenge — Day 81/365

264. Ugly Number II (Medium)
https://leetcode.com/problems/ugly-number-ii/

Testing each integer in turn is hopeless here - the 1690th ugly number is 2,123,366,400, so almost everything you test is a miss. Stop testing candidates and start generating them.

Read the definition forwards: the only way to produce an ugly number is to multiply a smaller one by 2, 3 or 5. So the sequence builds out of itself.

The obvious structure for "repeatedly take the smallest" is a heap, which this whole arc has been about. Push 1, pop the smallest, push its three multiples. Correct, O(n log n) - and it needs a SET alongside it, because 6 arrives as both 2x3 and 3x2 and would otherwise be emitted twice.

That set is the tell. The structure is generating work it then throws away.

The three-pointer version keeps one pointer per multiplier and takes the smallest of the three candidates. When 6 comes up as both 2x3 and 3x2, BOTH pointers advance - three separate ifs, not an else-if. That is the deduplication: no set, just the observation that a value produced twice has consumed both producers.

O(n) instead of O(n log n), with no heap and no set.

Days 76 to 78 built the case for the heap. Day 79 found a problem where sorting reads better. Today the heap is the obvious answer and an array with three counters is strictly better. Knowing when to put a technique down is half the skill.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #DynamicProgramming #Heap #CodingInterview #Algorithms
