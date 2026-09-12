// Package filebreadcrumbs compares two ways of producing the full path of
// every node in a tree - the breadcrumb each row of a file browser shows.
package filebreadcrumbs

import "strings"

// Node is a folder or a file. Parent is nil at the root, which is what makes
// walking upward terminate.
type Node struct {
	Name     string
	Parent   *Node
	Children []*Node
}

// Sep is the separator between path segments.
const Sep = "/"

// PathsByWalkingUp returns every node's full path by climbing from it to the
// root and reversing what it collected.
//
// This is the version that gets written first, and it is a fair one. Each
// node's path is derivable from the node alone: no traversal order to get
// right, no state threaded through a recursion, and it works just as well on
// one node as on the whole tree - which matters, because "show me the path of
// the file I clicked" is a real request.
func PathsByWalkingUp(all []*Node) map[*Node]string {
	out := make(map[*Node]string, len(all))
	for _, n := range all {
		var parts []string
		for cur := n; cur != nil; cur = cur.Parent {
			parts = append(parts, cur.Name)
		}
		// collected root-last, so reverse into reading order
		for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
			parts[i], parts[j] = parts[j], parts[i]
		}
		out[n] = strings.Join(parts, Sep)
	}
	return out
}

// PathsByCarryingDown returns the same map in one traversal, passing each
// node's path to its children as it goes.
//
// A child's path is its parent's path plus one segment, and the parent's is
// finished before any child is visited, so it is never rebuilt.
func PathsByCarryingDown(root *Node) map[*Node]string {
	out := make(map[*Node]string)
	var visit func(n *Node, prefix string)
	visit = func(n *Node, prefix string) {
		path := n.Name
		if prefix != "" {
			path = prefix + Sep + n.Name
		}
		out[n] = path
		for _, c := range n.Children {
			// path is a local string and strings are immutable, so siblings
			// cannot disturb each other's prefix - no undo step is needed.
			visit(c, path)
		}
	}
	if root != nil {
		visit(root, "")
	}
	return out
}
