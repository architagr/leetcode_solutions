package binary_tree_postorder_traversal

import (
	"reflect"
	"testing"
)

// buildTree rebuilds a tree from the level-order form the examples use,
// where nil marks a missing child.
func buildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) {
			if vals[i] != nil {
				node.Left = &TreeNode{Val: vals[i].(int)}
				queue = append(queue, node.Left)
			}
			i++
		}
		if i < len(vals) {
			if vals[i] != nil {
				node.Right = &TreeNode{Val: vals[i].(int)}
				queue = append(queue, node.Right)
			}
			i++
		}
	}
	return root
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root []any
		want []int
	}{
		{name: "example 1", root: []any{1, nil, 2, 3}, want: []int{3, 2, 1}},
		{name: "example 2", root: []any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}, want: []int{4, 6, 7, 5, 2, 9, 8, 3, 1}},
		{name: "example 3", root: []any{}, want: []int{}},
		{name: "example 4", root: []any{1}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PostorderTraversal(buildTree(tt.root)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
