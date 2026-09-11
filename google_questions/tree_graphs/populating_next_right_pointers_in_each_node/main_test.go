package populatingnextrightpointersineachnode

import (
	"reflect"
	"testing"
)

// buildPerfect builds the perfect binary tree the examples use, filling
// it level by level from vals.
func buildPerfect(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*Node, len(vals))
	for i, v := range vals {
		nodes[i] = &Node{Val: v}
	}
	for i := range nodes {
		if 2*i+1 < len(nodes) {
			nodes[i].Left = nodes[2*i+1]
		}
		if 2*i+2 < len(nodes) {
			nodes[i].Right = nodes[2*i+2]
		}
	}
	return nodes[0]
}

// byLevel follows Next along each level, which is what the problem asks
// the solution to wire up. The statement prints the same thing with a
// "#" ending each level.
func byLevel(root *Node) [][]int {
	out := [][]int{}
	for level := root; level != nil; level = level.Left {
		vals := []int{}
		for n := level; n != nil; n = n.Next {
			vals = append(vals, n.Val)
		}
		out = append(out, vals)
	}
	return out
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestConnect(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want [][]int
	}{
		{name: "example 1", vals: []int{1, 2, 3, 4, 5, 6, 7}, want: [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
		{name: "example 2 - empty", vals: []int{}, want: [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := byLevel(connect(buildPerfect(tt.vals)))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() levels = %v, want %v", got, tt.want)
			}
		})
	}
}
