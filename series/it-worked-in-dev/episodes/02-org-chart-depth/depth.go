// Package orgchartdepth compares two ways of labelling every person in an org
// chart with their level, which is what an indented list needs to render.
package orgchartdepth

// Person is a node in the reporting tree. ManagerID is empty for the one at
// the top, which is what makes walking upward terminate.
type Person struct {
	ID        string
	Name      string
	ManagerID string
}

// LevelsByWalkingUp returns each person's depth by climbing from them to the
// top and counting the hops.
//
// This is the version that gets written first, and it is a reasonable one. It
// needs nothing but the list you already have - every row carries its own
// manager, so each answer is derivable on its own, with no ordering
// requirement and no tree to build. It is also the only version you can write
// if you are handed one person and asked how deep they are.
func LevelsByWalkingUp(people []Person) map[string]int {
	byID := make(map[string]Person, len(people))
	for _, p := range people {
		byID[p.ID] = p
	}
	out := make(map[string]int, len(people))
	for _, p := range people {
		level, cur := 0, p
		for cur.ManagerID != "" {
			next, ok := byID[cur.ManagerID]
			if !ok {
				break // a dangling manager reference: stop rather than loop
			}
			cur = next
			level++
		}
		out[p.ID] = level
	}
	return out
}

// LevelsByLevelOrder returns the same map by walking down from the top once,
// a level at a time.
//
// Everyone on a level is one deeper than the level above, so the depth is
// known on arrival and never computed per person.
func LevelsByLevelOrder(people []Person) map[string]int {
	reports := make(map[string][]string, len(people))
	var roots []string
	for _, p := range people {
		if p.ManagerID == "" {
			roots = append(roots, p.ID)
			continue
		}
		reports[p.ManagerID] = append(reports[p.ManagerID], p.ID)
	}

	out := make(map[string]int, len(people))
	queue := roots
	for level := 0; len(queue) > 0; level++ {
		// Everything currently in the queue sits at this level, so the whole
		// row is labelled before any of their reports are queued. That is what
		// makes the level a property of the pass rather than of each person.
		next := make([]string, 0, len(queue))
		for _, id := range queue {
			out[id] = level
			next = append(next, reports[id]...)
		}
		queue = next
	}
	return out
}
