365 Days of LeetCode Challenge — Day 78/365

347. Top K Frequent Elements (Medium)
https://leetcode.com/problems/top-k-frequent-elements/

Two steps: count the values, then select the top k by count. Only the second has a decision in it.

Sorting the counts works and is O(d log d) in the number of distinct values, but the follow-up asks for better than O(n log n) - and sorting computes a total ordering when only the top k is wanted.

A heap capped at k does it in O(d log k). Push each pair; if the heap exceeds k, evict.

And now the part that reads backwards. To keep the k MOST frequent, you want a MIN-heap.

Ask what the structure does on each arrival: decide which incumbent the new entry displaces. The one displaced is always the LEAST frequent of those currently kept - the weakest member. That is the entry you need instant access to, so it has to be at the root.

A max-heap of size k holds the most frequent at the root. That entry is in no danger at all, and finding the one to evict would mean scanning.

Get it backwards and nothing crashes. It silently returns the k least frequent elements.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Heap #PriorityQueue #CodingInterview #Algorithms
