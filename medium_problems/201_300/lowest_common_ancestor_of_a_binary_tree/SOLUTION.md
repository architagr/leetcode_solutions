## Solution walkthrough

The implementation is `lowestCommonAncestor(root, p, q *TreeNode) *TreeNode` wrapping the
recursive `f`, in `main.go`. `f` returns two things: the answer once found, and a count of
how many targets its subtree contains.

![Example 1](images/1.png)

We'll trace `[3,5,1,6,2,0,8,null,null,7,4]` with `p = 5` and `q = 1`. Expected answer: `3`.

1. **Two return values, doing different jobs.** `f` returns `(n *TreeNode, count int)`.
   `count` is what a parent needs in order to add up its side; `n` is the answer, non-nil
   only once some node has seen both targets. Yesterday's BST version needed neither — it
   walked down and returned directly.

2. **The base case reports nothing found.** `if root == nil { return nil, 0 }`.

   ![Step 1: a leaf that is not a target returns 0](images/walkthrough-1.png)

3. **Recurse left, and bail out early if it already found the answer.**

   ```go
   ln, lcount := f(root.Left, p, q)
   if ln != nil {
       n = ln
       return
   }
   ```

   This early return is what keeps the answer *lowest*. Without it, the recursion would
   carry on, and every ancestor above the real LCA would also see a count of two and
   overwrite the answer with itself — producing a common ancestor, but the highest one
   rather than the lowest.

   ![Step 2: the subtree under 2 holds neither target](images/walkthrough-2.png)

4. **Same for the right side.** `rn, rcount := f(root.Right, p, q)` with the same guard.

5. **The node counts itself.**

   ```go
   if root.Val == p.Val || root.Val == q.Val {
       lcount++
   }
   count = lcount + rcount
   ```

   Folding the node's own match into `lcount` rather than a separate variable is what makes
   "a node can be a descendant of itself" work. At node `5`, which is `p`, the subtrees
   report `0` and `0`, and the node itself adds one, giving `1`.

   ![Step 3: node 5 is a target and counts toward its own total](images/walkthrough-3.png)

   If `p` were an ancestor of `q`, this is the mechanism that makes `p` the answer: one for
   being `p`, one from the subtree holding `q`, total two, right there.

6. **Two means this node is the answer.** `if count == 2 { n = root }`.

   ![Step 4: the right subtree reports 1 for node 1](images/walkthrough-4.png)

   At the root, `lcount` is `1` from the left subtree and `rcount` is `1` from the right.
   Two, so the root is the answer — and because a post-order traversal reaches nodes
   bottom-up, the first node to hit two is the lowest one that can.

   ![Step 5: the root reaches 2 and becomes the answer](images/walkthrough-5.png)

7. **Comparison is on `Val`, not pointer identity.** Same choice as yesterday, and it
   relies on the problem's guarantee that values are unique. With duplicates, a count of
   two could be reached by two nodes that merely share a value with the targets; comparing
   pointers would be the fix.

**Complexity:** O(n) worst case — every node may be visited. The early returns prune work
once the answer is found but don't change the bound. Space is O(h) for the recursion
stack. Compare yesterday's O(h) *time*: that's the whole cost of losing the ordering.
