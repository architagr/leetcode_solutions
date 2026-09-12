365 Days of LeetCode Challenge — Day 80/365

23. Merge k Sorted Lists (Hard)
https://leetcode.com/problems/merge-k-sorted-lists/

This is the problem the last two months have been pointing at. It needs the linked list arc from day 22 and the heap arc from this week, and it needs both at once.

Day 22 merged two sorted lists: compare the heads, take the smaller, advance. That comparison was an if, because with exactly two candidates that is all a comparison needs to be.

With k lists, "take the smallest head" stops being one comparison. It becomes a repeated query for the minimum of a collection that changes after every answer - which is the definition of what a heap is for.

So this is day 22's merge with the if replaced by a heap. Dummy head, splicing, loop shape: all unchanged.

The detail the complexity rests on: the heap holds one node per list, never every node. A list's second element cannot be the next smallest while its first is unplaced, because the list is sorted. So the heap is size k, not size N - O(N log k) rather than O(N log N).

And splicing has a consequence. Every reused node arrives with a Next still pointing into its original list, so the merged list has to be terminated explicitly or you can hand back a cycle.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Heap #LinkedList #CodingInterview #Algorithms
