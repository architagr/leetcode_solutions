package whichdepthisheaviest

import "testing"

// The shapes a category tree takes. store_7k is the size of a real
// department-store taxonomy; chain_2k is what an importer produces when it
// nests every path segment.
var shapes = []struct {
	name string
	root *Category
}{
	{"store_400", catalog(3, 7, 40)},  // 400 categories: a small shop
	{"store_7k", catalog(4, 9, 40)},   // 7,381: a department store
	{"market_56k", catalog(6, 6, 40)}, // 55,987: a marketplace
	{"flat_50k", flat(50000, 3)},      // nobody organised it
	{"chain_2k", chain(2000)},         // an importer that nested every segment
}

func run(b *testing.B, fn func(*Category) int) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.root)
			}
		})
	}
}

func BenchmarkHeaviestByMap(b *testing.B)         { run(b, HeaviestByMap) }
func BenchmarkHeaviestByMapTieBreak(b *testing.B) { run(b, HeaviestByMapTieBreak) }
func BenchmarkHeaviestBySlice(b *testing.B)       { run(b, HeaviestBySlice) }
func BenchmarkHeaviestByLevel(b *testing.B)       { run(b, HeaviestByLevel) }
