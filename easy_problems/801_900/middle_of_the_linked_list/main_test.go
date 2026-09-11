package middleofthelinkedlist

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
func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{name: "example 1", head: []int{1, 2, 3, 4, 5}, want: []int{3, 4, 5}},
		{name: "example 2", head: []int{1, 2, 3, 4, 5, 6}, want: []int{4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listVals(middleNode(buildList(tt.head))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
