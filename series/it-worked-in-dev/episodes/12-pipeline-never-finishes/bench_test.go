package pipelineneverfinishes

import "testing"

// The shapes a "what runs next" chain actually takes: a hand-written pipeline
// of a dozen steps, a generated one of a few hundred, and the long chains that
// turn up when the steps are per-tenant, per-redirect or per-retry.
var shapes = []struct {
	name string
	head *Stage
}{
	{"pipeline_12", chain(12, -1)},            // a hand-written pipeline
	{"pipeline_200", chain(200, -1)},          // generated, one step per rule
	{"chain_100k", chain(100000, -1)},         // long, and it terminates
	{"chain_100k_loop", chain(100000, 1)},     // long, and it loops near the top
	{"chain_100k_tail", chain(100000, 99000)}, // loops in the last 1%
}

func BenchmarkSeen(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				HasCycleBySeen(s.head)
			}
		})
	}
}

func BenchmarkTwoPointers(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				HasCycleByTwoPointers(s.head)
			}
		})
	}
}

func BenchmarkEntryBySeen(b *testing.B) {
	for _, s := range shapes[3:] {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				EntryBySeen(s.head)
			}
		})
	}
}

func BenchmarkEntryByTwoPointers(b *testing.B) {
	for _, s := range shapes[3:] {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				EntryByTwoPointers(s.head)
			}
		})
	}
}

func TestShapeSizes(t *testing.T) {
	for _, s := range shapes {
		if HasCycleByTwoPointers(s.head) {
			t.Logf("%-16s loops, entry %s", s.name, EntryByTwoPointers(s.head))
			continue
		}
		t.Logf("%-16s %6d stages, terminates", s.name, Length(s.head))
	}
}
