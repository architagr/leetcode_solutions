package univaluedbinarytree

import "testing"

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
func TestIsUnivalTree(t *testing.T) {
	tests := []struct {
		name string
		root []any
		want bool
	}{
		{name: "example 1", root: []any{1, 1, 1, 1, 1, nil, 1}, want: true},
		{name: "example 2", root: []any{2, 2, 2, 5, 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnivalTree(buildTree(tt.root)); got != tt.want {
				t.Errorf("isUnivalTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
