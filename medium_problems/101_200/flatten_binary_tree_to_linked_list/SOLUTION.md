## Solution walkthrough

The implementation is `Flatten(root *TreeNode)` plus its helper `getPreOrderArray`, in
`flatten_binary_tree_to_linked_list.go`. `Flatten` returns nothing — it mutates the tree
it was handed, which is what the problem asks for.

![Example 1](images/1.jpg)

We'll trace it on the example above, `root = [1,2,5,3,4,null,6]`: node `1` has left child
`2` and right child `5`; `2` has children `3` and `4`; `5` has a right child `6`. The
expected flattened order is `1, 2, 3, 4, 5, 6`.

1. **The empty case, first.** `if root == nil { return }`. There's nothing to flatten and
   nothing to return, so this is the entire base case. It also means every line below can
   assume `root` is a real node — worth noting, because step 3 indexes `arr[0]` without
   checking, and that's only safe because a non-nil root guarantees at least one collected
   value.

2. **Collect the pre-order values.** `arr := make([]TreeNode, 0)` and then
   `getPreOrderArray(temp, &arr)`. The helper is a plain pre-order recursion: append the
   current node, recurse left, recurse right.

   ```go
   (*arr) = append((*arr), TreeNode{
       Val: root.Val,
   })
   getPreOrderArray(root.Left, arr)
   getPreOrderArray(root.Right, arr)
   ```

   Two things worth pointing at. It appends a `TreeNode` value carrying only `Val` — not
   the node itself, and not a pointer to it — so `Left` and `Right` on the collected
   copies are `nil` and are never read. Only the number matters. And it takes `arr` as a
   `*[]TreeNode` and assigns through the pointer rather than returning the slice, which is
   the other way to handle `append` reallocating: Day 8 threaded the slice through as a
   return value, this one hands down a pointer to the slice header so every call mutates
   the same one. Both work. The pointer version reads a little heavier at each call site.

   After this call, `arr` holds values `[1, 2, 3, 4, 5, 6]`.

3. **Reset the root.** `root.Val = arr[0].Val` and `temp.Left = nil`. The root keeps its
   identity — the caller is holding that pointer, so it has to be the same node — but
   its value is written from the collected list and its left pointer is cleared. Here
   `arr[0].Val` is `1`, which the root already held, but that isn't guaranteed in general
   and the assignment is what makes it not matter.

4. **Rebuild as a right-leaning chain.** The loop runs from index 1 to the end:

   ```go
   for i := 1; i < len(arr); i++ {
       y := new(TreeNode)
       y.Val = arr[i].Val
       temp.Right = y
       temp.Left = nil
       temp = temp.Right
   }
   ```

   Each iteration allocates a brand-new node, gives it the next collected value, hangs it
   off `temp.Right`, clears `temp.Left`, and steps `temp` forward onto the node just
   created. So the chain is built front to back, and `temp` is always the tail.

   Tracing it: `i=1` attaches a new node holding `2`; `i=2` attaches `3`; `i=3` attaches
   `4`; `i=4` attaches `5`; `i=5` attaches `6`. The result is
   `1 -> 2 -> 3 -> 4 -> 5 -> 6`, every link through `Right`, every `Left` nil.

   The `temp.Left = nil` inside the loop is worth a second look. On the first iteration
   it's re-clearing the root's `Left`, which step 3 already did. On every later iteration
   `temp` is a node this loop allocated moments ago, whose `Left` is already nil because
   `new(TreeNode)` zeroes it. So the line is never actually load-bearing. It's harmless,
   and it makes the loop body state its invariant explicitly — every node this loop leaves
   behind has a nil `Left` — rather than relying on the reader knowing what `new` does.

5. **The last node.** The loop ends with `temp` pointing at the node holding `6`, whose
   `Right` was never assigned and is therefore nil from `new(TreeNode)`. That nil is the
   end of the list, and it comes for free rather than needing a line to terminate it.

**Complexity:** O(n) time — the traversal visits every node once and the rebuild loop runs
once per collected value. O(n) space for `arr` and O(n) again for the newly allocated
chain, plus O(h) for the traversal's recursion stack, where h is the tree's height.
