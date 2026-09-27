package matchingbracketsconfig

import "testing"

var shapes = []struct {
	name string
	text string
}{
	{"rules_100", rules(100, 11)},              // a small rules file
	{"rules_10k", rules(10000, 12)},            // the whole pricing engine
	{"nested_500", nested(500)},                // one generated expression, 500 deep
	{"nested_5000", nested(5000)},              // the same generator, 5,000 deep
	{"crossed_10k", crossed(rules(10000, 13))}, // the big file, one pair crossed
}

func run(b *testing.B, fn func(string) bool) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(s.text)))
			for i := 0; i < b.N; i++ {
				fn(s.text)
			}
		})
	}
}

func BenchmarkBalancedByCounting(b *testing.B) { run(b, BalancedByCounting) }
func BenchmarkBalancedByErasing(b *testing.B)  { run(b, BalancedByErasing) }
func BenchmarkBalancedByStack(b *testing.B)    { run(b, BalancedByStack) }
