package slowestendpoints

import "testing"

var shapes = []struct {
	name string
	log  []Req
}{
	{"page", traffic(1000, 11)},           // the last thousand requests
	{"minute", traffic(60000, 12)},        // a minute at 1,000 requests a second
	{"hour", traffic(3600000, 13)},        // an hour at 1,000 a second
	{"backlog", backlog(3600000, 14)},     // an hour in which a queue backs up
	{"ascending", ascending(3600000, 15)}, // an hour exported fastest first
	{"climbing", climbing(3600000, 16)},   // every request slower than the last
}

func run(b *testing.B, fn func([]Req, int) []Req) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.log, 10)
			}
		})
	}
}

func BenchmarkSlowestBySort(b *testing.B) { run(b, SlowestBySort) }
func BenchmarkSlowestByHeap(b *testing.B) { run(b, SlowestByHeap) }
