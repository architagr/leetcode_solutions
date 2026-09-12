365 Days of LeetCode Challenge — Day 24/365

19. Remove Nth Node From End of List (Medium)
https://leetcode.com/problems/remove-nth-node-from-end-of-list/

"nth from the end" is a position measured from a place you cannot start from. A singly linked list gives you the head; the end only exists once you have walked to it. So the first job is translating it into count - n, a position measured from the front.

The part that actually catches people is smaller than that. To remove a node from a singly linked list you need the node before it. A node cannot remove itself, because it has no pointer to whatever refers to it. Removal is always performed by the predecessor.

So the second walk stops one node early, on purpose. That off-by-one is the whole problem.

Removing the head is its own branch here, because the head has no predecessor. Two days ago a dummy node erased exactly that special case. Worth noticing when you are writing the branch that there was another option.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #LinkedList #TwoPointers #CodingInterview #Algorithms
