## Intuition

The "depth" of a binary tree is just its number of levels. If you could see the whole
tree drawn out, the answer is the number of rows before the tree runs out of nodes.

The natural way to count rows is a **level-order (breadth-first) traversal**: visit all
nodes at depth 1, then all nodes at depth 2, then depth 3, and so on, counting how many
full rounds you complete before there's nothing left to visit.

A BFS naturally processes nodes queue-by-queue, but a plain queue doesn't tell you
*where one level ends and the next begins* — you'd just see a flat stream of nodes. The
trick used here is to push a `nil` "sentinel" value right after the root, marking the
end of the current level. Every time you pop a `nil` off the queue, you know you've
just finished a full level: increment the depth counter, and (if there are still real
nodes left in the queue) push a fresh `nil` sentinel to mark the end of the *next*
level.

This turns "how many levels are there" into "how many sentinels did I pop", which is
easy to track with a single counter and no extra bookkeeping about level sizes.

**Complexity:**
- Time: O(n) — every node is enqueued and dequeued exactly once.
- Space: O(n) — in the worst case (a very wide, shallow tree) the queue can hold up to
  roughly half the nodes at once, so the queue size is O(n) in the worst case.
