package whichdepthisheaviest

// A category tree: departments, then aisles, then shelves, and so on down.
// SKUs counts the products filed directly in a category, not in its children.
type Category struct {
	ID       string
	SKUs     int
	Children []*Category
}

// HeaviestByMap is what I would write. One recursive walk, add each
// category's SKUs into a map keyed by depth, then pick the biggest entry.
//
// It is short, it has no queue to get wrong, and on any catalogue without a
// tie it returns the right answer.
func HeaviestByMap(root *Category) int {
	sums := map[int]int{}
	var walk func(c *Category, depth int)
	walk = func(c *Category, depth int) {
		if c == nil {
			return
		}
		sums[depth] += c.SKUs
		for _, ch := range c.Children {
			walk(ch, depth+1)
		}
	}
	walk(root, 0)

	best, bestDepth := -1, -1
	for depth, sum := range sums {
		// Strict, so the first depth to reach the maximum keeps it. "First"
		// here means first in map iteration order, which Go randomises.
		if sum > best {
			best, bestDepth = sum, depth
		}
	}
	return bestDepth
}

// HeaviestBySlice is the same walk with the map replaced by a slice indexed
// by depth, and the comparison done by walking that slice in order.
func HeaviestBySlice(root *Category) int {
	var sums []int
	var walk func(c *Category, depth int)
	walk = func(c *Category, depth int) {
		if c == nil {
			return
		}
		// Depths arrive in a preorder walk without gaps: depth d is only
		// reached from depth d-1, so the slice never needs to skip an index.
		if len(sums) == depth {
			sums = append(sums, 0)
		}
		sums[depth] += c.SKUs
		for _, ch := range c.Children {
			walk(ch, depth+1)
		}
	}
	walk(root, 0)
	return firstMax(sums)
}

// firstMax is the smallest index holding the largest value. The strict > is
// the tie-break, and it only works because the indices are read in order.
func firstMax(sums []int) int {
	best, bestDepth := -1, -1
	for depth, sum := range sums {
		if sum > best {
			best, bestDepth = sum, depth
		}
	}
	return bestDepth
}

// HeaviestByLevel walks one level at a time and compares each total the moment
// its level is finished, so nothing per level is kept at all.
func HeaviestByLevel(root *Category) int {
	if root == nil {
		return -1
	}
	best, bestDepth := -1, -1
	level := []*Category{root}
	var next []*Category
	for depth := 0; len(level) > 0; depth++ {
		sum := 0
		next = next[:0]
		for _, c := range level {
			sum += c.SKUs
			next = append(next, c.Children...)
		}
		// Levels close in increasing depth, so strict > keeps the shallowest
		// of any tie without a word of code about ties.
		if sum > best {
			best, bestDepth = sum, depth
		}
		level, next = next, level
	}
	return bestDepth
}

// HeaviestByMapTieBreak keeps the map and says the tie-break out loud. It is
// correct, and it is the fix to reach for when the key really is sparse - a
// category name, a price band ID. Depth is not sparse.
func HeaviestByMapTieBreak(root *Category) int {
	sums := map[int]int{}
	var walk func(c *Category, depth int)
	walk = func(c *Category, depth int) {
		if c == nil {
			return
		}
		sums[depth] += c.SKUs
		for _, ch := range c.Children {
			walk(ch, depth+1)
		}
	}
	walk(root, 0)

	best, bestDepth := -1, -1
	for depth, sum := range sums {
		if sum > best || (sum == best && depth < bestDepth) {
			best, bestDepth = sum, depth
		}
	}
	return bestDepth
}
