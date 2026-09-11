package delete_node_in_a_linked_list

import (
	"reflect"
	"testing"
)

// The function is handed the node to delete, not the head, and it has no
// way to reach the node before it. The case therefore builds the list,
// finds the node holding the given value, and checks the list afterwards.
func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		node int
		want []int
	}{
		{name: "example 1", vals: []int{4, 5, 1, 9}, node: 5, want: []int{4, 1, 9}},
		{name: "example 2", vals: []int{4, 5, 1, 9}, node: 1, want: []int{4, 5, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head, tail *ListNode
			for _, v := range tt.vals {
				n := &ListNode{Val: v}
				if head == nil {
					head = n
				} else {
					tail.Next = n
				}
				tail = n
			}
			target := head
			for target != nil && target.Val != tt.node {
				target = target.Next
			}
			if target == nil {
				t.Fatalf("value %d is not in the list", tt.node)
			}
			DeleteNode(target)

			got := []int{}
			for n := head; n != nil; n = n.Next {
				got = append(got, n.Val)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("list after DeleteNode = %v, want %v", got, tt.want)
			}
		})
	}
}
