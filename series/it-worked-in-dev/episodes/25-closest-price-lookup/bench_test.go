package closestpricelookup

import "testing"

var small, smallPrices = catalog(1000, 7)
var big, bigPrices = catalog(200000, 8)

// Targets in the cheap end, the middle and the dear end of each catalogue:
// the streaming walk's cost depends on where the target falls.
var shapes = []struct {
	name   string
	root   *Node
	target int
}{
	{"1k_middle", small, smallPrices[500] + 3},
	{"200k_cheap", big, bigPrices[2000] + 3},
	{"200k_middle", big, bigPrices[100000] + 3},
	{"200k_dear", big, bigPrices[198000] + 3},
}

func run(b *testing.B, fn func(*Node, int) int) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.root, s.target)
			}
		})
	}
}

func BenchmarkClosestBySorting(b *testing.B)    { run(b, ClosestBySorting) }
func BenchmarkClosestBySortedFlat(b *testing.B) { run(b, ClosestBySortedFlat) }
func BenchmarkClosestByWalk(b *testing.B)       { run(b, ClosestByWalk) }
func BenchmarkClosestByDescent(b *testing.B)    { run(b, ClosestByDescent) }
