## Solution walkthrough

The implementation is `maxLevelSum(root *TreeNode) int` in `main.go`. It's a BFS over an
explicit queue with `push`/`pop` closures, carrying four pieces of state rather than
storing any levels.

![Example 1](images/1.jpg)

We'll trace `[1,7,0,7,-8,null,null]`: root `1` with children `7` and `0`; the left `7` has
children `7` and `-8`. Level sums are `1`, `7`, and `-1`, so the answer is level `2`.

1. **Four variables, no level storage.** `currentLevel` (1-based, matching the problem),
   `maxLevel`, `maxSum`, and a running `sum`. Nothing accumulates the nodes of a level —
   once a level's total is folded in, its values are gone.

2. **`maxSum` starts at `root.Val`, not zero.** This matters because node values can be
   negative. On a tree where every level sums negative, a zero starting point would never
   be beaten and `maxLevel` would keep its initial value without any comparison having
   decided it. Seeding from the root guarantees the first real comparison is against an
   actual level total.

   ![Step 1: state before the walk, seeded from the root](images/walkthrough-1.png)

3. **Seed the queue with the root and a sentinel.** `push(root)` then `push(nil)`. Same
   marker as Day 30: everything ahead of the `nil` is the current level.

4. **A real node adds and enqueues.**

   ```go
   sum += node.Val
   if node.Left != nil {
       push(node.Left)
   }
   if node.Right != nil {
       push(node.Right)
   }
   ```

   Only non-nil children are pushed, so the only `nil` ever in the queue is a sentinel —
   which is what lets `node == nil` mean "level boundary" unambiguously.

5. **The sentinel closes the level and does the comparison.**

   ```go
   if node == nil {
       if sum > maxSum {
           maxLevel = currentLevel
           maxSum = sum
       }
       currentLevel++
       sum = 0
       if len(queue) > 0 {
           push(nil)
       }
       continue
   }
   ```

   Level 1's total is `1`, which is not strictly greater than the seeded `maxSum` of `1`,
   so nothing changes — correctly, since level 1 is already the answer-so-far.

   ![Step 2: level 1 drains, the comparison is not strict-greater](images/walkthrough-2.png)

6. **The strict `>` is the tie-break.** The problem asks for the *smallest* level whose sum
   is maximal. There's no code for that anywhere; it's this one character. A later level
   that merely matches the best doesn't replace it, so `maxLevel` keeps the earlier value.
   Change it to `>=` and the function still returns a level with the maximum sum — just the
   wrong one whenever two levels tie.

   Level 2 sums to `7`, which does beat `1`, so `maxLevel` becomes `2`.

   ![Step 3: level 2 beats the best and takes maxLevel](images/walkthrough-3.png)

7. **Negative levels lose without special handling.** Level 3 sums to `-1`, not greater
   than `7`, so nothing moves.

   ![Step 4: level 3 loses the comparison](images/walkthrough-4.png)

8. **The re-push guard ends the loop.** As in Day 30, the fresh sentinel only goes on if
   the queue still holds nodes. Without the guard the final sentinel would be re-added
   forever.

   ![Step 5: the queue empties and the answer is level 2](images/walkthrough-5.png)

Note the function returns `maxLevel`, not `maxSum` — the problem asks which level, not how
much. Easy to get backwards on a first read.

**Complexity:** O(n) time — every node is pushed and popped exactly once and contributes a
single addition. Space is O(w) for the queue, where w is the width of the widest level.
The output is a single int, and no level is ever materialised.
