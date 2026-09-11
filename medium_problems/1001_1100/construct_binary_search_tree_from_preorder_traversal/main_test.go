package constructbinaryserchtreefrompreordertraversal

import (
	"reflect"
	"testing"
)

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
func TestBstFromPreorder(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		want     []any
	}{
		{name: "example 1", preorder: []int{8, 5, 1, 7, 10, 12}, want: []any{8, 5, 10, 1, 7, nil, 12}},
		{name: "example 2", preorder: []int{1, 3}, want: []any{1, nil, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeVals(bstFromPreorder(tt.preorder)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
