## Solution walkthrough

The implementation is `maxDepth(root *TreeNode) int` in `main.go`.

1. **Base case.** If `root == nil`, the tree is empty, so its depth is `0`. Handled
   immediately at the top.

2. **Seed the queue with a level-end sentinel.** `queue` starts as `[root, nil]`. The
   `nil` right after `root` marks "end of level 1" before any traversal happens — since
   level 1 only ever contains the root itself.

3. **`pop`/`push` helpers.** These are small closures over `queue`, kept local to
   `maxDepth` since they're only meaningful in the context of this one traversal:
   - `pop()` takes the front element off the queue (`queue[0]`) and re-slices `queue`
     past it. This is a plain FIFO dequeue.
   - `push(node)` appends to the back of `queue`.

4. **The main loop.** While the queue isn't empty:
   - Pop the front node.
   - **If it's the sentinel (`node == nil`)**: this level is done. `max++` records that
     another full level was completed. Then, only if the queue still has real nodes
     left (`len(queue) > 0`), push a new `nil` sentinel — this marks where the *next*
     level will end. Skipping this when the queue is empty is what stops the loop from
     spinning forever pushing sentinels after the last real level.
   - **If it's a real node**: push its non-nil children (`Left`/`Right`) onto the back
     of the queue, so they'll be processed as part of the next level.

5. **Return `max`** once the queue is fully drained — it now holds exactly the count of
   completed levels, i.e. the tree's maximum depth.

Walking it through the test case (`[3,9,20,null,null,15,7]`, expected depth `3`):
- Queue starts `[3, nil]`. Pop `3` (real): push its children `9, 20` → queue `[nil, 9, 20]`.
- Pop `nil` (level 1 done, `max=1`): queue not empty, push new sentinel → `[9, 20, nil]`.
- Pop `9` (real, no children): queue unchanged → `[20, nil]`.
- Pop `20` (real): push its children `15, 7` → `[nil, 15, 7]`.
- Pop `nil` (level 2 done, `max=2`): push new sentinel → `[15, 7, nil]`.
- Pop `15`, then `7` (both real, no children) → queue `[nil]`.
- Pop `nil` (level 3 done, `max=3`): queue is now empty, so no new sentinel is pushed.
- Loop ends, return `max = 3`. ✓
