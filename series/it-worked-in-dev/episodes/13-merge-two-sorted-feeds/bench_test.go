package mergetwosortedfeeds

import "testing"

// The shapes a two-source feed actually takes. tick=1 means about one event
// in two shares a timestamp with another, which is what a coarse clock and
// bulk inserts produce - and it is the case the tie-breaking argument is
// about, so the benchmark runs on it rather than on a convenient input.
var shapes = []struct {
	name string
	n, m int
}{
	{"page_20_20", 20, 20},               // one screen of a feed
	{"day_500_500", 500, 500},            // a day of activity for an active account
	{"month_20k_20k", 20000, 20000},      // the "load everything" month view
	{"export_200k_200k", 200000, 200000}, // an export, or an audit log page that forgot to paginate
	{"lopsided_200k_50", 200000, 50},     // one busy source, one quiet one
}

func inputs() [][2][]Event {
	out := make([][2][]Event, len(shapes))
	for i, s := range shapes {
		own, followed := feeds(s.n, s.m, 1, int64(i)+1)
		out[i] = [2][]Event{own, followed}
	}
	return out
}

func run(b *testing.B, fn func(a, c []Event) []Event) {
	in := inputs()
	for i, s := range shapes {
		own, followed := in[i][0], in[i][1]
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				fn(own, followed)
			}
		})
	}
}

func BenchmarkSorting(b *testing.B)       { run(b, MergeBySorting) }
func BenchmarkStableSorting(b *testing.B) { run(b, MergeByStableSorting) }
func BenchmarkSortFunc(b *testing.B)      { run(b, MergeBySortFunc) }
func BenchmarkWalking(b *testing.B)       { run(b, MergeByWalking) }
