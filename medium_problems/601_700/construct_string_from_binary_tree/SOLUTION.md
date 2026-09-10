## Solution walkthrough

The implementation is `Tree2str(root *TreeNode) string`, which is a one-line wrapper over
the recursive `parse`, in `construct_string_from_binary_tree.go`. All the work is in
`parse`.

![Example 2](images/2.jpg)

We'll trace `[1,2,3,null,4]` — the second example, chosen over the first because it's the
one that exercises the empty-parentheses rule. Node `1` has children `2` and `3`; node `2`
has no left child and a right child `4`. Expected output: `"1(2()(4))(3)"`.

1. **The base case returns an empty string, not a placeholder.** `if node == nil { return
   "" }`. This is worth pausing on, because it's doing more work than a base case usually
   does. It isn't only stopping the recursion — the empty string it returns is what turns
   into the `()` the formatting rules demand, once a caller wraps it in parentheses. The
   rule and the base case are the same mechanism.

2. **Every node contributes its own value first.** `s := strconv.Itoa(node.Val)`. That's
   the "pre" in pre-order: the node is written before either subtree is looked at.

   ![Step 1: parse(1) writes its value and opens the left pair](images/walkthrough-1.png)

3. **Childless nodes stop here.** The whole parentheses block is guarded by
   `if node.Right != nil || node.Left != nil`. A leaf skips it entirely and returns just
   its digits, which is why `4` becomes `"4"` and not `"4()"`.

4. **The left pair is unconditional once any child exists.**

   ```go
   s += "(" + parse(node.Left) + ")"
   ```

   No test for whether `Left` is nil. If it exists, `parse` returns its subtree's string
   and this produces `(2(4))`-style output. If it doesn't, `parse(nil)` returns `""` and
   the same line produces `()` — exactly the placeholder the exception calls for.

   That's the whole trick. There is no branch anywhere in this function testing for "right
   child but no left child", because the asymmetry in the code already produces it.

   ![Step 2: parse(2) has only a right child, so parse(nil) yields the empty pair](images/walkthrough-2.png)

5. **The right pair is conditional.**

   ```go
   if node.Right != nil {
       s += "(" + parse(node.Right) + ")"
   }
   ```

   The mirror case needs no placeholder. A node with a left child and no right child emits
   `(left)` and stops, and nothing about that string is ambiguous — there's no later
   position for the reader to be confused about. So this branch simply doesn't fire.

   ![Step 3: parse(4) is a leaf, so node 2 closes as "2()(4)"](images/walkthrough-3.png)

6. **Unwinding.** `parse(4)` returns `"4"`, so `parse(2)` finishes as `"2" + "()" + "(4)"`
   = `"2()(4)"`. `parse(3)` is a leaf and returns `"3"`. Back at the root,
   `"1" + "(2()(4))" + "(3)"` gives `"1(2()(4))(3)"`.

   ![Step 4: the root closes both pairs for the final string](images/walkthrough-4.png)

   Run the first example through the same code and the asymmetry shows from the other
   side: `[1,2,3,4]` gives `"1(2(4))(3)"`, where node `2` has a left child and no right
   one, so step 5 never fires and no placeholder appears.

**Complexity:** O(n) node visits, each converting one value. The catch is `+=` on strings
— Go strings are immutable, so every concatenation allocates and copies what came before,
which on a skewed tree degrades to O(n^2) total copying. Threading a `strings.Builder`
through the recursion instead would hold it at O(n). Space is O(h) for the recursion
stack plus O(n) for the string itself.
