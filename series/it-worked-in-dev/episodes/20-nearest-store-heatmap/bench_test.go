package neareststoreheatmap

import "testing"

// One 500 x 500 city, 250,000 blocks, with more and more stores open.
var shapes = []struct {
	name       string
	rows, cols int
	stores     []Point
}{
	{"stores_1", 500, 500, scatter(500, 500, 1, 11)},       // one warehouse
	{"stores_10", 500, 500, scatter(500, 500, 10, 12)},     // a chain starting out
	{"stores_100", 500, 500, scatter(500, 500, 100, 13)},   // a city chain
	{"stores_1000", 500, 500, scatter(500, 500, 1000, 14)}, // dark stores everywhere
}

func run(b *testing.B, fn func(int, int, []Point) [][]int) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.rows, s.cols, s.stores)
			}
		})
	}
}

func BenchmarkHeatmapByScan(b *testing.B)   { run(b, HeatmapByScan) }
func BenchmarkHeatmapByBFS(b *testing.B)    { run(b, HeatmapByBFS) }
func BenchmarkHeatmapBySweeps(b *testing.B) { run(b, HeatmapBySweeps) }
