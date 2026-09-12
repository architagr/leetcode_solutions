# Reorder List — intuition

## Builds on

- [Day 19: Reverse Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/reverse_linked_list/) — the in-place reversal, applied to the back half
- [Day 20: Middle of the Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/middle_of_the_linked_list/) — the fast/slow walk, in the variant that stops on the last node of the first half
- [Day 22: Merge Two Sorted Lists](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/merge_two_sorted_lists/) — taking alternately from two lists and relinking rather than copying

## The problem in one line

Turn `L0 -> L1 -> ... -> Ln-1 -> Ln` into `L0 -> Ln -> L1 -> Ln-1 -> L2 -> ...`.

## Reading the target order properly

Write the target out and it looks like a strange shuffle. Read it again and it is two sequences interleaved:

- `L0, L1, L2, ...` walking forward from the start
- `Ln, Ln-1, Ln-2, ...` walking backward from the end

So the reordering is a merge. One node from the front, one node from the back, alternating.

Once you see that, the only obstacle left is the familiar one: a singly linked list cannot be walked backward. And the fix is the one from day 23 — reverse the back half, and walking it forward walks the original backward.

## The three steps

1. Find where the middle is.
2. Cut there, and reverse the second half.
3. Merge the two halves alternately.

Every one of those is a problem from earlier in this arc. This is the capstone in the literal sense: there is nothing new in it, and it is only hard if the previous six days did not happen.

## Why the split has to land on the first half's last node

The fast/slow walk here uses the variant that stops one node early, the one day 20 flagged and day 23 used. It has to, because the node whose `Next` gets set to `nil` is the last node of the first half. Stop a node later and you cut in the wrong place.

For an odd-length list this also puts the extra node in the first half, which is what the expected output wants: `[1,2,3,4,5]` becomes `[1,5,2,4,3]`, where 3 ends up last.

## Why the cut matters

Without `mid.Next = nil`, the first half does not terminate. It runs straight into the second half, which the merge is also walking, and the two loops walk over each other. The cut is one assignment and the whole thing is wrong without it.

## The signature returns nothing

`reorderList` returns no value, and it does not need to. Every step relinks existing nodes, so the node the caller's `head` points at is still the first node afterwards. Nothing is allocated and nothing moves.

## Complexity

- **Time: O(n).** Find the middle, reverse half, merge. Three linear passes.
- **Space: O(1).** Pointers only.
