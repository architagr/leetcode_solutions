package isthisconfigatree

// A config declares n components, numbered 0..n-1, and a list of links
// between them. It is only valid if the links form one tree: every component
// reachable from every other, and no loops.

// PathExists is the function the config tooling already had, the one behind
// "why can service 12 reach service 40": build the adjacency list from the
// links, then BFS from src. Tested, and used elsewhere.
func PathExists(n int, links [][2]int, src, dst int) bool {
	adj := make([][]int, n)
	for _, l := range links {
		adj[l[0]] = append(adj[l[0]], l[1])
		adj[l[1]] = append(adj[l[1]], l[0])
	}
	seen := make([]bool, n)
	seen[src] = true
	queue := []int{src}
	for i := 0; i < len(queue); i++ {
		for _, t := range adj[queue[i]] {
			if !seen[t] {
				seen[t] = true
				queue = append(queue, t)
			}
		}
	}
	return seen[dst]
}

// IsTreeByLinkCheck is what I would write. Take the links one at a time, the
// way a person reads the file. A link between two components that are already
// connected closes a loop, so reject it and say which line. At the end,
// everything must be connected to component 0.
//
// It returns -1 for a valid tree, the index of the first link that closes a
// loop, or len(links) if the config is valid so far but not connected.
func IsTreeByLinkCheck(n int, links [][2]int) int {
	for i, l := range links {
		if PathExists(n, links[:i], l[0], l[1]) {
			return i
		}
	}
	for c := 1; c < n; c++ {
		if !PathExists(n, links, 0, c) {
			return len(links)
		}
	}
	return -1
}

// IsTreeByCount uses the fact that a tree on n nodes has exactly n-1 edges.
// With that count, connected implies acyclic, so one traversal settles it and
// no loop is ever looked for.
func IsTreeByCount(n int, links [][2]int) bool {
	// Any other count is not a tree, and deciding that costs nothing.
	if len(links) != n-1 {
		return false
	}
	adj := make([][]int, n)
	for _, l := range links {
		adj[l[0]] = append(adj[l[0]], l[1])
		adj[l[1]] = append(adj[l[1]], l[0])
	}
	seen := make([]bool, n)
	seen[0] = true
	queue := make([]int, 1, n)
	for i := 0; i < len(queue); i++ {
		for _, t := range adj[queue[i]] {
			if !seen[t] {
				seen[t] = true
				queue = append(queue, t)
			}
		}
	}
	// n-1 links and all n reached: connected, so there is no room for a loop.
	return len(queue) == n
}

// IsTreeByUnionFind keeps the per-link answer - which line closes the loop -
// and makes each check nearly constant instead of a full search. It returns
// the same codes as IsTreeByLinkCheck.
func IsTreeByUnionFind(n int, links [][2]int) int {
	root := make([]int, n)
	for i := range root {
		root[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if root[x] != x {
			root[x] = find(root[x]) // point straight at the root next time
		}
		return root[x]
	}
	groups := n
	for i, l := range links {
		a, b := find(l[0]), find(l[1])
		if a == b {
			return i // already connected: this link closes a loop
		}
		root[a] = b
		groups--
	}
	if groups != 1 {
		return len(links)
	}
	return -1
}
