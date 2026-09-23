package orgchartbylevel

import "testing"

// The shapes an org chart takes. Two of them are here because the queue wins
// nothing on them, and an episode that only benchmarks the shapes it wins on
// is an advertisement.
var shapes = []struct {
	name string
	root *Person
}{
	{"org_5k", org(6, 4)},     // 5,461 people, six levels of four
	{"org_56k", org(6, 6)},    // 55,987 people, six levels of six
	{"flat_50k", flat(50000)}, // one person at the top, everybody else under them
	{"chain_2k", chain(2000)}, // the reorg nobody meant to ship
}

func runShapes(b *testing.B, fn func(root *Person)) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.root)
			}
		})
	}
}

func BenchmarkAllLevelsByDepthKeyedWalk(b *testing.B) {
	runShapes(b, func(root *Person) { LevelsByDepthKeyedWalk(root) })
}

func BenchmarkAllLevelsByQueue(b *testing.B) {
	runShapes(b, func(root *Person) { LevelsByQueue(root) })
}

// What the page actually asks for on first paint: the top three levels.
func BenchmarkFirstThreeByDepthKeyedWalk(b *testing.B) {
	runShapes(b, func(root *Person) { FirstLevelsByDepthKeyedWalk(root, 3) })
}

func BenchmarkFirstThreeByQueue(b *testing.B) {
	runShapes(b, func(root *Person) { FirstLevelsByQueue(root, 3) })
}
