package circulardependency

import "testing"

var shapes = []struct {
	name    string
	imports [][]int
}{
	{"deps_first_10", layered(5000, 10, "deps-first", 31)},   // libraries listed first, 10 layers
	{"apps_first_10", layered(5000, 10, "apps-first", 32)},   // same repo shape, apps listed first
	{"apps_first_100", layered(5000, 100, "apps-first", 33)}, // 100 layers deep
	{"apps_first_1000", layered(5000, 1000, "apps-first", 34)},
	{"chain_5000", layered(5000, 5000, "apps-first", 35)},        // one long line: every package imports the last
	{"cycle_10", withCycle(layered(5000, 10, "apps-first", 36))}, // one import pointing back up
}

func BenchmarkOrderByPasses(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				OrderByPasses(s.imports)
			}
		})
	}
}

func BenchmarkOrderByKahn(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				OrderByKahn(s.imports)
			}
		})
	}
}

func BenchmarkOrderByKahnFlat(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				OrderByKahnFlat(s.imports)
			}
		})
	}
}
