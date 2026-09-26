package circulardependency

// imports[p] lists the packages p imports. Every import has to be built
// before p is.

// OrderByPasses is what I would write. Walk the package list; build anything
// whose imports are all built; repeat until a pass builds nothing. Whatever is
// left over cannot be built, which means a cycle.
//
// It reads like the build tool's own log, and it is obviously correct.
// It returns the build order and whether every package made it in.
func OrderByPasses(imports [][]int) ([]int, bool, int) {
	n := len(imports)
	built := make([]bool, n)
	order := make([]int, 0, n)
	passes := 0
	for len(order) < n {
		passes++
		progress := false
		for p := 0; p < n; p++ {
			if built[p] {
				continue
			}
			ready := true
			for _, d := range imports[p] {
				if !built[d] {
					ready = false
					break
				}
			}
			if ready {
				built[p] = true
				order = append(order, p)
				progress = true
			}
		}
		if !progress {
			return order, false, passes
		}
	}
	return order, true, passes
}

// OrderByKahn counts, for every package, how many of its imports are not built
// yet, and queues it the moment that count reaches zero. Each import edge is
// looked at once, when the import it points at is built.
func OrderByKahn(imports [][]int) ([]int, bool) {
	n := len(imports)
	waiting := make([]int, n) // imports not yet built
	users := make([][]int, n) // reverse edges: who imports d
	for p, ds := range imports {
		waiting[p] = len(ds)
		for _, d := range ds {
			users[d] = append(users[d], p)
		}
	}
	queue := make([]int, 0, n)
	for p := 0; p < n; p++ {
		if waiting[p] == 0 {
			queue = append(queue, p)
		}
	}
	for i := 0; i < len(queue); i++ {
		for _, u := range users[queue[i]] {
			waiting[u]--
			// Reaching a package only counts down. It is queued when the
			// count hits zero: everything it imports is built.
			if waiting[u] == 0 {
				queue = append(queue, u)
			}
		}
	}
	// Anything missing never reached zero: it is in a cycle, or imports
	// something that is.
	return queue, len(queue) == n
}

// Cycle names one import cycle among the packages Kahn could not build, so
// the error can say which packages to look at instead of "cycle somewhere".
// Every stuck package imports at least one stuck package, so following stuck
// imports must revisit something; the loop from there is a cycle.
func Cycle(imports [][]int, built []int) []int {
	n := len(imports)
	done := make([]bool, n)
	for _, p := range built {
		done[p] = true
	}
	start := -1
	for p := 0; p < n; p++ {
		if !done[p] {
			start = p
			break
		}
	}
	if start < 0 {
		return nil
	}
	pos := map[int]int{}
	path := []int{}
	for p := start; ; {
		if i, seen := pos[p]; seen {
			return path[i:]
		}
		pos[p] = len(path)
		path = append(path, p)
		for _, d := range imports[p] {
			if !done[d] {
				p = d
				break
			}
		}
	}
}

// OrderByKahnFlat is OrderByKahn with the reverse edges packed into one
// array instead of a slice per package: count each package's users, lay the
// users out end to end, then fill. Same algorithm, three allocations.
func OrderByKahnFlat(imports [][]int) ([]int, bool) {
	n := len(imports)
	waiting := make([]int, n)
	start := make([]int, n+1) // users of d live in flat[start[d]:start[d+1]]
	for p, ds := range imports {
		waiting[p] = len(ds)
		for _, d := range ds {
			start[d+1]++
		}
	}
	for d := 0; d < n; d++ {
		start[d+1] += start[d]
	}
	flat := make([]int, start[n])
	fill := append([]int(nil), start[:n]...)
	for p, ds := range imports {
		for _, d := range ds {
			flat[fill[d]] = p
			fill[d]++
		}
	}
	queue := make([]int, 0, n)
	for p := 0; p < n; p++ {
		if waiting[p] == 0 {
			queue = append(queue, p)
		}
	}
	for i := 0; i < len(queue); i++ {
		d := queue[i]
		for _, u := range flat[start[d]:start[d+1]] {
			waiting[u]--
			if waiting[u] == 0 {
				queue = append(queue, u)
			}
		}
	}
	return queue, len(queue) == n
}
