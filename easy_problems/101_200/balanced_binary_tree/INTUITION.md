# Intuition

## The problem

A binary tree is "height-balanced" if, at **every** node, the heights of its left and
right subtrees differ by at most 1. We need to check this for the whole tree, not just
the root — a tree can look balanced at the top while a node buried three levels down
violates the rule.

## The naive approach (and why it's wasteful)

The first idea most people reach for is: write a `height(node)` function, then for every
node compare `height(node.Left)` and `height(node.Right)`, recursing into children to
check them too.

That works, but it's wasteful: computing `height()` for a node already walks that node's
entire subtree. If you then call `height()` again separately on the left and right
children while checking balance top-down, you end up recomputing the height of the same
nodes over and over — once for every ancestor above them. That's O(n) work repeated at
each of O(n) levels, giving **O(n²)** in the worst case (e.g. a skewed tree).

## The key insight

Height and balance-checking are naturally computed **bottom-up**, and they can be
computed **in the same pass**. A node's height only depends on its children's heights,
and a node's balance only depends on its children's heights too. So instead of two
separate traversals (one for height, one for balance), do a single post-order traversal
that returns *both* pieces of information at once for every node:

- how tall is this subtree, and
- is this subtree (and everything under it) still balanced so far?

If a subtree deep down is already unbalanced, there's no point computing anything above
it — the whole tree is unbalanced. So the moment an imbalance is found, the recursion can
short-circuit and bail out immediately instead of finishing the rest of the tree.

## Why this technique applies

This is the classic **"augment the return value"** pattern for tree recursion: instead of
a function that only answers one question (e.g. "what's the height?"), have it return a
small bundle of everything a parent needs to make its own decision. Each node computes
its own height using only its children's already-computed heights (post-order — children
before parent), and each node validates its own balance using only its children's
already-computed ok-ness. No node ever needs to look further down than its immediate
children, and no subtree's height is ever computed more than once.

## Complexity

- **Time: O(n)** — every node is visited exactly once, and each visit does O(1) work
  besides the two recursive calls.
- **Space: O(h)** — where `h` is the height of the tree, for the recursion call stack.
  That's O(log n) for a balanced tree and O(n) for a completely skewed one (worst case).
