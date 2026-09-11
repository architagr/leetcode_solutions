package findelementsinacontaminatedbinarytree

// The tree arrives with every value set to -1. Constructor recovers them
// from the shape: the root is 0, a left child is 2x+1, a right child
// 2x+2. The cases are the worked examples from the statement.

import "testing"

func contaminated(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: -1}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) {
			if vals[i] != nil {
				node.Left = &TreeNode{Val: -1}
				queue = append(queue, node.Left)
			}
			i++
		}
		if i < len(vals) {
			if vals[i] != nil {
				node.Right = &TreeNode{Val: -1}
				queue = append(queue, node.Right)
			}
			i++
		}
	}
	return root
}

func TestFindElements(t *testing.T) {
	tests := []struct {
		name   string
		root   []any
		probes map[int]bool
	}{
		{name: "example 1", root: []any{-1, nil, -1}, probes: map[int]bool{1: false, 2: true}},
		{name: "example 2", root: []any{-1, -1, -1, -1, -1}, probes: map[int]bool{1: true, 3: true, 5: false}},
		{name: "example 3", root: []any{-1, nil, -1, -1, nil, -1}, probes: map[int]bool{2: true, 3: false, 4: false, 5: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fe := Constructor(contaminated(tt.root))
			for target, want := range tt.probes {
				if got := fe.Find(target); got != want {
					t.Errorf("Find(%d) = %v, want %v", target, got, want)
				}
			}
		})
	}
}
