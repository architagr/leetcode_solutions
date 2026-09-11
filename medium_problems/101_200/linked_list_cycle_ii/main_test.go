package linked_list_cycle_2

import "testing"

// buildCycle links vals into a list and, when pos is not -1, points the
// tail back at index pos. The examples describe the cycle that way
// because the bracket form cannot show it.
func buildCycle(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &ListNode{Val: v}
	}
	for i := 0; i+1 < len(nodes); i++ {
		nodes[i].Next = nodes[i+1]
	}
	if pos >= 0 && pos < len(nodes) {
		nodes[len(nodes)-1].Next = nodes[pos]
	}
	return nodes[0]
}

// The answer is a node, so the case checks which node came back by its
// index in the original list rather than comparing pointers.
func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
	}{
		{name: "example 1", vals: []int{3, 2, 0, -4}, pos: 1},
		{name: "example 2", vals: []int{1, 2}, pos: 0},
		{name: "example 3", vals: []int{1}, pos: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildCycle(tt.vals, tt.pos)
			got := DetectCycle(head)
			if tt.pos < 0 {
				if got != nil {
					t.Errorf("DetectCycle() = %v, want nil", got.Val)
				}
				return
			}
			var want *ListNode
			for n, i := head, 0; i <= tt.pos; n, i = n.Next, i+1 {
				want = n
			}
			if got != want {
				t.Errorf("DetectCycle() returned the wrong node, want the one at index %d (value %d)", tt.pos, want.Val)
			}
		})
	}
}
