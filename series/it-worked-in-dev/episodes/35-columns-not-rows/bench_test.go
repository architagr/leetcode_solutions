package columnsnotrows

import "testing"

var shapes = []struct {
	name string
	root *Node
}{
	{"stump", full(6)},          // a 63-node tree, what the viewer shows by default
	{"model", full(14)},         // a complete depth-14 tree, 16,383 nodes
	{"grown", grown(200000, 3)}, // an unpruned tree, 200,000 nodes, lopsided
	{"deep", full(20)},          // a complete depth-20 tree, 1,048,575 nodes
}

func run(b *testing.B, fn func(*Node) [][]int) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.root)
			}
		})
	}
}

func BenchmarkColumnsBySort(b *testing.B)  { run(b, ColumnsBySort) }
func BenchmarkColumnsByLevel(b *testing.B) { run(b, ColumnsByLevel) }
