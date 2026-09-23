package collapsedtreeview

import "testing"

// The shapes a comment thread takes. wide_50k is here because it is the one
// where holding every level costs almost nothing extra.
var shapes = []struct {
	name string
	root *Comment
}{
	{"thread_1k", thread(5, 4)},  // 1,365 comments, five levels deep
	{"thread_5k", thread(6, 4)},  // 5,461 comments
	{"thread_56k", thread(6, 6)}, // 55,987 comments
	{"wide_50k", wide(50000)},    // a popular post: one comment, 50,000 replies
	{"chain_2k", chain(2000)},    // two people arguing
}

func run(b *testing.B, fn func(root *Comment)) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.root)
			}
		})
	}
}

func BenchmarkVisibleByAllLevels(b *testing.B) {
	run(b, func(root *Comment) { VisibleByAllLevels(root) })
}

func BenchmarkVisibleByDepthMap(b *testing.B) {
	run(b, func(root *Comment) { VisibleByDepthMap(root) })
}

func BenchmarkVisibleByFirstArrival(b *testing.B) {
	run(b, func(root *Comment) { VisibleByFirstArrival(root) })
}

func BenchmarkPreviewsByDepthMap(b *testing.B) {
	run(b, func(root *Comment) { PreviewsByDepthMap(root) })
}

func BenchmarkPreviewsByFirstArrival(b *testing.B) {
	run(b, func(root *Comment) { PreviewsByFirstArrival(root) })
}
