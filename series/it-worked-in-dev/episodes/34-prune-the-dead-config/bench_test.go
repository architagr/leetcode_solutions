package prunethedeadconfig

import (
	"math/rand"
	"testing"
)

var shapes = func() []struct {
	name string
	cfg  *Node
} {
	r := rand.New(rand.NewSource(7))
	return []struct {
		name string
		cfg  *Node
	}{
		{"service", config(r, 5000, 8, 6, 0.3)},    // one service, 30% of keys still read
		{"platform", config(r, 500000, 8, 8, 0.3)}, // every service's config in one tree
		{"stale", config(r, 500000, 8, 8, 0.002)},  // the same, after years: 0.2% still read
		{"rules", rules(5000)},                     // a pricing rule chain, only the fallback read
	}
}()

func run(b *testing.B, fn func(*Node) *Node) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				c := clone(s.cfg) // both prune in place
				b.StartTimer()
				fn(c)
			}
		})
	}
}

func BenchmarkPruneByAsking(b *testing.B) { run(b, PruneByAsking) }
func BenchmarkPruneBottomUp(b *testing.B) { run(b, PruneBottomUp) }
