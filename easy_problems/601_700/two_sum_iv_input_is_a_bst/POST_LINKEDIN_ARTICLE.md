## 365 Days of LeetCode Challenge — Day 17/365

# Two Sum IV - Input is a BST

🔗 https://leetcode.com/problems/two-sum-iv-input-is-a-bst/ · Difficulty: Easy

### The problem

Given the root of a binary search tree and an integer `k`, return `true` if two
elements in the tree add up to `k`, or `false` otherwise.

### The intuition

Strip away the "BST" label for a moment and this is just the classic **Two Sum**
problem: find two elements that add up to `k`. The one-pass array solution keeps a
hash set of values seen so far and, for each new value `v`, checks whether `k - v`
is already in the set. If it is, the matching pair has been found; otherwise `v` (or
its complement) gets recorded for a future match.

This solution ports that exact idea onto a tree by walking it with plain recursion
(preorder: visit the node, then its left subtree, then its right subtree) instead of
iterating an array. The one twist is *what* gets stored: instead of recording each
visited value directly, the recursive helper stores that value's **complement**,
`k - root.Val`, in a shared map. Then, for every new node, it checks whether the
node's *own* value is already sitting in the map as someone else's complement. If it
is, some earlier node `A` previously computed `k - A.Val == root.Val`, which
rearranges to `A.Val + root.Val == k` — exactly the pair being searched for.

Because the map is a single object passed down through every recursive call (maps
are reference types in Go), a match can be found across *any* two nodes in the tree
— the complement recorded while visiting one branch is still visible when a
completely different branch is visited later.

Worth calling out: this approach completely ignores the fact that the tree is a
**binary search tree**. A plain hash-set two-sum check works on any binary tree,
sorted or not — the BST ordering isn't used anywhere. That's a valid, simpler choice
than an in-order-traversal-plus-two-pointer approach that *does* exploit the
ordering, but it trades away the ability to do the search with O(h) extra space
instead of O(n).

**Complexity:** O(n) time — every node visited once. O(n) space for the hash map in
the worst case, plus O(h) for the recursion stack, where `h` is the tree's height.

### The solution

![Example 1](images/1.jpg "Example1")

```go
func findTarget(root *TreeNode, k int) bool {
	// One map shared across the whole recursion (maps are reference types
	// in Go), so a complement stored while visiting one branch is still
	// visible when a completely different branch is visited later.
	var hashMap map[int]bool = make(map[int]bool)

	return find(root, k, hashMap)
}

// find walks the tree preorder (this node, then left subtree, then right
// subtree). It's the classic single-pass Two Sum hash-set trick ported onto
// a tree traversal instead of an array iteration -- it never relies on the
// BST ordering, so it would work on any binary tree.
func find(root *TreeNode, k int, hashMap map[int]bool) bool {
	if root == nil {
		// Empty subtree can't contain a matching pair.
		return false
	}
	if _, ok := hashMap[root.Val]; ok {
		// Some earlier-visited node already recorded root.Val as the
		// complement it needed (i.e. that node's value + root.Val == k).
		// Return immediately without recursing into this node's subtree.
		return true
	}
	// No match yet: remember what value this node would need to see later
	// in order to complete a pair.
	hashMap[k-root.Val] = true
	// Note: left and right are separate statements, not a single
	// short-circuited `find(...) || find(...)` expression, so right is
	// always evaluated even when left already found a match.
	left := find(root.Left, k, hashMap)
	right := find(root.Right, k, hashMap)
	return left || right
}
```

Walking it through `root = [5,3,6,2,4,null,7]`, `k = 9` (expected `true`):

- **At node `5` (root).** The map is empty, so `5` isn't in it. Store the complement
  it needs: `map[9-5] = map[4] = true`.

  ![Step 1: at root 5, map is empty, store complement 4](images/walkthrough-1.svg)

- **Recurse left, into `3`.** `3` isn't in `{4}`. Store `map[9-3] = map[6] = true`.

  ![Step 2: at node 3, store complement 6](images/walkthrough-2.svg)

  - **Recurse left, into `2` (a leaf).** `2` isn't in `{4, 6}`. Store
    `map[9-2] = map[7] = true`. Both children are `nil`, so `find(2, ...)` returns
    `false`.

    ![Step 3: at leaf 2, store complement 7, returns false](images/walkthrough-3.svg)

  - **Recurse right, into `4`.** `4` **is already in the map** (`{4, 6, 7}`) — it was
    stored two levels up, while visiting `5`. `find(4, ...)` returns `true`
    immediately, without recursing into `4`'s (nil) children. This is the actual pair:
    `5 + 4 = 9`.

    ![Step 4: at node 4, 4 is already in the map -> match, returns true](images/walkthrough-4.svg)

- **Recurse right, into `6`.** Because `left` and `right` are computed as two
  separate statements at the root, `find(6, ...)` still runs even though the left
  subtree already produced `true`. `6` **is also already in the map** (stored while
  visiting `3`, since `9 - 3 = 6`) — another valid pair, `3 + 6 = 9`. `find(6, ...)`
  returns `true` immediately, so `6`'s own right child `7` is never visited at all.

  ![Step 5: at node 6, 6 is already in the map -> also a match, 7 is never visited](images/walkthrough-5.svg)

- Back at the root: `left = true`, `right = true` → `return left || right` → `true`.
  ✓ Matches the expected output.

This also shows why `findTarget([5,3,6,2,4,null,7], 28)` (the problem's second
example) comes back `false`: the preorder walk visits every node — `5, 3, 2, 4, 6,
7` — storing each one's complement (`23, 25, 26, 24, 22, 21`), and none of those six
complements ever equals a node value that's visited afterward, so the match check
never succeeds.

Full code: `easy_problems/601_700/two_sum_iv_input_is_a_bst/` in the repo.
