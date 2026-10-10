package whoholdsthetime

import "testing"

var shapes = []struct {
	name   string
	events []Event
}{
	{"request", walk(2000, 12, 0.5, 8, 11)},    // one HTTP request, a dozen layers
	{"batch", walk(200000, 16, 0.5, 8, 12)},    // a batch job, still shallow
	{"parser", walk(20000, 4000, 0.52, 8, 13)}, // a recursive-descent parser on nested input
	{"recursion", chain(5000, 14)},             // one recursion 5,000 deep
}

func run(b *testing.B, fn func(int, []Event) []int64) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(8, s.events)
			}
		})
	}
}

func BenchmarkSelfBySpans(b *testing.B) { run(b, SelfBySpans) }
func BenchmarkSelfByStack(b *testing.B) { run(b, SelfByStack) }
