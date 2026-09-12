365 Days of LeetCode Challenge — Day 58/365

20. Valid Parentheses (Easy)
https://leetcode.com/problems/valid-parentheses/

New topic. Stacks start today.

Start with what does not work. Counting openers against closers accepts )( - balanced by count, wrong by order. Tracking a depth that rises and falls accepts ([)] - depth never goes negative and ends at zero.

Both fail for the same reason, and it is worth naming precisely: a single number records HOW MANY brackets are open and never WHICH. This problem is entirely about which.

What has to be remembered is the list of brackets opened and not yet closed, in order. And when a closer arrives only one candidate can match it - the most recent opener. In "([", a ")" is wrong not because "(" is unmatched but because "[" is still open and has to close first.

Most recent in, first out. That is a stack, and seeing that this problem is one is the whole exercise.

One detail worth stealing: key the map by the CLOSING bracket. One lookup then tells you both whether the character closes anything and what must be on top if it does.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Stack #CodingInterview #SoftwareEngineering #Algorithms
