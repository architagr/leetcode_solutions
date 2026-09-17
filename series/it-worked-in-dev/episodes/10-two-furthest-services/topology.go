// Package twofurthestservices compares two ways of finding the two service
// instances furthest apart in a deployment topology - the pair whose traffic
// has to climb the most levels and come back down again.
package twofurthestservices

// Node is one element of a deployment topology: a region, a zone, a rack, or
// a service instance at the leaf. Every node has one parent and nothing loops
// back, because that is what a topology label is - region/zone/rack, read left
// to right.
type Node struct {
	Name     string
	Children []*Node
}

// index assigns every node a position and returns the nodes in that order,
// so the neighbour lists below can be slices rather than a map keyed by
// pointer. The scan version is slow enough without map lookups in its inner
// loop, and measuring Go's map is not the point of this episode.
func index(root *Node) []*Node {
	if root == nil {
		return nil
	}
	all := []*Node{root}
	for i := 0; i < len(all); i++ {
		all = append(all, all[i].Children...)
	}
	return all
}

// neighbours builds the undirected adjacency of the topology: a node's
// children and its parent, because a packet between two instances goes up to
// their common ancestor and back down again.
func neighbours(root *Node) [][]int {
	all := index(root)
	pos := make(map[*Node]int, len(all))
	for i, n := range all {
		pos[n] = i
	}
	adj := make([][]int, len(all))
	for i, n := range all {
		for _, c := range n.Children {
			j := pos[c]
			adj[i] = append(adj[i], j)
			adj[j] = append(adj[j], i)
		}
	}
	return adj
}

// FurthestFrom returns the hop count from start to the node furthest from it.
// A plain breadth-first walk: every node is one hop further out than the node
// that reached it, so the last level reached is the answer.
func FurthestFrom(adj [][]int, start int) int {
	seen := make([]bool, len(adj))
	seen[start] = true
	frontier := []int{start}
	hops := 0
	for len(frontier) > 0 {
		var next []int
		for _, n := range frontier {
			for _, m := range adj[n] {
				if !seen[m] {
					seen[m] = true
					next = append(next, m)
				}
			}
		}
		if len(next) > 0 {
			hops++
		}
		frontier = next
	}
	return hops
}

// WidestByScan answers "the two furthest apart" the way the question is
// phrased: the largest distance over every starting point.
//
// FurthestFrom is a helper you already needed and already trust, and this is
// the one composition of it that cannot be wrong. Nothing here is careless.
func WidestByScan(root *Node) int {
	adj := neighbours(root)
	best := 0
	for i := range adj {
		if d := FurthestFrom(adj, i); d > best {
			best = d
		}
	}
	return best
}

// WidestFromRoot is the shortcut that looks like it should work: the two
// deepest branches of the root, added together.
//
// It is wrong, and TestRootOnlyIsWrong shows the tree it is wrong on. It is
// here because it is the version I wrote before the benchmark, not as a
// strawman - the path it misses is one that never touches the root.
func WidestFromRoot(root *Node) int {
	if root == nil {
		return 0
	}
	top1, top2 := 0, 0
	for _, c := range root.Children {
		if h := height(c) + 1; h > top1 {
			top1, top2 = h, top1
		} else if h > top2 {
			top2 = h
		}
	}
	return top1 + top2
}

// height is the hop count from n down to the leaf furthest below it.
func height(n *Node) int {
	h := 0
	for _, c := range n.Children {
		if d := height(c) + 1; d > h {
			h = d
		}
	}
	return h
}

// WidestByOnePass walks the topology once, bottom up.
//
// Every node gets its turn as the highest point of the path, and the longest
// path bending at a node is its two tallest branches joined. Both of those
// numbers arrive from the children, so one pass produces all of them.
func WidestByOnePass(root *Node) int {
	if root == nil {
		return 0
	}
	best := 0
	var tallest func(n *Node) int
	tallest = func(n *Node) int {
		top1, top2 := 0, 0 // the two tallest branches hanging off n
		for _, c := range n.Children {
			h := tallest(c) + 1
			if h > top1 {
				top1, top2 = h, top1
			} else if h > top2 {
				top2 = h
			}
		}
		if top1+top2 > best { // the path that bends here, if it is the winner
			best = top1 + top2
		}
		return top1 // what the parent needs: how far down this branch reaches
	}
	tallest(root)
	return best
}

// Pair is the answer with its endpoints, which is what an operator actually
// wants to read: these two, that far apart.
type Pair struct {
	A, B string
	Hops int
}

// WidestPair is WidestByOnePass carrying the name of the deepest leaf up
// alongside the height, so the winning pair can be named rather than just
// counted.
func WidestPair(root *Node) Pair {
	if root == nil {
		return Pair{}
	}
	best := Pair{A: root.Name, B: root.Name}
	// reach returns how far the deepest leaf under n is, and which leaf it is.
	var reach func(n *Node) (int, string)
	reach = func(n *Node) (int, string) {
		h1, n1 := 0, n.Name
		h2, n2 := 0, n.Name
		for _, c := range n.Children {
			h, name := reach(c)
			h++
			if h > h1 {
				h1, n1, h2, n2 = h, name, h1, n1
			} else if h > h2 {
				h2, n2 = h, name
			}
		}
		if h1+h2 > best.Hops {
			best = Pair{A: n1, B: n2, Hops: h1 + h2}
		}
		return h1, n1
	}
	reach(root)
	return best
}

// Size is how many nodes the topology holds, used by the benchmark to report
// the shape it is running on.
func Size(root *Node) int {
	return len(index(root))
}

// WidestByScanReusing is the scan with its garbage removed: one visited slice
// and two frontier buffers, allocated once and reused for every starting
// point instead of once per starting point.
//
// It is the first thing a profile suggests, and the benchmark is what says
// whether the allocations were the problem or just the loudest symptom.
func WidestByScanReusing(root *Node) int {
	adj := neighbours(root)
	seen := make([]bool, len(adj))
	frontier := make([]int, 0, len(adj))
	next := make([]int, 0, len(adj))
	best := 0
	for start := range adj {
		for i := range seen {
			seen[i] = false
		}
		seen[start] = true
		frontier = append(frontier[:0], start)
		hops := 0
		for len(frontier) > 0 {
			next = next[:0]
			for _, n := range frontier {
				for _, m := range adj[n] {
					if !seen[m] {
						seen[m] = true
						next = append(next, m)
					}
				}
			}
			if len(next) > 0 {
				hops++
			}
			frontier = append(frontier[:0], next...)
		}
		if hops > best {
			best = hops
		}
	}
	return best
}
