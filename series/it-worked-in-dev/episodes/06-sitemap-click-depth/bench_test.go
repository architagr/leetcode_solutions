package sitemapclickdepth

import (
	"fmt"
	"testing"
)

var shapes = []struct {
	name  string
	build func() *Page
}{
	// A marketing site: a few shallow pages beside a large docs section. The
	// shallowest dead end is one click away and the rest is irrelevant.
	{"site_with_docs", func() *Page {
		return &Page{URL: "/", Children: []*Page{
			balanced(6, 4, "/docs"), {URL: "/pricing"}, {URL: "/contact"},
		}}
	}},
	// Everything is the same depth, so nothing can be skipped.
	{"uniform_5x4", func() *Page { return balanced(5, 4, "/") }},
	// One long route and no shallow dead end at all - the level walk has to
	// go all the way down, paying for a queue it never benefits from.
	{"chain_400", func() *Page { return chain(400, "/") }},
	// Deep docs section, shallow page last in the list.
	{"deep_then_shallow", func() *Page {
		return &Page{URL: "/", Children: []*Page{chain(2000, "/docs"), {URL: "/x"}}}
	}},
}

func BenchmarkRecursion(b *testing.B) {
	for _, s := range shapes {
		root := s.build()
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ClicksByRecursion(root)
			}
		})
	}
}

func BenchmarkByLevel(b *testing.B) {
	for _, s := range shapes {
		root := s.build()
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ClicksByLevel(root)
			}
		})
	}
}

func TestShapeSizes(t *testing.T) {
	var count func(*Page) int
	count = func(p *Page) int {
		n := 1
		for _, c := range p.Children {
			n += count(c)
		}
		return n
	}
	for _, s := range shapes {
		r := s.build()
		t.Logf("%-18s %6d pages, answer %d", s.name, count(r), ClicksByLevel(r))
	}
	_ = fmt.Sprint()
}
