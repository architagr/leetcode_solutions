package findacorrespondingnodeofabinarytreeinclone

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// getTargetCopy walks original and clones in lockstep: every recursive call
// advances both pointers together, so wherever the traversal is standing in
// original, clones is standing at the structurally-matching node in the
// cloned tree. That's guaranteed because clones is an exact copy of
// original, which lets the search find target's counterpart purely by
// position/value instead of needing to compare node identities.
func getTargetCopy(original, clones, target *TreeNode) *TreeNode {
	// clones mirrors original node-for-node, so clones == nil alone is
	// enough to detect "ran out of tree" along this path; original doesn't
	// need its own nil check.
	if clones == nil {
		return nil
	}
	// Values are guaranteed unique, so a value match means clones is
	// standing exactly where target stands in original — return this
	// cloned-tree node as the answer.
	if clones.Val == target.Val {
		return clones
	}
	// Search the left subtrees of both trees together first.
	n := getTargetCopy(original.Left, clones.Left, target)
	if n != nil {
		return n
	}
	// Not found on the left; whatever the right subtree search finds (or
	// doesn't) is this call's result.
	return getTargetCopy(original.Right, clones.Right, target)
}
