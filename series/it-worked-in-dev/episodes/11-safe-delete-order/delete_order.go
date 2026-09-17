// Package safedeleteorder compares two ways of working out the order a
// hierarchy can be deleted in when every child holds a foreign key to its
// parent - what can go first, what has to wait, and how few round trips it
// takes.
package safedeleteorder

// Record is one row in a hierarchy: a workspace, a project inside it, a board
// inside that, a card, a comment on the card. Children reference the parent,
// so a parent cannot be deleted while any child is still there.
type Record struct {
	ID       string
	Children []*Record
}

// Waves is the delete plan: one batch of ids per round trip, in the order
// they must be issued.
type Waves [][]string

// WavesByRepeatedScan builds the plan the way the constraint reads: whatever
// has nothing pointing at it can go now, so delete that, then look again.
//
// Each round walks the whole hierarchy, collects every record whose children
// are already gone, and removes them. It is the definition, executed.
func WavesByRepeatedScan(root *Record) Waves {
	if root == nil {
		return nil
	}
	gone := map[*Record]bool{}
	var waves Waves
	for {
		var wave []string
		var removed []*Record
		var visit func(r *Record)
		visit = func(r *Record) {
			for _, c := range r.Children {
				visit(c)
			}
			if gone[r] {
				return
			}
			for _, c := range r.Children {
				if !gone[c] { // something still points at this row
					return
				}
			}
			wave = append(wave, r.ID)
			removed = append(removed, r)
		}
		visit(root)
		if len(wave) == 0 {
			return waves
		}
		for _, r := range removed {
			gone[r] = true
		}
		waves = append(waves, wave)
	}
}

// WavesByDepth groups by distance from the root and deletes the deepest level
// first. It is safe - no parent is ever deleted before its children - and it
// answers a different question from the one that was asked.
//
// TestDepthAnswersADifferentQuestion holds the hierarchy where it is visibly
// the wrong answer to "what can I delete first".
func WavesByDepth(root *Record) Waves {
	if root == nil {
		return nil
	}
	var byDepth Waves
	var visit func(r *Record, depth int)
	visit = func(r *Record, depth int) {
		for len(byDepth) <= depth {
			byDepth = append(byDepth, nil)
		}
		byDepth[depth] = append(byDepth[depth], r.ID)
		for _, c := range r.Children {
			visit(c, depth+1)
		}
	}
	visit(root, 0)
	// deepest first, because a parent has to wait for its children
	for i, j := 0, len(byDepth)-1; i < j; i, j = i+1, j-1 {
		byDepth[i], byDepth[j] = byDepth[j], byDepth[i]
	}
	return byDepth
}

// WavesByOnePass computes every record's wave in a single walk, bottom up.
//
// A record's wave is one more than the largest wave among its children, which
// is its height - how far it is from the leaf furthest below it, not how far
// it is from the root.
func WavesByOnePass(root *Record) Waves {
	if root == nil {
		return nil
	}
	var waves Waves
	// wave returns the round in which r can be deleted, and files r into it
	// on the way past.
	var wave func(r *Record) int
	wave = func(r *Record) int {
		w := 0
		for _, c := range r.Children {
			if cw := wave(c) + 1; cw > w { // wait for the slowest child
				w = cw
			}
		}
		for len(waves) <= w {
			waves = append(waves, nil)
		}
		waves[w] = append(waves[w], r.ID)
		return w
	}
	wave(root)
	return waves
}

// Count is how many records the hierarchy holds.
func Count(r *Record) int {
	if r == nil {
		return 0
	}
	n := 1
	for _, c := range r.Children {
		n += Count(c)
	}
	return n
}
