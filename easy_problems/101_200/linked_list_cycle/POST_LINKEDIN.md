365 Days of LeetCode Challenge — Day 21/365

141. Linked List Cycle (Easy)
https://leetcode.com/problems/linked-list-cycle/

The answer most people write first keeps a set of visited nodes. Correct, and O(n) space.

The constant-space version is yesterday's walk with nothing changed: one pointer moving a node at a time, one moving two. What changes is the question. If the list has an end, the fast pointer finds it and the answer is no. If it does not, neither pointer can ever leave.

Why they must meet is one sentence. Inside the loop, the forward gap between them shrinks by exactly one every iteration, and a non-negative integer that drops by one per step cannot skip past zero. Gap zero means the same node.

That is also why the steps are 1 and 2 rather than 1 and 3. A gap shrinking by two can step straight over zero, and the pointers pass without landing together.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #LinkedList #TwoPointers #CodingInterview #Algorithms
