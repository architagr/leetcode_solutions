365 Days of LeetCode Challenge — Day 55/365

Binary Search Tree Iterator (Medium)
🔗 https://leetcode.com/problems/binary-search-tree-iterator/

An iterator promises an ordering and hides when the work happens. That second half is the
only real decision in this problem.

This solution takes the blunt end of it: the constructor walks the whole tree in-order and
flattens it into a slice, so next() becomes a slice read and hasNext() a bounds check.
Both O(1), no traversal state to restore.

What it costs is O(n) memory held for the iterator's lifetime, and a full walk before the
caller has asked for a single value. The follow-up asks for the O(h) version instead.

Neither is wrong — but you should be able to say which you wrote and why.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinarySearchTree #DataStructures #Golang
