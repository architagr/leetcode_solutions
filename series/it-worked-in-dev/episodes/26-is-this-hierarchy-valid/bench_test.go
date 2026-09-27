package isthishierarchyvalid

import "testing"

var shapes = []struct {
	name string
	root *Node
}{
	{"valid_32k", balanced(0, 327660)},                         // 32,767 products, balanced
	{"valid_262k", balanced(0, 2621430)},                       // 262,143 products, balanced
	{"corrupt_deep_262k", corruptDeep(balanced(0, 2621430))},   // one leaf on the wrong side of the root
	{"corrupt_early_262k", corruptEarly(balanced(0, 2621430))}, // broken at the root's left child
	{"chain_5k", chain(5000)},                                  // valid, from a sorted import
}

func run(b *testing.B, fn func(*Node) bool) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.root)
			}
		})
	}
}

func BenchmarkValidByChildren(b *testing.B) { run(b, ValidByChildren) }
func BenchmarkValidBySubtrees(b *testing.B) { run(b, ValidBySubtrees) }
func BenchmarkValidBySorting(b *testing.B)  { run(b, ValidBySorting) }
func BenchmarkValidByBounds(b *testing.B)   { run(b, ValidByBounds) }
