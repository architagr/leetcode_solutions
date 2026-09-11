package rangeadditionii

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMaxCount(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		ops  [][]int
		want int
	}{
		{name: "example 1", m: 3, n: 3, ops: [][]int{[]int{2, 2}, []int{3, 3}}, want: 4},
		{name: "example 2", m: 3, n: 3, ops: [][]int{[]int{2, 2}, []int{3, 3}, []int{3, 3}, []int{3, 3}, []int{2, 2}, []int{3, 3}, []int{3, 3}, []int{3, 3}, []int{2, 2}, []int{3, 3}, []int{3, 3}, []int{3, 3}}, want: 4},
		{name: "example 3", m: 3, n: 3, ops: [][]int{}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCount(tt.m, tt.n, tt.ops); got != tt.want {
				t.Errorf("maxCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
