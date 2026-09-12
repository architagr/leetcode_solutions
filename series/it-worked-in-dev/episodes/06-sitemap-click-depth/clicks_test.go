package sitemapclickdepth

import (
	"fmt"
	"math/rand"
	"testing"
)

func chain(n int, prefix string) *Page {
	root := &Page{URL: prefix}
	cur := root
	for i := 1; i < n; i++ {
		next := &Page{URL: fmt.Sprintf("%s/%d", prefix, i)}
		cur.Children = []*Page{next}
		cur = next
	}
	return root
}

func balanced(depth, width int, prefix string) *Page {
	p := &Page{URL: prefix}
	if depth == 0 {
		return p
	}
	for w := 0; w < width; w++ {
		p.Children = append(p.Children, balanced(depth-1, width, fmt.Sprintf("%s/%d", prefix, w)))
	}
	return p
}

func TestBothAgree(t *testing.T) {
	// a shallow page beside a very deep section: the case the whole episode
	// is about
	lopsided := &Page{URL: "/", Children: []*Page{chain(400, "/docs"), {URL: "/contact"}}}

	cases := []struct {
		name string
		root *Page
		want int
	}{
		{"homepage only", &Page{URL: "/"}, 1},
		{"one level", balanced(1, 4, "/"), 2},
		{"balanced 3x3", balanced(3, 3, "/"), 4},
		{"a chain of 10", chain(10, "/"), 10},
		{"shallow beside deep", lopsided, 2},
		{"deep section first", &Page{URL: "/", Children: []*Page{chain(400, "/docs"), {URL: "/x"}}}, 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ClicksByRecursion(c.root); got != c.want {
				t.Errorf("recursion = %d, want %d", got, c.want)
			}
			if got := ClicksByLevel(c.root); got != c.want {
				t.Errorf("by level = %d, want %d", got, c.want)
			}
		})
	}
}

// Day 7 warns that on a binary tree a missing child returns 0 and wins the
// comparison, reporting a dead end that is not there. This model cannot have
// that bug - a page with one child has one entry in Children, not one entry
// and one nil - and the test exists to prove the claim rather than assert it.
func TestSingleChildPagesAreNotDeadEnds(t *testing.T) {
	// every page on this route has exactly one child, so a buggy version that
	// let a missing sibling count as depth 0 would answer 1 or 2
	root := chain(6, "/")
	if got, want := ClicksByRecursion(root), 6; got != want {
		t.Errorf("recursion = %d, want %d", got, want)
	}
	if got, want := ClicksByLevel(root), 6; got != want {
		t.Errorf("by level = %d, want %d", got, want)
	}
}

func TestBothAgreeOnRandomSites(t *testing.T) {
	rng := rand.New(rand.NewSource(37))
	for trial := 0; trial < 2000; trial++ {
		root := &Page{URL: "/"}
		all := []*Page{root}
		for i := 1; i < rng.Intn(70)+2; i++ {
			p := all[rng.Intn(len(all))]
			c := &Page{URL: fmt.Sprintf("/p%d", i)}
			p.Children = append(p.Children, c)
			all = append(all, c)
		}
		if a, b := ClicksByRecursion(root), ClicksByLevel(root); a != b {
			t.Fatalf("trial %d: recursion=%d byLevel=%d", trial, a, b)
		}
	}
}

// The claim in the write-up is that the level walk stops early. Counted here:
// on a site with a dead end two clicks away, it must not touch the deep branch.
func TestLevelWalkStopsEarly(t *testing.T) {
	deep := chain(500, "/docs")
	root := &Page{URL: "/", Children: []*Page{deep, {URL: "/contact"}}}

	visited := 0
	var countRecursion func(*Page) int
	countRecursion = func(p *Page) int {
		visited++
		if len(p.Children) == 0 {
			return 1
		}
		best := -1
		for _, c := range p.Children {
			if d := countRecursion(c); best == -1 || d < best {
				best = d
			}
		}
		return best + 1
	}
	countRecursion(root)
	if visited != 502 {
		t.Fatalf("recursion visited %d pages, expected every one (502)", visited)
	}

	seen := 0
	queue := []*Page{root}
	for len(queue) > 0 {
		var next []*Page
		stop := false
		for _, p := range queue {
			seen++
			if len(p.Children) == 0 {
				stop = true
				break
			}
			next = append(next, p.Children...)
		}
		if stop {
			break
		}
		queue = next
	}
	if seen > 4 {
		t.Errorf("level walk touched %d pages, expected at most 4", seen)
	}
	t.Logf("recursion visited %d pages, level walk visited %d", visited, seen)
}
