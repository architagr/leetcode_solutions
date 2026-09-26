package invalidationwavefront

// A cache cluster as a gossip graph: peers[n] holds the nodes n forwards an
// invalidation to. Forwarding is two-way and takes one round per hop.
type Cluster [][]int

// HopsFrom is the function the ops dashboard already had: how many rounds an
// invalidation starting at origin takes to reach each node, or -1 for a node
// it never reaches. A plain BFS.
func HopsFrom(c Cluster, origin int) []int {
	dist := make([]int, len(c))
	for i := range dist {
		dist[i] = -1
	}
	dist[origin] = 0
	queue := []int{origin}
	for i := 0; i < len(queue); i++ {
		n := queue[i]
		for _, p := range c[n] {
			if dist[p] < 0 {
				dist[p] = dist[n] + 1
				queue = append(queue, p)
			}
		}
	}
	return dist
}

// RoundsBySearchEach is what I would write with HopsFrom in hand. A node is
// invalidated by whichever origin reaches it first, so take the minimum over
// the origins at each node, then the maximum over the nodes.
//
// Returns -1 if some node is never reached.
func RoundsBySearchEach(c Cluster, origins []int) int {
	best := make([]int, len(c))
	for i := range best {
		best[i] = -1
	}
	for _, o := range origins {
		for n, d := range HopsFrom(c, o) {
			if d >= 0 && (best[n] < 0 || d < best[n]) {
				best[n] = d
			}
		}
	}
	rounds := 0
	for _, d := range best {
		if d < 0 {
			return -1
		}
		if d > rounds {
			rounds = d
		}
	}
	return rounds
}

// RoundsByOneWave puts every origin in the queue at round 0 before the walk
// starts, and runs one BFS. The queue does not care where its contents came
// from, so the frontier it expands is the combined one.
func RoundsByOneWave(c Cluster, origins []int) int {
	dist := make([]int, len(c))
	for i := range dist {
		dist[i] = -1
	}
	queue := make([]int, 0, len(c))
	for _, o := range origins {
		// Every origin is at round 0, and all of them are in the queue
		// before anything is expanded. That ordering is the whole idea.
		if dist[o] < 0 {
			dist[o] = 0
			queue = append(queue, o)
		}
	}
	rounds, reached := 0, len(queue)
	for i := 0; i < len(queue); i++ {
		n := queue[i]
		// No comparison: a FIFO queue hands nodes out in non-decreasing
		// round order, so the last one assigned is the largest.
		rounds = dist[n]
		for _, p := range c[n] {
			if dist[p] < 0 {
				dist[p] = dist[n] + 1
				queue = append(queue, p)
				reached++
			}
		}
	}
	if reached < len(c) {
		return -1
	}
	return rounds
}
