package orgchartbylevel

// VisitsBeforeLevelComplete counts, for level k:
//
//	lastSeen - the visit on which the depth-first walk reaches the last member
//	           of level k, which is when it has the answer
//	declared - the visit on which it can say so, which is the end of the walk,
//	           because nothing else rules out another node at depth k
//	queued   - what the queue looks at before it finishes level k
//
// The gap between the first two is the episode. These are counts rather than
// durations, so they do not move between runs.
func VisitsBeforeLevelComplete(root *Person, k int) (lastSeen, declared, queued int) {
	// The depth-first walk, in the order it actually visits people. The last
	// member of level k can be anywhere in that order, because the walk
	// descends one branch to the bottom before starting the next.
	seen := 0
	lastAtK := 0
	var walk func(p *Person, depth int)
	walk = func(p *Person, depth int) {
		if p == nil {
			return
		}
		seen++
		if depth == k {
			lastAtK = seen
		}
		for _, r := range p.Reports {
			walk(r, depth+1)
		}
	}
	walk(root, 0)
	declared = seen // the walk can only declare a level finished once it ends

	// The queue finishes level k after it has looked at levels 0 through k
	// and nothing else.
	queued = 0
	level := 0
	if root != nil {
		frontier := []*Person{root}
		for len(frontier) > 0 && level <= k {
			queued += len(frontier)
			var next []*Person
			for _, p := range frontier {
				next = append(next, p.Reports...)
			}
			frontier = next
			level++
		}
	}
	return lastAtK, declared, queued
}
