package howmanynetworks

// A service mesh as an adjacency list: adj[s] holds every service that s has
// a link with. Links are two-way - if a can reach b over the mesh, b can
// reach a - which is the property this whole episode turns on.
type Mesh [][]int

// Reachable is the blast-radius function: every service a failure in s could
// reach. It already existed, it is tested, and a page in the admin UI calls it.
func Reachable(m Mesh, s int) []int {
	// A map, so a small blast radius costs a small search. A slice sized to
	// the whole mesh would make every call pay for every service.
	seen := map[int]bool{s: true}
	out := []int{s}
	for i := 0; i < len(out); i++ {
		for _, t := range m[out[i]] {
			if !seen[t] {
				seen[t] = true
				out = append(out, t)
			}
		}
	}
	return out
}

// NetworksByReach is what I would write with Reachable already in hand: every
// service's network is the set it can reach, so name each network by the
// smallest service in it and count the distinct names.
func NetworksByReach(m Mesh) int {
	names := map[int]bool{}
	for s := range m {
		reach := Reachable(m, s)
		smallest := reach[0]
		for _, t := range reach {
			if t < smallest {
				smallest = t
			}
		}
		names[smallest] = true
	}
	return len(names)
}

// NetworksBySweep scans the services once, and every time it meets one that no
// earlier search has reached, counts a network and consumes all of it.
func NetworksBySweep(m Mesh) int {
	// One visited set for the whole scan, not one per search. A service
	// reached from s is in s's network, so it can never start a new one.
	seen := make([]bool, len(m))
	queue := make([]int, 0, len(m))
	count := 0
	for s := range m {
		if seen[s] {
			continue
		}
		count++
		seen[s] = true
		queue = append(queue[:0], s)
		for i := 0; i < len(queue); i++ {
			for _, t := range m[queue[i]] {
				if !seen[t] {
					seen[t] = true
					queue = append(queue, t)
				}
			}
		}
	}
	return count
}
