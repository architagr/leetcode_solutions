365 Days of LeetCode Challenge — Day 79/365

692. Top K Frequent Words (Medium)
https://leetcode.com/problems/top-k-frequent-words/

One sentence separates this from yesterday.

Day 78: return the k most frequent elements, in any order.
Today: return them sorted by frequency, and sort words with the same frequency lexicographically.

That addition is enough to change which tool is right.

Yesterday's size-k min-heap still works. It gives the right set in ascending frequency, so you reverse it and put the tie-break in the comparator. And that is where it turns unpleasant.

A heap evicts its root, which must be the WORST entry. For counts, worst means least frequent. For a tie, worst means the word that should lose - the lexicographically later one. So the two fields point in opposite directions: inverted for the tie-break, not inverted for the count.

That line is correct and it is exactly the kind of thing that sits wrong in a codebase for months, because nothing about it looks unusual and it only misbehaves on ties.

Sorting is O(d log d) against the heap's O(d log k), and with the input capped at 500 words that difference is unmeasurable. The comparator reading like the spec is worth more.

Days 76 to 78 built the case for the heap. This is the day to notice the technique is not the goal.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Sorting #Heap #CodingInterview #Algorithms
