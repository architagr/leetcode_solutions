365 Days of LeetCode Challenge — Day 63/365

341. Flatten Nested List Iterator (Medium)
https://leetcode.com/problems/flatten-nested-list-iterator/

The obvious solution walks the whole structure in the constructor, copies every integer into a slice, and indexes it in Next.

It returns the right values. It is not an iterator, and that distinction is the entire problem. A flatten-then-index does all the work up front whether the caller wants one element or all of them, and keeps a copy of every integer alive for the object's lifetime. Fine on a small input, wrong in exactly the situations iterators exist for.

An iterator does nothing until asked, and keeps only enough state to resume. For a nested list that state is the path from the outer list down to the cursor - which list you are in and how far into it, at every level. A stack of cursors.

Which is exactly what a BST iterator holds, so day 55 and today are the same design wearing different data.

The part that surprises people: HasNext is the method that does the work. It pops finished lists, descends into nested ones, and stops when the cursor reaches an integer. Next only reads and steps past.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Stack #Iterators #CodingInterview #Algorithms
