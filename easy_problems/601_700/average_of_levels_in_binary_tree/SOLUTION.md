## Solution walkthrough

The implementation is `AverageOfLevel(root *TreeNode) []float64`, backed by the helper
`getSumAndCount(node *TreeNode, result, count *[]float64, level int)`, both in
`average_of_levels_in_binary_tree.go`.

![Example 1](images/1.jpg "Example1")

We'll trace it on the example above, `root = [3,9,20,null,null,15,7]`.

1. **Set up the running totals.** `AverageOfLevel` allocates two empty slices,
   `result` and `count`, both `[]float64`. `result[i]` will end up holding the sum of
   all node values at level `i`; `count[i]` holds how many nodes exist at level `i`.
   Both are passed to `getSumAndCount` by pointer so every recursive call can grow and
   mutate the same underlying slices.

2. **Base case.** `if node == nil { return }` — an empty subtree adds nothing to any
   level's total.

3. **Grow the slices lazily when a new level is reached.**
   `if len(*result) < level+1 { *result = append(*result, 0.0); *count = append(*count, 0.0) }`
   — the first time the recursion arrives at a given depth, it appends a fresh `0.0`
   slot for that level to both slices. Every later visit to the same depth finds the
   slot already there and skips straight to updating it.

4. **Accumulate into this node's level slot.**
   `(*result)[level] += float64(node.Val)` adds the current node's value into its
   level's running sum, and `(*count)[level]++` bumps that level's node count. Because
   the slot is indexed by `level` rather than by visit order, nodes from completely
   different branches of the tree still land in the same slot as long as they share a
   depth.

5. **Recurse into both children, one level deeper.**
   `getSumAndCount(node.Left, result, count, level+1)` then
   `getSumAndCount(node.Right, result, count, level+1)` — left subtree fully explored
   before the right one, in classic preorder DFS fashion.

6. **Convert sums into averages.** Back in `AverageOfLevel`, once the recursion has
   returned, `for i := range result { result[i] /= count[i] }` divides each level's
   accumulated sum by that level's node count, turning `result` from "sum per level"
   into "average per level" in place. That's the array that gets returned.

Walking it through `root = [3,9,20,null,null,15,7]` (expected `[3.0, 14.5, 11.0]`):

- `getSumAndCount(3, level=0)`: no slot exists yet, so slot 0 is created; `result=[3.0]`,
  `count=[1.0]`. Then it recurses left into `9` and right into `20`.

  ![Step 1: visit 3 (level 0), create slot 0, result[0]=3, count[0]=1](images/walkthrough-1.svg)

- `getSumAndCount(9, level=1)`: no slot 1 yet, so it's created; `result=[3.0, 9.0]`,
  `count=[1.0, 1.0]`. `9` has no children, so both recursive calls hit the base case
  immediately.

  ![Step 2: visit 9 (level 1), create slot 1, result[1]=9, count[1]=1](images/walkthrough-2.svg)

- `getSumAndCount(20, level=1)`: slot 1 already exists from the `9` visit, so no new
  slot is appended — it's updated in place: `result[1] += 20` → `29.0`,
  `count[1]++` → `2.0`. This is the step that shows why indexing by `level` matters:
  `20` is in a completely different branch than `9`, yet they both write into the same
  slot because they share a depth.

  ![Step 3: visit 20 (level 1), slot 1 exists, result[1] becomes 29, count[1] becomes 2](images/walkthrough-3.svg)

- `getSumAndCount(15, level=2)`: no slot 2 yet, so it's created; `result=[3.0, 29.0,
  15.0]`, `count=[1.0, 2.0, 1.0]`.

  ![Step 4: visit 15 (level 2), create slot 2, result[2]=15, count[2]=1](images/walkthrough-4.svg)

- `getSumAndCount(7, level=2)`: slot 2 already exists, updated in place:
  `result[2] += 7` → `22.0`, `count[2]++` → `2.0`.

  ![Step 5: visit 7 (level 2), slot 2 exists, result[2] becomes 22, count[2] becomes 2](images/walkthrough-5.svg)

- Recursion is done: `result=[3.0, 29.0, 22.0]`, `count=[1.0, 2.0, 2.0]`. Back in
  `AverageOfLevel`, the final loop divides each slot: `3.0/1.0 = 3.0`,
  `29.0/2.0 = 14.5`, `22.0/2.0 = 11.0`, giving `[3.0, 14.5, 11.0]`. ✓

  ![Step 6: divide result[i] by count[i] for every level, giving 3.0, 14.5, 11.0](images/walkthrough-6.svg)

**Complexity:**
- Time: O(n) — `getSumAndCount` visits every node exactly once; the final division
  loop is O(number of levels), which is bounded by O(n).
- Space: O(h) for the recursion stack (h = tree height) plus O(L) for the `result`/
  `count` slices and the returned array, where L is the number of levels (L <= h + 1).
