# Solution walkthrough

Two functions, and the interesting one is four lines long.

```go
func reverseOddLevels(root *TreeNode) *TreeNode {
	rev(root.Left, root.Right, 1)
	return root
}
```

The entry point doesn't start at the root. It starts at the root's two children, passing them as a pair along with depth `1`.

Starting one level down is what makes the rest work. `rev` operates on pairs of nodes, not single nodes, and the root has no partner to pair with. It also never needs one: level 0 holds a single node, and reversing a one-element level does nothing. Passing `1` as the initial depth means the argument always equals the real level number, so the odd-level test is a plain parity check with no offset to remember.

If the tree is a single node, `root.Left` is `nil` and the guard inside `rev` returns immediately.

## The pair-wise recursion

```go
func rev(l, r *TreeNode, d int) {
	if l == nil {
		return
	}
	// At odd levels, swap the mirrored nodes' values
	if d%2 == 1 {
		l.Val, r.Val = r.Val, l.Val
	}
	// Recurse into children, swapping pointers to maintain mirror
	rev(l.Left, r.Right, d+1)
	rev(l.Right, r.Left, d+1)
}
```

The invariant is that `l` and `r` are always mirror images of each other across the centre of the tree, at the same level.

Checking only `l == nil` is enough. The tree is perfect, so `l` and `r` are at the same depth and are either both leaves' children (both `nil`) or both real nodes. There's no state where one is `nil` and the other isn't, so a second check would be dead code.

The swap is a Go tuple assignment, so no temporary is needed. Only the values move; the nodes stay where they are and no pointers are rewired. Rearranging the values is enough because the problem only asks about values.

## Why the recursive calls cross

The two calls at the bottom are the crux:

![How the child pairs are formed](images/walkthrough-2.png)

Given the mirror pair `(2, 3)`, their four children have to be regrouped into two new mirror pairs. Node 4 is the leftmost of the four and node 7 is the rightmost, so those two mirror each other: `rev(l.Left, r.Right, d+1)`. Nodes 5 and 6 are the middle two: `rev(l.Right, r.Left, d+1)`.

The second call is the one that crosses. Pairing left-with-left and right-with-right would hand node 4 to node 6, which are not mirrors, and the level would come out shuffled instead of reversed.

## Tracing it

A four-level perfect tree numbered 1 to 15, so every position is distinguishable:

![Mirror pairs on each odd level](images/walkthrough-1.png)

Levels 1 and 3 are odd and get reversed. Levels 0 and 2 are only passed through.

**`rev(2, 3, d=1)`** — depth is odd, so 2 and 3 swap. Level 1 goes from `[2, 3]` to `[3, 2]`, which is that level reversed. Then it recurses twice.

**`rev(4, 7, d=2)`** — even, no swap. It still recurses, because level 3 is underneath:

- `rev(8, 15, d=3)` — odd, swap. These are the outermost two nodes on the level.
- `rev(9, 14, d=3)` — odd, swap.

**`rev(5, 6, d=2)`** — even, no swap. Recurses:

- `rev(10, 13, d=3)` — odd, swap.
- `rev(11, 12, d=3)` — odd, swap.

Level 3 started as `[8, 9, 10, 11, 12, 13, 14, 15]`. The four swaps exchange 8 with 15, 9 with 14, 10 with 13, and 11 with 12, leaving `[15, 14, 13, 12, 11, 10, 9, 8]`. Fully reversed, and no node was moved.

The calls at depth 4 receive `nil` from the leaves and return at the guard.

## Complexity

**Time: O(n).** Each node is reached exactly once as half of some pair, and the work per pair is a parity check plus at most one swap.

**Space: O(log n).** The recursion stack goes as deep as the tree is tall, and a perfect binary tree with n nodes has height log2(n). Nothing is allocated per level, which is the advantage over the BFS approach that buffers each level into a slice.
