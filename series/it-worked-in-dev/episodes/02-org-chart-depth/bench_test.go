package orgchartdepth

import (
	"fmt"
	"testing"
)

// Two shapes on purpose. A real company is wide and shallow - thousands of
// people, maybe a dozen levels - and that is the case where the naive version
// is fine. The deep shapes are where it is not, and they are what a category
// tree or a threaded discussion actually looks like.
var shapes = []struct {
	name   string
	people []Person
}{
	{"flat_1000", bushy(1, 999)},         // 1000 people, one level
	{"company_6x4", bushy(6, 4)},         // 5461 people, 6 levels - a real org
	{"deep_200", chain(200)},             // a 200-deep thread
	{"deep_2000", chain(2000)},           // ten times deeper
}

func BenchmarkWalkingUp(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				LevelsByWalkingUp(s.people)
			}
		})
	}
}

func BenchmarkLevelOrder(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				LevelsByLevelOrder(s.people)
			}
		})
	}
}

func TestShapeSizes(t *testing.T) {
	for _, s := range shapes {
		t.Logf("%-12s %5d people", s.name, len(s.people))
	}
	_ = fmt.Sprint()
}
