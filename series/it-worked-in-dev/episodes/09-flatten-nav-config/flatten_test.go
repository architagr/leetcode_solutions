package flattennavconfig

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func menu(depth, width int, n *int) *Item {
	it := &Item{Label: fmt.Sprintf("item%d", *n)}
	*n++
	if depth == 0 {
		return it
	}
	for w := 0; w < width; w++ {
		it.Children = append(it.Children, menu(depth-1, width, n))
	}
	return it
}

func TestAllThreeAgree(t *testing.T) {
	n := 0
	cases := []struct {
		name string
		root *Item
	}{
		{"one item", &Item{Label: "home"}},
		{"one level", func() *Item { n = 0; return menu(1, 4, &n) }()},
		{"balanced 3x3", func() *Item { n = 0; return menu(3, 3, &n) }()},
		{"deep chain", func() *Item { n = 0; return menu(40, 1, &n) }()},
		{"wide", func() *Item { n = 0; return menu(1, 50, &n) }()},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			a := FlattenByConcat(c.root)
			b := FlattenByAppending(c.root)
			d := FlattenPresized(c.root)
			if !reflect.DeepEqual(a, b) {
				t.Fatalf("concat and appending disagreed:\n %v\n %v", a, b)
			}
			if !reflect.DeepEqual(a, d) {
				t.Fatalf("concat and presized disagreed")
			}
		})
	}
}

func TestKnownRows(t *testing.T) {
	root := &Item{Label: "products", Children: []*Item{
		{Label: "laptops", Children: []*Item{{Label: "gaming"}}},
		{Label: "phones"},
	}}
	want := []Row{{"products", 0}, {"laptops", 1}, {"gaming", 2}, {"phones", 1}}
	for name, fn := range map[string]func(*Item) []Row{
		"concat": FlattenByConcat, "appending": FlattenByAppending, "presized": FlattenPresized,
	} {
		if got := fn(root); !reflect.DeepEqual(got, want) {
			t.Errorf("%s = %v, want %v", name, got, want)
		}
	}
}

func TestNilIsEmpty(t *testing.T) {
	for name, fn := range map[string]func(*Item) []Row{
		"concat": FlattenByConcat, "appending": FlattenByAppending, "presized": FlattenPresized,
	} {
		if got := fn(nil); len(got) != 0 {
			t.Errorf("%s(nil) = %v", name, got)
		}
	}
}

func TestAllAgreeOnRandomMenus(t *testing.T) {
	rng := rand.New(rand.NewSource(43))
	for trial := 0; trial < 1000; trial++ {
		root := &Item{Label: "root"}
		all := []*Item{root}
		for i := 1; i < rng.Intn(70)+2; i++ {
			p := all[rng.Intn(len(all))]
			c := &Item{Label: fmt.Sprintf("i%d", i)}
			p.Children = append(p.Children, c)
			all = append(all, c)
		}
		a, b, d := FlattenByConcat(root), FlattenByAppending(root), FlattenPresized(root)
		if !reflect.DeepEqual(a, b) || !reflect.DeepEqual(a, d) {
			t.Fatalf("trial %d disagreed", trial)
		}
	}
}

// The write-up claims a row is copied once per ancestor, so on a chain the
// total copies are the sum of the depths. Counted here rather than claimed.
func TestConcatRowCopyCount(t *testing.T) {
	for _, d := range []int{10, 50, 200} {
		n := 0
		root := menu(d, 1, &n)
		copies := 0
		var count func(it *Item) int
		count = func(it *Item) int {
			rows := 1
			for _, c := range it.Children {
				child := count(c)
				copies += child // every row of the child is copied into the parent
				rows += child
			}
			return rows
		}
		count(root)
		// items sit at depths 0..d; the one at depth k is copied k times
		want := d * (d + 1) / 2
		if copies != want {
			t.Errorf("depth %d: %d row copies, want %d", d, copies, want)
		}
	}
}
