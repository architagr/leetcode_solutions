// Package directorysizes computes the total size of every directory in a tree,
// which is what `du` prints.
//
// It holds two implementations of the same function. They return identical
// results and one of them stops working as the tree gets deeper.
package directorysizes

// File is a leaf with a size in bytes.
type File struct {
	Name string
	Size int64
}

// Dir is a directory: its own files, and its subdirectories.
type Dir struct {
	Path     string
	Files    []File
	Children []*Dir
}

// ---------------------------------------------------------------------------
// The version you would write first
// ---------------------------------------------------------------------------

// SizeOf returns the total bytes under dir, including everything nested.
//
// There is nothing wrong with this function. It is correct, it is clear, and
// it is the obvious way to answer "how big is this directory".
func SizeOf(dir *Dir) int64 {
	var total int64
	for _, f := range dir.Files {
		total += f.Size
	}
	for _, child := range dir.Children {
		total += SizeOf(child)
	}
	return total
}

// AllSizesBrute reports the total size of every directory in the tree.
//
// This is where it goes wrong, and the reason is not visible from here: the
// walk visits each directory once, but SizeOf re-walks that directory's entire
// subtree every time it is called.
//
// So a file nested d levels deep is added into the running total d separate
// times, once for each ancestor that asks. The deeper the tree, the more times
// the same bytes are counted.
func AllSizesBrute(root *Dir) map[string]int64 {
	out := make(map[string]int64)
	var walk func(*Dir)
	walk = func(d *Dir) {
		out[d.Path] = SizeOf(d)
		for _, child := range d.Children {
			walk(child)
		}
	}
	walk(root)
	return out
}

// ---------------------------------------------------------------------------
// The version that scales
// ---------------------------------------------------------------------------

// AllSizesPostorder reports the same thing in a single pass.
//
// The change is the order of two lines. A directory's total is computed only
// after its children have returned theirs, so each child's bytes are added once
// and then carried upward rather than recomputed by every ancestor.
//
// Children before parent is postorder traversal.
func AllSizesPostorder(root *Dir) map[string]int64 {
	out := make(map[string]int64)
	var visit func(*Dir) int64
	visit = func(d *Dir) int64 {
		var total int64
		for _, f := range d.Files {
			total += f.Size
		}
		// Recurse first. Each child hands back a number it computed once.
		for _, child := range d.Children {
			total += visit(child)
		}
		// Only now is this directory's own answer known.
		out[d.Path] = total
		return total
	}
	visit(root)
	return out
}
