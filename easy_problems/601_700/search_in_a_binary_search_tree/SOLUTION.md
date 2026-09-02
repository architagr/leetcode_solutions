## Solution walkthrough

The implementation is `searchBST(root *TreeNode, val int) *TreeNode` in `main.go`.

```go
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return root
}
```

Four lines of logic, four things to trace:

1. **Base case.** `if root == nil { return nil }` — the recursion has walked off the
   tree without finding `val`, so there's nothing left to check.
2. **Target is smaller.** `if root.Val > val { return searchBST(root.Left, val) }` —
   everything in `root`'s right subtree is even bigger than `root.Val`, so it's
   impossible for `val` to be there. Only the left subtree can still contain it.
3. **Target is bigger.** `if root.Val < val { return searchBST(root.Right, val) }` — the
   mirror image of step 2; only the right subtree is still worth checking.
4. **Match.** If neither comparison fired, `root.Val == val`, and `return root` hands
   back the matching node — along with its `Left`/`Right` pointers, which is exactly the
   "subtree rooted at that node" the problem asks for.

### Example 1: value found

!["Example 1"](tree1.jpg "Example 1")

```
Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]
```

- `searchBST(4, val=2)`: `root.Val(4) > val(2)` → recurse left, into node `2`. The right
  subtree (`7`) is eliminated in this single comparison — it's never even visited.

  ![Step 1: at node 4, 4 > 2, recurse left, right subtree eliminated](images/walkthrough-1.svg)

- `searchBST(2, val=2)`: `root.Val(2) == val(2)` → neither `if` fires, so `return root`
  returns node `2`. Its `Left` (`1`) and `Right` (`3`) pointers are still attached,
  giving the expected output subtree `[2,1,3]`.

  ![Step 2: at node 2, match found, returns subtree [2,1,3]](images/walkthrough-2.svg)

### Example 2: value not found

!["Example 2"](tree2.jpg "Example 2")

```
Input: root = [4,2,7,1,3], val = 5
Output: []
```

- `searchBST(4, val=5)`: `root.Val(4) < val(5)` → recurse right, into node `7`. This time
  the *left* subtree (`2,1,3`) is the one eliminated.

  ![Step 3: at node 4, 4 < 5, recurse right, left subtree eliminated](images/walkthrough-3.svg)

- `searchBST(7, val=5)`: `root.Val(7) > val(5)` → recurse left, into `7.Left`, which is
  `nil`. That call hits the base case immediately: `searchBST(nil, val=5)` returns
  `nil`, and that `nil` propagates straight back up as the final answer — matching the
  expected empty output.

  ![Step 4: at node 7, 7 > 5, recurse left into nil, base case returns nil](images/walkthrough-4.svg)

Both traces show the same shape: one comparison per level, one subtree eliminated per
comparison, no branch ever explored twice.

**Complexity:** O(h) time and O(h) space, where h is the tree's height — each recursive
call descends exactly one level along a single path.
