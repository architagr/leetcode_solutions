package mergepaginatedfeeds

import "testing"

var shapes = []struct {
	name  string
	feeds [][]Post
	n     int
}{
	{"new", following(10, 100, 11), 50},        // a new account following ten
	{"typical", following(300, 100, 12), 50},   // following 300, first page
	{"heavy", following(5000, 100, 13), 50},    // following 5,000, first page
	{"export", following(300, 100, 14), 30000}, // following 300, everything
}

func run(b *testing.B, fn func([][]Post, int) []Post) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.feeds, s.n)
			}
		})
	}
}

func BenchmarkPageBySort(b *testing.B)    { run(b, PageBySort) }
func BenchmarkPageByFolding(b *testing.B) { run(b, PageByFolding) }
func BenchmarkPageByHeap(b *testing.B)    { run(b, PageByHeap) }
