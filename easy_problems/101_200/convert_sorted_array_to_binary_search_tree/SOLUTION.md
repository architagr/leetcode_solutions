## Solution walkthrough

The implementation is `SortedArrayToBST(nums []int) *TreeNode` in
`convert_sorted_array_to_binary_search_tree.go`.

1. **Base case.** `if len(nums) == 0 { return nil }` — an empty slice means there's no
   subtree to build at this position, so the function returns `nil` and the recursion
   bottoms out here.

2. **Pick the middle element as this subtree's root.** `mid := len(nums) / 2` finds the
   midpoint index of the current slice. `root.Val = nums[mid]` makes that value the
   node here.

3. **Recurse left.** `root.Left = SortedArrayToBST(nums[:mid])` builds the left subtree
   from everything strictly before `mid` — all values smaller than `root.Val`, since
   the array is sorted ascending.

4. **Recurse right.** `root.Right = SortedArrayToBST(nums[mid+1:])` builds the right
   subtree from everything strictly after `mid` — all values larger than `root.Val`.
   (The `if mid < len(nums)` guard here is always true for a non-empty slice, since
   `mid = len(nums)/2` is strictly less than `len(nums)` whenever `len(nums) > 0` — the
   empty-slice case is already handled by the base case above. `nums[mid+1:]` is safe
   and correctly evaluates to an empty slice on its own when `mid+1 == len(nums)`, so
   the right-subtree recursion naturally returns `nil` in that case regardless.)

5. **Return `root`.**

Walking it through `nums = [-10,-3,0,5,9]`:
- `mid = 2`, `root.Val = 0`.
- Left: `SortedArrayToBST([-10,-3])` → `mid=1`, root `-3`, left `SortedArrayToBST([-10])`
  → single-node tree `-10`, right `SortedArrayToBST([])` → `nil`. So `-3` has left
  child `-10`.
- Right: `SortedArrayToBST([5,9])` → `mid=1`, root `9`, left `SortedArrayToBST([5])` →
  single-node tree `5`, right `nil`. So `9` has left child `5`.
- Final tree: `0` with left subtree rooted at `-3` (which has left child `-10`) and
  right subtree rooted at `9` (which has left child `5`) — matching the example's
  accepted output shape `[0,-3,9,-10,null,5]`.
