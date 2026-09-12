# Linked List Cycle — intuition

## Builds on

- [Day 20: Middle of the Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/middle_of_the_linked_list/) — the identical two-pointer walk, used here for a completely different purpose

## The problem in one line

Does following `Next` from the head ever revisit a node?

## The version that obviously works

Keep a set of every node you have seen. At each step, check whether the current node is already in it. If it is, there is a cycle; if you reach `nil`, there is not.

That is correct, it is O(n) time, and it is O(n) space. The follow-up question on the problem page asks for O(1) space, which is the only reason to look further.

## The version that is interesting

Yesterday's walk, unchanged. One pointer moves a node at a time, the other moves two.

What changes is what the walk is being used for. Yesterday the point was that the slow pointer ends up halfway. Today the point is what happens when the list has no end.

If the list is straight, the fast pointer runs off it and returns `nil`, and that settles the question. If the list loops, neither pointer can ever leave, and then the interesting thing is the gap between them.

## Why they must meet

Once both pointers are inside the loop, think about the distance from fast to slow measured forward around the cycle. Every iteration, fast advances two and slow advances one, so that gap shrinks by exactly one.

A quantity that decreases by exactly one each step and cannot go below zero must hit zero. It cannot jump over it, because it changes by one, not two. Hitting zero means the pointers are on the same node.

So a cycle is not merely likely to be detected, it is guaranteed, and within one lap of the loop. That is the part I find genuinely nice: the argument is about a gap closing by one, and it needs nothing about list lengths or where the loop begins.

## The comparison that matters

```go
if slow == fast {
```

That compares nodes, not values. Two different nodes can hold the same number, and comparing `slow.Val == fast.Val` would report a cycle for the perfectly straight list `[1,1]`.

## Complexity

- **Time: O(n).** Without a cycle, fast reaches the end in about n/2 iterations. With one, the pointers meet within one lap after slow enters the loop.
- **Space: O(1).** Two pointers. That is the entire reason to prefer this over the set.
