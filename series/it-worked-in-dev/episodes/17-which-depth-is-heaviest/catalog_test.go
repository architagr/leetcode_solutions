package whichdepthisheaviest

import (
	"fmt"
	"testing"
)

// catalog builds a category tree `depth` levels below the root where every
// non-leaf has `branch` children. Products live on the leaves, the way a real
// catalogue files them on the shelf rather than the department.
func catalog(depth, branch, skusPerLeaf int) *Category {
	n := 0
	var build func(level int) *Category
	build = func(level int) *Category {
		n++
		c := &Category{ID: fmt.Sprintf("c%d", n)}
		if level == depth {
			c.SKUs = skusPerLeaf
			return c
		}
		c.SKUs = 1 // the odd product filed on a department page
		for i := 0; i < branch; i++ {
			c.Children = append(c.Children, build(level+1))
		}
		return c
	}
	return build(0)
}

// flat is a catalogue nobody organised: one root, every category directly
// under it.
func flat(width, skus int) *Category {
	root := &Category{ID: "root"}
	for i := 0; i < width; i++ {
		root.Children = append(root.Children, &Category{ID: fmt.Sprintf("c%d", i), SKUs: skus})
	}
	return root
}

// chain is a category tree built by an importer that nested every path
// segment: 2,000 levels, one category each.
func chain(depth int) *Category {
	root := &Category{ID: "c0", SKUs: 1}
	c := root
	for i := 1; i <= depth; i++ {
		ch := &Category{ID: fmt.Sprintf("c%d", i), SKUs: 1 + i%7}
		c.Children = []*Category{ch}
		c = ch
	}
	return root
}

// tied is the catalogue from the diagrams: two levels hold 12 SKUs each.
//
//	root (0)
//	├── home (6)       <- depth 1: 6 + 6 = 12
//	│   ├── kitchen (3)
//	│   └── bath (3)   <- depth 2: 3 + 3 + 4 + 2 = 12
//	└── garden (6)
//	    ├── tools (4)
//	    └── seeds (2)
func tied() *Category {
	return &Category{ID: "root", Children: []*Category{
		{ID: "home", SKUs: 6, Children: []*Category{{ID: "kitchen", SKUs: 3}, {ID: "bath", SKUs: 3}}},
		{ID: "garden", SKUs: 6, Children: []*Category{{ID: "tools", SKUs: 4}, {ID: "seeds", SKUs: 2}}},
	}}
}

func count(c *Category) int {
	n := 1
	for _, ch := range c.Children {
		n += count(ch)
	}
	return n
}

// The two ordered versions agree everywhere. The map version agrees only where
// no two depths tie, so it is checked on those shapes alone - chain_2k is
// excluded because its depths hold 1 to 7 SKUs each and tie constantly.
func TestAllVersionsAgree(t *testing.T) {
	noTie := []*Category{
		{ID: "only", SKUs: 3},
		catalog(1, 3, 5), catalog(3, 4, 9), catalog(4, 6, 2), flat(300, 2),
	}
	all := append([]*Category{chain(50)}, noTie...)
	for _, s := range shapes {
		all = append(all, s.root)
		if s.name != "chain_2k" {
			noTie = append(noTie, s.root)
		}
	}
	for _, tr := range append(all, tied()) {
		want := HeaviestByLevel(tr)
		if got := HeaviestBySlice(tr); got != want {
			t.Fatalf("slice says depth %d, level walk says %d", got, want)
		}
		for i := 0; i < 50; i++ {
			if got := HeaviestByMapTieBreak(tr); got != want {
				t.Fatalf("map with tie-break says depth %d, level walk says %d", got, want)
			}
		}
	}
	for _, tr := range noTie {
		if got, want := HeaviestByMap(tr), HeaviestByLevel(tr); got != want {
			t.Fatalf("map says depth %d, level walk says %d", got, want)
		}
	}
}

// chain_2k is a real catalogue shape with ties in it, and the map version
// answers it differently from call to call.
func TestMapOnTheChain(t *testing.T) {
	var tr *Category
	for _, s := range shapes {
		if s.name == "chain_2k" {
			tr = s.root
		}
	}
	seen := map[int]int{}
	for i := 0; i < 200; i++ {
		seen[HeaviestByMap(tr)]++
	}
	t.Logf("chain_2k, correct answer depth %d; map version over 200 calls gave %d different depths",
		HeaviestByLevel(tr), len(seen))
}

// On a tie the problem wants the shallowest depth. The ordered versions give
// it every time. The map version gives whichever depth Go's randomised
// iteration happened to reach first.
func TestMapAnswerChangesBetweenRunsOnATie(t *testing.T) {
	tr := tied()
	if got := HeaviestByLevel(tr); got != 1 {
		t.Fatalf("level walk: depth %d, want 1", got)
	}
	if got := HeaviestBySlice(tr); got != 1 {
		t.Fatalf("slice: depth %d, want 1", got)
	}
	seen := map[int]int{}
	for i := 0; i < 1000; i++ {
		seen[HeaviestByMap(tr)]++
	}
	t.Logf("map version over 1,000 calls on the same catalogue: %v", seen)
	if len(seen) < 2 {
		t.Fatal("the map version was stable here; the episode says it is not")
	}
}

func TestShapes(t *testing.T) {
	for _, s := range shapes {
		sums := map[int]int{}
		var walk func(c *Category, d int)
		walk = func(c *Category, d int) {
			sums[d] += c.SKUs
			for _, ch := range c.Children {
				walk(ch, d+1)
			}
		}
		walk(s.root, 0)
		t.Logf("%-12s %6d categories, %5d depths, heaviest depth %d",
			s.name, count(s.root), len(sums), HeaviestByLevel(s.root))
	}
}
