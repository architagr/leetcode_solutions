## Intuition

The most common way to solve this problem is a breadth-first search: process the tree
one level at a time with a queue, average each level's node values as you drain it, and
move to the next level. This solution takes a different route — a **depth-first search**
that still produces the same per-level averages, by carrying the current depth as a
parameter and using it as an index into two running-total arrays.

The key trick is: a DFS doesn't visit nodes level by level, it visits them branch by
branch. So instead of finishing one level before starting the next, this solution keeps
a running `sum` and `count` **per level**, stored as `result[level]` and `count[level]`,
and lets every visit — no matter which branch of the tree it comes from — add its value
into the slot for its own depth. Whichever node happens to reach depth 2 first, second,
or last, they all land in `result[2]` and `count[2]` and get summed together correctly,
because the accumulation is keyed by level, not by visit order.

The two arrays start empty and grow lazily: the first time the recursion reaches a new
depth, it appends a fresh `0.0` slot to both `result` and `count` for that level. Every
subsequent visit to that same depth reuses the existing slot instead of appending again.
Once the whole tree has been visited, `sum` has become each level's total and `count`
has become each level's node count, so a final single pass divides `result[i] /=
count[i]` to convert every level's running sum into that level's average.

**Complexity:**
- Time: O(n) — every node is visited exactly once by the recursion, and the final
  division pass is O(number of levels), which is at most O(n).
- Space: O(h + L) — O(h) for the recursion stack (h = tree height), plus O(L) for the
  `result`/`count` slices, where L is the number of levels (L <= h + 1, so this is
  effectively O(h)). The output array itself is also O(L).
