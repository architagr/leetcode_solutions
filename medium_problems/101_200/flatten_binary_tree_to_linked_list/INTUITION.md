## Intuition

The shape being asked for is a linked list wearing `TreeNode`: every node's `Left` is
`nil`, every node's `Right` points at whatever comes next, and "next" means the order a
pre-order traversal would visit them in.

That last clause is the whole problem, and it's worth reading twice. The order isn't
sorted, it isn't level by level, it's specifically root, then left subtree, then right
subtree. Which means the traversal that produces the answer is one already written on
Day 8.

## Builds on

- [Day 8: Binary Tree Preorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/) — the exact visit order this problem's output is defined by, and the slice-threading that collects it

Once you see that, the problem splits into two pieces that don't interact:

1. Walk the tree in pre-order and record the values in that order.
2. Walk that flat list and rebuild the tree as a right-leaning chain.

Neither half is hard on its own. Keeping them apart is what makes this version easy to
read, and it's a deliberate trade rather than an oversight — the problem's follow-up asks
for an in-place version with O(1) extra space, and that one is genuinely fiddly: you're
rewiring the pointers you're still navigating by, so you have to stash the right subtree
before you overwrite `Right` with the left one, then walk to the end of the newly attached
run to hang the old right subtree off it.

This solution takes the O(n) space instead and gets code you can read in one pass. Worth
knowing both exist and which one you wrote.

One detail in the implementation that isn't obvious from the description: the second phase
doesn't reuse the original nodes. It overwrites the root's value from the collected list,
then allocates a fresh `TreeNode` for every subsequent value and chains those. The
original nodes below the root become garbage. That's why the traversal collects values
rather than pointers — nothing from the first phase survives into the second except the
numbers.

**Complexity:**
- Time: O(n) — one traversal to collect, one pass to rebuild, each visiting every node once.
- Space: O(n) for the collected slice, plus O(h) for the recursion stack during the
  traversal, plus O(n) for the freshly allocated chain. The follow-up's in-place version
  is what removes the first and last of those.
