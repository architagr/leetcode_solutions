## Intuition

An iterator promises an ordering and hides when the work happens. That second half is the
only real decision in this problem, and the solution here takes the blunt end of it: do
everything up front.

The ordering is in-order, which for a BST is ascending — the property Days 29, 30 and 32
all leaned on.

## Builds on

- [Day 47: Minimum Absolute Difference in BST](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/minimum_absolute_difference_in_bst/) — in-order on a BST yields sorted values, which is exactly the sequence this iterator has to produce
- [Day 46: Find Mode in Binary Search Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/find_mode_in_binary_search_tree/) — the same property again, with the same choice to collect first and compute afterwards

So the constructor walks the tree once in-order, flattens it into a slice, and keeps an
index. After that `Next` is a slice read and an increment, and `HasNext` is a bounds check.
Both are O(1) with no traversal state to maintain at all.

What that buys is the simplest possible implementation of the two methods that get called
repeatedly. What it costs is O(n) memory held for the iterator's whole lifetime, and a
constructor that walks the entire tree before the caller has asked for a single value. If
the caller only ever wants the first three of a million nodes, this solution did all the
work anyway.

The problem's follow-up asks for exactly that improvement: O(h) memory and average O(1)
`next()`. The standard answer is a stack holding the path down the leftmost spine — push
left children on construction, and on each `next()` pop a node, then push the left spine of
its right child. That keeps only one root-to-node path in memory, and the amortised cost
works out because every node is pushed and popped exactly once across the whole iteration.

Neither is wrong. Precomputing is a fine answer when the tree is small or the caller will
consume most of it; the stack version is the one to reach for when either of those isn't
true. Worth being able to say which you wrote and why.

**Complexity:**
- `Constructor`: O(n) time, O(n) space.
- `Next`: O(1). `HasNext`: O(1).
- Total space held: O(n). The follow-up's stack version is O(h).
