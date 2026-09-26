package isthisconfigatree

import "testing"

var shapes = []struct {
	name  string
	n     int
	links [][2]int
}{
	{"tree_100", 100, randomTree(100, 21)},  // a small service's config
	{"tree_1k", 1000, randomTree(1000, 22)}, // a platform team's
	{"tree_5k", 5000, randomTree(5000, 23)}, // a whole org's
	{"loop_5k", 5000, withLoop(5000, 24)},   // one extra link, last line
	{"swapped_5k", 5000, swapped(5000, 25)}, // right count, a loop anyway
}

func BenchmarkIsTreeByLinkCheck(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				IsTreeByLinkCheck(s.n, s.links)
			}
		})
	}
}

func BenchmarkIsTreeByCount(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				IsTreeByCount(s.n, s.links)
			}
		})
	}
}

func BenchmarkIsTreeByUnionFind(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				IsTreeByUnionFind(s.n, s.links)
			}
		})
	}
}
