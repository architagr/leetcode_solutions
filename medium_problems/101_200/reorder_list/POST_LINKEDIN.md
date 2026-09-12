365 Days of LeetCode Challenge — Day 26/365

143. Reorder List (Medium)
https://leetcode.com/problems/reorder-list/

Turn L0 -> L1 -> ... -> Ln into L0 -> Ln -> L1 -> Ln-1 -> ...

Written out that looks like an arbitrary shuffle. Written as two sequences it is not:

L0, L1, L2, ... walking forward from the start
Ln, Ln-1, Ln-2, ... walking backward from the end

Interleaved. So it is a merge of the list with its own reverse, and the only obstacle left is that a singly linked list has no backward walk. Reverse the back half and walking it forward walks the original backward.

Find the middle. Cut and reverse. Merge alternately. That is day 20, day 19 and day 22, called in order.

This closes the linked list arc, and it closes it on a problem with nothing new in it. That is the point. Most problems that feel hard are not a new idea, they are two or three old ideas standing next to each other, and the skill being tested is recognising them.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #LinkedList #TwoPointers #CodingInterview #Algorithms
