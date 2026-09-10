## Solution walkthrough

The implementation is the `BSTIterator` struct plus `Constructor`, `Next` and `HasNext`, in
`main.go`.

![Example 1](images/1.png)

We'll trace the example tree `[7,3,15,null,null,9,20]`, whose in-order sequence is
`[3,7,9,15,20]`.

1. **The struct holds a flattened sequence and a cursor.**

   ```go
   type BSTIterator struct {
       inorder []int
       root    *TreeNode
       index   int
   }
   ```

   No stack, no current node, no parent pointers. Once the constructor has run, the tree
   isn't consulted again — `root` is kept but never read by `Next` or `HasNext`.

2. **The constructor does all the work.** `obj.inOrder(root)` walks the whole tree before
   returning, appending values left-node-right.

   ```go
   func (this *BSTIterator) inOrder(node *TreeNode) {
       if node == nil {
           return
       }
       this.inOrder(node.Left)
       this.inorder = append(this.inorder, node.Val)
       this.inOrder(node.Right)
   }
   ```

   This is a method rather than a free function, so it appends into `this.inorder` directly
   instead of threading a slice through as a parameter and return value the way earlier days
   did. The receiver is a pointer, so the append is visible on the struct.

   ![Step 1: the constructor flattens the entire tree](images/walkthrough-1.png)

3. **`index` starts at 0, pointing at the first value to hand out.** The problem describes
   a pointer "initialized to a non-existent number smaller than any element", with `next()`
   moving it and then returning. Same behaviour, expressed as an index that's read and then
   advanced.

   ![Step 2: after construction, index is 0 and the tree is done with](images/walkthrough-2.png)

4. **`Next` is a read and an increment.**

   ```go
   func (this *BSTIterator) Next() int {
       val := this.inorder[this.index]
       this.index++
       return val
   }
   ```

   O(1), with no traversal state to restore. There's no bounds check, which is safe because
   the problem guarantees `next()` is only called when a next value exists.

   ![Step 3: two calls have advanced the index to 2](images/walkthrough-3.png)

5. **`HasNext` is a bounds check.** `return this.index < len(this.inorder)`.

   ![Step 4: the index reaches the end and HasNext goes false](images/walkthrough-4.png)

6. **What this design trades.** Both methods that get called repeatedly are as simple as
   they can be, and that's the win. The cost is O(n) memory held for the iterator's whole
   lifetime, and a constructor that walks every node before the caller has asked for
   anything. A caller wanting the first three values of a million-node tree pays for all
   million.

   The problem's follow-up asks for O(h) memory and average O(1) `next()`. The standard
   answer keeps a stack of the leftmost spine: push left children on construction, and on
   each `next()` pop a node then push the left spine of its right child. Only one
   root-to-node path is ever in memory, and the amortised cost works out because each node
   is pushed and popped exactly once across the whole iteration.

**Complexity:** `Constructor` is O(n) time and O(n) space. `Next` and `HasNext` are both
O(1). Total memory held is O(n), against O(h) for the follow-up's stack version.
