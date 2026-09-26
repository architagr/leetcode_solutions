## Intuition

Any ancestor/descendant pair lies on a single root-to-leaf path. So the question becomes:
on which path is the gap between two of its nodes the widest? And the widest gap on a path
is just its largest value minus its smallest.

That turns a pairs problem into the Count Good Nodes pattern from two days ago, carrying
two values down instead of one: the max and the min seen on the path so far. Each call
folds its own value into both and hands them to its children.

## Builds on

- [Day 87: Count Good Nodes in Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1401_1500/count_good_nodes_in_binary_tree/) — passing the running max down as a parameter, so each branch keeps its own path's value; this carries the min alongside it

The detail I liked is *where* the answer gets recorded. The code only computes
`ancestorMax - ancestorMin` at leaves. That's enough, because going down a path the max
can only go up and the min can only go down, so the gap at the bottom of a path is at least
as large as the gap anywhere along it. Every node sits on some root-to-leaf path, so every
pair gets covered by at least one leaf's check.

The seeds are `math.MinInt` and `math.MaxInt`, which the root immediately replaces with its
own value. And the problem's "different nodes" rule looks like it needs a special case,
but it doesn't: a max and min coming from the same node give a gap of 0, which never beats
a real answer.

One thing I'd change: `res` is a package-level variable. `maxAncestorDiff` resets it
before each run, so sequential calls are fine, but two calls running at once on different
goroutines would share it and trample each other. Having `diff` return the best gap in its
subtree, or keeping `res` in a closure, removes the shared state.

**Complexity:**
- Time: O(n), each node visited once.
- Space: O(h) for the recursion stack.
