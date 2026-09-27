package paginatewithoutloading

import "testing"

var big = catalog(200000, 7) // 200,000 products, prices 10 .. 2,000,000

const size = 20

var shapes = []struct {
	name   string
	page   int
	cursor int // the last price on the previous page
}{
	{"page_1", 0, 0},
	{"page_100", 99, 99 * size * 10},
	{"page_5000", 4999, 4999 * size * 10},
	{"last_page", 9999, 9999 * size * 10},
}

func BenchmarkPageByList(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				PageByList(big, s.page, size)
			}
		})
	}
}

func BenchmarkPageByStack(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				PageByStack(big, s.page, size)
			}
		})
	}
}

func BenchmarkPageAfter(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				PageAfter(big, s.cursor, size)
			}
		})
	}
}

// The export: every product, in order. The list is built once either way.
func BenchmarkExportAllByList(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		it := NewIterator(big)
		for it.HasNext() {
			it.Next()
		}
	}
}

func BenchmarkExportAllByStack(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		it := NewStackIterator(big)
		for it.HasNext() {
			it.Next()
		}
	}
}
