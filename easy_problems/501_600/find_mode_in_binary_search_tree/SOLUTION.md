## Solution walkthrough

The implementation is `findMode(root *TreeNode) []int`, backed by the helper
`getCnt(root *TreeNode, data map[int]int) map[int]int`, in `main.go`.

![Example 1](images/1.jpg "Example1")

We'll trace it on the example above, `root = [1,null,2,2]` — node `1` at the root with
no left child, its right child is a `2`, and that `2`'s left child is another `2`.

### Phase 1 — `getCnt`: tally every value into a map

1. **Guard the empty case.** `findMode` starts with `if root == nil { return []int{} }` —
   an empty tree has no modes.

2. **Build the frequency map.** `data := make(map[int]int)` creates an empty map, then
   `data = getCnt(root, data)` threads it through a recursive traversal. Inside
   `getCnt`: `if root == nil { return nil }` stops recursion at empty subtrees, then
   `data[root.Val]++` bumps the current node's count before recursing into
   `getCnt(root.Left, data)` and `getCnt(root.Right, data)`. Because `data` is a map
   (a reference type), every recursive call mutates the *same* underlying map — the
   `map[int]int` return value only matters for the top-level caller.

   Walking the traversal in the order the code visits nodes (`root` itself, then
   `root.Left`, then `root.Right`):
   - `getCnt(1, data)`: `data[1]++` → `data = {1: 1}`.

     ![Step 1: getCnt visits the root, data becomes {1: 1}](images/walkthrough-1.svg)

   - `1.Left` is `nil`, so that branch returns immediately without touching `data`.
   - `getCnt(2a, data)` (the right child of `1`): `data[2]++` → `data = {1: 1, 2: 1}`.

     ![Step 2: getCnt visits 1's right child, data becomes {1: 1, 2: 1}](images/walkthrough-2.svg)

   - `getCnt(2b, data)` (the left child of `2a`): `data[2]++` again →
     `data = {1: 1, 2: 2}`. `2b` has no children, so both of its recursive calls hit the
     `nil` base case and return.

     ![Step 3: getCnt visits 2a's left child, data becomes {1: 1, 2: 2}](images/walkthrough-3.svg)

   By the time `getCnt` unwinds back to `findMode`, `data` holds the full frequency
   table: `{1: 1, 2: 2}`.

### Phase 2 — `findMode`: find the max, then collect the ties

3. **First pass — find the highest count.** `cnt := 0` starts at zero, then
   `for _, v := range data { if v > cnt { cnt = v } }` scans every value in the map and
   keeps the largest one seen. For `{1: 1, 2: 2}`, that leaves `cnt = 2`.

4. **Second pass — collect every key that hits that count.**
   `res := make([]int, 0)` starts empty, then
   `for k, v := range data { if v == cnt { res = append(res, k) } }` walks the map again
   and appends every key whose count equals the max — this is what makes the function
   correctly return *all* modes when there's a tie, not just one. Here only `k = 2` has
   `v == cnt`, so `res = [2]`.

   ![Step 4: findMode scans for the max count, then collects every key matching it, producing [2]](images/walkthrough-4.svg)

5. **Return the result.** `return res` hands back `[2]`, matching the expected output.

**Complexity:**
- Time: O(n) — `getCnt` visits every node exactly once, and each of the two scans over
  `data` in `findMode` is O(n) in the worst case (every value distinct).
- Space: O(n) for the `data` map (up to one entry per distinct value) plus O(h) for the
  recursion stack during `getCnt`, where h is the tree's height.
