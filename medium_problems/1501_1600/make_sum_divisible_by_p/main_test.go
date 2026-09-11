package makesumdivisiblebyp

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		p    int
		want int
	}{
		{name: "example 1", nums: []int{3, 1, 4, 2}, p: 6, want: 1},
		{name: "example 2", nums: []int{6, 3, 5, 2}, p: 9, want: 2},
		{name: "example 3", nums: []int{1, 2, 3}, p: 3, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubarray(tt.nums, tt.p); got != tt.want {
				t.Errorf("minSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
