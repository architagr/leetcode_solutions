package flattennestedlistiterator

// frame is one list being walked, plus how far into it we are. The stack of
// frames is the iterator's whole memory - it is the path from the outermost
// list down to wherever the cursor currently sits.
type frame struct {
	list []*NestedInteger
	i    int
}

type NestedIterator struct {
	stack []*frame
}

// Constructor does no work beyond wrapping the outer list. Nothing is walked
// and nothing is flattened, so construction is O(1) however deep or large the
// structure is.
func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{stack: []*frame{{list: nestedList}}}
}

// HasNext does the advancing, which is what makes this lazy: it moves the
// stack forward only until the cursor is sitting on an integer, then stops.
//
// Three cases, and the loop keeps going until one of them settles it:
// a finished list is popped, a nested list is descended into, and an integer
// means there is something to return.
func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		top := this.stack[len(this.stack)-1]

		if top.i == len(top.list) {
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}

		item := top.list[top.i]
		if item.IsInteger() {
			return true
		}

		// Step past the nested list in the parent before descending, so the
		// parent resumes after it rather than re-entering it forever.
		top.i++
		this.stack = append(this.stack, &frame{list: item.GetList()})
	}
	return false
}

// Next positions the cursor itself rather than assuming HasNext was just
// called. The problem only promises that a value exists when Next is called,
// not that the caller asked first, and HasNext is idempotent and cheap once
// the cursor is already on an integer - so paying for it here costs nothing
// and removes the ordering requirement between the two methods.
func (this *NestedIterator) Next() int {
	this.HasNext()
	top := this.stack[len(this.stack)-1]
	v := top.list[top.i].GetInteger()
	top.i++
	return v
}
