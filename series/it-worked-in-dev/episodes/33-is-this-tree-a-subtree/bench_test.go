package isthistreeasubtree

import (
	"math/rand"
	"testing"
)

type shape struct {
	name       string
	page, snip *Node
	want       bool
}

var shapes = func() []shape {
	r := rand.New(rand.NewSource(7))
	page := random(r, 5000)
	// the last piece of the page with 40 to 60 components, so the walk
	// has to get most of the way through the page before it finds it
	var pick *Node
	var w func(*Node)
	w = func(x *Node) {
		if size(x) >= 40 && size(x) <= 60 {
			pick = x
		}
		for _, k := range x.Kids {
			w(k)
		}
	}
	w(page)
	gridPage, gridSnip := grid(r, 2000, 60)
	bigPage, bigSnip := grid(r, 200, 2000)
	deepPage, deepSnip := thread(5000, 1000)
	return []shape{
		{"page", page, clone(pick), true},     // a normal page, a card that is on it
		{"grid", gridPage, gridSnip, false},   // 2,000 cards from one template, a new variant
		{"sections", bigPage, bigSnip, false}, // 200 large sections from one template
		{"thread", deepPage, deepSnip, false}, // one reply chain 5,000 deep
	}
}()

func run(b *testing.B, fn func(*Node, *Node) bool) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.page, s.snip)
			}
		})
	}
}

func BenchmarkContainsByWalk(b *testing.B) { run(b, ContainsByWalk) }
func BenchmarkContainsByText(b *testing.B) { run(b, ContainsByText) }
func BenchmarkContainsByHash(b *testing.B) { run(b, ContainsByHash) }
