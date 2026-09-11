package convertsortedlisttobinarysearchtree

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

// buildList turns the array form the examples use into a linked list.
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestSortedListToBST(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []any
	}{
		{name: "example 1", head: []int{-10, -3, 0, 5, 9}, want: []any{0, -3, 9, -10, nil, 5}},
		{name: "example 2", head: []int{}, want: []any{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeVals(sortedListToBST(buildList(tt.head))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
