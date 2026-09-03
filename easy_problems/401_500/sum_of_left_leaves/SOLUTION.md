## Solution walkthrough

The implementation is `sumOfLeftLeaves(root *TreeNode) int` in `main.go`.

![Example 1](images/1.jpg "Example1")

We'll trace it on the example above, `root = [3,9,20,null,null,15,7]`.

1. **Base case.** `if root == nil { return 0 }` — an empty subtree contributes nothing.

2. **Recurse into both subtrees first.** `l := sumOfLeftLeaves(root.Left)` and
   `r := sumOfLeftLeaves(root.Right)` collect whatever left-leaf sums already exist
   deeper in each subtree, independent of the current node.

3. **Check if THIS node's left child is a left leaf.**
   `if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil` — this is
   the key check, and it's evaluated from `root`'s perspective, looking at its `Left`
   child: does that child exist, and does it have no children of its own (i.e. is it a
   leaf)? If so, it's a left leaf by definition (it's `root`'s *left* child, and it's a
   leaf), so its value is added: `l += root.Left.Val`.

4. **Return the combined total.** `return l + r` — the (possibly augmented) left-subtree
   sum plus the right-subtree sum.

Walking it through `root = [3,9,20,null,null,15,7]` (`9` is a leaf and `3`'s left
child; `15` is a leaf and `20`'s left child; `7` is a leaf but `20`'s *right* child, so
it doesn't count):
- `sumOfLeftLeaves(3)`: recurse left into `9`, recurse right into `20`.
  - `sumOfLeftLeaves(9)`: both children nil, `l=0, r=0`. `9`'s own left child is nil,
    so no addition here. Returns `0`. (Note: `9` being a leaf itself is only detected
    and counted by its *parent*, `3` — not by this call.)
  - `sumOfLeftLeaves(20)`: recurse left into `15`, recurse right into `7`.
    - `sumOfLeftLeaves(15)` → `0` (leaf, but detected by its parent `20`, not here).
    - `sumOfLeftLeaves(7)` → `0` (same reasoning).

    ![Step 1: 9, 15, and 7 bottom out as leaves, each returning 0](images/walkthrough-1.png)

    - Back in `sumOfLeftLeaves(20)`: `l=0, r=0`. Check `20.Left` (`15`): it exists and
      has no children → it's a left leaf → `l += 15` → `l=15`. Returns `15`.

    ![Step 2: at node 20, left child 15 is a left leaf, l becomes 15](images/walkthrough-2.png)

  - Back in `sumOfLeftLeaves(3)`: `l=0` (from the `9` call), `r=15` (from the `20`
    call). Check `3.Left` (`9`): it exists and has no children → left leaf →
    `l += 9` → `l=9`. Returns `l + r = 9 + 15 = 24`. ✓

  ![Step 3: at node 3, left child 9 is a left leaf, l becomes 9, r is 15, returns 24](images/walkthrough-3.png)
