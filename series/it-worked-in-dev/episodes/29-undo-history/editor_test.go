package undohistory

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
)

type editor interface {
	Insert(pos int, text string)
	Delete(pos, n int)
	Undo() bool
	Redo() bool
	Held() int
}

func doc(e editor) []byte {
	switch x := e.(type) {
	case *SnapshotEditor:
		return x.Doc
	case *OpEditor:
		return x.Doc
	}
	return nil
}

// session is a writing session: mostly typing a word or two, sometimes
// deleting a few characters, now and then pasting a paragraph.
type step struct {
	kind   byte // 'i', 'd', 'u', 'r'
	pos, n int
	text   string
}

func session(start, edits int, seed int64) (string, []step) {
	r := rand.New(rand.NewSource(seed))
	base := make([]byte, start)
	for i := range base {
		base[i] = "abcdefghij klmnopqrst\n"[r.Intn(22)]
	}
	size := start
	var steps []step
	for i := 0; i < edits; i++ {
		switch k := r.Intn(20); {
		case k < 14 || size < 20:
			t := "word" + fmt.Sprint(i%97) + " "
			steps = append(steps, step{kind: 'i', pos: r.Intn(size + 1), text: t})
			size += len(t)
		case k < 19:
			n := 1 + r.Intn(8)
			steps = append(steps, step{kind: 'd', pos: r.Intn(size - n), n: n})
			size -= n
		default:
			t := string(bytes.Repeat([]byte("pasted paragraph. "), 12))
			steps = append(steps, step{kind: 'i', pos: r.Intn(size + 1), text: t})
			size += len(t)
		}
	}
	return string(base), steps
}

func play(e editor, steps []step) {
	for _, s := range steps {
		switch s.kind {
		case 'i':
			e.Insert(s.pos, s.text)
		case 'd':
			e.Delete(s.pos, s.n)
		}
	}
}

func TestBothEditorsAgree(t *testing.T) {
	base, steps := session(2000, 400, 1)
	a := &SnapshotEditor{Doc: []byte(base)}
	b := &OpEditor{Doc: []byte(base)}
	c := &OpEditor{Doc: []byte(base), Copying: true}
	history := [][]byte{[]byte(base)}
	for _, s := range steps {
		play(a, []step{s})
		play(b, []step{s})
		play(c, []step{s})
		if !bytes.Equal(a.Doc, b.Doc) || !bytes.Equal(a.Doc, c.Doc) {
			t.Fatal("documents diverged while editing")
		}
		history = append(history, append([]byte(nil), a.Doc...))
	}
	// undo everything, checking each step against the recorded history
	for i := len(history) - 2; i >= 0; i-- {
		a.Undo()
		b.Undo()
		c.Undo()
		if !bytes.Equal(a.Doc, history[i]) || !bytes.Equal(b.Doc, history[i]) || !bytes.Equal(c.Doc, history[i]) {
			t.Fatalf("undo to step %d gave the wrong document", i)
		}
	}
	if a.Undo() || b.Undo() {
		t.Fatal("undo past the start should do nothing")
	}
	// redo half, then a new edit must drop the rest of the redo history
	for i := 0; i < 200; i++ {
		a.Redo()
		b.Redo()
	}
	a.Insert(0, "fork ")
	b.Insert(0, "fork ")
	if a.Redo() || b.Redo() {
		t.Fatal("a new edit should clear redo")
	}
	if !bytes.Equal(a.Doc, b.Doc) {
		t.Fatal("documents diverged after redo and fork")
	}
}

func TestHeld(t *testing.T) {
	for _, s := range shapes {
		base, steps := session(s.start, s.edits, s.seed)
		a := &SnapshotEditor{Doc: []byte(base)}
		b := &OpEditor{Doc: []byte(base)}
		play(a, steps)
		play(b, steps)
		t.Logf("%-10s document %7d bytes after %5d edits  |  history held: snapshots %11d bytes, operations %7d bytes",
			s.name, len(a.Doc), s.edits, a.Held(), b.Held())
	}
}
