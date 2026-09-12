package lopsidedcategories

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

// chainUnder gives c a single branch `n` levels deep, which is how a category
// tree becomes lopsided in practice.
func chainUnder(c *Category, n int, prefix string) {
	cur := c
	for i := 0; i < n; i++ {
		next := &Category{Name: fmt.Sprintf("%s-%d", prefix, i)}
		cur.Children = append(cur.Children, next)
		cur = next
	}
}

func balanced(depth, width int, prefix string) *Category {
	root := &Category{Name: prefix}
	if depth == 0 {
		return root
	}
	for w := 0; w < width; w++ {
		root.Children = append(root.Children,
			balanced(depth-1, width, fmt.Sprintf("%s.%d", prefix, w)))
	}
	return root
}

func sorted(s []string) []string {
	out := append([]string(nil), s...)
	sort.Strings(out)
	return out
}

// The two functions must find the same categories. They report them in
// different orders, which is checked separately below.
func TestBothFindTheSameCategories(t *testing.T) {
	lop := &Category{Name: "shop"}
	lop.Children = []*Category{{Name: "books"}, {Name: "electronics"}}
	chainUnder(lop.Children[1], 6, "e")

	cases := []struct {
		name string
		root *Category
	}{
		{"perfectly balanced", balanced(3, 3, "root")},
		{"one deep branch", lop},
		{"single node", &Category{Name: "only"}},
		{"one child is never lopsided", &Category{Name: "r", Children: []*Category{{Name: "a"}}}},
		{"wide and even", balanced(1, 20, "wide")},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			a := sorted(LopsidedByAskingTwice(c.root))
			b := sorted(LopsidedInOnePass(c.root))
			if !reflect.DeepEqual(a, b) {
				t.Fatalf("disagreed:\n  twice   %v\n  onepass %v", a, b)
			}
		})
	}
}

func TestKnownResult(t *testing.T) {
	shop := &Category{Name: "shop"}
	shop.Children = []*Category{{Name: "books"}, {Name: "electronics"}}
	chainUnder(shop.Children[1], 4, "e")

	want := []string{"shop"}
	if got := sorted(LopsidedByAskingTwice(shop)); !reflect.DeepEqual(got, want) {
		t.Errorf("asking twice = %v, want %v", got, want)
	}
	if got := sorted(LopsidedInOnePass(shop)); !reflect.DeepEqual(got, want) {
		t.Errorf("one pass = %v, want %v", got, want)
	}
}

// The orders genuinely differ, and the write-up says so, so it is pinned here
// rather than left as a surprise: asking twice reports a parent before the
// children it contains, and one pass cannot, because a parent's verdict is not
// known until its children have returned.
func TestReportingOrderDiffers(t *testing.T) {
	root := &Category{Name: "root"}
	inner := &Category{Name: "inner"}
	root.Children = []*Category{{Name: "shallow"}, inner}
	chainUnder(inner, 5, "i")
	inner.Children = append(inner.Children, &Category{Name: "stub"})

	twice := LopsidedByAskingTwice(root)
	once := LopsidedInOnePass(root)
	if len(twice) < 2 || len(once) < 2 {
		t.Skipf("need at least two lopsided nodes, got %v / %v", twice, once)
	}
	if reflect.DeepEqual(twice, once) {
		t.Errorf("expected the orders to differ; both were %v", twice)
	}
	if !reflect.DeepEqual(sorted(twice), sorted(once)) {
		t.Errorf("same set expected: %v vs %v", twice, once)
	}
}

func TestBothAgreeOnRandomTrees(t *testing.T) {
	rng := rand.New(rand.NewSource(31))
	for trial := 0; trial < 1000; trial++ {
		root := &Category{Name: "c0"}
		all := []*Category{root}
		for i := 1; i < rng.Intn(60)+2; i++ {
			p := all[rng.Intn(len(all))]
			c := &Category{Name: fmt.Sprintf("c%d", i)}
			p.Children = append(p.Children, c)
			all = append(all, c)
		}
		if a, b := sorted(LopsidedByAskingTwice(root)), sorted(LopsidedInOnePass(root)); !reflect.DeepEqual(a, b) {
			t.Fatalf("trial %d disagreed:\n  %v\n  %v", trial, a, b)
		}
	}
}

// The write-up claims depth() is called once per node per ancestor, so a chain
// of n costs n(n-1)/2 calls below the root. Counted here.
func TestDepthCallCount(t *testing.T) {
	for _, n := range []int{10, 50, 200} {
		root := &Category{Name: "r"}
		chainUnder(root, n, "c")
		calls := 0
		var counted func(*Category) int
		counted = func(c *Category) int {
			calls++
			best := 0
			for _, ch := range c.Children {
				if d := counted(ch); d > best {
					best = d
				}
			}
			return best + 1
		}
		var walk func(*Category)
		walk = func(c *Category) {
			for _, ch := range c.Children {
				counted(ch)
			}
			for _, ch := range c.Children {
				walk(ch)
			}
		}
		walk(root)
		if want := n * (n + 1) / 2; calls != want {
			t.Errorf("chain of %d: %d depth calls, want %d", n, calls, want)
		}
	}
}
