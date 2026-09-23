package collapsedtreeview

// A collapsed comment thread shows one line per nesting level: the most recent
// reply at that depth, whatever branch it is in.
//
// Replies are stored oldest first, so "most recent at this depth" is the last
// comment at that depth in reading order.
type Comment struct {
	ID      string
	Replies []*Comment
}

// VisibleByAllLevels is what I would write, because I would already have the
// level-order function from last time: build every level, take the last of
// each.
//
// Two lines on top of something that exists, and obviously correct.
func VisibleByAllLevels(root *Comment) []string {
	levels := AllLevels(root)
	out := make([]string, 0, len(levels))
	for _, level := range levels {
		out = append(out, level[len(level)-1])
	}
	return out
}

// AllLevels is the level-order walk, kept whole so the episode benchmarks the
// real thing rather than a version trimmed to look bad.
func AllLevels(root *Comment) [][]string {
	if root == nil {
		return nil
	}
	var out [][]string
	queue := []*Comment{root}
	for len(queue) > 0 {
		width := len(queue)
		level := make([]string, 0, width)
		next := make([]*Comment, 0, width)
		for i := 0; i < width; i++ {
			level = append(level, queue[i].ID)
			next = append(next, queue[i].Replies...)
		}
		out = append(out, level)
		queue = next
	}
	return out
}

// VisibleByDepthMap is the other thing people write: one pass, and every
// comment overwrites the entry for its own depth, so the last writer wins.
//
// It keeps h entries rather than n, and it is correct only because the walk
// visits replies in order.
func VisibleByDepthMap(root *Comment) []string {
	var out []string
	var walk func(c *Comment, depth int)
	walk = func(c *Comment, depth int) {
		if c == nil {
			return
		}
		if len(out) == depth {
			out = append(out, "")
		}
		out[depth] = c.ID // later arrivals at this depth overwrite earlier ones
		for _, r := range c.Replies {
			walk(r, depth+1)
		}
	}
	walk(root, 0)
	return out
}

// VisibleByFirstArrival walks the replies in reverse and keeps only the first
// comment it meets at each depth.
//
// Reversing the order is the whole idea: it turns "the last comment at this
// depth" into "the first one I see at this depth", and a first is something
// you can record and stop thinking about.
func VisibleByFirstArrival(root *Comment) []string {
	var out []string
	var walk func(c *Comment, depth int)
	walk = func(c *Comment, depth int) {
		if c == nil {
			return
		}
		// True only for the first comment reached at this depth, because out
		// holds one entry per depth filled so far.
		if len(out) == depth {
			out = append(out, c.ID)
		}
		// Newest reply first. Swap this to forward order and the function
		// returns the oldest reply per level instead.
		for i := len(c.Replies) - 1; i >= 0; i-- {
			walk(c.Replies[i], depth+1)
		}
	}
	walk(root, 0)
	return out
}

// VisibleByNewestBranch is the version that looks obviously right and is not:
// follow the newest reply all the way down.
//
// It is wrong whenever the newest branch is shallower than one beside it,
// which is the common case in a thread where somebody replied once to an old
// comment and twenty times to a new one. TestNewestBranchMissesDeeperSiblings
// holds the thread it gets wrong.
func VisibleByNewestBranch(root *Comment) []string {
	var out []string
	for c := root; c != nil; {
		out = append(out, c.ID)
		if len(c.Replies) == 0 {
			break
		}
		c = c.Replies[len(c.Replies)-1]
	}
	return out
}
