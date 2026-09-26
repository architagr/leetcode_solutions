package circulardependency

import (
	"math/rand"
	"testing"
)

// layered is a monorepo of n packages in `layers` layers. Every package in
// layer L imports up to three packages from layer L-1, so the longest import
// chain is `layers` long. order controls how the packages are listed:
// "deps-first" lists libraries before the apps that use them; "apps-first" is
// the reverse, which is what an alphabetical apps/ ... libs/ tree gives you.
func layered(n, layers int, order string, seed int64) [][]int {
	r := rand.New(rand.NewSource(seed))
	per := n / layers
	layerOf := func(i int) int { return i / per }
	imports := make([][]int, n)
	for i := per; i < n; i++ {
		l := layerOf(i)
		if l >= layers {
			l = layers - 1
		}
		lo := (l - 1) * per
		k := 1 + r.Intn(3)
		for j := 0; j < k; j++ {
			imports[i] = append(imports[i], lo+r.Intn(per))
		}
	}
	if order == "apps-first" {
		return relabel(imports, func(i int) int { return n - 1 - i })
	}
	return imports
}

// relabel renumbers the packages so the list is in a different order without
// changing which package imports which.
func relabel(imports [][]int, to func(int) int) [][]int {
	out := make([][]int, len(imports))
	for p, ds := range imports {
		q := to(p)
		for _, d := range ds {
			out[q] = append(out[q], to(d))
		}
	}
	return out
}

// withCycle adds one import that points back up the layers: a low-level
// package importing an app that depends on it.
func withCycle(imports [][]int) [][]int {
	out := make([][]int, len(imports))
	for i := range imports {
		out[i] = append([]int(nil), imports[i]...)
	}
	// find a package with imports, follow its first import down to the
	// bottom, and make that bottom package import the top one
	top := len(out) - 1
	for len(out[top]) == 0 {
		top--
	}
	bottom := top
	for len(out[bottom]) > 0 {
		bottom = out[bottom][0]
	}
	out[bottom] = append(out[bottom], top)
	return out
}

func validOrder(t *testing.T, imports [][]int, order []int) {
	t.Helper()
	at := make([]int, len(imports))
	for i := range at {
		at[i] = -1
	}
	for i, p := range order {
		at[p] = i
	}
	for p, ds := range imports {
		if at[p] < 0 {
			continue
		}
		for _, d := range ds {
			if at[d] < 0 || at[d] > at[p] {
				t.Fatalf("package %d is built before its import %d", p, d)
			}
		}
	}
}

func TestBothBuildValidOrders(t *testing.T) {
	cases := [][][]int{
		{}, {{}}, {{1}, {}}, {{1}, {0}}, {{1}, {2}, {0}, {}},
		layered(300, 10, "deps-first", 1), layered(300, 10, "apps-first", 2),
		withCycle(layered(300, 10, "apps-first", 3)),
	}
	for _, s := range shapes {
		cases = append(cases, s.imports)
	}
	for _, im := range cases {
		a, okA, _ := OrderByPasses(im)
		b, okB := OrderByKahn(im)
		f, okF := OrderByKahnFlat(im)
		if okF != okB || len(f) != len(b) {
			t.Fatalf("flat Kahn built %d, Kahn built %d", len(f), len(b))
		}
		validOrder(t, im, f)
		if okA != okB || len(a) != len(b) {
			t.Fatalf("%d packages: passes built %d (ok=%v), Kahn built %d (ok=%v)",
				len(im), len(a), okA, len(b), okB)
		}
		validOrder(t, im, a)
		validOrder(t, im, b)
		if !okB {
			c := Cycle(im, b)
			if len(c) == 0 {
				t.Fatal("stuck packages but no cycle named")
			}
			for i, p := range c {
				next := c[(i+1)%len(c)]
				found := false
				for _, d := range im[p] {
					if d == next {
						found = true
					}
				}
				if !found {
					t.Fatalf("named cycle %v: %d does not import %d", c, p, next)
				}
			}
		}
	}
}

func TestPasses(t *testing.T) {
	for _, s := range shapes {
		order, ok, passes := OrderByPasses(s.imports)
		b, _ := OrderByKahn(s.imports)
		msg := ""
		if !ok {
			msg = "  cycle: " + itoa(len(Cycle(s.imports, b))) + " packages"
		}
		t.Logf("%-18s %5d packages, built %5d, passes %4d%s",
			s.name, len(s.imports), len(order), passes, msg)
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for ; n > 0; n /= 10 {
		s = string(rune('0'+n%10)) + s
	}
	return s
}
