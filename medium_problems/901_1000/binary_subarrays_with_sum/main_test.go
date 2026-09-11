package binarysubarrayswithsum

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestNumSubarraysWithSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		goal int
		want int
	}{
		{name: "example 1", nums: []int{1, 0, 1, 0, 1}, goal: 2, want: 4},
		{name: "example 2", nums: []int{0, 0, 0, 0, 0}, goal: 0, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarraysWithSum(tt.nums, tt.goal); got != tt.want {
				t.Errorf("numSubarraysWithSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
