---
meta_title: "Why two pointers always meet inside a loop"
meta_description: "Floyd's cycle detection is yesterday's walk with a different question attached, and the proof it works is one sentence about a gap shrinking by one."
---

![Day 21](HERO.png)

## 365 Days of LeetCode Challenge — Day 21/365

**[141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)** (Easy)

Does following `Next` from the head ever come back to a node you have already visited?

## The answer everyone writes first

Keep a set of the nodes you have seen. At each step check whether the current node is in it. If it is, there is a cycle. If you reach `nil`, there is not.

Correct, O(n) time, O(n) space. The follow-up on the problem page asks for constant space, and that is the only reason to keep going.

## The same walk as yesterday

Day 20 ran two pointers, one moving a node per step and one moving two, to find a list's middle. Today's solution is that walk, unchanged. What changes is the question attached to it.

```go
func HasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
```

Yesterday those two loop conditions were only there to stop the fast pointer dereferencing nothing. Here, reaching them *is* the answer: a list that has an end cannot have a cycle, so falling out of the loop returns false.

And a cyclic list never falls out. Neither pointer can reach `nil` because there is no `nil` to reach. The only exit is the return inside.

## Watching it run

`[3,2,0,-4]`, with the tail pointing back at index 1.

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

One iteration: `slow` on 2, `fast` on 0.

![Step 3](images/walkthrough-3.png)

Two iterations, and `fast` has gone round through the back edge to sit *behind* `slow`. Worth pausing here, because "fast is ahead" stops meaning anything once the structure wraps. What still means something is the distance forward around the cycle.

![Step 4](images/walkthrough-4.png)

Three iterations and both land on the same node.

## Why they must meet

This is the part worth actually knowing, because it is one sentence.

Once both pointers are inside the loop, measure the gap as the number of steps forward from `fast` to `slow` around the cycle. Every iteration `fast` moves two and `slow` moves one, so the gap shrinks by exactly one.

A non-negative integer that decreases by exactly one per step must reach zero. It cannot jump over zero, because it changes by one, not two. Gap zero means the same node.

So detection is guaranteed rather than probable, and it happens within one lap.

It also explains the step sizes. With 1 and 3, the gap shrinks by two per iteration and can step straight over zero, and the pointers pass each other without ever landing together.

## One line that is easy to get wrong

```go
if slow == fast {
```

These are pointers, so this asks whether they are the same node. That is the actual question.

`slow.Val == fast.Val` is a different and wrong program: it reports a cycle for `[1,1]`, a straight two-node list, because after one iteration both pointers sit on nodes holding 1.

The check also has to come after both pointers move. At the top of the first iteration they are both on the head, and comparing there returns true for every non-empty list.

## Complexity

- **Time: O(n)**. No cycle, the fast pointer hits the end in about n/2 iterations. Cycle, they meet within a lap.
- **Space: O(1)**. Two pointers, which is the whole point next to the hash set.

## Builds on

- [Day 20: Middle of the Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/middle_of_the_linked_list/) — the identical two-pointer walk, used here for a completely different purpose

Full code and the step-by-step walkthrough:
[linked_list_cycle](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/linked_list_cycle/SOLUTION.md)

#DSA #LeetCode #Golang #LinkedList #TwoPointers #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
