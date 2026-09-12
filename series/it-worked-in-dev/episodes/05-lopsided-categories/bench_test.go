package lopsidedcategories

import (
	"fmt"
	"testing"
)

// A navigation tree is wide and shallow. The deep shapes are what a category
// tree turns into when one section keeps growing and nobody flattens it.
var shapes = []struct {
	name  string
	build func() *Category
	nodes int
}{
	{"nav_3x5", func() *Category { return balanced(3, 5, "n") }, 156},
	{"nav_4x5", func() *Category { return balanced(4, 5, "n") }, 781},
	{"deep_200", func() *Category {
		r := &Category{Name: "r"}
		chainUnder(r, 200, "c")
		return r
	}, 201},
	{"deep_1000", func() *Category {
		r := &Category{Name: "r"}
		chainUnder(r, 1000, "c")
		return r
	}, 1001},
}

func BenchmarkAskingTwice(b *testing.B) {
	for _, s := range shapes {
		root := s.build()
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				LopsidedByAskingTwice(root)
			}
		})
	}
}

func BenchmarkOnePass(b *testing.B) {
	for _, s := range shapes {
		root := s.build()
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				LopsidedInOnePass(root)
			}
		})
	}
}

func TestShapeSizes(t *testing.T) {
	var count func(*Category) int
	count = func(c *Category) int {
		n := 1
		for _, ch := range c.Children {
			n += count(ch)
		}
		return n
	}
	for _, s := range shapes {
		t.Logf("%-11s %5d nodes", s.name, count(s.build()))
	}
	_ = fmt.Sprint()
}
