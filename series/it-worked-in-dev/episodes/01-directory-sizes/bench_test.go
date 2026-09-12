package directorysizes

import (
	"fmt"
	"testing"
)

// Benchmarked at several depths on purpose. The interesting fact is not that
// one wins - it is where they stop being the same.
var shapes = []struct {
	name    string
	depth   int
	breadth int
}{
	{"shallow_d3_b4", 3, 4}, // 85 dirs, the shape of a small project
	{"medium_d6_b3", 6, 3},  // 1093 dirs
	{"deep_d12_b2", 12, 2},  // 8191 dirs, deep and narrow
	{"chain_d200", 200, 1},  // a single nested path, node_modules at its worst
	{"chain_d800", 800, 1},  // the same shape, four times deeper
}

func BenchmarkBrute(b *testing.B) {
	for _, s := range shapes {
		root := buildTree(s.depth, s.breadth)
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				AllSizesBrute(root)
			}
		})
	}
}

func BenchmarkPostorder(b *testing.B) {
	for _, s := range shapes {
		root := buildTree(s.depth, s.breadth)
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				AllSizesPostorder(root)
			}
		})
	}
}

// How many directories each shape actually contains, so the numbers above can
// be read against a size rather than a label.
func TestShapeSizes(t *testing.T) {
	for _, s := range shapes {
		n := len(AllSizesPostorder(buildTree(s.depth, s.breadth)))
		t.Logf("%-16s %6d directories", s.name, n)
	}
	_ = fmt.Sprint()
}
