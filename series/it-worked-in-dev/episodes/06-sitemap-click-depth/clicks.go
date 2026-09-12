// Package sitemapclickdepth compares ways of answering "how many clicks from
// the homepage before a visitor reaches a page with nothing below it" - the
// shallowest dead end in a site's navigation tree.
package sitemapclickdepth

// Page is a node in the navigation hierarchy. A page with no children is where
// a visitor stops: a product, an article, a form.
type Page struct {
	URL      string
	Children []*Page
}

// ClicksByRecursion returns the click depth of the shallowest dead end by
// asking every page how deep its own shallowest dead end is.
//
// This is the version that gets written first, and it is correct. It states
// the definition directly, it is four lines, and it needs no queue, no extra
// structure and no thought about ordering.
func ClicksByRecursion(root *Page) int {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1 // a dead end: one click to get here
	}
	best := -1
	for _, c := range root.Children {
		if d := ClicksByRecursion(c); best == -1 || d < best {
			best = d
		}
	}
	return best + 1
}

// ClicksByLevel returns the same number by walking the site a level at a time
// and stopping at the first dead end it meets.
//
// The first dead end found is the shallowest one, because every page one click
// away is examined before any page two clicks away. Nothing below that level is
// ever visited.
func ClicksByLevel(root *Page) int {
	if root == nil {
		return 0
	}
	queue := []*Page{root}
	for clicks := 1; len(queue) > 0; clicks++ {
		next := make([]*Page, 0, len(queue))
		for _, p := range queue {
			if len(p.Children) == 0 {
				// Everything shallower has already been checked, so this is
				// the answer - and the rest of the site is never touched.
				return clicks
			}
			next = append(next, p.Children...)
		}
		queue = next
	}
	return 0
}
