365 Days of LeetCode Challenge — Day 77/365

215. Kth Largest Element in an Array (Medium)
https://leetcode.com/problems/kth-largest-element-in-an-array/

First, the definition: "kth largest" means position k in sorted-descending order, not the kth distinct value. In [5,5,4] the 2nd largest is 5, not 4. A solution that deduplicates passes every simple example and fails there.

Now the cost. Pushing all n into a heap and popping k times is O(n log n) - the same as sorting and indexing. What the heap buys is stopping after k pops rather than ordering everything: a constant-factor saving, not an asymptotic one. This is a heap performing a sort, interrupted early.

The version the question is really asking for keeps a MIN-heap capped at k. Push each element; if the heap exceeds k, pop the smallest. That is O(n log k), and with k of 5 and n of a million, log k is 2 and log n is 20.

Why a min-heap when you want the largest? Because to retain the k largest, the element you need at your fingertips is the smallest of the ones you kept - that is what an arriving element must beat.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Heap #PriorityQueue #CodingInterview #Algorithms
