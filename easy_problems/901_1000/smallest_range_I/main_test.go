package smallest_range_I

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestSmallestRangeI(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{name: "example 1", nums: []int{1}, k: 0, want: 0},
		{name: "example 2", nums: []int{0, 10}, k: 2, want: 6},
		{name: "example 3", nums: []int{1, 3, 6}, k: 3, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestRangeI(tt.nums, tt.k); got != tt.want {
				t.Errorf("SmallestRangeI() = %v, want %v", got, tt.want)
			}
		})
	}
}
