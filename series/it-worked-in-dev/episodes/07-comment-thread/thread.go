// Package commentthread compares two ways of putting a threaded discussion
// into display order - the order a page renders it in, parent immediately
// above its replies.
package commentthread

import (
	"fmt"
	"sort"
	"strings"
)

// Comment is one node of a discussion. Replies are already in the order they
// should appear under their parent.
type Comment struct {
	ID      int
	Author  string
	Replies []*Comment
}

// Rendered is a comment paired with how far to indent it.
type Rendered struct {
	ID     int
	Indent int
}

// OrderBySortKey builds a sort key for every comment and sorts the flat list.
//
// This is the version that gets written first, and it is the one the database
// suggests: give every row a materialised path, then ORDER BY it. It has real
// advantages - it works on a flat result set with no tree in memory, it
// survives being handed to a template unchanged, and it is one line to change
// if the ordering rule changes.
func OrderBySortKey(root *Comment) []Rendered {
	type keyed struct {
		key    string
		indent int
		id     int
	}
	var rows []keyed

	// Every comment carries the path of positions that reaches it, zero-padded
	// so a plain string comparison orders 2 before 10.
	var collect func(c *Comment, path []string)
	collect = func(c *Comment, path []string) {
		rows = append(rows, keyed{strings.Join(path, "."), len(path) - 1, c.ID})
		for i, r := range c.Replies {
			collect(r, append(append([]string{}, path...), fmt.Sprintf("%06d", i)))
		}
	}
	collect(root, []string{"000000"})

	sort.Slice(rows, func(a, b int) bool { return rows[a].key < rows[b].key })

	out := make([]Rendered, 0, len(rows))
	for _, r := range rows {
		out = append(out, Rendered{ID: r.id, Indent: r.indent})
	}
	return out
}

// OrderByWalking emits the same list by walking the thread in display order.
//
// A comment is rendered, then its replies, then the next sibling - which is
// exactly the order the page shows. Nothing is compared to anything.
func OrderByWalking(root *Comment) []Rendered {
	out := make([]Rendered, 0, 16)
	var visit func(c *Comment, indent int)
	visit = func(c *Comment, indent int) {
		// Record this comment BEFORE its replies. That is the whole ordering
		// rule, and it is the only place order is decided.
		out = append(out, Rendered{ID: c.ID, Indent: indent})
		for _, r := range c.Replies {
			visit(r, indent+1)
		}
	}
	if root != nil {
		visit(root, 0)
	}
	return out
}
