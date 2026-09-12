# Remove Nth Node From End of List — intuition

## Builds on

- [Day 20: Middle of the Linked List](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/middle_of_the_linked_list/) — the same problem shape: a position defined relative to the end of a list you can only walk forward

## The problem in one line

Remove the nth node counting from the end, and return the head.

## The awkward bit

"nth from the end" is a position measured from a place you cannot start at. A singly linked list gives you the head and nothing else, and the end is only reachable by walking there.

So the first job is turning "nth from the end" into a position measured from the front, which is `count - n`. Once you have that number, the problem is ordinary.

## Two passes, and why that is a fine answer

Walk the list to learn its length. Walk it again to `count - n`. Unlink. Done.

That is what this solution does, and it is O(n) time and O(1) space. The follow-up asks whether it can be done in one pass, and it can, but I want to be clear that the follow-up is asking for a nicer solution, not a faster one. Two passes over n nodes is still O(n), and the one-pass version does not touch fewer nodes overall.

## The thing you actually have to get right

To remove a node from a singly linked list you need the node *before* it, because removal means pointing the previous node past this one. A node cannot remove itself; it has no way to reach whatever points at it.

So the second walk deliberately stops one short of the target. That off-by-one is the entire problem, and it is where most wrong submissions live.

## Removing the head

If `n` equals the length, the node to remove is the head, and there is no previous node to relink. That case gets handled separately by returning `head.Next`.

This is exactly the situation day 22's dummy node exists to erase. Putting a dummy in front of the head would make the head an ordinary node with a predecessor like any other, and the special case would disappear. This solution handles it with an explicit branch instead, which is a legitimate choice and a slightly longer one.

## Complexity

- **Time: O(n).** One pass to count, at most one more to reach the target.
- **Space: O(1).** Two integers and a pointer.
