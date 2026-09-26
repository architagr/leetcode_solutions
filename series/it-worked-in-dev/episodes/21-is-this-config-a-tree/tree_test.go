package isthisconfigatree

import (
	"math/rand"
	"testing"
)

// randomTree is a valid config: each component links to one earlier one, in
// shuffled order, the way a config file is written by hand over time.
func randomTree(n int, seed int64) [][2]int {
	r := rand.New(rand.NewSource(seed))
	links := make([][2]int, 0, n-1)
	for i := 1; i < n; i++ {
		links = append(links, [2]int{i, r.Intn(i)})
	}
	r.Shuffle(len(links), func(i, j int) { links[i], links[j] = links[j], links[i] })
	return links
}

// withLoop is a valid tree with one extra link added at the end of the file.
func withLoop(n int, seed int64) [][2]int {
	links := randomTree(n, seed)
	return append(links, [2]int{n - 1, n / 2})
}

// split is a valid tree with one link removed: two islands.
func split(n int, seed int64) [][2]int {
	links := randomTree(n, seed)
	return links[:len(links)-1]
}

// swapped has the right count and a loop: one link is replaced with one that
// closes a cycle, so the count shortcut has to do its traversal to catch it.
func swapped(n int, seed int64) [][2]int {
	links := randomTree(n, seed)
	links[len(links)-1] = [2]int{links[0][0], links[0][1]} // a duplicate link: a loop of two
	return links
}

func TestAllThreeAgree(t *testing.T) {
	type tc struct {
		n     int
		links [][2]int
	}
	cases := []tc{
		{1, nil}, {2, [][2]int{{0, 1}}}, {2, nil}, {3, [][2]int{{0, 1}, {1, 2}, {2, 0}}},
		{4, [][2]int{{0, 1}, {2, 3}}},
		{200, randomTree(200, 1)}, {200, withLoop(200, 2)}, {200, split(200, 3)}, {200, swapped(200, 4)},
	}
	for _, s := range shapes {
		if s.n <= 5000 {
			cases = append(cases, tc{s.n, s.links})
		}
	}
	for _, x := range cases {
		a := IsTreeByLinkCheck(x.n, x.links)
		b := IsTreeByUnionFind(x.n, x.links)
		c := IsTreeByCount(x.n, x.links)
		if a != b {
			t.Fatalf("n=%d: link check says %d, union-find says %d", x.n, a, b)
		}
		if (a == -1) != c {
			t.Fatalf("n=%d: link check says %d, count says tree=%v", x.n, a, c)
		}
	}
}

// The loop is on the last line, so the link check reads the whole file before
// it can say so. The count shortcut decides from the length alone.
func TestLoopReportedAtItsLine(t *testing.T) {
	links := withLoop(1000, 9)
	if got := IsTreeByLinkCheck(1000, links); got != len(links)-1 {
		t.Fatalf("link check reports line %d, the loop is on line %d", got, len(links)-1)
	}
	if got := IsTreeByUnionFind(1000, links); got != len(links)-1 {
		t.Fatalf("union-find reports line %d, the loop is on line %d", got, len(links)-1)
	}
}

func TestSearches(t *testing.T) {
	for _, s := range shapes {
		searches := 0
		code := IsTreeByUnionFind(s.n, s.links)
		switch {
		case code == -1:
			searches = len(s.links) + s.n - 1
		case code == len(s.links):
			searches = len(s.links) + 1 // the first unreachable component ends it early at best
		default:
			searches = code + 1
		}
		t.Logf("%-12s n=%6d links=%6d  valid=%-5v  link check runs at least %6d full searches",
			s.name, s.n, len(s.links), code == -1, searches)
	}
}

// Links in this config are parent -> child. Component 2 names two parents,
// 0 and 1, and nothing else links anywhere: three components, two links, one
// connected piece with no loop. Every undirected check calls it a tree. As a
// hierarchy it has two roots and a child with two parents.
func TestTwoParentsPassesTheUndirectedCheck(t *testing.T) {
	links := [][2]int{{0, 2}, {1, 2}}
	t.Logf("count says tree=%v, union-find says %d (-1 is a tree)",
		IsTreeByCount(3, links), IsTreeByUnionFind(3, links))
	if !IsTreeByCount(3, links) || IsTreeByUnionFind(3, links) != -1 {
		t.Fatal("an undirected check rejected it; the episode says it does not")
	}
	parents := map[int]int{}
	for _, l := range links {
		parents[l[1]]++
	}
	if parents[2] != 2 {
		t.Fatal("component 2 should have two parents")
	}
}
