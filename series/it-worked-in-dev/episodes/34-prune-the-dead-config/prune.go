package prunethedeadconfig

// A service's config is a tree of sections and keys. Before shipping, the
// build drops every branch nothing in the code reads: a key no code
// references, and a section left with nothing under it. A section stays if
// anything below it is still read, whether or not its own name is.

// Node is a section or a key.
type Node struct {
	Key  string
	Used bool // something in the code reads this key
	Kids []*Node
}

// PruneByAsking is what I would write: at each section, ask whether anything
// under it is used; drop it if not, otherwise keep it and do the same for its
// children.
func PruneByAsking(n *Node) *Node {
	if !anythingUsed(n) {
		return nil
	}
	kept := n.Kids[:0]
	for _, k := range n.Kids {
		if p := PruneByAsking(k); p != nil {
			kept = append(kept, p)
		}
	}
	n.Kids = kept
	return n
}

// anythingUsed reports whether n or anything below it is used. It stops at
// the first used key it finds.
func anythingUsed(n *Node) bool {
	if n.Used {
		return true
	}
	for _, k := range n.Kids {
		if anythingUsed(k) {
			return true
		}
	}
	return false
}

// PruneBottomUp lets the children answer first. Each one comes back pruned,
// or nil if nothing under it survived, and the section decides from that.
func PruneBottomUp(n *Node) *Node {
	kept := n.Kids[:0]
	for _, k := range n.Kids {
		if p := PruneBottomUp(k); p != nil { // the child has already decided
			kept = append(kept, p)
		}
	}
	n.Kids = kept
	if !n.Used && len(kept) == 0 {
		return nil
	}
	return n
}
