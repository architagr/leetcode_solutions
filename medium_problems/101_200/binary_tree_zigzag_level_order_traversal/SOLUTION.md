## Solution walkthrough

The implementation is `zigzagLevelOrder(root *TreeNode) [][]int` plus the helper
`reverseArr`, in `main.go`. The traversal is BFS over an explicit queue, with two closures
(`pop` and `push`) standing in for queue operations.

![Example 1](images/1.jpg)

We'll trace `[3,9,20,null,null,15,7]`: root `3` with children `9` and `20`, and `20` with
children `15` and `7`. Expected output: `[[3],[20,9],[15,7]]` — level 1 reversed, the
others not.

1. **The empty tree returns early.** `if root == nil { return [][]int{} }`. An empty slice
   rather than nil, matching what the tests expect.

2. **Seed the queue with the root and a sentinel.** `queue = append(queue, root, nil)`.
   That `nil` is the level marker: everything ahead of it belongs to level 0. This is the
   same trick as Day 1's depth counter, reused for a different purpose.

   ![Step 1: root and the sentinel that closes level 0](images/walkthrough-1.png)

3. **Popping a real node collects it and queues its children.**

   ```go
   arr = append(arr, node.Val)
   if node.Left != nil {
       push(node.Left)
   }
   if node.Right != nil {
       push(node.Right)
   }
   ```

   Children go on the back in left-then-right order, so the next level arrives left to
   right regardless of what will be done to it afterwards. `arr` is the current level
   under construction.

   ![Step 2: 3 is collected, its children queued, and the sentinel surfaces](images/walkthrough-2.png)

4. **Popping the sentinel closes the level.** This branch is where everything interesting
   happens:

   ```go
   if node == nil {
       if len(queue) > 0 {
           push(nil)
       }
       x := make([]int, len(arr))
       copy(x, arr)
       if !leftToRight {
           x = reverseArr(x)
       }
       result = append(result, x)
       arr = make([]int, 0)
       leftToRight = !leftToRight
       continue
   }
   ```

   The re-push is guarded on `len(queue) > 0`, and that guard is load-bearing: without it
   the last sentinel would be re-added to an otherwise empty queue and the loop would spin
   on it forever.

   ![Step 3: the flag flips and a fresh sentinel closes level 1](images/walkthrough-3.png)

5. **The copy is not optional.** `x := make([]int, len(arr))` then `copy(x, arr)`. `arr`
   gets reused for the next level, and `reverseArr` mutates in place. Appending `arr`
   directly would put the same backing array into `result` several times, and the answer
   would come out as repeats of the last level.

6. **Reverse every second level.** `leftToRight` starts `true` and flips at each boundary,
   so level 0 is stored as collected, level 1 reversed, level 2 as collected. Here `arr`
   for level 1 is `[9,20]`, and the flag is false, so it's stored as `[20,9]`.

   ![Step 4: level 1 is reversed before being stored](images/walkthrough-4.png)

7. **The last level drains without a new sentinel.** After `15` and `7` are collected the
   queue holds only the sentinel; popping it finds `len(queue) == 0`, so no replacement
   goes on, the level is stored, and the loop ends.

   ![Step 5: the final level is stored and the queue empties](images/walkthrough-5.png)

`reverseArr` itself is a two-pointer swap with `l <= r`. The `<=` rather than `<` means an
odd-length slice swaps its middle element with itself once — harmless, and it keeps the
loop condition simple.

**Complexity:** O(n) time — every node is pushed and popped exactly once, and the
reversals total O(n) across the whole tree because each element is swapped at most once.
Space is O(w) for the queue, where w is the width of the widest level, plus O(n) for the
output. That width is BFS's real memory cost, and it's the opposite trade to Day 15's
recursion, whose peak was the tree's height.
