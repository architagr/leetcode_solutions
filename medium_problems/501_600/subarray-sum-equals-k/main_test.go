package subarraysumequalsk

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{name: "example 1", nums: []int{1, 1, 1}, k: 2, want: 2},
		{name: "example 2", nums: []int{1, 2, 3}, k: 3, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.nums, tt.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
