package paginatewithoutloading

// The price index again: a binary search tree keyed by price. The catalogue
// page lists products sorted by price, 20 to a page.
type Node struct {
	Price       int
	Left, Right *Node
}

// Iterator is day 56's: the whole in-order sequence built up front, then
// handed out one at a time.
type Iterator struct {
	all []int
	i   int
}

func NewIterator(root *Node) *Iterator {
	it := &Iterator{}
	var walk func(n *Node)
	walk = func(n *Node) {
		if n == nil {
			return
		}
		walk(n.Left)
		it.all = append(it.all, n.Price)
		walk(n.Right)
	}
	walk(root)
	return it
}

func (it *Iterator) HasNext() bool { return it.i < len(it.all) }
func (it *Iterator) Next() int     { v := it.all[it.i]; it.i++; return v }

// PageByList is what I would write with that iterator: ask for page p, skip
// p*size, take size.
func PageByList(root *Node, page, size int) []int {
	it := NewIterator(root)
	for skip := page * size; skip > 0 && it.HasNext(); skip-- {
		it.Next()
	}
	out := make([]int, 0, size)
	for len(out) < size && it.HasNext() {
		out = append(out, it.Next())
	}
	return out
}

// StackIterator keeps only the path it has not finished: the left spine
// below wherever it is. Next pops the smallest, then pushes the left spine of
// that node's right subtree. It never holds more than the tree's depth.
type StackIterator struct {
	stack []*Node
}

func NewStackIterator(root *Node) *StackIterator {
	it := &StackIterator{}
	it.pushLeft(root)
	return it
}

func (it *StackIterator) pushLeft(n *Node) {
	for ; n != nil; n = n.Left {
		it.stack = append(it.stack, n)
	}
}

func (it *StackIterator) HasNext() bool { return len(it.stack) > 0 }

func (it *StackIterator) Next() int {
	n := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	it.pushLeft(n.Right) // everything in n's right subtree comes next, smallest first
	return n.Price
}

// PageByStack is the same page-number API on the stack iterator: nothing is
// built up front, but page p still has to step past p*size prices.
func PageByStack(root *Node, page, size int) []int {
	it := NewStackIterator(root)
	for skip := page * size; skip > 0 && it.HasNext(); skip-- {
		it.Next()
	}
	out := make([]int, 0, size)
	for len(out) < size && it.HasNext() {
		out = append(out, it.Next())
	}
	return out
}

// SeekAfter starts the stack iterator just after a price, the one the last
// page ended on. It descends once, keeping only the nodes still to come:
// those greater than the cursor, whose left side it went into.
func SeekAfter(root *Node, after int) *StackIterator {
	it := &StackIterator{}
	for n := root; n != nil; {
		if n.Price > after {
			it.stack = append(it.stack, n) // n and its right side are still to come
			n = n.Left
		} else {
			n = n.Right // n and everything on its left are already shown
		}
	}
	return it
}

// PageAfter is the cursor API: "the 20 prices after the last one you saw".
func PageAfter(root *Node, after, size int) []int {
	it := SeekAfter(root, after)
	out := make([]int, 0, size)
	for len(out) < size && it.HasNext() {
		out = append(out, it.Next())
	}
	return out
}
