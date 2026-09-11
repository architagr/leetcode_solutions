package clonegraph

import (
	"reflect"
	"testing"
)

// buildGraph reads the adjacency list the examples use, where entry i
// lists the 1-based neighbours of node i+1.
func buildGraph(adj [][]int) *Node {
	if len(adj) == 0 {
		return nil
	}
	nodes := make([]*Node, len(adj))
	for i := range adj {
		nodes[i] = &Node{Val: i + 1}
	}
	for i, neighbours := range adj {
		for _, n := range neighbours {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[n-1])
		}
	}
	return nodes[0]
}

// adjacency walks a graph back into the example's list form.
func adjacency(start *Node) [][]int {
	seen := map[int]*Node{}
	var walk func(*Node)
	walk = func(n *Node) {
		if n == nil || seen[n.Val] != nil {
			return
		}
		seen[n.Val] = n
		for _, nb := range n.Neighbors {
			walk(nb)
		}
	}
	walk(start)
	out := make([][]int, len(seen))
	for val, n := range seen {
		vals := []int{}
		for _, nb := range n.Neighbors {
			vals = append(vals, nb.Val)
		}
		out[val-1] = vals
	}
	return out
}

// The clone must have the same shape and share no nodes with the
// original, which is the whole point of the problem.
func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name string
		adj  [][]int
	}{
		{name: "example 1", adj: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}},
		{name: "example 2 - single node", adj: [][]int{{}}},
		{name: "example 3 - empty", adj: [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adj)
			got := cloneGraph(original)
			if original == nil {
				if got != nil {
					t.Fatal("cloneGraph(nil) should be nil")
				}
				return
			}
			if got == original {
				t.Fatal("cloneGraph returned the original node, not a copy")
			}
			if !reflect.DeepEqual(adjacency(got), adjacency(original)) {
				t.Errorf("clone = %v, want %v", adjacency(got), adjacency(original))
			}
			// No node of the clone may be a node of the original.
			origNodes := map[*Node]bool{}
			var walk func(*Node, map[int]bool)
			walk = func(n *Node, seen map[int]bool) {
				if n == nil || seen[n.Val] {
					return
				}
				seen[n.Val] = true
				origNodes[n] = true
				for _, nb := range n.Neighbors {
					walk(nb, seen)
				}
			}
			walk(original, map[int]bool{})
			seen := map[int]bool{}
			var check func(*Node)
			check = func(n *Node) {
				if n == nil || seen[n.Val] {
					return
				}
				seen[n.Val] = true
				if origNodes[n] {
					t.Errorf("clone reuses the original node %d", n.Val)
				}
				for _, nb := range n.Neighbors {
					check(nb)
				}
			}
			check(got)
		})
	}
}
