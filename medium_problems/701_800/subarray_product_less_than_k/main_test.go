package subarrayproductlessthank

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestNumSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{name: "example 1", nums: []int{10, 5, 2, 6}, k: 100, want: 8},
		{name: "example 2", nums: []int{1, 2, 3}, k: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.nums, tt.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
