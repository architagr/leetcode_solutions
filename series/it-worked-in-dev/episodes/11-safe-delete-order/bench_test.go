package safedeleteorder

import "testing"

// Hierarchies as they actually come: wide and shallow when the schema is
// workspace/project/board/card, and deep when the thing nests into itself -
// a comment reply, a subtask, a category under a category.
var shapes = []struct {
	name string
	root *Record
}{
	{"board_20", hierarchy([]int{20})},            // one board, 20 cards
	{"workspace_319", hierarchy([]int{3, 5, 20})}, // 3 projects, 5 boards, 20 cards
	{"tenant_1245", hierarchy([]int{4, 10, 30})},  // one level wider again
	{"categories_2047", hierarchy(dup(10, 2))},    // 10 levels of nesting
	{"thread_2000", thread(1999)},                 // a reply to a reply to a reply
}

func dup(n, v int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = v
	}
	return out
}

func BenchmarkRepeatedScan(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				WavesByRepeatedScan(s.root)
			}
		})
	}
}

func BenchmarkOnePass(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				WavesByOnePass(s.root)
			}
		})
	}
}

func TestShapeSizes(t *testing.T) {
	for _, s := range shapes {
		t.Logf("%-16s %5d records, %4d waves",
			s.name, Count(s.root), len(WavesByOnePass(s.root)))
	}
}
