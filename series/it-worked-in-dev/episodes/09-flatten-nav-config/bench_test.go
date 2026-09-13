package flattennavconfig

import (
	"fmt"
	"testing"
)

// A nav menu is wide and two or three deep. The deep shapes are what a
// category tree or a docs sidebar becomes when nobody flattens it.
var shapes = []struct {
	name         string
	depth, width int
}{
	{"menu_2x6", 2, 6},      // 43 items, a real nav menu
	{"sidebar_3x5", 3, 5},   // 156 items
	{"deep_100", 100, 1},    // a 100-level chain
	{"deep_600", 600, 1},    // six times deeper
}

func build(depth, width int) *Item {
	n := 0
	return menu(depth, width, &n)
}

func BenchmarkConcat(b *testing.B) {
	for _, s := range shapes {
		root := build(s.depth, s.width)
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				FlattenByConcat(root)
			}
		})
	}
}

func BenchmarkAppending(b *testing.B) {
	for _, s := range shapes {
		root := build(s.depth, s.width)
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				FlattenByAppending(root)
			}
		})
	}
}

func BenchmarkPresized(b *testing.B) {
	for _, s := range shapes {
		root := build(s.depth, s.width)
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				FlattenPresized(root)
			}
		})
	}
}

func TestShapeSizes(t *testing.T) {
	for _, s := range shapes {
		t.Logf("%-13s %5d items", s.name, Count(build(s.depth, s.width)))
	}
	_ = fmt.Sprint()
}
