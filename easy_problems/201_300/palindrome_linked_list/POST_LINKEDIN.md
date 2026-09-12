365 Days of LeetCode Challenge — Day 23/365

234. Palindrome Linked List (Easy)
https://leetcode.com/problems/palindrome-linked-list/

On an array this is nothing: an index at each end, walk inward, compare. A singly linked list has no index and no pointer backwards, so that algorithm cannot run at all.

Copying the values into a slice works and costs O(n) space. The follow-up asks for O(1).

Here is where this week pays off. If the problem is that you cannot walk the second half backwards, reverse the second half, and then walking it forwards walks the original backwards.

Find the middle with day 20's fast/slow walk. Reverse the back half with day 19's in-place reversal. Compare inward. No new technique at all, just noticing the problem is two problems you already solved.

One detail worth stealing: drive the comparison loop with the second half. On odd lengths the middle node has no partner, and letting the shorter half end the loop skips it automatically. No parity check needed.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #LinkedList #TwoPointers #CodingInterview #Algorithms
