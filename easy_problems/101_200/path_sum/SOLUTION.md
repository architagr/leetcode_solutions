## Solution walkthrough

The implementation is `hasPathSum(root *TreeNode, targetSum int) bool` in `path_sum.go`,
which is just a thin wrapper around the recursive helper `sum(root *TreeNode, target,
current int) bool`.

![Example 1](images/1.jpg)

We'll trace it on the example above, `root = [5,4,8,11,null,13,4,7,2,null,null,null,1]`,
`targetSum = 22`. As a tree:

```
            5
          /   \
         4     8
        /     / \
      11     13  4
      / \          \
     7   2          1
```

1. **Entry point.** `hasPathSum` just calls `sum(root, targetSum, 0)` — `current` starts
   at `0`, since no value has been accumulated yet.

2. **Base case: empty subtree.** `if root == nil { return false }` — this handles the
   fully empty tree (Example 3) directly, and also means a nonexistent child never gets
   mistaken for a valid path.

3. **Both children exist → try both.**
   `else if root.Left != nil && root.Right != nil { return sum(root.Left, target,
   current+root.Val) || sum(root.Right, target, current+root.Val) }` — a node with two
   children isn't a leaf, so it isn't checked against `target` here. Instead it adds its
   own value to `current` and hands that running total down to *both* children,
   short-circuiting on `||` the moment either side finds a match.

4. **Only a left child → descend left.**
   `else if root.Left != nil { return sum(root.Left, target, current+root.Val) }` —
   a node with just one child is still not a leaf. It's important this is a distinct
   branch from step 3: without it, a node with only a left child would otherwise reach
   the final `return target == root.Val+current` and be wrongly evaluated as if it were
   a leaf.

5. **Only a right child → descend right.**
   `else if root.Right != nil { return sum(root.Right, target, current+root.Val) }` —
   the mirror of step 4.

6. **Leaf → this is the only place a comparison happens.**
   `return target == root.Val+current` — control only reaches this line when both
   `root.Left` and `root.Right` are nil, i.e. `root` is a genuine leaf. `root.Val +
   current` is the full root-to-leaf sum, checked against `target` exactly once, right
   here.

Walking the trace on `[5,4,8,11,null,13,4,7,2,null,null,null,1]`, `targetSum = 22`:

- `sum(5, 22, 0)`: `5` has both children → recurse left into `4` with `current=5`, OR
  recurse right into `8` (only if the left side comes back `false`).
- `sum(4, 22, 5)`: `4` has only a left child (`11`) → recurse into `11` with
  `current=5+4=9`.
- `sum(11, 22, 9)`: `11` has both children → recurse into `7` and `2`, each with
  `current=9+11=20`.

  ![Step 1: current sum descends 5 -> 4 -> 11, current becomes 9](images/walkthrough-1.svg)

- `sum(7, 22, 20)`: `7` is a leaf. Check: `22 == 7+20` (`27`)? No → returns `false`.

  ![Step 2: at leaf 7, 22 != 27, returns false](images/walkthrough-2.svg)

- `sum(2, 22, 20)`: `2` is a leaf. Check: `22 == 2+20` (`22`)? Yes → returns `true`.

  ![Step 3: at leaf 2, 22 == 22, returns true](images/walkthrough-3.svg)

- Back in `sum(11, 22, 9)`: `false || true` → `true`.
- Back in `sum(4, 22, 5)`: `true` (nothing else to combine with; only the left branch
  existed).
- Back in `sum(5, 22, 0)`: `true || ...` — the `||` **short-circuits**, so
  `sum(8, 22, 5)` (the entire right subtree, `13`/`4`/`1`) is never evaluated at all.
  `hasPathSum` returns `true`. ✓

  ![Step 4: true bubbles up 2 -> 11 -> 4 -> 5, right subtree never visited](images/walkthrough-4.svg)

**Complexity:** O(n) time — every node is visited at most once (fewer, when `||`
short-circuits). O(h) space for the recursion stack, where h is the tree's height.
