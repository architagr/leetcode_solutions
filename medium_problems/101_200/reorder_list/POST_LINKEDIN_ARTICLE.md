---
meta_title: "The reorder that is really three problems you already solved"
meta_description: "L0, Ln, L1, Ln-1 looks like a strange shuffle until you notice it is two sequences interleaved. Then it is find the middle, reverse, merge."
---

![Day 26](HERO.png)

## 365 Days of LeetCode Challenge — Day 26/365

**[143. Reorder List](https://leetcode.com/problems/reorder-list/)** (Medium)

Turn `L0 -> L1 -> ... -> Ln-1 -> Ln` into `L0 -> Ln -> L1 -> Ln-1 -> L2 -> ...`.

This is the last day of the linked list arc, and it is the capstone in a literal sense: there is nothing new in it. Every piece is a problem from the last week.

## Read the target order again

Written out, the required order looks like an arbitrary shuffle. Written as two sequences, it is not:

- `L0, L1, L2, ...` walking forward from the start
- `Ln, Ln-1, Ln-2, ...` walking backward from the end

Interleaved, one from each, alternating.

So this is a merge. The list merged with its own reverse.

Once you see that, the only obstacle left is the one this arc has been circling all week: a singly linked list cannot be walked backward. Day 23 already answered it. Reverse the back half, and walking it forward walks the original backward.

## Three steps, all of them old

```go
func reorderList(head *ListNode) {
	mid := getMid(head)
	head2 := mid.Next

	mid.Next = nil
	head2 = reverseList(head2)

	head = mergeList(head, head2)
}
```

Find the middle. Cut and reverse. Merge alternately.

`getMid` is day 20. `reverseList` is day 19. `mergeList` is day 22 with the comparison removed, because the alternation here is unconditional rather than value-driven.

## Watching it run

`[1,2,3,4]`.

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

## The split has to stop a node early

`getMid` tests `fast.Next` and `fast.Next.Next`, not `fast` and `fast.Next`.

That is day 20's flagged variant, and this is the second problem in four days that needs it. The reason is the same both times: the node being returned is the one whose `Next` gets set to `nil`, so it must be the last node of the *first* half, not the first node of the second.

It also decides where the odd node goes. `[1,2,3,4,5]` has to produce `[1,5,2,4,3]`, with 3 last, and that only works if the first half is the longer one.

## One assignment the whole thing depends on

```go
mid.Next = nil
```

Without it the first half never terminates. It runs straight into the second half, which the merge is walking at the same time, and the two walks trample each other.

It is one line, it looks like tidying up, and the solution is broken without it.

## The merge saves two pointers, not one

```go
t, t1, t2 := a1, a1.Next, a2.Next
```

Day 19's lesson was: save the next pointer before the assignment that destroys it. Here there are two lists being relinked simultaneously, so there are two saved pointers.

The `if t1 != nil` guards exist because the two halves can differ in length by one, and advancing the shorter one's pointer past its end would dereference `nil`.

## It returns nothing, and that is correct

Every step relinks nodes that already exist. Nothing is allocated, nothing moves, and the node the caller's `head` points at is still the first node when the function finishes. There is no new head to hand back.

## Complexity

- **Time: O(n)**. Find the middle, reverse half, merge. Three linear passes.
- **Space: O(1)**. Pointers only.

## Builds on

- [Day 19: Reverse Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/reverse_linked_list/) — the in-place reversal, applied to the back half
- [Day 20: Middle of the Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/middle_of_the_linked_list/) — the fast/slow walk, in the variant that stops on the last node of the first half
- [Day 22: Merge Two Sorted Lists](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/merge_two_sorted_lists/) — taking alternately from two lists and relinking rather than copying

Full code and the step-by-step walkthrough:
[reorder_list](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/reorder_list/SOLUTION.md)

#DSA #LeetCode #Golang #LinkedList #TwoPointers #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
