**365 Days of LeetCode Challenge — Day 25/365**
**Reverse Linked List II** (Medium)
🔗 https://leetcode.com/problems/reverse-linked-list-ii/

Reverse only the nodes between positions left and right.

The reversal is day 19's function, character for character. What makes this a Medium is the sewing: a reversed run in the middle of a list has two seams, and they are different problems. The front seam is a node you walked past and remembered. The back seam is a node you can only find AFTER reversing, because reversing decides which node lands there.

The move that keeps it clean:

```go
rightNode = tempHead.Next
tempHead.Next = nil
leftSide = reverse(leftSide)
```

Cut the run loose first. Once it is nil-terminated it is an ordinary list, so day 19's reverse works on it untouched - no bounds, no counter, nothing inside it knows it is part of something bigger.

Then walk the reversed run to its end to find the new tail, because you could not have saved that pointer beforehand.

Third time this week that "the head has no predecessor" has come up. Day 22 erased it with a dummy node; day 24 and today take an explicit branch.

O(n) time, O(1) space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1_100/reverse_linked_list_ii/SOLUTION.md
