package reorderlist

import (
	"reflect"
	"testing"
)

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

// listVals walks a list back into the array form the examples use.
func listVals(head *ListNode) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

// reorderList works in place, so the case compares the argument after the call.
// Cases are the worked examples from the problem statement on LeetCode.
func TestReorderList(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{name: "example 1", head: []int{1, 2, 3, 4}, want: []int{1, 4, 2, 3}},
		{name: "example 2", head: []int{1, 2, 3, 4, 5}, want: []int{1, 5, 2, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := buildList(tt.head)
			reorderList(arg)
			got := listVals(arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderList() = %v, want %v", got, tt.want)
			}
		})
	}
}
