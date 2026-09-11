package maximumuniquesubarraysumafterdeletion

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMaxSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 2, 3, 4, 5}, want: 15},
		{name: "example 2", nums: []int{1, 1, 0, 1, 1}, want: 1},
		{name: "example 3", nums: []int{1, 2, -1, -2, 1, 0, -1}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.nums); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
