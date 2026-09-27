## Intuition

The values are gone, but the rule that produced them is known, and the shape of the tree
survived. Root is 0; a node `x` has children `2x+1` and `2x+2`. Walk down from the root
applying the rule, and every node gets its label back.

That rule is the numbering an array-backed heap uses: node `i` has children at `2i+1`
and `2i+2`. So a node's value is where it would sit if the tree were a heap laid out in an
array.

The constructor recovers the values with a top-down walk and puts every label into a set.
`Find` is then a single set lookup. With up to `10^4` calls, doing the work once up front
and making each query O(1) is a sensible trade.

## Builds on

This one introduces something rather than reusing it: labels that come from a node's
position. Tomorrow's problem, Path In Zigzag Labelled Binary Tree, is built on exactly this
numbering.

## Without the set

The labels encode the path from the root. Add 1 to a label and write it in binary: after the
leading 1, each bit is a turn, 0 for left and 1 for right. For label 4, `4 + 1 = 101`, so
the path is left (0) then right (1). That means `Find` can walk down from the root instead
of looking in a set:

```go
func findByPath(root *TreeNode, target int) bool {
	x := target + 1
	node := root
	for i := bits.Len(uint(x)) - 2; i >= 0 && node != nil; i-- {
		if x>>i&1 == 0 {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return node != nil
}
```

That's O(log target) per query with no extra memory. I checked it against the set version
on a few hundred random trees. The set is simpler and faster per query; the bit walk is
worth knowing because it explains *why* the labelling works.

The constructor also mutates the tree it was given and keeps a reference to it that
`Find` never uses. Harmless for LeetCode; worth tidying in real code.

**Complexity:**
- Constructor: O(n) time, O(n) space for the set.
- Find: O(1).
