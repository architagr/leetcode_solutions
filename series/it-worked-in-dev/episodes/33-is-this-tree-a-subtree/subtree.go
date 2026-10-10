package isthistreeasubtree

import "strings"

// A rendered page is a tree of components. The design-system linter asks
// whether a given component tree - a cookie banner, a pricing card - already
// appears somewhere in the page, exactly: same kinds, same children, same order.

// Node is one component.
type Node struct {
	Kind string
	Kids []*Node
}

// ContainsByWalk is what I would write: try every component of the right kind
// as the place the snippet starts, and compare from there in lockstep.
func ContainsByWalk(page, snip *Node) bool {
	if page == nil {
		return snip == nil
	}
	if page.Kind == snip.Kind && same(page, snip) {
		return true
	}
	for _, k := range page.Kids {
		if ContainsByWalk(k, snip) {
			return true
		}
	}
	return false
}

func same(a, b *Node) bool {
	if a.Kind != b.Kind || len(a.Kids) != len(b.Kids) {
		return false
	}
	for i := range a.Kids {
		if !same(a.Kids[i], b.Kids[i]) {
			return false
		}
	}
	return true
}

// ContainsByText flattens both trees to text and searches one in the other.
// Every component is written as "(kind" then its children then ")", so the
// text of a subtree is exactly the text of that component and nothing else.
func ContainsByText(page, snip *Node) bool {
	var p, s strings.Builder
	write(&p, page)
	write(&s, snip)
	return strings.Contains(p.String(), s.String())
}

func write(b *strings.Builder, n *Node) {
	b.WriteByte('(')
	b.WriteString(n.Kind)
	for _, k := range n.Kids {
		write(b, k)
	}
	b.WriteByte(')')
}

// ContainsByHash gives every subtree a fingerprint made from its kind and its
// children's fingerprints, computed bottom-up in one pass. Two equal subtrees
// always have equal fingerprints; only a match has to be confirmed node by node.
func ContainsByHash(page, snip *Node) bool {
	want := fingerprint(snip)
	found := false
	var walk func(n *Node) uint64
	walk = func(n *Node) uint64 {
		h := kindHash(n.Kind)
		for _, k := range n.Kids {
			h = mix(h, walk(k)) // the children answer first
		}
		h = mix(h, uint64(len(n.Kids)))
		if !found && h == want && same(n, snip) { // equal fingerprints, then make sure
			found = true
		}
		return h
	}
	walk(page)
	return found
}

func fingerprint(n *Node) uint64 {
	h := kindHash(n.Kind)
	for _, k := range n.Kids {
		h = mix(h, fingerprint(k))
	}
	return mix(h, uint64(len(n.Kids)))
}

func kindHash(s string) uint64 {
	h := uint64(14695981039346656037)
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= 1099511628211
	}
	return h
}

func mix(h, v uint64) uint64 {
	h ^= v + 0x9e3779b97f4a7c15 + h<<6 + h>>2
	h ^= h >> 31
	h *= 0xbf58476d1ce4e5b9
	return h ^ h>>29
}
