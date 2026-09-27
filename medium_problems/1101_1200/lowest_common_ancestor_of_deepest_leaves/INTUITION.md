## Intuition

Yesterday's problem, word for word; LeetCode lists them as duplicates. What makes it worth a
second day is that the repo solves it differently, with a trick that needs an argument.

Yesterday's code kept *every* deepest path and compared them all. This one keeps only the
*first* deepest path the DFS meets. Each later leaf at the same depth is compared against that
first path, and their lowest common ancestor becomes the answer, overwriting the previous one.
A deeper leaf resets everything.

Why is comparing with just the first path enough? A depth-first search visits every subtree as
one contiguous run. So if the first deepest leaf and the latest one both sit under some node,
every leaf the DFS visited between them sits under it too. The LCA of the first and the latest
deepest leaf is therefore the LCA of all the deepest leaves seen so far. I didn't take that on
faith: the solution agrees with the postorder version from yesterday on five thousand random
trees.

## Builds on

- [Day 134: Smallest Subtree with all the Deepest Nodes](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/801_900/smallest_subtree_with_all_the_deepest_nodes/) — the same question, solved there by keeping every deepest path and comparing them all

The paths stored here exclude the leaf itself: `path` is the list of ancestors on the way
down, and `newPath` (a fresh copy plus the current node) is what the children receive. A
single-node tree never enters the tie branch, and `resFinal` stays at the root, which is the
correct answer.

Like yesterday's code, this uses package-level variables for its state. And copying the path
at every node costs O(n · h). The postorder pass from yesterday's post is still the O(n)
answer.

**Complexity:**
- Time: O(n · h) for the path copies.
- Space: O(n · h) in the worst case across the recursion.
