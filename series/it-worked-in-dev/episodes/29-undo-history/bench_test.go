package undohistory

import "testing"

var shapes = []struct {
	name         string
	start, edits int
	seed         int64
}{
	{"note_2k", 2000, 500, 11},      // a note: 2 KB, 500 edits
	{"doc_50k", 50000, 2000, 12},    // a long document: 50 KB, 2,000 edits
	{"spec_200k", 200000, 1000, 13}, // a spec: 200 KB, 1,000 edits
}

// The whole session: every edit, then every one of them undone.
func run(b *testing.B, mk func(string) editor) {
	for _, s := range shapes {
		base, steps := session(s.start, s.edits, s.seed)
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				e := mk(base)
				play(e, steps)
				for e.Undo() {
				}
			}
		})
	}
}

func BenchmarkSnapshotEditor(b *testing.B) {
	run(b, func(s string) editor { return &SnapshotEditor{Doc: []byte(s)} })
}

func BenchmarkOpEditor(b *testing.B) {
	run(b, func(s string) editor { return &OpEditor{Doc: []byte(s)} })
}

func BenchmarkOpEditorCopying(b *testing.B) {
	run(b, func(s string) editor { return &OpEditor{Doc: []byte(s), Copying: true} })
}
