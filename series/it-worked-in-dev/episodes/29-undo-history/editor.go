package undohistory

// An editor over one document, with undo and redo. Edits are inserts and
// deletes at a byte position; that is what a keystroke, a paste or a cut
// turns into.

func splice(doc []byte, pos, del int, ins []byte) []byte {
	out := make([]byte, 0, len(doc)-del+len(ins))
	out = append(out, doc[:pos]...)
	out = append(out, ins...)
	return append(out, doc[pos+del:]...)
}

// SnapshotEditor is what I would write. Before every edit, keep a copy of the
// document. Undo puts the last copy back; redo does the same the other way.
// It cannot get an undo wrong, because it never works out what an edit did.
type SnapshotEditor struct {
	Doc        []byte
	undo, redo [][]byte
}

func (e *SnapshotEditor) Insert(pos int, text string) {
	e.undo = append(e.undo, e.Doc)
	e.redo = e.redo[:0] // a new edit forks history: the old future is gone
	e.Doc = splice(e.Doc, pos, 0, []byte(text))
}

func (e *SnapshotEditor) Delete(pos, n int) {
	e.undo = append(e.undo, e.Doc)
	e.redo = e.redo[:0]
	e.Doc = splice(e.Doc, pos, n, nil)
}

func (e *SnapshotEditor) Undo() bool {
	if len(e.undo) == 0 {
		return false
	}
	e.redo = append(e.redo, e.Doc)
	e.Doc = e.undo[len(e.undo)-1]
	e.undo = e.undo[:len(e.undo)-1]
	return true
}

func (e *SnapshotEditor) Redo() bool {
	if len(e.redo) == 0 {
		return false
	}
	e.undo = append(e.undo, e.Doc)
	e.Doc = e.redo[len(e.redo)-1]
	e.redo = e.redo[:len(e.redo)-1]
	return true
}

// Held is the history's memory: every snapshot kept for undo or redo.
func (e *SnapshotEditor) Held() int {
	n := 0
	for _, s := range e.undo {
		n += len(s)
	}
	for _, s := range e.redo {
		n += len(s)
	}
	return n
}

// op is one edit and everything needed to reverse it: where, and the text
// that went in or came out. An insert is undone by deleting that text; a
// delete by inserting it back.
type op struct {
	pos    int
	text   []byte
	insert bool
}

// OpEditor keeps a stack of edits instead of a stack of documents.
type OpEditor struct {
	Doc        []byte
	undo, redo []op

	// Copying makes every edit build a new document, as SnapshotEditor must.
	// It is here only so the episode can measure what editing in place is
	// worth on its own.
	Copying bool
}

// apply edits the document in place. Nothing else holds a reference to the
// document - history is edits, not copies - so there is no reason to copy it.
func (e *OpEditor) apply(o op) {
	if e.Copying {
		if o.insert {
			e.Doc = splice(e.Doc, o.pos, 0, o.text)
		} else {
			e.Doc = splice(e.Doc, o.pos, len(o.text), nil)
		}
		return
	}
	if o.insert {
		n, at := len(o.text), o.pos
		e.Doc = append(e.Doc, o.text...) // grow by len(text); contents fixed below
		copy(e.Doc[at+n:], e.Doc[at:len(e.Doc)-n])
		copy(e.Doc[at:], o.text)
	} else {
		n, at := len(o.text), o.pos
		copy(e.Doc[at:], e.Doc[at+n:])
		e.Doc = e.Doc[:len(e.Doc)-n]
	}
}

func (o op) inverse() op { return op{pos: o.pos, text: o.text, insert: !o.insert} }

func (e *OpEditor) Insert(pos int, text string) {
	o := op{pos: pos, text: []byte(text), insert: true}
	e.apply(o)
	e.undo = append(e.undo, o)
	e.redo = e.redo[:0]
}

func (e *OpEditor) Delete(pos, n int) {
	// Keep the deleted text: it is the only way to put it back.
	o := op{pos: pos, text: append([]byte(nil), e.Doc[pos:pos+n]...), insert: false}
	e.apply(o)
	e.undo = append(e.undo, o)
	e.redo = e.redo[:0]
}

func (e *OpEditor) Undo() bool {
	if len(e.undo) == 0 {
		return false
	}
	o := e.undo[len(e.undo)-1]
	e.undo = e.undo[:len(e.undo)-1]
	e.apply(o.inverse()) // the most recent edit is the only one safe to reverse
	e.redo = append(e.redo, o)
	return true
}

func (e *OpEditor) Redo() bool {
	if len(e.redo) == 0 {
		return false
	}
	o := e.redo[len(e.redo)-1]
	e.redo = e.redo[:len(e.redo)-1]
	e.apply(o)
	e.undo = append(e.undo, o)
	return true
}

func (e *OpEditor) Held() int {
	n := 0
	for _, o := range e.undo {
		n += len(o.text) + 24
	}
	for _, o := range e.redo {
		n += len(o.text) + 24
	}
	return n
}
