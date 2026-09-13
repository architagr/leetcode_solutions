// Package flattennavconfig compares two ways of turning a nested navigation
// config into the flat list a renderer wants - one row per item, in display
// order, each carrying its depth.
package flattennavconfig

// Item is a nav entry as it appears in config: nested, because that is how a
// human writes it.
type Item struct {
	Label    string
	Children []*Item
}

// Row is a nav entry as the renderer wants it: flat, with its depth.
type Row struct {
	Label string
	Depth int
}

// FlattenByConcat builds each subtree's rows and joins them.
//
// This is the version that gets written first, and it is the one functional
// style suggests: a function that turns an item into its rows, defined in terms
// of itself. Every call is independent, nothing is shared, and there is no
// accumulator to thread through or forget to pass.
func FlattenByConcat(root *Item) []Row {
	if root == nil {
		return nil
	}
	out := []Row{{Label: root.Label, Depth: 0}}
	for _, c := range root.Children {
		for _, r := range FlattenByConcat(c) {
			// each child's rows are rebuilt one level deeper
			out = append(out, Row{Label: r.Label, Depth: r.Depth + 1})
		}
	}
	return out
}

// FlattenByAppending writes every row into one slice as it walks.
//
// The slice is created once and passed down. Nothing is joined, and no row is
// ever copied out of one slice into another.
func FlattenByAppending(root *Item) []Row {
	if root == nil {
		return nil
	}
	out := make([]Row, 0, 16)
	var visit func(it *Item, depth int)
	visit = func(it *Item, depth int) {
		out = append(out, Row{Label: it.Label, Depth: depth})
		for _, c := range it.Children {
			visit(c, depth+1)
		}
	}
	visit(root, 0)
	return out
}

// Count returns how many items are in the tree, so a caller can size the
// destination exactly and skip the growth entirely.
func Count(it *Item) int {
	if it == nil {
		return 0
	}
	n := 1
	for _, c := range it.Children {
		n += Count(c)
	}
	return n
}

// FlattenPresized is FlattenByAppending with the slice allocated once at the
// right size, which costs one extra pass over the tree.
func FlattenPresized(root *Item) []Row {
	if root == nil {
		return nil
	}
	out := make([]Row, 0, Count(root))
	var visit func(it *Item, depth int)
	visit = func(it *Item, depth int) {
		out = append(out, Row{Label: it.Label, Depth: depth})
		for _, c := range it.Children {
			visit(c, depth+1)
		}
	}
	visit(root, 0)
	return out
}
