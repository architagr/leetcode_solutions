package remove_duplicate_from_sorted_list

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
func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{name: "example 1", head: []int{1, 1, 2}, want: []int{1, 2}},
		{name: "example 2", head: []int{1, 1, 2, 3, 3}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listVals(DeleteDuplicates(buildList(tt.head))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
