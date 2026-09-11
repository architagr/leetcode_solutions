package remove_nth_node_from_end_of_list

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

// Cases are the worked examples from the problem statement on LeetCode.
func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head []int
		n    int
		want []int
	}{
		{name: "example 1", head: []int{1, 2, 3, 4, 5}, n: 2, want: []int{1, 2, 3, 5}},
		{name: "example 2", head: []int{1}, n: 1, want: []int{}},
		{name: "example 3", head: []int{1, 2}, n: 1, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listVals(RemoveNthFromEnd(buildList(tt.head), tt.n)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
