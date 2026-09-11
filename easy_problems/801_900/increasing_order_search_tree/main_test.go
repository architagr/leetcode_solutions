package increaseordersearchtree

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

// treeVals writes a tree back out in the same level-order form, with
// trailing nils dropped, so it can be compared against an example.
func treeVals(root *TreeNode) []any {
	if root == nil {
		return []any{}
	}
	out := []any{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			out = append(out, nil)
			continue
		}
		out = append(out, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	for len(out) > 0 && out[len(out)-1] == nil {
		out = out[:len(out)-1]
	}
	return out
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestIncreasingBST(t *testing.T) {
	tests := []struct {
		name string
		root []any
		want []any
	}{
		{name: "example 1", root: []any{5, 3, 6, 2, 4, nil, 8, 1, nil, nil, nil, 7, 9}, want: []any{1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6, nil, 7, nil, 8, nil, 9}},
		{name: "example 2", root: []any{5, 1, 7}, want: []any{1, nil, 5, nil, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeVals(increasingBST(buildTree(tt.root))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("increasingBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
