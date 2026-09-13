package commentthread

import (
	"fmt"
	"testing"
)

// A real discussion is wide and shallow - many top-level comments, a couple of
// replies deep. The deep shapes are what a long argument between two people
// looks like.
var shapes = []struct {
	name         string
	depth, width int
}{
	{"discussion_2x8", 2, 8},  // 73 comments, the common case
	{"busy_3x6", 3, 6},        // 259 comments
	{"argument_60", 60, 1},    // two people, sixty replies deep
	{"argument_400", 400, 1},  // the same, much longer
}

func build(depth, width int) *Comment {
	n := 0
	return thread(depth, width, &n)
}

func BenchmarkSortKey(b *testing.B) {
	for _, s := range shapes {
		root := build(s.depth, s.width)
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				OrderBySortKey(root)
			}
		})
	}
}

func BenchmarkWalking(b *testing.B) {
	for _, s := range shapes {
		root := build(s.depth, s.width)
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				OrderByWalking(root)
			}
		})
	}
}

func TestShapeSizes(t *testing.T) {
	for _, s := range shapes {
		t.Logf("%-15s %5d comments", s.name, len(OrderByWalking(build(s.depth, s.width))))
	}
	_ = fmt.Sprint()
}
