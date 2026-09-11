package intersectionoftwolinkedlist

import "testing"

// buildIntersecting builds the two lists the examples describe: skipA
// and skipB nodes of their own, then a shared tail both end in.
func buildIntersecting(a, b, shared []int) (*ListNode, *ListNode, *ListNode) {
	var tail *ListNode
	for i := len(shared) - 1; i >= 0; i-- {
		tail = &ListNode{Val: shared[i], Next: tail}
	}
	build := func(vals []int) *ListNode {
		head := tail
		for i := len(vals) - 1; i >= 0; i-- {
			head = &ListNode{Val: vals[i], Next: head}
		}
		return head
	}
	return build(a), build(b), tail
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		name         string
		a, b, shared []int
	}{
		{name: "example 1", a: []int{4, 1}, b: []int{5, 6, 1}, shared: []int{8, 4, 5}},
		{name: "example 2", a: []int{1, 9, 1}, b: []int{3}, shared: []int{2, 4}},
		{name: "example 3 - no intersection", a: []int{2, 6, 4}, b: []int{1, 5}, shared: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA, headB, want := buildIntersecting(tt.a, tt.b, tt.shared)
			got := getIntersectionNode(headA, headB)
			if got != want {
				if want == nil {
					t.Errorf("getIntersectionNode() = %v, want nil", got.Val)
				} else {
					t.Errorf("getIntersectionNode() returned the wrong node, want the one holding %d", want.Val)
				}
			}
		})
	}
}
