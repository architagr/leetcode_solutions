package nextbiggerspike

import "testing"

var shapes = []struct {
	name string
	load []int
}{
	{"week_hours", daily(168, 24, 50, 11)},        // a week, hourly
	{"year_hours", daily(8760, 24, 10, 12)},       // a growing year, hourly
	{"year_minutes", daily(525600, 1440, 10, 14)}, // a growing year, per minute
	{"launch_hours", launch(8760, 24, 16)},        // the year after a launch, hourly
	{"launch_minutes", launch(43200, 1440, 17)},   // the month after a launch, per minute
}

func run(b *testing.B, fn func([]int) []int) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.load)
			}
		})
	}
}

func BenchmarkWaitByScan(b *testing.B)  { run(b, WaitByScan) }
func BenchmarkWaitByStack(b *testing.B) { run(b, WaitByStack) }
