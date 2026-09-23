package orgchartbylevel

// The org chart page renders one level at a time: the person at the top, then
// their directs, then the level below that, each appearing as it is ready.
type Person struct {
	Name    string
	Reports []*Person
}

// LevelsByDepthKeyedWalk is the version I would write, and it is the version
// day 29 of the challenge publishes: one depth-first walk, with every node
// appending into the slot for its own depth.
//
// It is O(n), it allocates one slice per level, and there is no queue to
// reason about. Which branch a node arrived from stops mattering, because the
// accumulator is keyed on depth rather than on visit order.
func LevelsByDepthKeyedWalk(root *Person) [][]string {
	var out [][]string
	var walk func(p *Person, depth int)
	walk = func(p *Person, depth int) {
		if p == nil {
			return
		}
		if len(out) == depth {
			out = append(out, nil) // first node to reach this depth opens it
		}
		out[depth] = append(out[depth], p.Name)
		for _, r := range p.Reports {
			walk(r, depth+1)
		}
	}
	walk(root, 0)
	return out
}

// LevelsByQueue produces the identical answer by finishing each level before
// starting the next one.
//
// The trick is the single read of len(queue) at the top of the loop: whatever
// is in the queue at that moment is exactly this level, because everything
// appended during the loop belongs to the next one.
func LevelsByQueue(root *Person) [][]string {
	if root == nil {
		return nil
	}
	var out [][]string
	queue := []*Person{root}
	for len(queue) > 0 {
		width := len(queue) // this level, and nothing else
		level := make([]string, 0, width)
		for i := 0; i < width; i++ {
			p := queue[i]
			level = append(level, p.Name)
			queue = append(queue, p.Reports...)
		}
		queue = queue[width:]
		out = append(out, level)
	}
	return out
}

// StreamLevels hands each level to emit the moment that level is complete,
// and stops walking when emit returns false.
//
// This is the shape the page actually wants, and it is the one the recursive
// version cannot offer: a level is finished here because the queue said so,
// not because the whole tree ran out.
func StreamLevels(root *Person, emit func(level []string) bool) {
	if root == nil {
		return
	}
	queue := []*Person{root}
	for len(queue) > 0 {
		width := len(queue)
		level := make([]string, 0, width)
		for i := 0; i < width; i++ {
			level = append(level, queue[i].Name)
		}
		if !emit(level) {
			return
		}
		next := make([]*Person, 0, width)
		for i := 0; i < width; i++ {
			next = append(next, queue[i].Reports...)
		}
		queue = next
	}
}

// FirstLevelsByQueue is what a handler for "give me the top k levels" looks
// like once StreamLevels exists.
func FirstLevelsByQueue(root *Person, k int) [][]string {
	out := make([][]string, 0, k)
	StreamLevels(root, func(level []string) bool {
		out = append(out, level)
		return len(out) < k
	})
	return out
}

// FirstLevelsByDepthKeyedWalk is the same request answered by the recursive
// version, which has no way to stop early: every node has to be visited
// before any level can be called finished.
func FirstLevelsByDepthKeyedWalk(root *Person, k int) [][]string {
	all := LevelsByDepthKeyedWalk(root)
	if len(all) > k {
		all = all[:k]
	}
	return all
}
