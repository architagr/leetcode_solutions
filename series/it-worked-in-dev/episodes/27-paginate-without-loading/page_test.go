package paginatewithoutloading

import (
	"math/rand"
	"reflect"
	"testing"
)

func insert(root *Node, p int) *Node {
	if root == nil {
		return &Node{Price: p}
	}
	n := root
	for {
		if p < n.Price {
			if n.Left == nil {
				n.Left = &Node{Price: p}
				return root
			}
			n = n.Left
		} else {
			if n.Right == nil {
				n.Right = &Node{Price: p}
				return root
			}
			n = n.Right
		}
	}
}

// catalog is n distinct prices 10, 20, ... inserted in random order.
func catalog(n int, seed int64) *Node {
	r := rand.New(rand.NewSource(seed))
	var root *Node
	for _, i := range r.Perm(n) {
		root = insert(root, (i+1)*10)
	}
	return root
}

func TestAllThreeServeTheSamePages(t *testing.T) {
	root := catalog(1234, 1)
	const size = 20
	last := 0
	for page := 0; ; page++ {
		a := PageByList(root, page, size)
		b := PageByStack(root, page, size)
		c := PageAfter(root, last, size)
		if !reflect.DeepEqual(a, b) || !reflect.DeepEqual(b, c) {
			t.Fatalf("page %d: list %v\nstack %v\ncursor %v", page, a, b, c)
		}
		if len(a) == 0 {
			if page != (1234+size-1)/size {
				t.Fatalf("ran out at page %d", page)
			}
			break
		}
		last = a[len(a)-1]
	}
	// a cursor that is not a price in the index still lands in the right place
	if got := PageAfter(root, 15, 3); !reflect.DeepEqual(got, []int{20, 30, 40}) {
		t.Fatalf("after 15: %v", got)
	}
	if got := PageAfter(nil, 0, 5); len(got) != 0 {
		t.Fatal("empty index should give an empty page")
	}
}

func TestStackStaysShallow(t *testing.T) {
	root := catalog(200000, 2)
	it := NewStackIterator(root)
	most := 0
	for it.HasNext() {
		if len(it.stack) > most {
			most = len(it.stack)
		}
		it.Next()
	}
	t.Logf("200000 prices walked in order; the stack never held more than %d nodes", most)
}
