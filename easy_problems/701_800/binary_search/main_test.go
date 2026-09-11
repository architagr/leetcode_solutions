package binary_search

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "example 1", nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, want: 4},
		{name: "example 2", nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.nums, tt.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
