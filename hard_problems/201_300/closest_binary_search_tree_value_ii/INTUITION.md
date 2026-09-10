## Intuition

The second hard of the challenge, and like the first one it isn't hard because of the
traversal. It's hard because of a claim you have to notice and trust.

Flatten the BST in-order and you get the values in ascending order — the property this
batch has now used five times. Yesterday's iterator did exactly that flattening.

## Builds on

- [Day 33: Binary Search Tree Iterator](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_search_tree_iterator/) — flattening a BST in-order into a slice, which is the first half of this solution verbatim

Now the claim: **in a sorted array, the k values closest to a target are always
contiguous.**

That's worth a second, because it's what turns the problem from "pick k things" into "pick
a window." Suppose the answer skipped some value `v` and instead included a value `w`
further from the target, with `v` sitting between the closest element and `w`. Because the
array is sorted, `v` lies between them in value too, so `v` is at least as close to the
target as `w` is. Swapping `w` for `v` gives an answer at least as good. So a
non-contiguous answer is never strictly better than the contiguous one covering the same
span, and the k closest values can always be taken as a run.

Once that's settled, the algorithm falls out. Find the single closest value, then grow
outward from it, always taking whichever neighbour — the one to the left or the one to the
right — is nearer the target. Stop at k. Because distance increases monotonically as you
move away from the target in either direction, a greedy choice at each step is safe.

Two pointers walking outward from a centre, repeatedly taking the smaller of two
candidates, is the merge step of a merge sort read backwards — the code's own comment says
as much.

The closest index doesn't need a separate search. The traversal is already visiting every
value in order, so it tracks the running minimum difference as it goes and records where it
was found.

The two trailing loops handle running out of array. If one side is exhausted before k
values are collected, the other side supplies the rest, and only one of the two can ever
execute.

**Complexity:**
- Time: O(n) for the traversal plus O(k) for the expansion.
- Space: O(n) for the flattened slice plus O(h) for the recursion stack. The known better
  answer is O(k + h): run two stack-based iterators outward from the target, one going
  backwards and one forwards, and merge k values off them — yesterday's follow-up, applied
  twice. This solution takes the simpler route.
