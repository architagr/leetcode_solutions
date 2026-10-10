package columnsnotrows

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

// full is a complete tree of the given depth, ids in level order: node i
// has children 2i+1 and 2i+2.
func full(depth int) *Node {
	var build func(i, d int) *Node
	build = func(i, d int) *Node {
		if d == 0 {
			return nil
		}
		return &Node{ID: i, Left: build(2*i+1, d-1), Right: build(2*i+2, d-1)}
	}
	return build(0, depth)
}

// grown is a tree grown by inserting n random keys into a binary search
// tree, the way an unpruned decision tree ends up lopsided.
func grown(n int, seed int64) *Node {
	r := rand.New(rand.NewSource(seed))
	type kn struct {
		k int
		n *Node
	}
	var root *kn
	nodes := map[*Node]*kn{}
	left, right := map[*kn]*kn{}, map[*kn]*kn{}
	for i := 0; i < n; i++ {
		k := r.Int()
		x := &kn{k, &Node{ID: i}}
		nodes[x.n] = x
		if root == nil {
			root = x
			continue
		}
		cur := root
		for {
			if k < cur.k {
				if left[cur] == nil {
					left[cur] = x
					cur.n.Left = x.n
					break
				}
				cur = left[cur]
			} else {
				if right[cur] == nil {
					right[cur] = x
					cur.n.Right = x.n
					break
				}
				cur = right[cur]
			}
		}
	}
	return root.n
}

// The example from the write-up: 0 has children 1 and 2; 1 has 3 and 4; 2
// has 5 and 6. 4 and 5 share column 0 and row 2, and 4 is on the left.
func TestExample(t *testing.T) {
	want := [][]int{{3}, {1}, {0, 4, 5}, {2}, {6}}
	for name, fn := range funcs {
		if got := fn(full(3)); !reflect.DeepEqual(got, want) {
			t.Fatalf("%s: got %v, want %v", name, got, want)
		}
	}
}

var funcs = map[string]func(*Node) [][]int{"sort": ColumnsBySort, "level": ColumnsByLevel}

// ColumnsBySortUnstable is the first version I wrote: sort.Slice, which does
// not promise to keep equal elements in order.
func ColumnsBySortUnstable(root *Node) [][]int {
	var all []placed
	var walk func(n *Node, col, depth int)
	walk = func(n *Node, col, depth int) {
		if n == nil {
			return
		}
		all = append(all, placed{n.ID, col, depth})
		walk(n.Left, col-1, depth+1)
		walk(n.Right, col+1, depth+1)
	}
	walk(root, 0, 0)
	sort.Slice(all, func(i, j int) bool {
		if all[i].col != all[j].col {
			return all[i].col < all[j].col
		}
		return all[i].depth < all[j].depth
	})
	var out [][]int
	for i, p := range all {
		if i == 0 || p.col != all[i-1].col {
			out = append(out, nil)
		}
		out[len(out)-1] = append(out[len(out)-1], p.id)
	}
	return out
}

// sort.Slice is an insertion sort below 12 elements, which is stable, so the
// unstable version passes on any tree small enough to check by eye.
func TestUnstableSortPassesSmallFailsLarge(t *testing.T) {
	if !reflect.DeepEqual(ColumnsBySortUnstable(full(3)), ColumnsByLevel(full(3))) {
		t.Fatal("expected the unstable sort to be right on 7 nodes")
	}
	wrong := 0
	for d := 4; d <= 12; d++ {
		if !reflect.DeepEqual(ColumnsBySortUnstable(full(d)), ColumnsByLevel(full(d))) {
			wrong++
		}
	}
	if wrong == 0 {
		t.Fatal("expected the unstable sort to break a tie somewhere on a bigger tree")
	}
	t.Logf("unstable sort wrong on %d of 9 complete trees of depth 4 to 12", wrong)
}

func TestBothAgree(t *testing.T) {
	for i := 0; i < 300; i++ {
		root := grown(1+i*7, int64(i))
		if !reflect.DeepEqual(ColumnsBySort(root), ColumnsByLevel(root)) {
			t.Fatalf("tree %d: sort and level disagree", i)
		}
	}
	for _, s := range shapes {
		if !reflect.DeepEqual(ColumnsBySort(s.root), ColumnsByLevel(s.root)) {
			t.Fatalf("%s: disagree", s.name)
		}
	}
}

func TestShapes(t *testing.T) {
	for _, s := range shapes {
		cols := ColumnsByLevel(s.root)
		n, widest := 0, 0
		for _, c := range cols {
			n += len(c)
			if len(c) > widest {
				widest = len(c)
			}
		}
		t.Logf("%-8s %8d nodes, %5d columns, widest %7d", s.name, n, len(cols), widest)
	}
}

// ColumnsBySortThreeKeys fixes the tie with a key instead of with stability:
// the walk's own order breaks it, so sort.Slice is enough.
func ColumnsBySortThreeKeys(root *Node) [][]int {
	type seq struct{ id, col, depth, at int }
	var all []seq
	var walk func(n *Node, col, depth int)
	walk = func(n *Node, col, depth int) {
		if n == nil {
			return
		}
		all = append(all, seq{n.ID, col, depth, len(all)})
		walk(n.Left, col-1, depth+1)
		walk(n.Right, col+1, depth+1)
	}
	walk(root, 0, 0)
	sort.Slice(all, func(i, j int) bool {
		a, b := all[i], all[j]
		if a.col != b.col {
			return a.col < b.col
		}
		if a.depth != b.depth {
			return a.depth < b.depth
		}
		return a.at < b.at
	})
	var out [][]int
	for i, p := range all {
		if i == 0 || p.col != all[i-1].col {
			out = append(out, nil)
		}
		out[len(out)-1] = append(out[len(out)-1], p.id)
	}
	return out
}

func TestThreeKeysAgrees(t *testing.T) {
	for _, s := range shapes {
		if !reflect.DeepEqual(ColumnsBySortThreeKeys(s.root), ColumnsByLevel(s.root)) {
			t.Fatalf("%s: three keys disagree", s.name)
		}
	}
}

func BenchmarkColumnsBySortThreeKeys(b *testing.B) { run(b, ColumnsBySortThreeKeys) }

// counted is the same sort through sort.Stable, so the swaps can be counted.
type counted struct {
	p           []placed
	less, swaps int
}

func (c *counted) Len() int { return len(c.p) }
func (c *counted) Less(i, j int) bool {
	c.less++
	if c.p[i].col != c.p[j].col {
		return c.p[i].col < c.p[j].col
	}
	return c.p[i].depth < c.p[j].depth
}
func (c *counted) Swap(i, j int) { c.swaps++; c.p[i], c.p[j] = c.p[j], c.p[i] }

func TestSortWork(t *testing.T) {
	for _, s := range shapes {
		var all []placed
		var walk func(n *Node, col, depth int)
		walk = func(n *Node, col, depth int) {
			if n == nil {
				return
			}
			all = append(all, placed{n.ID, col, depth})
			walk(n.Left, col-1, depth+1)
			walk(n.Right, col+1, depth+1)
		}
		walk(s.root, 0, 0)
		c := &counted{p: all}
		sort.Stable(c)
		n := len(all)
		t.Logf("%-6s %8d nodes  |  stable sort: %10d comparisons (%2.0f per node), %11d swaps (%3.0f per node)",
			s.name, n, c.less, float64(c.less)/float64(n), c.swaps, float64(c.swaps)/float64(n))
	}
}
func BenchmarkColumnsBySortUnstable(b *testing.B) { run(b, ColumnsBySortUnstable) }
