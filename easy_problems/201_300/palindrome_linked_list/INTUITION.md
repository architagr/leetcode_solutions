# Palindrome Linked List — intuition

## Builds on

- [Day 19: Reverse Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/reverse_linked_list/) — the in-place reversal, used here on half the list
- [Day 20: Middle of the Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/middle_of_the_linked_list/) — the fast/slow walk that finds where to split

## The problem in one line

Does the list read the same forwards and backwards?

## Why it is harder than the array version

On an array this is trivial: an index at each end, walk them toward each other, compare. Ten seconds of work.

A singly linked list has no way to walk backwards. There is no index, and a node has no pointer to the thing before it. So the array algorithm, which depends entirely on being able to move right-to-left, simply cannot run.

## The answer that sidesteps it

Copy every value into a slice and run the array algorithm on that. Correct, O(n) time, and O(n) space.

The follow-up on the problem page asks for O(n) time and O(1) space, and the whole interest of this problem lives in that follow-up.

## The answer this arc has been building toward

If the problem is that you cannot walk the second half backwards, then reverse the second half so that walking it forwards goes backwards through the original.

That is it. And both pieces are already done:

1. Find where the middle is, with day 20's fast/slow walk.
2. Reverse everything after it, with day 19's in-place reversal.
3. Walk the two halves inward, comparing as you go.

No new technique. The only new thing is noticing that the problem decomposes into two problems already solved, which is a skill worth more than either of them individually.

## The detail that makes the comparison loop simple

Stop when the *second* half runs out, not the first.

On an odd-length list the two halves are unequal: the middle node belongs to the first half and has no partner. If the loop is driven by the second half, that middle node is never compared, which is exactly right, since a palindrome's middle element is compared with itself and always matches. No parity check anywhere.

## Putting the list back

The function reverses half of the caller's list. Unless it undoes that, the caller's list is silently corrupted by a function whose name promises only to answer a question.

Reversing the second half again restores it. It costs another O(n/2) pass and changes no complexity class. I think it is worth doing, and I think an interviewer asking this question is at least as interested in whether you notice the problem as in whether you solve it.

## Complexity

- **Time: O(n).** A pass to find the middle, a pass to reverse, a pass to compare, a pass to restore. Four halves, still linear.
- **Space: O(1).** A handful of pointers. Nothing allocated.
