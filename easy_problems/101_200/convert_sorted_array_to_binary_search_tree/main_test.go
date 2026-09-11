package convert_sorted_array_to_binary_search_tree

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
func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []any
	}{
		{name: "example 1", nums: []int{-10, -3, 0, 5, 9}, want: []any{0, -3, 9, -10, nil, 5}},
		{name: "example 2", nums: []int{1, 3}, want: []any{3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeVals(SortedArrayToBST(tt.nums)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
