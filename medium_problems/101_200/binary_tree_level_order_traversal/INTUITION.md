## Intuition

"Level order" is the textbook use for BFS: a queue, drain one level, move to the next.
This solution doesn't use one. It's a depth-first recursion that produces level-ordered
output anyway, and the reason it works is the same one from Day 14.

## Builds on

- [Day 14: Average of Levels in Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/average_of_levels_in_binary_tree/) — the same idea, that a DFS can answer a per-level question by keying its accumulator on depth instead of on visit order

A DFS visits nodes branch by branch. It goes all the way down the left side before it
touches anything on the right, so the order it arrives at nodes has nothing to do with
levels. That sounds fatal for a problem whose entire output is grouped by level.

It isn't, because the grouping doesn't depend on arrival order. Each recursive call
carries its own depth as a parameter, and every node appends into `arr[level]` — the slot
for its own depth. Whichever node reaches depth 2 first, second or last, they all land in
`arr[2]`, and they land there in left-to-right order because the recursion always descends
into `Left` before `Right`.

That last clause is the part that makes the output correct rather than merely grouped. The
problem wants each level read left to right; a DFS that visits `Left` first appends the
leftmost node of every level before anything to its right, at every depth simultaneously.

The one piece of bookkeeping is growing the outer slice. `arr` starts empty, so the first
time the recursion reaches a new depth there's no slot to append into yet, and the
`if (len(arr) - 1) < level` check creates one. Every later visit to that depth finds the
slot already there.

What you get in exchange for not writing BFS is a shorter function and no explicit queue.
What you give up is the level boundaries as a thing the code knows about — BFS knows when
a level ends, this doesn't, and for a problem that needed that (a zigzag, say, or stopping
early at a given depth) BFS would be the better shape.

**Complexity:**
- Time: O(n) — every node is visited exactly once and does O(1) work.
- Space: O(h) for the recursion stack, where h is the tree's height, plus O(n) for the
  output itself. Note the recursion depth here is the tree's height, where BFS's peak
  memory would instead be the width of the widest level.
