365 Days of LeetCode Challenge — Day 76/365

1046. Last Stone Weight (Easy)
https://leetcode.com/problems/last-stone-weight/

Heaps start today.

Why sorting is the wrong shape here: sort descending, take the first two, smash them, put the remainder back - in the right place, which is an insertion. O(n) per round, paid every round.

The deeper mismatch is that sorting answers a question nobody asked. This problem never needs the third-heaviest stone or the ordering of the rest. It needs the maximum, repeatedly, from a collection that keeps changing.

A heap keeps exactly one promise: the largest is at the root. It says nothing about the order of anything else. Read its backing array left to right and it looks close to random - and that is the point. Maintaining less is what makes it cheaper.

One detail that catches people with Go's container/heap: it always builds a MIN-heap. You get a max-heap by lying to it, defining Less as greater-than. That single inverted comparison is the whole difference.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Heap #PriorityQueue #CodingInterview #Algorithms
