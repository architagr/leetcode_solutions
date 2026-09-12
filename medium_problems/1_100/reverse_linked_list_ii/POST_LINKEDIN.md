365 Days of LeetCode Challenge — Day 25/365

92. Reverse Linked List II (Medium)
https://leetcode.com/problems/reverse-linked-list-ii/

Reverse only the nodes between positions left and right.

The reversal is day 19's function, character for character. What makes this a Medium is everything around it: a reversed run sitting in the middle of a list has to be sewn back in on both sides, and the two seams are different problems.

The front seam is a node you walked past and had to remember. The back seam is a node you can only find after reversing, because reversing is what decides which node ends up there.

The move that keeps it clean is cutting the run loose before reversing it. Once it is nil-terminated it is an ordinary list, so the Easy problem's function works on it untouched. No bounds, no counter, no special cases inside the reversal.

Third time this week that "the head has no predecessor" has come up. Day 22 erased it with a dummy node, day 24 took a branch, this takes a branch. Worth noticing the pattern repeating.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #LinkedList #CodingInterview #SoftwareEngineering #Algorithms
