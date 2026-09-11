package linked_list_cycle

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

// Cases are the worked examples from the problem statement on LeetCode.
func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
		want bool
	}{
		{name: "example 1", vals: []int{3, 2, 0, -4}, pos: 1, want: true},
		{name: "example 2", vals: []int{1, 2}, pos: 0, want: true},
		{name: "example 3", vals: []int{1}, pos: -1, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycle(buildCycle(tt.vals, tt.pos)); got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
