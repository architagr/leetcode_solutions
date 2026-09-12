package filebreadcrumbs

import (
	"fmt"
	"testing"
)

// Real file trees are wide and shallow; a document hierarchy or a category
// tree is narrower and deeper. Both are here, because they behave differently.
var shapes = []struct {
	name         string
	depth, width int
}{
	{"flat_500", 1, 500},   // one folder, 500 files
	{"browser_5x4", 5, 4},  // 1365 nodes, 5 levels - a project tree
	{"deep_100", 100, 1},   // a 100-deep chain
	{"deep_600", 600, 1},   // six times deeper
}

func BenchmarkWalkingUp(b *testing.B) {
	for _, s := range shapes {
		_, all := tree(s.depth, s.width)
		b.Run(fmt.Sprintf("%s_%d", s.name, len(all)), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				PathsByWalkingUp(all)
			}
		})
	}
}

func BenchmarkCarryingDown(b *testing.B) {
	for _, s := range shapes {
		root, _ := tree(s.depth, s.width)
		b.Run(fmt.Sprintf("%s", s.name), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				PathsByCarryingDown(root)
			}
		})
	}
}
