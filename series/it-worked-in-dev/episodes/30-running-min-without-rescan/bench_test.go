package runningminwithoutrescan

import "testing"

var shapes = []struct {
	name  string
	steps []int
}{
	{"month_100", session(100, 11)},    // a month of transactions
	{"year_2k", session(2000, 12)},     // a year
	{"decade_20k", session(20000, 13)}, // a decade, or a small business
	{"log_100k", session(100000, 14)},  // a small business's plan, its change log replayed on load
}

func run(b *testing.B, mk func() plan) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				play(mk(), s.steps)
			}
		})
	}
}

func BenchmarkPlanByRescan(b *testing.B) {
	run(b, func() plan { return &PlanByRescan{Opening: 100000} })
}
func BenchmarkPlanByEntry(b *testing.B) { run(b, func() plan { return &PlanByEntry{Opening: 100000} }) }
