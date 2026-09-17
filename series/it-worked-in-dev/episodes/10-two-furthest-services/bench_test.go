package twofurthestservices

import "testing"

// The shapes a topology actually takes. A fleet is wide and four or five
// levels deep; the chain is what you get when the hierarchy is really a
// relay path, or when somebody labelled every host its own zone.
var shapes = []struct {
	name string
	root *Node
}{
	{"rack_24", fanout([]int{24})},              // one rack, 24 instances
	{"fleet_1245", fanout([]int{4, 10, 30})},    // 4 zones, 10 racks, 30 each
	{"fleet_4197", fanout([]int{4, 8, 10, 12})}, // a level deeper again
	{"chain_2000", chain(1999)},                 // a topology that is a list
}

func BenchmarkScan(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				WidestByScan(s.root)
			}
		})
	}
}

func BenchmarkOnePass(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				WidestByOnePass(s.root)
			}
		})
	}
}

func BenchmarkWidestPair(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				WidestPair(s.root)
			}
		})
	}
}

func TestShapeSizes(t *testing.T) {
	for _, s := range shapes {
		t.Logf("%-11s %5d nodes, widest pair %d hops",
			s.name, Size(s.root), WidestByOnePass(s.root))
	}
}

func BenchmarkScanReusing(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				WidestByScanReusing(s.root)
			}
		})
	}
}
