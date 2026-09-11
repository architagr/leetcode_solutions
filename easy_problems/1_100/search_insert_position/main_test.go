package searchinsertposition

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "example 1", nums: []int{1, 3, 5, 6}, target: 5, want: 2},
		{name: "example 2", nums: []int{1, 3, 5, 6}, target: 2, want: 1},
		{name: "example 3", nums: []int{1, 3, 5, 6}, target: 7, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.nums, tt.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
