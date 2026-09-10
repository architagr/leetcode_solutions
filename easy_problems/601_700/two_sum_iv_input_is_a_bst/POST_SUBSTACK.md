---
meta_title: "Two Sum on a tree: store the complement, not the value"
meta_description: "Strip the BST label and this is classic Two Sum. Store each node's complement rather than its value and every later node checks itself in one lookup."
tags: [golang, binary-search-tree, hashmap, two-pointers, leetcode]
---

# Two Sum IV - Input is a BST

*365 Days of LeetCode Challenge — Day 37/365*

🔗 [LeetCode #653](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/) · Difficulty: Easy

The BST framing invites a clever tree solution. The straightforward one ignores the ordering
entirely and ports the classic array Two Sum onto a traversal.

The only interesting decision is what goes into the hash set, and the non-obvious choice makes
the check at each node a single lookup with no arithmetic.

### The problem

Given the root of a binary search tree and an integer `k`, return `true` if two
elements in the tree add up to `k`, or `false` otherwise.

### The intuition

Take away the "BST" part and this is just Two Sum. The one-pass array version keeps
a hash set of values seen so far, and for each new value `v` checks whether `k - v`
is already in the set. If it is, you've got your pair. If not, you record `v` (or
its complement) and move on.

This solution runs that same idea over a tree instead of an array, walking it
preorder: visit the node, then the left subtree, then the right. The twist is what
gets stored. Instead of remembering each value directly, the recursive helper stores
that value's complement, `k - root.Val`, in a shared map. Then for every new node it
checks whether the node's own value is already sitting in the map as someone else's
complement. If it is, some earlier node `A` computed `k - A.Val == root.Val`, which
rearranges to `A.Val + root.Val == k`. That's the pair.

The map gets passed down through every recursive call instead of getting recreated,
and since Go maps are reference types, every call is reading and writing the same
underlying storage. That's what lets a match surface between two branches that have
nothing to do with each other. The complement recorded while visiting one branch is
still sitting there when the walk reaches a totally different branch later on.

What I like about this one: it flat out ignores that the tree is a binary search
tree. The ordering never gets used, anywhere. A plain hash-set two-sum check works
just as well on an unsorted binary tree. It's a simpler approach than the
in-order-traversal-plus-two-pointer trick that actually exploits the ordering, at
the cost of O(n) space instead of O(h).

**Complexity:** O(n) time, since every node gets visited once. O(n) space for the
hash map in the worst case, plus O(h) for the recursion stack, where `h` is the
tree's height.

### The solution

![Example 1](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/two_sum_iv_input_is_a_bst/images/1.jpg "Example1")

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

- At node `5`, the root. The map is empty, so `5` isn't in it. Store the complement
  it needs: `map[9-5] = map[4] = true`.

  ![Step 1: at root 5, map is empty, store complement 4](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/two_sum_iv_input_is_a_bst/images/walkthrough-1.png)

- Recurse left into `3`. Not in `{4}`. Store `map[9-3] = map[6] = true`.

  ![Step 2: at node 3, store complement 6](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/two_sum_iv_input_is_a_bst/images/walkthrough-2.png)

  - Recurse left into `2`, a leaf. Not in `{4, 6}`. Store `map[9-2] = map[7] = true`.
    Both children are `nil`, so `find(2, ...)` returns `false`.

    ![Step 3: at leaf 2, store complement 7, returns false](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/two_sum_iv_input_is_a_bst/images/walkthrough-3.png)

  - Recurse right into `4`. This one's already sitting in the map (`{4, 6, 7}`),
    stored two levels up while visiting `5`. `find(4, ...)` returns `true` right
    away, no need to look at `4`'s (nil) children. That's the actual pair: `5 + 4 = 9`.

    ![Step 4: at node 4, 4 is already in the map -> match, returns true](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/two_sum_iv_input_is_a_bst/images/walkthrough-4.png)

- Recurse right into `6`. `left` and `right` are two separate statements at the
  root, not a short-circuited `||`, so `find(6, ...)` runs anyway even though the
  left subtree already found a match. Turns out `6` is also already in the map
  (stored while visiting `3`, since `9 - 3 = 6`), a second valid pair: `3 + 6 = 9`.
  `find(6, ...)` returns `true` immediately, so `7`, `6`'s own right child, never
  gets visited.

  ![Step 5: at node 6, 6 is already in the map -> also a match, 7 is never visited](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/two_sum_iv_input_is_a_bst/images/walkthrough-5.png)

- Back at the root: `left = true`, `right = true`, so `return left || right` gives
  `true`. Matches what the problem expects.

Same walk explains why `findTarget([5,3,6,2,4,null,7], 28)` (the problem's second
example) comes back `false`: the preorder walk hits every node (`5, 3, 2, 4, 6, 7`),
storing each one's complement (`23, 25, 26, 24, 22, 21`), and none of those six ever
equals a value the walk reaches afterward, so the match check never fires.

---

Storing the complement instead of the value is a small inversion that shows up well beyond this
problem. It moves the arithmetic to write time, where you're doing it once per element anyway,
and leaves the read a plain membership test.

Full code and the step-by-step walkthrough:
[two_sum_iv_input_is_a_bst](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/two_sum_iv_input_is_a_bst/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
