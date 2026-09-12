package directorysizes

import (
	"fmt"
	"reflect"
	"testing"
)

// buildTree makes a directory tree `depth` levels deep where every directory
// has `breadth` subdirectories and one 1-byte file.
//
// One byte per file keeps the arithmetic checkable by hand: a directory's total
// is exactly the number of directories at or below it.
func buildTree(depth, breadth int) *Dir {
	var build func(level int, path string) *Dir
	build = func(level int, path string) *Dir {
		d := &Dir{Path: path, Files: []File{{Name: "f", Size: 1}}}
		if level == 0 {
			return d
		}
		for i := 0; i < breadth; i++ {
			d.Children = append(d.Children, build(level-1, fmt.Sprintf("%s/%d", path, i)))
		}
		return d
	}
	return build(depth, "/root")
}

// A chain: one directory per level, no branching. The worst case for the
// brute-force version, and the shape node_modules tends toward.
func buildChain(depth int) *Dir {
	root := &Dir{Path: "/root", Files: []File{{Name: "f", Size: 1}}}
	cur := root
	for i := 0; i < depth; i++ {
		next := &Dir{Path: fmt.Sprintf("%s/%d", cur.Path, i), Files: []File{{Name: "f", Size: 1}}}
		cur.Children = []*Dir{next}
		cur = next
	}
	return root
}

// The whole argument of the episode rests on the two producing identical
// output, so that is the first thing tested.
func TestBothAgree(t *testing.T) {
	cases := []struct {
		name string
		root *Dir
	}{
		{"single directory", &Dir{Path: "/root", Files: []File{{Name: "a", Size: 10}}}},
		{"empty directory", &Dir{Path: "/root"}},
		{"flat, no nesting", buildTree(1, 5)},
		{"balanced", buildTree(4, 3)},
		{"deep chain", buildChain(50)},
		{"wide and shallow", buildTree(2, 20)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			brute := AllSizesBrute(c.root)
			fast := AllSizesPostorder(c.root)
			if !reflect.DeepEqual(brute, fast) {
				t.Fatalf("implementations disagree\n brute=%v\n fast =%v", brute, fast)
			}
		})
	}
}

// With one byte per file, a directory's total is the count of directories at or
// below it - so the answers can be checked against arithmetic rather than
// against the other implementation.
func TestKnownTotals(t *testing.T) {
	root := buildTree(2, 2) // 1 + 2 + 4 = 7 directories
	got := AllSizesPostorder(root)

	if got["/root"] != 7 {
		t.Errorf("root = %d, want 7", got["/root"])
	}
	if got["/root/0"] != 3 { // itself plus its two leaves
		t.Errorf("/root/0 = %d, want 3", got["/root/0"])
	}
	if got["/root/0/0"] != 1 { // a leaf
		t.Errorf("/root/0/0 = %d, want 1", got["/root/0/0"])
	}
	if len(got) != 7 {
		t.Errorf("reported %d directories, want 7", len(got))
	}
}

// The claim in the write-up is that a file d levels deep is added d times by
// the brute-force version. This counts the additions to show it is true.
func TestBruteForceRepeatsWork(t *testing.T) {
	for _, depth := range []int{5, 10, 20} {
		root := buildChain(depth)
		dirs := depth + 1

		var bruteAdds int
		var countingSizeOf func(*Dir) int64
		countingSizeOf = func(d *Dir) int64 {
			var total int64
			for _, f := range d.Files {
				bruteAdds++
				total += f.Size
			}
			for _, c := range d.Children {
				total += countingSizeOf(c)
			}
			return total
		}
		var walk func(*Dir)
		walk = func(d *Dir) {
			countingSizeOf(d)
			for _, c := range d.Children {
				walk(c)
			}
		}
		walk(root)

		// Every directory sums itself and all its descendants, so on a chain of
		// n directories the additions are n + (n-1) + ... + 1.
		want := dirs * (dirs + 1) / 2
		if bruteAdds != want {
			t.Errorf("depth %d: brute force made %d additions, expected %d", depth, bruteAdds, want)
		}
		// The postorder version touches each file exactly once.
		if dirs >= bruteAdds {
			t.Errorf("depth %d: expected brute force (%d) to exceed one-per-file (%d)", depth, bruteAdds, dirs)
		}
	}
}
