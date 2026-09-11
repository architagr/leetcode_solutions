package n_ary_tree_preorder_traversal

import (
	"reflect"
	"testing"
)

// buildNary reads the level-order form the examples use, where nil ends
// a node's list of children:
//
//	[1,null,3,2,4,null,5,6]
func buildNary(vals []any) *Node {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &Node{Val: vals[0].(int)}
	queue := []*Node{root}
	i := 1
	if i < len(vals) && vals[i] == nil {
		i++ // the separator that follows the root
	}
	for len(queue) > 0 && i < len(vals) {
		parent := queue[0]
		queue = queue[1:]
		for i < len(vals) && vals[i] != nil {
			child := &Node{Val: vals[i].(int)}
			parent.Children = append(parent.Children, child)
			queue = append(queue, child)
			i++
		}
		i++ // skip the nil separating this parent from the next
	}
	return root
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestPreorder(t *testing.T) {
	tests := []struct {
		name string
		root []any
		want []int
	}{
		{name: "example 1", root: []any{1, nil, 3, 2, 4, nil, 5, 6}, want: []int{1, 3, 5, 6, 2, 4}},
		{name: "example 2", root: []any{1, nil, 2, 3, 4, 5, nil, nil, 6, 7, nil, 8, nil, 9, 10, nil, nil, 11, nil, 12, nil, 13, nil, nil, 14}, want: []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Preorder(buildNary(tt.root)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
