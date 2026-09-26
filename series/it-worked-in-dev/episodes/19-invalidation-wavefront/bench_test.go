package invalidationwavefront

import "testing"

var cluster10k = gossip(10000, 10000, 7)

// The same 10,000-node cluster with more and more origins. The one-origin
// shape is the one where searching from each origin is exactly one search.
var shapes = []struct {
	name    string
	c       Cluster
	origins []int
}{
	{"origins_1", cluster10k, spread(10000, 1)},       // one node took the write
	{"origins_10", cluster10k, spread(10000, 10)},     // one per availability zone
	{"origins_100", cluster10k, spread(10000, 100)},   // a bulk write across keys
	{"origins_1000", cluster10k, spread(10000, 1000)}, // a region-wide deploy
	{"ring_2k_20", ring(2000), spread(2000, 20)},      // a ring, 20 origins
}

func run(b *testing.B, fn func(Cluster, []int) int) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.c, s.origins)
			}
		})
	}
}

func BenchmarkRoundsBySearchEach(b *testing.B) { run(b, RoundsBySearchEach) }
func BenchmarkRoundsByOneWave(b *testing.B)    { run(b, RoundsByOneWave) }
