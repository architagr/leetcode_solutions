// Package lopsidedcategories compares two ways of finding which parts of a
// category tree are lopsided - one branch far deeper than its siblings, which
// is what makes a single navigation page slow while the rest are fine.
package lopsidedcategories

// Category is a node in a navigation tree.
type Category struct {
	Name     string
	Children []*Category
}

// MaxTilt is how much deeper one child may be than another before the node is
// called lopsided.
const MaxTilt = 1

// depth returns how many levels sit below this category, inclusive.
//
// This is the obvious helper, and nothing is wrong with it on its own.
func depth(c *Category) int {
	if c == nil {
		return 0
	}
	best := 0
	for _, ch := range c.Children {
		if d := depth(ch); d > best {
			best = d
		}
	}
	return best + 1
}

// LopsidedByAskingTwice returns every category whose children differ in depth
// by more than MaxTilt.
//
// It reads exactly like the definition: for each category, ask how deep each
// child is, and compare. Each answer comes from a helper that is obviously
// correct, and the two concerns - measuring depth, and judging balance - stay
// separate, which is usually the thing you want.
func LopsidedByAskingTwice(root *Category) []string {
	var out []string
	var walk func(*Category)
	walk = func(c *Category) {
		if c == nil {
			return
		}
		lo, hi := -1, -1
		for _, ch := range c.Children {
			d := depth(ch)
			if lo == -1 || d < lo {
				lo = d
			}
			if d > hi {
				hi = d
			}
		}
		if len(c.Children) > 1 && hi-lo > MaxTilt {
			out = append(out, c.Name)
		}
		for _, ch := range c.Children {
			walk(ch)
		}
	}
	walk(root)
	return out
}

// LopsidedInOnePass returns the same list from a single traversal.
//
// The recursion hands back two facts instead of one: how deep this subtree is,
// and which categories inside it were lopsided. The depth a parent needs is the
// value its child already returned.
func LopsidedInOnePass(root *Category) []string {
	var out []string
	var visit func(*Category) int
	visit = func(c *Category) int {
		if c == nil {
			return 0
		}
		lo, hi := -1, -1
		for _, ch := range c.Children {
			// visit returns the child's depth AND records anything lopsided
			// beneath it, so the parent never asks a second time.
			d := visit(ch)
			if lo == -1 || d < lo {
				lo = d
			}
			if d > hi {
				hi = d
			}
		}
		if len(c.Children) > 1 && hi-lo > MaxTilt {
			out = append(out, c.Name)
		}
		if hi < 0 {
			hi = 0
		}
		return hi + 1
	}
	visit(root)
	return out
}
